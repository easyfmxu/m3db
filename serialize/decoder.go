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

package serialize

import (
	"errors"

	"github.com/m3db/m3db/x/xpool"
	"github.com/m3db/m3x/checked"
	"github.com/m3db/m3x/ident"
)

var (
	errIncorrectHeader               = errors.New("header magic number does not match expected value")
	errInvalidByteStreamIDDecoding   = errors.New("internal error, invalid byte stream while decoding ID")
	errInvalidByteStreamUintDecoding = errors.New("internal error, invalid byte stream while decoding uint")
)

type decoder struct {
	checkedData checked.Bytes
	data        []byte
	remaining   int
	err         error
	hasCurrent  bool
	current     ident.Tag

	pool        DecoderPool
	wrapperPool xpool.CheckedBytesWrapperPool
}

func newDecoder(pool DecoderPool, wrapperPool xpool.CheckedBytesWrapperPool) Decoder {
	return &decoder{
		pool:        pool,
		wrapperPool: wrapperPool,
	}
}

func (d *decoder) Reset(b checked.Bytes) {
	d.checkedData = b
	d.checkedData.IncRef()
	d.data = d.checkedData.Get()

	header, err := d.decodeUInt16()
	if err != nil {
		d.err = err
		return
	}

	if header != HeaderMagicNumber {
		d.err = errIncorrectHeader
		return
	}

	remain, err := d.decodeUInt16()
	if err != nil {
		d.err = err
		return
	}

	d.remaining = int(remain)
}

func (d *decoder) Next() bool {
	if d.err != nil || d.remaining <= 0 {
		return false
	}

	d.release()
	t, err := d.decodeTag()
	if err != nil {
		d.err = err
		return false
	}

	d.remaining--
	d.current = t
	d.hasCurrent = true
	return true
}

func (d *decoder) Current() ident.Tag {
	return d.current
}

func (d *decoder) decodeTag() (ident.Tag, error) {
	name, err := d.decodeID()
	if err != nil {
		return ident.Tag{}, err
	}

	value, err := d.decodeID()
	if err != nil {
		name.Finalize()
		return ident.Tag{}, err
	}

	return ident.Tag{name, value}, nil
}

func (d *decoder) decodeID() (ident.ID, error) {
	l, err := d.decodeUInt16()
	if err != nil {
		return nil, err
	}

	if len(d.data) < int(l) {
		return nil, errInvalidByteStreamIDDecoding
	}

	// incRef to indicate another checked.Bytes has a
	// reference to the original bytes
	d.checkedData.IncRef()
	b := d.wrapperPool.Get(d.data[:l])
	d.data = d.data[l:]

	return ident.BinaryID(b), nil
}

func (d *decoder) decodeUInt16() (uint16, error) {
	if len(d.data) < 2 {
		return 0, errInvalidByteStreamUintDecoding
	}

	n := decodeUInt16(d.data)
	d.data = d.data[2:]
	return n, nil
}

func (d *decoder) Err() error {
	return d.err
}

func (d *decoder) release() {
	if n := d.current.Name; n != nil {
		n.Finalize()
		d.checkedData.DecRef() // indicate we've released the extra ref
		d.current.Name = nil
	}
	if v := d.current.Value; v != nil {
		v.Finalize()
		d.checkedData.DecRef() // indicate we've released the extra ref
		d.current.Value = nil
	}
}

func (d *decoder) Remaining() int {
	return d.remaining
}

func (d *decoder) Close() {
	d.release()
	d.data = nil
	d.err = nil
	d.remaining = 0
	if d.checkedData != nil {
		d.checkedData.DecRef()
		d.checkedData = nil
	}
}

func (d *decoder) Finalize() {
	d.Close()
	if d.pool == nil {
		return
	}
	d.pool.Put(d)
}

func (d *decoder) Duplicate() ident.TagIterator {
	copy := *d
	if copy.hasCurrent {
		// TODO(prateek): use idPool.CloneTag
		copy.current = ident.StringTag(
			copy.current.Name.String(),
			copy.current.Value.String(),
		)
	}
	if copy.checkedData != nil {
		copy.checkedData.IncRef()
	}
	return &copy
}
