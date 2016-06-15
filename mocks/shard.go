// Automatically generated by MockGen. DO NOT EDIT!
// Source: shard.go

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

package mocks

import (
	time0 "time"

	gomock "github.com/golang/mock/gomock"
	memtsdb "github.com/m3db/m3db"
	time "github.com/m3db/m3db/x/time"
)

// Mock of databaseShard interface
type MockdatabaseShard struct {
	ctrl     *gomock.Controller
	recorder *_MockdatabaseShardRecorder
}

// Recorder for MockdatabaseShard (not exported)
type _MockdatabaseShardRecorder struct {
	mock *MockdatabaseShard
}

func NewMockdatabaseShard(ctrl *gomock.Controller) *MockdatabaseShard {
	mock := &MockdatabaseShard{ctrl: ctrl}
	mock.recorder = &_MockdatabaseShardRecorder{mock}
	return mock
}

func (_m *MockdatabaseShard) EXPECT() *_MockdatabaseShardRecorder {
	return _m.recorder
}

func (_m *MockdatabaseShard) ShardNum() uint32 {
	ret := _m.ctrl.Call(_m, "ShardNum")
	ret0, _ := ret[0].(uint32)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) ShardNum() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShardNum")
}

func (_m *MockdatabaseShard) Tick() {
	_m.ctrl.Call(_m, "Tick")
}

func (_mr *_MockdatabaseShardRecorder) Tick() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tick")
}

func (_m *MockdatabaseShard) Write(ctx memtsdb.Context, id string, timestamp time0.Time, value float64, unit time.Unit, annotation []byte) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, id, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) Write(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockdatabaseShard) ReadEncoded(ctx memtsdb.Context, id string, start time0.Time, end time0.Time) (memtsdb.ReaderSliceReader, error) {
	ret := _m.ctrl.Call(_m, "ReadEncoded", ctx, id, start, end)
	ret0, _ := ret[0].(memtsdb.ReaderSliceReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockdatabaseShardRecorder) ReadEncoded(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadEncoded", arg0, arg1, arg2, arg3)
}

func (_m *MockdatabaseShard) Bootstrap(writeStart time0.Time) error {
	ret := _m.ctrl.Call(_m, "Bootstrap", writeStart)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) Bootstrap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Bootstrap", arg0)
}

func (_m *MockdatabaseShard) FlushToDisk(blockStart time0.Time) error {
	ret := _m.ctrl.Call(_m, "FlushToDisk", blockStart)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockdatabaseShardRecorder) FlushToDisk(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushToDisk", arg0)
}
