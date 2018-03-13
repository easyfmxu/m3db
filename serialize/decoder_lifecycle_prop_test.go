// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package serialize

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/m3db/m3x/checked"
	"github.com/m3db/m3x/ident"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/commands"
	"github.com/leanovate/gopter/gen"
)

func TestDecoderLifecycle(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	properties := gopter.NewProperties(parameters)

	comms := decoderCommandsFunctor(t)
	properties.Property("Decoder Lifecycle Invariants", commands.Prop(comms))

	reporter := gopter.NewFormatedReporter(true, 160, os.Stdout)
	if !properties.Run(reporter) {
		t.Fail()
	}
}

var decoderCommandsFunctor = func(t *testing.T) *commands.ProtoCommands {
	return &commands.ProtoCommands{
		NewSystemUnderTestFunc: func(initialState commands.State) commands.SystemUnderTest {
			sut := initialState.(*multiDecoderState)
			d := newTestTagDecoder()
			d.Reset(sut.initBytes)
			if err := d.Err(); err != nil {
				panic(err)
			}
			return &multiDecoderSystem{
				primary: d,
			}
		},
		DestroySystemUnderTestFunc: func(s commands.SystemUnderTest) {
			sys := s.(*multiDecoderSystem)
			sys.primary.Finalize()
			for _, dupe := range sys.duplicates {
				dupe.Finalize()
			}
		},
		InitialStateGen: newDecoderState(),
		InitialPreConditionFunc: func(s commands.State) bool {
			return s != nil
		},
		GenCommandFunc: func(state commands.State) gopter.Gen {
			return gen.OneGenOf(
				gen.Const(nextCmd),
				gen.Const(errCmd),
				gen.Const(remainingCmd),
				gen.Const(duplicateCmd),
				gen.Const(swapToDuplicateCmd),
			)
		},
	}
}

func newDecoderState() gopter.Gen {
	return anyASCIITags().Map(
		func(tags ident.Tags) *multiDecoderState {
			enc := newTestTagEncoder()
			if err := enc.Encode(ident.NewTagSliceIterator(tags)); err != nil {
				return nil
			}
			b, ok := enc.Data()
			if !ok {
				return nil
			}
			data := checked.NewBytes(b.Get(), nil)
			return &multiDecoderState{
				tags:      tags,
				initBytes: data,
				numRefs:   1,
				primary: decoderState{
					numTags: len(tags),
				},
			}
		},
	)
}

type multiDecoderSystem struct {
	primary    TagDecoder
	duplicates []TagDecoder
}

type multiDecoderState struct {
	initBytes  checked.Bytes
	numRefs    int
	tags       ident.Tags
	primary    decoderState
	duplicates []decoderState
}

type decoderState struct {
	numTags           int
	idx               int
	nextHasBeenCalled bool
}

func (d decoderState) String() string {
	return fmt.Sprintf("[ nextCalled=%v, idx=%d ]", d.nextHasBeenCalled, d.idx)
}

func (d *decoderState) numRemaining() int {
	if !d.nextHasBeenCalled {
		return d.numTags
	}
	remain := d.numTags - (d.idx + 1)
	if remain >= 0 {
		return remain
	}
	return 0
}

func (d *multiDecoderState) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("[ numRefs=%d, tags=[%s], primary=%s ",
		d.numRefs, tagsToString(d.tags), d.primary.String()))

	for i, dupe := range d.duplicates {
		buf.WriteString(fmt.Sprintf(", dupe_%d=%s ", (i + 1), dupe.String()))
	}

	buf.WriteString("]")
	return buf.String()
}

func tagsToString(tags ident.Tags) string {
	var tagBuffer bytes.Buffer
	for i, t := range tags {
		if i != 0 {
			tagBuffer.WriteString(", ")
		}
		tagBuffer.WriteString(t.Name.String())
		tagBuffer.WriteString("=")
		tagBuffer.WriteString(t.Value.String())
	}
	return tagBuffer.String()
}

type systemAndResult struct {
	system *multiDecoderSystem
	result commands.Result
}

var errCmd = &commands.ProtoCommand{
	Name: "Err",
	RunFunc: func(s commands.SystemUnderTest) commands.Result {
		sys := s.(*multiDecoderSystem)
		d := sys.primary
		return &systemAndResult{
			system: sys,
			result: d.Err(),
		}
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		res := result.(*systemAndResult)
		if res.result == nil {
			return &gopter.PropResult{Status: gopter.PropTrue}
		}

		err := res.result.(error)
		return &gopter.PropResult{
			Status: gopter.PropError,
			Error:  fmt.Errorf("received error [ err = %v, state = [%s] ]", err, state.(decoderState)),
		}
	},
}

var nextCmd = &commands.ProtoCommand{
	Name: "Next",
	RunFunc: func(s commands.SystemUnderTest) commands.Result {
		sys := s.(*multiDecoderSystem)
		d := sys.primary
		return &systemAndResult{
			system: sys,
			result: d.Next(),
		}
	},
	NextStateFunc: func(state commands.State) commands.State {
		s := state.(*multiDecoderState)
		if !s.primary.nextHasBeenCalled {
			s.primary.nextHasBeenCalled = true
			if len(s.tags) > 0 {
				// i.e. only increment tag references the first time we allocate
				s.numRefs += 2 // tagName & tagValue
			}
		} else {
			s.primary.idx++
		}
		// when we have gone past the end, remove references to tagName/tagValue
		if s.primary.numRemaining() == -1 {
			if len(s.tags) > 0 {
				s.numRefs -= 2
			}
		}
		return s
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		res := result.(*systemAndResult)
		decState := state.(*multiDecoderState)
		if decState.primary.numRemaining() > 0 && !res.result.(bool) {
			// ensure we were told the correct value for Next()
			return &gopter.PropResult{
				Status: gopter.PropError,
				Error:  fmt.Errorf("received invalid Next()"),
			}
		}
		// ensure hold the correct number of references for underlying bytes
		dec := res.system.primary.(*decoder)
		if dec.checkedData.NumRef() != decState.numRefs {
			return &gopter.PropResult{
				Status: gopter.PropError,
				Error: fmt.Errorf(
					"received invalid number of refs [ expected = %d observed = %d state = [%s] ]",
					decState.numRefs, dec.checkedData.NumRef(), decState),
			}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
}

var remainingCmd = &commands.ProtoCommand{
	Name: "Remaining",
	RunFunc: func(s commands.SystemUnderTest) commands.Result {
		sys := s.(*multiDecoderSystem)
		d := sys.primary
		return &systemAndResult{
			system: sys,
			result: d.Remaining(),
		}
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		decState := state.(*multiDecoderState)
		remain := result.(*systemAndResult).result.(int)
		if remain != decState.primary.numRemaining() {
			return &gopter.PropResult{
				Status: gopter.PropError,
				Error: fmt.Errorf("received invalid Remain [ expected=%d, observed=%d ]",
					decState.primary.numRemaining(), remain),
			}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
}

var duplicateCmd = &commands.ProtoCommand{
	Name: "Duplicate",
	RunFunc: func(s commands.SystemUnderTest) commands.Result {
		sys := s.(*multiDecoderSystem)
		d := sys.primary
		dupe := d.Duplicate().(TagDecoder)
		sys.duplicates = append(sys.duplicates, dupe)
		return &systemAndResult{
			system: sys,
			result: dupe,
		}
	},
	NextStateFunc: func(state commands.State) commands.State {
		s := state.(*multiDecoderState)
		if s.numRefs > 0 {
			// i.e. we have a checked bytes still present, so we
			// atleast make another reference to it.
			s.numRefs++
		}
		// if we have any current tags, we should inc ref by 2 because of it
		if s.primary.nextHasBeenCalled && s.primary.numRemaining() >= 0 && len(s.tags) > 0 {
			s.numRefs += 2
		}
		s.duplicates = append(s.duplicates, s.primary)
		return s
	},
	PostConditionFunc: func(state commands.State, result commands.Result) *gopter.PropResult {
		decState := state.(*multiDecoderState)
		dec := result.(*systemAndResult).system.primary.(*decoder)
		// ensure hold the correct number of references for underlying bytes
		if dec.checkedData.NumRef() != decState.numRefs {
			return &gopter.PropResult{
				Status: gopter.PropError,
				Error: fmt.Errorf(
					"received invalid number of refs [ expected = %d observed = %d state = [%s] ]",
					decState.numRefs, dec.checkedData.NumRef(), decState),
			}
		}
		return &gopter.PropResult{Status: gopter.PropTrue}
	},
}

// swapToDuplicate swaps the current system under test to be operating on the most recent Duplicate
// we created (if any).
var swapToDuplicateCmd = &commands.ProtoCommand{
	Name: "swapToDuplicate",
	PreConditionFunc: func(s commands.State) bool {
		state := s.(*multiDecoderState)
		return len(state.duplicates) > 0
	},
	NextStateFunc: func(state commands.State) commands.State {
		s := state.(*multiDecoderState)
		x, y := s.primary, s.duplicates[len(s.duplicates)-1]
		s.duplicates[len(s.duplicates)-1] = x
		s.primary = y
		return s
	},
	RunFunc: func(sys commands.SystemUnderTest) commands.Result {
		s := sys.(*multiDecoderSystem)
		x, y := s.primary, s.duplicates[len(s.duplicates)-1]
		s.duplicates[len(s.duplicates)-1] = x
		s.primary = y
		return &systemAndResult{
			system: s,
		}
	},
}

// Close
// Duplicate
// Finalize

func anyASCIITag() gopter.Gen {
	return gopter.CombineGens(gen.Identifier(), gen.Identifier()).
		Map(func(values []interface{}) ident.Tag {
			name := values[0].(string)
			value := values[1].(string)
			return ident.StringTag(name, value)
		})
}

func anyASCIITags() gopter.Gen { return gen.SliceOf(anyASCIITag()) }
