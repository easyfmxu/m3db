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
// Source: github.com/m3db/m3db/topology/types.go

package topology

import (
	"reflect"
	"time"

	"github.com/m3db/m3cluster/client"
	"github.com/m3db/m3cluster/services"
	"github.com/m3db/m3db/sharding"
	"github.com/m3db/m3x/ident"
	"github.com/m3db/m3x/instrument"

	"github.com/golang/mock/gomock"
)

// MockHost is a mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *MockHostMockRecorder
}

// MockHostMockRecorder is the mock recorder for MockHost
type MockHostMockRecorder struct {
	mock *MockHost
}

// NewMockHost creates a new mock instance
func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &MockHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockHost) EXPECT() *MockHostMockRecorder {
	return _m.recorder
}

// ID mocks base method
func (_m *MockHost) ID() string {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (_mr *MockHostMockRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ID", reflect.TypeOf((*MockHost)(nil).ID))
}

// Address mocks base method
func (_m *MockHost) Address() string {
	ret := _m.ctrl.Call(_m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address
func (_mr *MockHostMockRecorder) Address() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Address", reflect.TypeOf((*MockHost)(nil).Address))
}

// String mocks base method
func (_m *MockHost) String() string {
	ret := _m.ctrl.Call(_m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (_mr *MockHostMockRecorder) String() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "String", reflect.TypeOf((*MockHost)(nil).String))
}

// MockHostShardSet is a mock of HostShardSet interface
type MockHostShardSet struct {
	ctrl     *gomock.Controller
	recorder *MockHostShardSetMockRecorder
}

// MockHostShardSetMockRecorder is the mock recorder for MockHostShardSet
type MockHostShardSetMockRecorder struct {
	mock *MockHostShardSet
}

// NewMockHostShardSet creates a new mock instance
func NewMockHostShardSet(ctrl *gomock.Controller) *MockHostShardSet {
	mock := &MockHostShardSet{ctrl: ctrl}
	mock.recorder = &MockHostShardSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockHostShardSet) EXPECT() *MockHostShardSetMockRecorder {
	return _m.recorder
}

// Host mocks base method
func (_m *MockHostShardSet) Host() Host {
	ret := _m.ctrl.Call(_m, "Host")
	ret0, _ := ret[0].(Host)
	return ret0
}

// Host indicates an expected call of Host
func (_mr *MockHostShardSetMockRecorder) Host() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Host", reflect.TypeOf((*MockHostShardSet)(nil).Host))
}

// ShardSet mocks base method
func (_m *MockHostShardSet) ShardSet() sharding.ShardSet {
	ret := _m.ctrl.Call(_m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

// ShardSet indicates an expected call of ShardSet
func (_mr *MockHostShardSetMockRecorder) ShardSet() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ShardSet", reflect.TypeOf((*MockHostShardSet)(nil).ShardSet))
}

// MockInitializer is a mock of Initializer interface
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return _m.recorder
}

// Init mocks base method
func (_m *MockInitializer) Init() (Topology, error) {
	ret := _m.ctrl.Call(_m, "Init")
	ret0, _ := ret[0].(Topology)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init
func (_mr *MockInitializerMockRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Init", reflect.TypeOf((*MockInitializer)(nil).Init))
}

// MockTopology is a mock of Topology interface
type MockTopology struct {
	ctrl     *gomock.Controller
	recorder *MockTopologyMockRecorder
}

// MockTopologyMockRecorder is the mock recorder for MockTopology
type MockTopologyMockRecorder struct {
	mock *MockTopology
}

// NewMockTopology creates a new mock instance
func NewMockTopology(ctrl *gomock.Controller) *MockTopology {
	mock := &MockTopology{ctrl: ctrl}
	mock.recorder = &MockTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockTopology) EXPECT() *MockTopologyMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockTopology) Get() Map {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Map)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockTopologyMockRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockTopology)(nil).Get))
}

// Watch mocks base method
func (_m *MockTopology) Watch() (MapWatch, error) {
	ret := _m.ctrl.Call(_m, "Watch")
	ret0, _ := ret[0].(MapWatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (_mr *MockTopologyMockRecorder) Watch() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Watch", reflect.TypeOf((*MockTopology)(nil).Watch))
}

// Close mocks base method
func (_m *MockTopology) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockTopologyMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockTopology)(nil).Close))
}

// MockDynamicTopology is a mock of DynamicTopology interface
type MockDynamicTopology struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicTopologyMockRecorder
}

// MockDynamicTopologyMockRecorder is the mock recorder for MockDynamicTopology
type MockDynamicTopologyMockRecorder struct {
	mock *MockDynamicTopology
}

// NewMockDynamicTopology creates a new mock instance
func NewMockDynamicTopology(ctrl *gomock.Controller) *MockDynamicTopology {
	mock := &MockDynamicTopology{ctrl: ctrl}
	mock.recorder = &MockDynamicTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDynamicTopology) EXPECT() *MockDynamicTopologyMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockDynamicTopology) Get() Map {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Map)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockDynamicTopologyMockRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockDynamicTopology)(nil).Get))
}

// Watch mocks base method
func (_m *MockDynamicTopology) Watch() (MapWatch, error) {
	ret := _m.ctrl.Call(_m, "Watch")
	ret0, _ := ret[0].(MapWatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (_mr *MockDynamicTopologyMockRecorder) Watch() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Watch", reflect.TypeOf((*MockDynamicTopology)(nil).Watch))
}

// Close mocks base method
func (_m *MockDynamicTopology) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockDynamicTopologyMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockDynamicTopology)(nil).Close))
}

// MarkShardAvailable mocks base method
func (_m *MockDynamicTopology) MarkShardAvailable(instanceID string, shardID uint32) error {
	ret := _m.ctrl.Call(_m, "MarkShardAvailable", instanceID, shardID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkShardAvailable indicates an expected call of MarkShardAvailable
func (_mr *MockDynamicTopologyMockRecorder) MarkShardAvailable(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MarkShardAvailable", reflect.TypeOf((*MockDynamicTopology)(nil).MarkShardAvailable), arg0, arg1)
}

// MockMapWatch is a mock of MapWatch interface
type MockMapWatch struct {
	ctrl     *gomock.Controller
	recorder *MockMapWatchMockRecorder
}

// MockMapWatchMockRecorder is the mock recorder for MockMapWatch
type MockMapWatchMockRecorder struct {
	mock *MockMapWatch
}

// NewMockMapWatch creates a new mock instance
func NewMockMapWatch(ctrl *gomock.Controller) *MockMapWatch {
	mock := &MockMapWatch{ctrl: ctrl}
	mock.recorder = &MockMapWatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockMapWatch) EXPECT() *MockMapWatchMockRecorder {
	return _m.recorder
}

// C mocks base method
func (_m *MockMapWatch) C() <-chan struct{} {
	ret := _m.ctrl.Call(_m, "C")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// C indicates an expected call of C
func (_mr *MockMapWatchMockRecorder) C() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "C", reflect.TypeOf((*MockMapWatch)(nil).C))
}

// Get mocks base method
func (_m *MockMapWatch) Get() Map {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Map)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockMapWatchMockRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockMapWatch)(nil).Get))
}

// Close mocks base method
func (_m *MockMapWatch) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockMapWatchMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockMapWatch)(nil).Close))
}

// MockMap is a mock of Map interface
type MockMap struct {
	ctrl     *gomock.Controller
	recorder *MockMapMockRecorder
}

// MockMapMockRecorder is the mock recorder for MockMap
type MockMapMockRecorder struct {
	mock *MockMap
}

// NewMockMap creates a new mock instance
func NewMockMap(ctrl *gomock.Controller) *MockMap {
	mock := &MockMap{ctrl: ctrl}
	mock.recorder = &MockMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockMap) EXPECT() *MockMapMockRecorder {
	return _m.recorder
}

// Hosts mocks base method
func (_m *MockMap) Hosts() []Host {
	ret := _m.ctrl.Call(_m, "Hosts")
	ret0, _ := ret[0].([]Host)
	return ret0
}

// Hosts indicates an expected call of Hosts
func (_mr *MockMapMockRecorder) Hosts() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Hosts", reflect.TypeOf((*MockMap)(nil).Hosts))
}

// HostShardSets mocks base method
func (_m *MockMap) HostShardSets() []HostShardSet {
	ret := _m.ctrl.Call(_m, "HostShardSets")
	ret0, _ := ret[0].([]HostShardSet)
	return ret0
}

// HostShardSets indicates an expected call of HostShardSets
func (_mr *MockMapMockRecorder) HostShardSets() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "HostShardSets", reflect.TypeOf((*MockMap)(nil).HostShardSets))
}

// LookupHostShardSet mocks base method
func (_m *MockMap) LookupHostShardSet(hostID string) (HostShardSet, bool) {
	ret := _m.ctrl.Call(_m, "LookupHostShardSet", hostID)
	ret0, _ := ret[0].(HostShardSet)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// LookupHostShardSet indicates an expected call of LookupHostShardSet
func (_mr *MockMapMockRecorder) LookupHostShardSet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "LookupHostShardSet", reflect.TypeOf((*MockMap)(nil).LookupHostShardSet), arg0)
}

// HostsLen mocks base method
func (_m *MockMap) HostsLen() int {
	ret := _m.ctrl.Call(_m, "HostsLen")
	ret0, _ := ret[0].(int)
	return ret0
}

// HostsLen indicates an expected call of HostsLen
func (_mr *MockMapMockRecorder) HostsLen() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "HostsLen", reflect.TypeOf((*MockMap)(nil).HostsLen))
}

// ShardSet mocks base method
func (_m *MockMap) ShardSet() sharding.ShardSet {
	ret := _m.ctrl.Call(_m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

// ShardSet indicates an expected call of ShardSet
func (_mr *MockMapMockRecorder) ShardSet() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ShardSet", reflect.TypeOf((*MockMap)(nil).ShardSet))
}

// Route mocks base method
func (_m *MockMap) Route(id ident.ID) (uint32, []Host, error) {
	ret := _m.ctrl.Call(_m, "Route", id)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].([]Host)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Route indicates an expected call of Route
func (_mr *MockMapMockRecorder) Route(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Route", reflect.TypeOf((*MockMap)(nil).Route), arg0)
}

// RouteForEach mocks base method
func (_m *MockMap) RouteForEach(id ident.ID, forEachFn RouteForEachFn) error {
	ret := _m.ctrl.Call(_m, "RouteForEach", id, forEachFn)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteForEach indicates an expected call of RouteForEach
func (_mr *MockMapMockRecorder) RouteForEach(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RouteForEach", reflect.TypeOf((*MockMap)(nil).RouteForEach), arg0, arg1)
}

// RouteShard mocks base method
func (_m *MockMap) RouteShard(shard uint32) ([]Host, error) {
	ret := _m.ctrl.Call(_m, "RouteShard", shard)
	ret0, _ := ret[0].([]Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteShard indicates an expected call of RouteShard
func (_mr *MockMapMockRecorder) RouteShard(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RouteShard", reflect.TypeOf((*MockMap)(nil).RouteShard), arg0)
}

// RouteShardForEach mocks base method
func (_m *MockMap) RouteShardForEach(shard uint32, forEachFn RouteForEachFn) error {
	ret := _m.ctrl.Call(_m, "RouteShardForEach", shard, forEachFn)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteShardForEach indicates an expected call of RouteShardForEach
func (_mr *MockMapMockRecorder) RouteShardForEach(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RouteShardForEach", reflect.TypeOf((*MockMap)(nil).RouteShardForEach), arg0, arg1)
}

// Replicas mocks base method
func (_m *MockMap) Replicas() int {
	ret := _m.ctrl.Call(_m, "Replicas")
	ret0, _ := ret[0].(int)
	return ret0
}

// Replicas indicates an expected call of Replicas
func (_mr *MockMapMockRecorder) Replicas() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Replicas", reflect.TypeOf((*MockMap)(nil).Replicas))
}

// MajorityReplicas mocks base method
func (_m *MockMap) MajorityReplicas() int {
	ret := _m.ctrl.Call(_m, "MajorityReplicas")
	ret0, _ := ret[0].(int)
	return ret0
}

// MajorityReplicas indicates an expected call of MajorityReplicas
func (_mr *MockMapMockRecorder) MajorityReplicas() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MajorityReplicas", reflect.TypeOf((*MockMap)(nil).MajorityReplicas))
}

// MockStaticOptions is a mock of StaticOptions interface
type MockStaticOptions struct {
	ctrl     *gomock.Controller
	recorder *MockStaticOptionsMockRecorder
}

// MockStaticOptionsMockRecorder is the mock recorder for MockStaticOptions
type MockStaticOptionsMockRecorder struct {
	mock *MockStaticOptions
}

// NewMockStaticOptions creates a new mock instance
func NewMockStaticOptions(ctrl *gomock.Controller) *MockStaticOptions {
	mock := &MockStaticOptions{ctrl: ctrl}
	mock.recorder = &MockStaticOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStaticOptions) EXPECT() *MockStaticOptionsMockRecorder {
	return _m.recorder
}

// Validate mocks base method
func (_m *MockStaticOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (_mr *MockStaticOptionsMockRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Validate", reflect.TypeOf((*MockStaticOptions)(nil).Validate))
}

// SetShardSet mocks base method
func (_m *MockStaticOptions) SetShardSet(value sharding.ShardSet) StaticOptions {
	ret := _m.ctrl.Call(_m, "SetShardSet", value)
	ret0, _ := ret[0].(StaticOptions)
	return ret0
}

// SetShardSet indicates an expected call of SetShardSet
func (_mr *MockStaticOptionsMockRecorder) SetShardSet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetShardSet", reflect.TypeOf((*MockStaticOptions)(nil).SetShardSet), arg0)
}

// ShardSet mocks base method
func (_m *MockStaticOptions) ShardSet() sharding.ShardSet {
	ret := _m.ctrl.Call(_m, "ShardSet")
	ret0, _ := ret[0].(sharding.ShardSet)
	return ret0
}

// ShardSet indicates an expected call of ShardSet
func (_mr *MockStaticOptionsMockRecorder) ShardSet() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ShardSet", reflect.TypeOf((*MockStaticOptions)(nil).ShardSet))
}

// SetReplicas mocks base method
func (_m *MockStaticOptions) SetReplicas(value int) StaticOptions {
	ret := _m.ctrl.Call(_m, "SetReplicas", value)
	ret0, _ := ret[0].(StaticOptions)
	return ret0
}

// SetReplicas indicates an expected call of SetReplicas
func (_mr *MockStaticOptionsMockRecorder) SetReplicas(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetReplicas", reflect.TypeOf((*MockStaticOptions)(nil).SetReplicas), arg0)
}

// Replicas mocks base method
func (_m *MockStaticOptions) Replicas() int {
	ret := _m.ctrl.Call(_m, "Replicas")
	ret0, _ := ret[0].(int)
	return ret0
}

// Replicas indicates an expected call of Replicas
func (_mr *MockStaticOptionsMockRecorder) Replicas() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Replicas", reflect.TypeOf((*MockStaticOptions)(nil).Replicas))
}

// SetHostShardSets mocks base method
func (_m *MockStaticOptions) SetHostShardSets(value []HostShardSet) StaticOptions {
	ret := _m.ctrl.Call(_m, "SetHostShardSets", value)
	ret0, _ := ret[0].(StaticOptions)
	return ret0
}

// SetHostShardSets indicates an expected call of SetHostShardSets
func (_mr *MockStaticOptionsMockRecorder) SetHostShardSets(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetHostShardSets", reflect.TypeOf((*MockStaticOptions)(nil).SetHostShardSets), arg0)
}

// HostShardSets mocks base method
func (_m *MockStaticOptions) HostShardSets() []HostShardSet {
	ret := _m.ctrl.Call(_m, "HostShardSets")
	ret0, _ := ret[0].([]HostShardSet)
	return ret0
}

// HostShardSets indicates an expected call of HostShardSets
func (_mr *MockStaticOptionsMockRecorder) HostShardSets() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "HostShardSets", reflect.TypeOf((*MockStaticOptions)(nil).HostShardSets))
}

// MockDynamicOptions is a mock of DynamicOptions interface
type MockDynamicOptions struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicOptionsMockRecorder
}

// MockDynamicOptionsMockRecorder is the mock recorder for MockDynamicOptions
type MockDynamicOptionsMockRecorder struct {
	mock *MockDynamicOptions
}

// NewMockDynamicOptions creates a new mock instance
func NewMockDynamicOptions(ctrl *gomock.Controller) *MockDynamicOptions {
	mock := &MockDynamicOptions{ctrl: ctrl}
	mock.recorder = &MockDynamicOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDynamicOptions) EXPECT() *MockDynamicOptionsMockRecorder {
	return _m.recorder
}

// Validate mocks base method
func (_m *MockDynamicOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (_mr *MockDynamicOptionsMockRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Validate", reflect.TypeOf((*MockDynamicOptions)(nil).Validate))
}

// SetConfigServiceClient mocks base method
func (_m *MockDynamicOptions) SetConfigServiceClient(c client.Client) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetConfigServiceClient", c)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetConfigServiceClient indicates an expected call of SetConfigServiceClient
func (_mr *MockDynamicOptionsMockRecorder) SetConfigServiceClient(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetConfigServiceClient", reflect.TypeOf((*MockDynamicOptions)(nil).SetConfigServiceClient), arg0)
}

// ConfigServiceClient mocks base method
func (_m *MockDynamicOptions) ConfigServiceClient() client.Client {
	ret := _m.ctrl.Call(_m, "ConfigServiceClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// ConfigServiceClient indicates an expected call of ConfigServiceClient
func (_mr *MockDynamicOptionsMockRecorder) ConfigServiceClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ConfigServiceClient", reflect.TypeOf((*MockDynamicOptions)(nil).ConfigServiceClient))
}

// SetServiceID mocks base method
func (_m *MockDynamicOptions) SetServiceID(s services.ServiceID) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetServiceID", s)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetServiceID indicates an expected call of SetServiceID
func (_mr *MockDynamicOptionsMockRecorder) SetServiceID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetServiceID", reflect.TypeOf((*MockDynamicOptions)(nil).SetServiceID), arg0)
}

// ServiceID mocks base method
func (_m *MockDynamicOptions) ServiceID() services.ServiceID {
	ret := _m.ctrl.Call(_m, "ServiceID")
	ret0, _ := ret[0].(services.ServiceID)
	return ret0
}

// ServiceID indicates an expected call of ServiceID
func (_mr *MockDynamicOptionsMockRecorder) ServiceID() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ServiceID", reflect.TypeOf((*MockDynamicOptions)(nil).ServiceID))
}

// SetServicesOverrideOptions mocks base method
func (_m *MockDynamicOptions) SetServicesOverrideOptions(opts services.OverrideOptions) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetServicesOverrideOptions", opts)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetServicesOverrideOptions indicates an expected call of SetServicesOverrideOptions
func (_mr *MockDynamicOptionsMockRecorder) SetServicesOverrideOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetServicesOverrideOptions", reflect.TypeOf((*MockDynamicOptions)(nil).SetServicesOverrideOptions), arg0)
}

// ServicesOverrideOptions mocks base method
func (_m *MockDynamicOptions) ServicesOverrideOptions() services.OverrideOptions {
	ret := _m.ctrl.Call(_m, "ServicesOverrideOptions")
	ret0, _ := ret[0].(services.OverrideOptions)
	return ret0
}

// ServicesOverrideOptions indicates an expected call of ServicesOverrideOptions
func (_mr *MockDynamicOptionsMockRecorder) ServicesOverrideOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ServicesOverrideOptions", reflect.TypeOf((*MockDynamicOptions)(nil).ServicesOverrideOptions))
}

// SetQueryOptions mocks base method
func (_m *MockDynamicOptions) SetQueryOptions(value services.QueryOptions) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetQueryOptions", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetQueryOptions indicates an expected call of SetQueryOptions
func (_mr *MockDynamicOptionsMockRecorder) SetQueryOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetQueryOptions", reflect.TypeOf((*MockDynamicOptions)(nil).SetQueryOptions), arg0)
}

// QueryOptions mocks base method
func (_m *MockDynamicOptions) QueryOptions() services.QueryOptions {
	ret := _m.ctrl.Call(_m, "QueryOptions")
	ret0, _ := ret[0].(services.QueryOptions)
	return ret0
}

// QueryOptions indicates an expected call of QueryOptions
func (_mr *MockDynamicOptionsMockRecorder) QueryOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryOptions", reflect.TypeOf((*MockDynamicOptions)(nil).QueryOptions))
}

// SetInstrumentOptions mocks base method
func (_m *MockDynamicOptions) SetInstrumentOptions(value instrument.Options) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetInstrumentOptions indicates an expected call of SetInstrumentOptions
func (_mr *MockDynamicOptionsMockRecorder) SetInstrumentOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetInstrumentOptions", reflect.TypeOf((*MockDynamicOptions)(nil).SetInstrumentOptions), arg0)
}

// InstrumentOptions mocks base method
func (_m *MockDynamicOptions) InstrumentOptions() instrument.Options {
	ret := _m.ctrl.Call(_m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

// InstrumentOptions indicates an expected call of InstrumentOptions
func (_mr *MockDynamicOptionsMockRecorder) InstrumentOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstrumentOptions", reflect.TypeOf((*MockDynamicOptions)(nil).InstrumentOptions))
}

// SetInitTimeout mocks base method
func (_m *MockDynamicOptions) SetInitTimeout(value time.Duration) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetInitTimeout", value)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetInitTimeout indicates an expected call of SetInitTimeout
func (_mr *MockDynamicOptionsMockRecorder) SetInitTimeout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetInitTimeout", reflect.TypeOf((*MockDynamicOptions)(nil).SetInitTimeout), arg0)
}

// InitTimeout mocks base method
func (_m *MockDynamicOptions) InitTimeout() time.Duration {
	ret := _m.ctrl.Call(_m, "InitTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// InitTimeout indicates an expected call of InitTimeout
func (_mr *MockDynamicOptionsMockRecorder) InitTimeout() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InitTimeout", reflect.TypeOf((*MockDynamicOptions)(nil).InitTimeout))
}

// SetHashGen mocks base method
func (_m *MockDynamicOptions) SetHashGen(h sharding.HashGen) DynamicOptions {
	ret := _m.ctrl.Call(_m, "SetHashGen", h)
	ret0, _ := ret[0].(DynamicOptions)
	return ret0
}

// SetHashGen indicates an expected call of SetHashGen
func (_mr *MockDynamicOptionsMockRecorder) SetHashGen(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetHashGen", reflect.TypeOf((*MockDynamicOptions)(nil).SetHashGen), arg0)
}

// HashGen mocks base method
func (_m *MockDynamicOptions) HashGen() sharding.HashGen {
	ret := _m.ctrl.Call(_m, "HashGen")
	ret0, _ := ret[0].(sharding.HashGen)
	return ret0
}

// HashGen indicates an expected call of HashGen
func (_mr *MockDynamicOptionsMockRecorder) HashGen() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "HashGen", reflect.TypeOf((*MockDynamicOptions)(nil).HashGen))
}
