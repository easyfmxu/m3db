// Copyright (c) 2016 Uber Technologies, Inc.
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

package encoding

import (
	"errors"
	"io"

	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3db/x/xio"
	xtime "github.com/m3db/m3x/time"
)

var (
	errOutOfOrderIterator = errors.New("values are out of order from inner iterator")
)

// multiReaderIterator is an iterator that iterates in order over a list of sets of
// internally ordered but not collectively in order readers, it also deduplicates datapoints.
type multiReaderIterator struct {
	iters            iterators
	slicesIter       xio.ReaderSliceOfSlicesIterator
	iteratorAlloc    ReaderIteratorAllocate
	singleSlicesIter singleSlicesOfSlicesIterator
	pool             MultiReaderIteratorPool
	err              error
	firstNext        bool
	closed           bool
}

// NewMultiReaderIterator creates a new multi-reader iterator.
func NewMultiReaderIterator(
	iteratorAlloc ReaderIteratorAllocate,
	pool MultiReaderIteratorPool,
) MultiReaderIterator {
	it := &multiReaderIterator{pool: pool, iteratorAlloc: iteratorAlloc}
	it.Reset(nil)
	return it
}

func (it *multiReaderIterator) Next() bool {
	if !it.firstNext {
		// When firstNext do not progress the first time
		if !it.hasNext() {
			return false
		}
		it.moveToNext()
	}
	it.firstNext = false
	return it.hasNext()
}

func (it *multiReaderIterator) Current() (ts.Datapoint, xtime.Unit, ts.Annotation) {
	return it.iters.current()
}

func (it *multiReaderIterator) hasError() bool {
	return it.err != nil
}

func (it *multiReaderIterator) isClosed() bool {
	return it.closed
}

func (it *multiReaderIterator) hasMore() bool {
	return it.iters.len() > 0 || it.slicesIter != nil
}

func (it *multiReaderIterator) hasNext() bool {
	return !it.hasError() && !it.isClosed() && it.hasMore()
}

func (it *multiReaderIterator) moveToNext() {
	if it.iters.len() > 0 {
		it.moveIteratorsToNext()
	}
	if it.iters.len() > 0 || it.hasError() {
		// Still have valid iters or has error
		return
	}

	// Move forward through slices of readers
	if !it.slicesIter.Next() {
		// No more readers, nil out so that hasMore reflects correctly
		it.slicesIter.Close()
		it.slicesIter = nil
		return
	}

	// Add all readers to current iterators heap
	currentLen := it.slicesIter.CurrentLen()
	for i := 0; i < currentLen; i++ {
		var (
			reader = it.slicesIter.CurrentAt(i)
			iter   = it.iteratorAlloc(reader)
		)
		if iter.Next() {
			// Only insert it if it has values
			it.iters.push(iter)
		} else {
			err := iter.Err()
			iter.Close()
			if it.err == nil && err != nil {
				it.err = err
			}
		}
	}

	if it.iters.len() == 0 && !it.hasError() {
		// No iterators were added, move to next
		it.moveToNext()
	}
}

func (it *multiReaderIterator) moveIteratorsToNext() {
	for {
		prev := it.iters.at()
		next, err := it.iters.moveToValidNext()
		if it.err == nil && err != nil {
			it.err = err
			return
		}
		if err != nil || !next {
			return
		}

		curr := it.iters.at()
		if !curr.Equal(prev) {
			return
		}

		// Dedupe by continuing
	}
}

func (it *multiReaderIterator) Err() error {
	return it.err
}

func (it *multiReaderIterator) Reset(readers []io.Reader) {
	it.singleSlicesIter.readers = readers
	it.singleSlicesIter.firstNext = true
	it.singleSlicesIter.closed = false
	it.ResetSliceOfSlices(&it.singleSlicesIter)
}

func (it *multiReaderIterator) ResetSliceOfSlices(slicesIter xio.ReaderSliceOfSlicesIterator) {
	it.iters.reset()
	it.slicesIter = slicesIter
	it.err = nil
	it.firstNext = true
	it.closed = false
	// Try moveToNext to load values for calls to Current before Next
	it.moveToNext()
}

func (it *multiReaderIterator) Close() {
	if it.isClosed() {
		return
	}
	it.closed = true
	it.iters.reset()
	if it.slicesIter != nil {
		it.slicesIter.Close()
	}
	it.slicesIter = nil
	if it.pool != nil {
		it.pool.Put(it)
	}
}

type singleSlicesOfSlicesIterator struct {
	readers   []io.Reader
	firstNext bool
	closed    bool
}

func (it *singleSlicesOfSlicesIterator) Next() bool {
	if !it.firstNext || it.closed {
		return false
	}
	it.firstNext = false
	return true
}

func (it *singleSlicesOfSlicesIterator) CurrentLen() int {
	return len(it.readers)
}

func (it *singleSlicesOfSlicesIterator) CurrentAt(idx int) io.Reader {
	return it.readers[idx]
}

func (it *singleSlicesOfSlicesIterator) Close() {
	if it.closed {
		return
	}
	it.closed = true
}
