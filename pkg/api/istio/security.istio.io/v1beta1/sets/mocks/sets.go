// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1beta1sets is a generated GoMock package.
package mock_v1beta1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1sets "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/sets"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockAuthorizationPolicySet is a mock of AuthorizationPolicySet interface.
type MockAuthorizationPolicySet struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationPolicySetMockRecorder
}

// MockAuthorizationPolicySetMockRecorder is the mock recorder for MockAuthorizationPolicySet.
type MockAuthorizationPolicySetMockRecorder struct {
	mock *MockAuthorizationPolicySet
}

// NewMockAuthorizationPolicySet creates a new mock instance.
func NewMockAuthorizationPolicySet(ctrl *gomock.Controller) *MockAuthorizationPolicySet {
	mock := &MockAuthorizationPolicySet{ctrl: ctrl}
	mock.recorder = &MockAuthorizationPolicySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationPolicySet) EXPECT() *MockAuthorizationPolicySetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockAuthorizationPolicySet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockAuthorizationPolicySetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Keys))
}

// List mocks base method.
func (m *MockAuthorizationPolicySet) List() []*v1beta1.AuthorizationPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1beta1.AuthorizationPolicy)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockAuthorizationPolicySetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).List))
}

// Map mocks base method.
func (m *MockAuthorizationPolicySet) Map() map[string]*v1beta1.AuthorizationPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1beta1.AuthorizationPolicy)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockAuthorizationPolicySetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Map))
}

// Insert mocks base method.
func (m *MockAuthorizationPolicySet) Insert(authorizationPolicy ...*v1beta1.AuthorizationPolicy) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range authorizationPolicy {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockAuthorizationPolicySetMockRecorder) Insert(authorizationPolicy ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Insert), authorizationPolicy...)
}

// Equal mocks base method.
func (m *MockAuthorizationPolicySet) Equal(authorizationPolicySet v1beta1sets.AuthorizationPolicySet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", authorizationPolicySet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockAuthorizationPolicySetMockRecorder) Equal(authorizationPolicySet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Equal), authorizationPolicySet)
}

// Has mocks base method.
func (m *MockAuthorizationPolicySet) Has(authorizationPolicy *v1beta1.AuthorizationPolicy) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", authorizationPolicy)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockAuthorizationPolicySetMockRecorder) Has(authorizationPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Has), authorizationPolicy)
}

// Delete mocks base method.
func (m *MockAuthorizationPolicySet) Delete(authorizationPolicy *v1beta1.AuthorizationPolicy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", authorizationPolicy)
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorizationPolicySetMockRecorder) Delete(authorizationPolicy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Delete), authorizationPolicy)
}

// Union mocks base method.
func (m *MockAuthorizationPolicySet) Union(set v1beta1sets.AuthorizationPolicySet) v1beta1sets.AuthorizationPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1beta1sets.AuthorizationPolicySet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockAuthorizationPolicySetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockAuthorizationPolicySet) Difference(set v1beta1sets.AuthorizationPolicySet) v1beta1sets.AuthorizationPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1beta1sets.AuthorizationPolicySet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockAuthorizationPolicySetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockAuthorizationPolicySet) Intersection(set v1beta1sets.AuthorizationPolicySet) v1beta1sets.AuthorizationPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1beta1sets.AuthorizationPolicySet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockAuthorizationPolicySetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockAuthorizationPolicySet)(nil).Intersection), set)
}
