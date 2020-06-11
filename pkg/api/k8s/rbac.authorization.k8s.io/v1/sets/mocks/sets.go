// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/rbac.authorization.k8s.io/v1/sets"
	v1 "k8s.io/api/rbac/v1"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockRoleSet is a mock of RoleSet interface.
type MockRoleSet struct {
	ctrl     *gomock.Controller
	recorder *MockRoleSetMockRecorder
}

// MockRoleSetMockRecorder is the mock recorder for MockRoleSet.
type MockRoleSetMockRecorder struct {
	mock *MockRoleSet
}

// NewMockRoleSet creates a new mock instance.
func NewMockRoleSet(ctrl *gomock.Controller) *MockRoleSet {
	mock := &MockRoleSet{ctrl: ctrl}
	mock.recorder = &MockRoleSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleSet) EXPECT() *MockRoleSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockRoleSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockRoleSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRoleSet)(nil).Keys))
}

// List mocks base method.
func (m *MockRoleSet) List() []*v1.Role {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.Role)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRoleSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleSet)(nil).List))
}

// Map mocks base method.
func (m *MockRoleSet) Map() map[string]*v1.Role {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Role)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRoleSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRoleSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockRoleSet) Insert(role ...*v1.Role) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range role {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockRoleSetMockRecorder) Insert(role ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRoleSet)(nil).Insert), role...)
}

// Equal mocks base method.
func (m *MockRoleSet) Equal(roleSet v1sets.RoleSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", roleSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockRoleSetMockRecorder) Equal(roleSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockRoleSet)(nil).Equal), roleSet)
}

// Has mocks base method.
func (m *MockRoleSet) Has(role *v1.Role) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", role)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockRoleSetMockRecorder) Has(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRoleSet)(nil).Has), role)
}

// Delete mocks base method.
func (m *MockRoleSet) Delete(role *v1.Role) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", role)
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleSetMockRecorder) Delete(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleSet)(nil).Delete), role)
}

// Union mocks base method.
func (m *MockRoleSet) Union(set v1sets.RoleSet) v1sets.RoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.RoleSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockRoleSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockRoleSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockRoleSet) Difference(set v1sets.RoleSet) v1sets.RoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.RoleSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockRoleSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockRoleSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockRoleSet) Intersection(set v1sets.RoleSet) v1sets.RoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.RoleSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockRoleSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockRoleSet)(nil).Intersection), set)
}

// MockRoleBindingSet is a mock of RoleBindingSet interface.
type MockRoleBindingSet struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingSetMockRecorder
}

// MockRoleBindingSetMockRecorder is the mock recorder for MockRoleBindingSet.
type MockRoleBindingSetMockRecorder struct {
	mock *MockRoleBindingSet
}

// NewMockRoleBindingSet creates a new mock instance.
func NewMockRoleBindingSet(ctrl *gomock.Controller) *MockRoleBindingSet {
	mock := &MockRoleBindingSet{ctrl: ctrl}
	mock.recorder = &MockRoleBindingSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleBindingSet) EXPECT() *MockRoleBindingSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockRoleBindingSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockRoleBindingSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockRoleBindingSet)(nil).Keys))
}

// List mocks base method.
func (m *MockRoleBindingSet) List() []*v1.RoleBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.RoleBinding)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRoleBindingSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleBindingSet)(nil).List))
}

// Map mocks base method.
func (m *MockRoleBindingSet) Map() map[string]*v1.RoleBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.RoleBinding)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockRoleBindingSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockRoleBindingSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockRoleBindingSet) Insert(roleBinding ...*v1.RoleBinding) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range roleBinding {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockRoleBindingSetMockRecorder) Insert(roleBinding ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRoleBindingSet)(nil).Insert), roleBinding...)
}

// Equal mocks base method.
func (m *MockRoleBindingSet) Equal(roleBindingSet v1sets.RoleBindingSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", roleBindingSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockRoleBindingSetMockRecorder) Equal(roleBindingSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockRoleBindingSet)(nil).Equal), roleBindingSet)
}

// Has mocks base method.
func (m *MockRoleBindingSet) Has(roleBinding *v1.RoleBinding) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", roleBinding)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockRoleBindingSetMockRecorder) Has(roleBinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockRoleBindingSet)(nil).Has), roleBinding)
}

// Delete mocks base method.
func (m *MockRoleBindingSet) Delete(roleBinding *v1.RoleBinding) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", roleBinding)
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleBindingSetMockRecorder) Delete(roleBinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleBindingSet)(nil).Delete), roleBinding)
}

// Union mocks base method.
func (m *MockRoleBindingSet) Union(set v1sets.RoleBindingSet) v1sets.RoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.RoleBindingSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockRoleBindingSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockRoleBindingSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockRoleBindingSet) Difference(set v1sets.RoleBindingSet) v1sets.RoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.RoleBindingSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockRoleBindingSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockRoleBindingSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockRoleBindingSet) Intersection(set v1sets.RoleBindingSet) v1sets.RoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.RoleBindingSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockRoleBindingSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockRoleBindingSet)(nil).Intersection), set)
}

// MockClusterRoleSet is a mock of ClusterRoleSet interface.
type MockClusterRoleSet struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleSetMockRecorder
}

// MockClusterRoleSetMockRecorder is the mock recorder for MockClusterRoleSet.
type MockClusterRoleSetMockRecorder struct {
	mock *MockClusterRoleSet
}

// NewMockClusterRoleSet creates a new mock instance.
func NewMockClusterRoleSet(ctrl *gomock.Controller) *MockClusterRoleSet {
	mock := &MockClusterRoleSet{ctrl: ctrl}
	mock.recorder = &MockClusterRoleSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRoleSet) EXPECT() *MockClusterRoleSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockClusterRoleSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockClusterRoleSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockClusterRoleSet)(nil).Keys))
}

// List mocks base method.
func (m *MockClusterRoleSet) List() []*v1.ClusterRole {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.ClusterRole)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockClusterRoleSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRoleSet)(nil).List))
}

// Map mocks base method.
func (m *MockClusterRoleSet) Map() map[string]*v1.ClusterRole {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.ClusterRole)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockClusterRoleSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockClusterRoleSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockClusterRoleSet) Insert(clusterRole ...*v1.ClusterRole) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range clusterRole {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockClusterRoleSetMockRecorder) Insert(clusterRole ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockClusterRoleSet)(nil).Insert), clusterRole...)
}

// Equal mocks base method.
func (m *MockClusterRoleSet) Equal(clusterRoleSet v1sets.ClusterRoleSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", clusterRoleSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockClusterRoleSetMockRecorder) Equal(clusterRoleSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockClusterRoleSet)(nil).Equal), clusterRoleSet)
}

// Has mocks base method.
func (m *MockClusterRoleSet) Has(clusterRole *v1.ClusterRole) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", clusterRole)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockClusterRoleSetMockRecorder) Has(clusterRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockClusterRoleSet)(nil).Has), clusterRole)
}

// Delete mocks base method.
func (m *MockClusterRoleSet) Delete(clusterRole *v1.ClusterRole) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", clusterRole)
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterRoleSetMockRecorder) Delete(clusterRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterRoleSet)(nil).Delete), clusterRole)
}

// Union mocks base method.
func (m *MockClusterRoleSet) Union(set v1sets.ClusterRoleSet) v1sets.ClusterRoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockClusterRoleSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockClusterRoleSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockClusterRoleSet) Difference(set v1sets.ClusterRoleSet) v1sets.ClusterRoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockClusterRoleSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockClusterRoleSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockClusterRoleSet) Intersection(set v1sets.ClusterRoleSet) v1sets.ClusterRoleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockClusterRoleSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockClusterRoleSet)(nil).Intersection), set)
}

// MockClusterRoleBindingSet is a mock of ClusterRoleBindingSet interface.
type MockClusterRoleBindingSet struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleBindingSetMockRecorder
}

// MockClusterRoleBindingSetMockRecorder is the mock recorder for MockClusterRoleBindingSet.
type MockClusterRoleBindingSetMockRecorder struct {
	mock *MockClusterRoleBindingSet
}

// NewMockClusterRoleBindingSet creates a new mock instance.
func NewMockClusterRoleBindingSet(ctrl *gomock.Controller) *MockClusterRoleBindingSet {
	mock := &MockClusterRoleBindingSet{ctrl: ctrl}
	mock.recorder = &MockClusterRoleBindingSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRoleBindingSet) EXPECT() *MockClusterRoleBindingSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockClusterRoleBindingSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockClusterRoleBindingSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Keys))
}

// List mocks base method.
func (m *MockClusterRoleBindingSet) List() []*v1.ClusterRoleBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.ClusterRoleBinding)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockClusterRoleBindingSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).List))
}

// Map mocks base method.
func (m *MockClusterRoleBindingSet) Map() map[string]*v1.ClusterRoleBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.ClusterRoleBinding)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockClusterRoleBindingSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockClusterRoleBindingSet) Insert(clusterRoleBinding ...*v1.ClusterRoleBinding) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range clusterRoleBinding {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockClusterRoleBindingSetMockRecorder) Insert(clusterRoleBinding ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Insert), clusterRoleBinding...)
}

// Equal mocks base method.
func (m *MockClusterRoleBindingSet) Equal(clusterRoleBindingSet v1sets.ClusterRoleBindingSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", clusterRoleBindingSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockClusterRoleBindingSetMockRecorder) Equal(clusterRoleBindingSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Equal), clusterRoleBindingSet)
}

// Has mocks base method.
func (m *MockClusterRoleBindingSet) Has(clusterRoleBinding *v1.ClusterRoleBinding) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", clusterRoleBinding)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockClusterRoleBindingSetMockRecorder) Has(clusterRoleBinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Has), clusterRoleBinding)
}

// Delete mocks base method.
func (m *MockClusterRoleBindingSet) Delete(clusterRoleBinding *v1.ClusterRoleBinding) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", clusterRoleBinding)
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterRoleBindingSetMockRecorder) Delete(clusterRoleBinding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Delete), clusterRoleBinding)
}

// Union mocks base method.
func (m *MockClusterRoleBindingSet) Union(set v1sets.ClusterRoleBindingSet) v1sets.ClusterRoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleBindingSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockClusterRoleBindingSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockClusterRoleBindingSet) Difference(set v1sets.ClusterRoleBindingSet) v1sets.ClusterRoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleBindingSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockClusterRoleBindingSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockClusterRoleBindingSet) Intersection(set v1sets.ClusterRoleBindingSet) v1sets.ClusterRoleBindingSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.ClusterRoleBindingSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockClusterRoleBindingSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockClusterRoleBindingSet)(nil).Intersection), set)
}
