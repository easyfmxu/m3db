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

package tchannelthrift

type options struct {
	blockMetadataPool        BlockMetadataPool
	blockMetadataV2Pool      BlockMetadataV2Pool
	blockMetadataSlicePool   BlockMetadataSlicePool
	blockMetadataV2SlicePool BlockMetadataV2SlicePool
	blocksMetadataPool       BlocksMetadataPool
	blocksMetadataSlicePool  BlocksMetadataSlicePool
}

// NewOptions creates new options
func NewOptions() Options {
	return &options{
		blockMetadataPool:        NewBlockMetadataPool(nil),
		blockMetadataV2Pool:      NewBlockMetadataV2Pool(nil),
		blockMetadataSlicePool:   NewBlockMetadataSlicePool(nil, 0),
		blockMetadataV2SlicePool: NewBlockMetadataV2SlicePool(nil, 0),
		blocksMetadataPool:       NewBlocksMetadataPool(nil),
		blocksMetadataSlicePool:  NewBlocksMetadataSlicePool(nil, 0),
	}
}

func (o *options) SetBlockMetadataPool(value BlockMetadataPool) Options {
	opts := *o
	opts.blockMetadataPool = value
	return &opts
}

func (o *options) BlockMetadataPool() BlockMetadataPool {
	return o.blockMetadataPool
}

func (o *options) SetBlockMetadataV2Pool(value BlockMetadataV2Pool) Options {
	opts := *o
	opts.blockMetadataV2Pool = value
	return &opts
}

func (o *options) BlockMetadataV2Pool() BlockMetadataV2Pool {
	return o.blockMetadataV2Pool
}

func (o *options) SetBlockMetadataSlicePool(value BlockMetadataSlicePool) Options {
	opts := *o
	opts.blockMetadataSlicePool = value
	return &opts
}

func (o *options) BlockMetadataSlicePool() BlockMetadataSlicePool {
	return o.blockMetadataSlicePool
}

func (o *options) SetBlockMetadataV2SlicePool(value BlockMetadataV2SlicePool) Options {
	opts := *o
	opts.blockMetadataV2SlicePool = value
	return &opts
}

func (o *options) BlockMetadataV2SlicePool() BlockMetadataV2SlicePool {
	return o.blockMetadataV2SlicePool
}

func (o *options) SetBlocksMetadataPool(value BlocksMetadataPool) Options {
	opts := *o
	opts.blocksMetadataPool = value
	return &opts
}

func (o *options) BlocksMetadataPool() BlocksMetadataPool {
	return o.blocksMetadataPool
}

func (o *options) SetBlocksMetadataSlicePool(value BlocksMetadataSlicePool) Options {
	opts := *o
	opts.blocksMetadataSlicePool = value
	return &opts
}

func (o *options) BlocksMetadataSlicePool() BlocksMetadataSlicePool {
	return o.blocksMetadataSlicePool
}
