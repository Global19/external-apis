// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha3 is a generated GoMock package.
package mock_v1alpha3

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	v1alpha30 "github.com/solo-io/external-apis/pkg/api/smi/specs.smi-spec.io/v1alpha3"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1alpha30.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha30.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// HTTPRouteGroups mocks base method
func (m *MockClientset) HTTPRouteGroups() v1alpha30.HTTPRouteGroupClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPRouteGroups")
	ret0, _ := ret[0].(v1alpha30.HTTPRouteGroupClient)
	return ret0
}

// HTTPRouteGroups indicates an expected call of HTTPRouteGroups
func (mr *MockClientsetMockRecorder) HTTPRouteGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPRouteGroups", reflect.TypeOf((*MockClientset)(nil).HTTPRouteGroups))
}

// MockHTTPRouteGroupReader is a mock of HTTPRouteGroupReader interface
type MockHTTPRouteGroupReader struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupReaderMockRecorder
}

// MockHTTPRouteGroupReaderMockRecorder is the mock recorder for MockHTTPRouteGroupReader
type MockHTTPRouteGroupReaderMockRecorder struct {
	mock *MockHTTPRouteGroupReader
}

// NewMockHTTPRouteGroupReader creates a new mock instance
func NewMockHTTPRouteGroupReader(ctrl *gomock.Controller) *MockHTTPRouteGroupReader {
	mock := &MockHTTPRouteGroupReader{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupReader) EXPECT() *MockHTTPRouteGroupReaderMockRecorder {
	return m.recorder
}

// GetHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupReader) GetHTTPRouteGroup(ctx context.Context, key client.ObjectKey) (*v1alpha3.HTTPRouteGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPRouteGroup", ctx, key)
	ret0, _ := ret[0].(*v1alpha3.HTTPRouteGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHTTPRouteGroup indicates an expected call of GetHTTPRouteGroup
func (mr *MockHTTPRouteGroupReaderMockRecorder) GetHTTPRouteGroup(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupReader)(nil).GetHTTPRouteGroup), ctx, key)
}

// ListHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupReader) ListHTTPRouteGroup(ctx context.Context, opts ...client.ListOption) (*v1alpha3.HTTPRouteGroupList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(*v1alpha3.HTTPRouteGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHTTPRouteGroup indicates an expected call of ListHTTPRouteGroup
func (mr *MockHTTPRouteGroupReaderMockRecorder) ListHTTPRouteGroup(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupReader)(nil).ListHTTPRouteGroup), varargs...)
}

// MockHTTPRouteGroupWriter is a mock of HTTPRouteGroupWriter interface
type MockHTTPRouteGroupWriter struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupWriterMockRecorder
}

// MockHTTPRouteGroupWriterMockRecorder is the mock recorder for MockHTTPRouteGroupWriter
type MockHTTPRouteGroupWriterMockRecorder struct {
	mock *MockHTTPRouteGroupWriter
}

// NewMockHTTPRouteGroupWriter creates a new mock instance
func NewMockHTTPRouteGroupWriter(ctrl *gomock.Controller) *MockHTTPRouteGroupWriter {
	mock := &MockHTTPRouteGroupWriter{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupWriter) EXPECT() *MockHTTPRouteGroupWriterMockRecorder {
	return m.recorder
}

// CreateHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupWriter) CreateHTTPRouteGroup(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHTTPRouteGroup indicates an expected call of CreateHTTPRouteGroup
func (mr *MockHTTPRouteGroupWriterMockRecorder) CreateHTTPRouteGroup(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupWriter)(nil).CreateHTTPRouteGroup), varargs...)
}

// DeleteHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupWriter) DeleteHTTPRouteGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHTTPRouteGroup indicates an expected call of DeleteHTTPRouteGroup
func (mr *MockHTTPRouteGroupWriterMockRecorder) DeleteHTTPRouteGroup(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupWriter)(nil).DeleteHTTPRouteGroup), varargs...)
}

// UpdateHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupWriter) UpdateHTTPRouteGroup(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHTTPRouteGroup indicates an expected call of UpdateHTTPRouteGroup
func (mr *MockHTTPRouteGroupWriterMockRecorder) UpdateHTTPRouteGroup(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupWriter)(nil).UpdateHTTPRouteGroup), varargs...)
}

// PatchHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupWriter) PatchHTTPRouteGroup(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchHTTPRouteGroup indicates an expected call of PatchHTTPRouteGroup
func (mr *MockHTTPRouteGroupWriterMockRecorder) PatchHTTPRouteGroup(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupWriter)(nil).PatchHTTPRouteGroup), varargs...)
}

// DeleteAllOfHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupWriter) DeleteAllOfHTTPRouteGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfHTTPRouteGroup indicates an expected call of DeleteAllOfHTTPRouteGroup
func (mr *MockHTTPRouteGroupWriterMockRecorder) DeleteAllOfHTTPRouteGroup(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupWriter)(nil).DeleteAllOfHTTPRouteGroup), varargs...)
}

// UpsertHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupWriter) UpsertHTTPRouteGroup(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, transitionFuncs ...v1alpha30.HTTPRouteGroupTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertHTTPRouteGroup indicates an expected call of UpsertHTTPRouteGroup
func (mr *MockHTTPRouteGroupWriterMockRecorder) UpsertHTTPRouteGroup(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupWriter)(nil).UpsertHTTPRouteGroup), varargs...)
}

// MockHTTPRouteGroupStatusWriter is a mock of HTTPRouteGroupStatusWriter interface
type MockHTTPRouteGroupStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupStatusWriterMockRecorder
}

// MockHTTPRouteGroupStatusWriterMockRecorder is the mock recorder for MockHTTPRouteGroupStatusWriter
type MockHTTPRouteGroupStatusWriterMockRecorder struct {
	mock *MockHTTPRouteGroupStatusWriter
}

// NewMockHTTPRouteGroupStatusWriter creates a new mock instance
func NewMockHTTPRouteGroupStatusWriter(ctrl *gomock.Controller) *MockHTTPRouteGroupStatusWriter {
	mock := &MockHTTPRouteGroupStatusWriter{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupStatusWriter) EXPECT() *MockHTTPRouteGroupStatusWriterMockRecorder {
	return m.recorder
}

// UpdateHTTPRouteGroupStatus mocks base method
func (m *MockHTTPRouteGroupStatusWriter) UpdateHTTPRouteGroupStatus(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateHTTPRouteGroupStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHTTPRouteGroupStatus indicates an expected call of UpdateHTTPRouteGroupStatus
func (mr *MockHTTPRouteGroupStatusWriterMockRecorder) UpdateHTTPRouteGroupStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHTTPRouteGroupStatus", reflect.TypeOf((*MockHTTPRouteGroupStatusWriter)(nil).UpdateHTTPRouteGroupStatus), varargs...)
}

// PatchHTTPRouteGroupStatus mocks base method
func (m *MockHTTPRouteGroupStatusWriter) PatchHTTPRouteGroupStatus(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchHTTPRouteGroupStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchHTTPRouteGroupStatus indicates an expected call of PatchHTTPRouteGroupStatus
func (mr *MockHTTPRouteGroupStatusWriterMockRecorder) PatchHTTPRouteGroupStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchHTTPRouteGroupStatus", reflect.TypeOf((*MockHTTPRouteGroupStatusWriter)(nil).PatchHTTPRouteGroupStatus), varargs...)
}

// MockHTTPRouteGroupClient is a mock of HTTPRouteGroupClient interface
type MockHTTPRouteGroupClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouteGroupClientMockRecorder
}

// MockHTTPRouteGroupClientMockRecorder is the mock recorder for MockHTTPRouteGroupClient
type MockHTTPRouteGroupClientMockRecorder struct {
	mock *MockHTTPRouteGroupClient
}

// NewMockHTTPRouteGroupClient creates a new mock instance
func NewMockHTTPRouteGroupClient(ctrl *gomock.Controller) *MockHTTPRouteGroupClient {
	mock := &MockHTTPRouteGroupClient{ctrl: ctrl}
	mock.recorder = &MockHTTPRouteGroupClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRouteGroupClient) EXPECT() *MockHTTPRouteGroupClientMockRecorder {
	return m.recorder
}

// GetHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupClient) GetHTTPRouteGroup(ctx context.Context, key client.ObjectKey) (*v1alpha3.HTTPRouteGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPRouteGroup", ctx, key)
	ret0, _ := ret[0].(*v1alpha3.HTTPRouteGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHTTPRouteGroup indicates an expected call of GetHTTPRouteGroup
func (mr *MockHTTPRouteGroupClientMockRecorder) GetHTTPRouteGroup(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).GetHTTPRouteGroup), ctx, key)
}

// ListHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupClient) ListHTTPRouteGroup(ctx context.Context, opts ...client.ListOption) (*v1alpha3.HTTPRouteGroupList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(*v1alpha3.HTTPRouteGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHTTPRouteGroup indicates an expected call of ListHTTPRouteGroup
func (mr *MockHTTPRouteGroupClientMockRecorder) ListHTTPRouteGroup(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).ListHTTPRouteGroup), varargs...)
}

// CreateHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupClient) CreateHTTPRouteGroup(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHTTPRouteGroup indicates an expected call of CreateHTTPRouteGroup
func (mr *MockHTTPRouteGroupClientMockRecorder) CreateHTTPRouteGroup(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).CreateHTTPRouteGroup), varargs...)
}

// DeleteHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupClient) DeleteHTTPRouteGroup(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHTTPRouteGroup indicates an expected call of DeleteHTTPRouteGroup
func (mr *MockHTTPRouteGroupClientMockRecorder) DeleteHTTPRouteGroup(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).DeleteHTTPRouteGroup), varargs...)
}

// UpdateHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupClient) UpdateHTTPRouteGroup(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHTTPRouteGroup indicates an expected call of UpdateHTTPRouteGroup
func (mr *MockHTTPRouteGroupClientMockRecorder) UpdateHTTPRouteGroup(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).UpdateHTTPRouteGroup), varargs...)
}

// PatchHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupClient) PatchHTTPRouteGroup(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchHTTPRouteGroup indicates an expected call of PatchHTTPRouteGroup
func (mr *MockHTTPRouteGroupClientMockRecorder) PatchHTTPRouteGroup(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).PatchHTTPRouteGroup), varargs...)
}

// DeleteAllOfHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupClient) DeleteAllOfHTTPRouteGroup(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfHTTPRouteGroup indicates an expected call of DeleteAllOfHTTPRouteGroup
func (mr *MockHTTPRouteGroupClientMockRecorder) DeleteAllOfHTTPRouteGroup(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).DeleteAllOfHTTPRouteGroup), varargs...)
}

// UpsertHTTPRouteGroup mocks base method
func (m *MockHTTPRouteGroupClient) UpsertHTTPRouteGroup(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, transitionFuncs ...v1alpha30.HTTPRouteGroupTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertHTTPRouteGroup", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertHTTPRouteGroup indicates an expected call of UpsertHTTPRouteGroup
func (mr *MockHTTPRouteGroupClientMockRecorder) UpsertHTTPRouteGroup(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertHTTPRouteGroup", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).UpsertHTTPRouteGroup), varargs...)
}

// UpdateHTTPRouteGroupStatus mocks base method
func (m *MockHTTPRouteGroupClient) UpdateHTTPRouteGroupStatus(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateHTTPRouteGroupStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHTTPRouteGroupStatus indicates an expected call of UpdateHTTPRouteGroupStatus
func (mr *MockHTTPRouteGroupClientMockRecorder) UpdateHTTPRouteGroupStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHTTPRouteGroupStatus", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).UpdateHTTPRouteGroupStatus), varargs...)
}

// PatchHTTPRouteGroupStatus mocks base method
func (m *MockHTTPRouteGroupClient) PatchHTTPRouteGroupStatus(ctx context.Context, obj *v1alpha3.HTTPRouteGroup, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchHTTPRouteGroupStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchHTTPRouteGroupStatus indicates an expected call of PatchHTTPRouteGroupStatus
func (mr *MockHTTPRouteGroupClientMockRecorder) PatchHTTPRouteGroupStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchHTTPRouteGroupStatus", reflect.TypeOf((*MockHTTPRouteGroupClient)(nil).PatchHTTPRouteGroupStatus), varargs...)
}

// MockMulticlusterHTTPRouteGroupClient is a mock of MulticlusterHTTPRouteGroupClient interface
type MockMulticlusterHTTPRouteGroupClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterHTTPRouteGroupClientMockRecorder
}

// MockMulticlusterHTTPRouteGroupClientMockRecorder is the mock recorder for MockMulticlusterHTTPRouteGroupClient
type MockMulticlusterHTTPRouteGroupClientMockRecorder struct {
	mock *MockMulticlusterHTTPRouteGroupClient
}

// NewMockMulticlusterHTTPRouteGroupClient creates a new mock instance
func NewMockMulticlusterHTTPRouteGroupClient(ctrl *gomock.Controller) *MockMulticlusterHTTPRouteGroupClient {
	mock := &MockMulticlusterHTTPRouteGroupClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterHTTPRouteGroupClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterHTTPRouteGroupClient) EXPECT() *MockMulticlusterHTTPRouteGroupClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterHTTPRouteGroupClient) Cluster(cluster string) (v1alpha30.HTTPRouteGroupClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha30.HTTPRouteGroupClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterHTTPRouteGroupClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterHTTPRouteGroupClient)(nil).Cluster), cluster)
}