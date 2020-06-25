// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha2sets is a generated GoMock package.
package mock_v1alpha2sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/linkerd/linkerd2/controller/gen/apis/serviceprofile/v1alpha2"
	v1alpha2sets "github.com/solo-io/external-apis/pkg/api/linkerd/linkerd.io/v1alpha2/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockServiceProfileSet is a mock of ServiceProfileSet interface.
type MockServiceProfileSet struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProfileSetMockRecorder
}

// MockServiceProfileSetMockRecorder is the mock recorder for MockServiceProfileSet.
type MockServiceProfileSetMockRecorder struct {
	mock *MockServiceProfileSet
}

// NewMockServiceProfileSet creates a new mock instance.
func NewMockServiceProfileSet(ctrl *gomock.Controller) *MockServiceProfileSet {
	mock := &MockServiceProfileSet{ctrl: ctrl}
	mock.recorder = &MockServiceProfileSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceProfileSet) EXPECT() *MockServiceProfileSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockServiceProfileSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockServiceProfileSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockServiceProfileSet)(nil).Keys))
}

// List mocks base method.
func (m *MockServiceProfileSet) List() []*v1alpha2.ServiceProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha2.ServiceProfile)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockServiceProfileSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceProfileSet)(nil).List))
}

// Map mocks base method.
func (m *MockServiceProfileSet) Map() map[string]*v1alpha2.ServiceProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha2.ServiceProfile)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockServiceProfileSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockServiceProfileSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockServiceProfileSet) Insert(serviceProfile ...*v1alpha2.ServiceProfile) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range serviceProfile {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockServiceProfileSetMockRecorder) Insert(serviceProfile ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockServiceProfileSet)(nil).Insert), serviceProfile...)
}

// Equal mocks base method.
func (m *MockServiceProfileSet) Equal(serviceProfileSet v1alpha2sets.ServiceProfileSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", serviceProfileSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockServiceProfileSetMockRecorder) Equal(serviceProfileSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockServiceProfileSet)(nil).Equal), serviceProfileSet)
}

// Has mocks base method.
func (m *MockServiceProfileSet) Has(serviceProfile *v1alpha2.ServiceProfile) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", serviceProfile)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockServiceProfileSetMockRecorder) Has(serviceProfile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockServiceProfileSet)(nil).Has), serviceProfile)
}

// Delete mocks base method.
func (m *MockServiceProfileSet) Delete(serviceProfile *v1alpha2.ServiceProfile) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", serviceProfile)
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceProfileSetMockRecorder) Delete(serviceProfile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceProfileSet)(nil).Delete), serviceProfile)
}

// Union mocks base method.
func (m *MockServiceProfileSet) Union(set v1alpha2sets.ServiceProfileSet) v1alpha2sets.ServiceProfileSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha2sets.ServiceProfileSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockServiceProfileSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockServiceProfileSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockServiceProfileSet) Difference(set v1alpha2sets.ServiceProfileSet) v1alpha2sets.ServiceProfileSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha2sets.ServiceProfileSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockServiceProfileSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockServiceProfileSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockServiceProfileSet) Intersection(set v1alpha2sets.ServiceProfileSet) v1alpha2sets.ServiceProfileSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha2sets.ServiceProfileSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockServiceProfileSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockServiceProfileSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockServiceProfileSet) Find(id ezkube.ResourceId) (*v1alpha2.ServiceProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha2.ServiceProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockServiceProfileSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockServiceProfileSet)(nil).Find), id)
}
