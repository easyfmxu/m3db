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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db/retention/types.go

package retention

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
)

// MockOptions is a mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsMockRecorder
}

// MockOptionsMockRecorder is the mock recorder for MockOptions
type MockOptionsMockRecorder struct {
	mock *MockOptions
}

// NewMockOptions creates a new mock instance
func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &MockOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockOptions) EXPECT() *MockOptionsMockRecorder {
	return _m.recorder
}

// Validate mocks base method
func (_m *MockOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (_mr *MockOptionsMockRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Validate", reflect.TypeOf((*MockOptions)(nil).Validate))
}

// Equal mocks base method
func (_m *MockOptions) Equal(value Options) bool {
	ret := _m.ctrl.Call(_m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (_mr *MockOptionsMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Equal", reflect.TypeOf((*MockOptions)(nil).Equal), arg0)
}

// SetRetentionPeriod mocks base method
func (_m *MockOptions) SetRetentionPeriod(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetRetentionPeriod", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetRetentionPeriod indicates an expected call of SetRetentionPeriod
func (_mr *MockOptionsMockRecorder) SetRetentionPeriod(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRetentionPeriod", reflect.TypeOf((*MockOptions)(nil).SetRetentionPeriod), arg0)
}

// RetentionPeriod mocks base method
func (_m *MockOptions) RetentionPeriod() time.Duration {
	ret := _m.ctrl.Call(_m, "RetentionPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RetentionPeriod indicates an expected call of RetentionPeriod
func (_mr *MockOptionsMockRecorder) RetentionPeriod() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RetentionPeriod", reflect.TypeOf((*MockOptions)(nil).RetentionPeriod))
}

// SetBlockSize mocks base method
func (_m *MockOptions) SetBlockSize(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetBlockSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockSize indicates an expected call of SetBlockSize
func (_mr *MockOptionsMockRecorder) SetBlockSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBlockSize", reflect.TypeOf((*MockOptions)(nil).SetBlockSize), arg0)
}

// BlockSize mocks base method
func (_m *MockOptions) BlockSize() time.Duration {
	ret := _m.ctrl.Call(_m, "BlockSize")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockSize indicates an expected call of BlockSize
func (_mr *MockOptionsMockRecorder) BlockSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BlockSize", reflect.TypeOf((*MockOptions)(nil).BlockSize))
}

// SetBufferFuture mocks base method
func (_m *MockOptions) SetBufferFuture(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetBufferFuture", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBufferFuture indicates an expected call of SetBufferFuture
func (_mr *MockOptionsMockRecorder) SetBufferFuture(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBufferFuture", reflect.TypeOf((*MockOptions)(nil).SetBufferFuture), arg0)
}

// BufferFuture mocks base method
func (_m *MockOptions) BufferFuture() time.Duration {
	ret := _m.ctrl.Call(_m, "BufferFuture")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BufferFuture indicates an expected call of BufferFuture
func (_mr *MockOptionsMockRecorder) BufferFuture() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BufferFuture", reflect.TypeOf((*MockOptions)(nil).BufferFuture))
}

// SetBufferPast mocks base method
func (_m *MockOptions) SetBufferPast(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetBufferPast", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBufferPast indicates an expected call of SetBufferPast
func (_mr *MockOptionsMockRecorder) SetBufferPast(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBufferPast", reflect.TypeOf((*MockOptions)(nil).SetBufferPast), arg0)
}

// BufferPast mocks base method
func (_m *MockOptions) BufferPast() time.Duration {
	ret := _m.ctrl.Call(_m, "BufferPast")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BufferPast indicates an expected call of BufferPast
func (_mr *MockOptionsMockRecorder) BufferPast() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BufferPast", reflect.TypeOf((*MockOptions)(nil).BufferPast))
}

// SetBlockDataExpiry mocks base method
func (_m *MockOptions) SetBlockDataExpiry(on bool) Options {
	ret := _m.ctrl.Call(_m, "SetBlockDataExpiry", on)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockDataExpiry indicates an expected call of SetBlockDataExpiry
func (_mr *MockOptionsMockRecorder) SetBlockDataExpiry(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBlockDataExpiry", reflect.TypeOf((*MockOptions)(nil).SetBlockDataExpiry), arg0)
}

// BlockDataExpiry mocks base method
func (_m *MockOptions) BlockDataExpiry() bool {
	ret := _m.ctrl.Call(_m, "BlockDataExpiry")
	ret0, _ := ret[0].(bool)
	return ret0
}

// BlockDataExpiry indicates an expected call of BlockDataExpiry
func (_mr *MockOptionsMockRecorder) BlockDataExpiry() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BlockDataExpiry", reflect.TypeOf((*MockOptions)(nil).BlockDataExpiry))
}

// SetBlockDataExpiryAfterNotAccessedPeriod mocks base method
func (_m *MockOptions) SetBlockDataExpiryAfterNotAccessedPeriod(period time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetBlockDataExpiryAfterNotAccessedPeriod", period)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetBlockDataExpiryAfterNotAccessedPeriod indicates an expected call of SetBlockDataExpiryAfterNotAccessedPeriod
func (_mr *MockOptionsMockRecorder) SetBlockDataExpiryAfterNotAccessedPeriod(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetBlockDataExpiryAfterNotAccessedPeriod", reflect.TypeOf((*MockOptions)(nil).SetBlockDataExpiryAfterNotAccessedPeriod), arg0)
}

// BlockDataExpiryAfterNotAccessedPeriod mocks base method
func (_m *MockOptions) BlockDataExpiryAfterNotAccessedPeriod() time.Duration {
	ret := _m.ctrl.Call(_m, "BlockDataExpiryAfterNotAccessedPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockDataExpiryAfterNotAccessedPeriod indicates an expected call of BlockDataExpiryAfterNotAccessedPeriod
func (_mr *MockOptionsMockRecorder) BlockDataExpiryAfterNotAccessedPeriod() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BlockDataExpiryAfterNotAccessedPeriod", reflect.TypeOf((*MockOptions)(nil).BlockDataExpiryAfterNotAccessedPeriod))
}
