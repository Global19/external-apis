// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	v1 "k8s.io/api/core/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockSecretEventHandler is a mock of SecretEventHandler interface.
type MockSecretEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSecretEventHandlerMockRecorder
}

// MockSecretEventHandlerMockRecorder is the mock recorder for MockSecretEventHandler.
type MockSecretEventHandlerMockRecorder struct {
	mock *MockSecretEventHandler
}

// NewMockSecretEventHandler creates a new mock instance.
func NewMockSecretEventHandler(ctrl *gomock.Controller) *MockSecretEventHandler {
	mock := &MockSecretEventHandler{ctrl: ctrl}
	mock.recorder = &MockSecretEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretEventHandler) EXPECT() *MockSecretEventHandlerMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method.
func (m *MockSecretEventHandler) CreateSecret(obj *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockSecretEventHandlerMockRecorder) CreateSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretEventHandler)(nil).CreateSecret), obj)
}

// UpdateSecret mocks base method.
func (m *MockSecretEventHandler) UpdateSecret(old, new *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockSecretEventHandlerMockRecorder) UpdateSecret(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockSecretEventHandler)(nil).UpdateSecret), old, new)
}

// DeleteSecret mocks base method.
func (m *MockSecretEventHandler) DeleteSecret(obj *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockSecretEventHandlerMockRecorder) DeleteSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretEventHandler)(nil).DeleteSecret), obj)
}

// GenericSecret mocks base method.
func (m *MockSecretEventHandler) GenericSecret(obj *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericSecret", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericSecret indicates an expected call of GenericSecret.
func (mr *MockSecretEventHandlerMockRecorder) GenericSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericSecret", reflect.TypeOf((*MockSecretEventHandler)(nil).GenericSecret), obj)
}

// MockSecretEventWatcher is a mock of SecretEventWatcher interface.
type MockSecretEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockSecretEventWatcherMockRecorder
}

// MockSecretEventWatcherMockRecorder is the mock recorder for MockSecretEventWatcher.
type MockSecretEventWatcherMockRecorder struct {
	mock *MockSecretEventWatcher
}

// NewMockSecretEventWatcher creates a new mock instance.
func NewMockSecretEventWatcher(ctrl *gomock.Controller) *MockSecretEventWatcher {
	mock := &MockSecretEventWatcher{ctrl: ctrl}
	mock.recorder = &MockSecretEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretEventWatcher) EXPECT() *MockSecretEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockSecretEventWatcher) AddEventHandler(ctx context.Context, h controller.SecretEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockSecretEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockSecretEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockServiceAccountEventHandler is a mock of ServiceAccountEventHandler interface.
type MockServiceAccountEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountEventHandlerMockRecorder
}

// MockServiceAccountEventHandlerMockRecorder is the mock recorder for MockServiceAccountEventHandler.
type MockServiceAccountEventHandlerMockRecorder struct {
	mock *MockServiceAccountEventHandler
}

// NewMockServiceAccountEventHandler creates a new mock instance.
func NewMockServiceAccountEventHandler(ctrl *gomock.Controller) *MockServiceAccountEventHandler {
	mock := &MockServiceAccountEventHandler{ctrl: ctrl}
	mock.recorder = &MockServiceAccountEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountEventHandler) EXPECT() *MockServiceAccountEventHandlerMockRecorder {
	return m.recorder
}

// CreateServiceAccount mocks base method.
func (m *MockServiceAccountEventHandler) CreateServiceAccount(obj *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceAccount", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateServiceAccount indicates an expected call of CreateServiceAccount.
func (mr *MockServiceAccountEventHandlerMockRecorder) CreateServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceAccount", reflect.TypeOf((*MockServiceAccountEventHandler)(nil).CreateServiceAccount), obj)
}

// UpdateServiceAccount mocks base method.
func (m *MockServiceAccountEventHandler) UpdateServiceAccount(old, new *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServiceAccount", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateServiceAccount indicates an expected call of UpdateServiceAccount.
func (mr *MockServiceAccountEventHandlerMockRecorder) UpdateServiceAccount(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServiceAccount", reflect.TypeOf((*MockServiceAccountEventHandler)(nil).UpdateServiceAccount), old, new)
}

// DeleteServiceAccount mocks base method.
func (m *MockServiceAccountEventHandler) DeleteServiceAccount(obj *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceAccount", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceAccount indicates an expected call of DeleteServiceAccount.
func (mr *MockServiceAccountEventHandlerMockRecorder) DeleteServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceAccount", reflect.TypeOf((*MockServiceAccountEventHandler)(nil).DeleteServiceAccount), obj)
}

// GenericServiceAccount mocks base method.
func (m *MockServiceAccountEventHandler) GenericServiceAccount(obj *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericServiceAccount", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericServiceAccount indicates an expected call of GenericServiceAccount.
func (mr *MockServiceAccountEventHandlerMockRecorder) GenericServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericServiceAccount", reflect.TypeOf((*MockServiceAccountEventHandler)(nil).GenericServiceAccount), obj)
}

// MockServiceAccountEventWatcher is a mock of ServiceAccountEventWatcher interface.
type MockServiceAccountEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountEventWatcherMockRecorder
}

// MockServiceAccountEventWatcherMockRecorder is the mock recorder for MockServiceAccountEventWatcher.
type MockServiceAccountEventWatcherMockRecorder struct {
	mock *MockServiceAccountEventWatcher
}

// NewMockServiceAccountEventWatcher creates a new mock instance.
func NewMockServiceAccountEventWatcher(ctrl *gomock.Controller) *MockServiceAccountEventWatcher {
	mock := &MockServiceAccountEventWatcher{ctrl: ctrl}
	mock.recorder = &MockServiceAccountEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountEventWatcher) EXPECT() *MockServiceAccountEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockServiceAccountEventWatcher) AddEventHandler(ctx context.Context, h controller.ServiceAccountEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockServiceAccountEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockServiceAccountEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockConfigMapEventHandler is a mock of ConfigMapEventHandler interface.
type MockConfigMapEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMapEventHandlerMockRecorder
}

// MockConfigMapEventHandlerMockRecorder is the mock recorder for MockConfigMapEventHandler.
type MockConfigMapEventHandlerMockRecorder struct {
	mock *MockConfigMapEventHandler
}

// NewMockConfigMapEventHandler creates a new mock instance.
func NewMockConfigMapEventHandler(ctrl *gomock.Controller) *MockConfigMapEventHandler {
	mock := &MockConfigMapEventHandler{ctrl: ctrl}
	mock.recorder = &MockConfigMapEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigMapEventHandler) EXPECT() *MockConfigMapEventHandlerMockRecorder {
	return m.recorder
}

// CreateConfigMap mocks base method.
func (m *MockConfigMapEventHandler) CreateConfigMap(obj *v1.ConfigMap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConfigMap", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateConfigMap indicates an expected call of CreateConfigMap.
func (mr *MockConfigMapEventHandlerMockRecorder) CreateConfigMap(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfigMap", reflect.TypeOf((*MockConfigMapEventHandler)(nil).CreateConfigMap), obj)
}

// UpdateConfigMap mocks base method.
func (m *MockConfigMapEventHandler) UpdateConfigMap(old, new *v1.ConfigMap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfigMap", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfigMap indicates an expected call of UpdateConfigMap.
func (mr *MockConfigMapEventHandlerMockRecorder) UpdateConfigMap(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfigMap", reflect.TypeOf((*MockConfigMapEventHandler)(nil).UpdateConfigMap), old, new)
}

// DeleteConfigMap mocks base method.
func (m *MockConfigMapEventHandler) DeleteConfigMap(obj *v1.ConfigMap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConfigMap", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConfigMap indicates an expected call of DeleteConfigMap.
func (mr *MockConfigMapEventHandlerMockRecorder) DeleteConfigMap(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfigMap", reflect.TypeOf((*MockConfigMapEventHandler)(nil).DeleteConfigMap), obj)
}

// GenericConfigMap mocks base method.
func (m *MockConfigMapEventHandler) GenericConfigMap(obj *v1.ConfigMap) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericConfigMap", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericConfigMap indicates an expected call of GenericConfigMap.
func (mr *MockConfigMapEventHandlerMockRecorder) GenericConfigMap(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericConfigMap", reflect.TypeOf((*MockConfigMapEventHandler)(nil).GenericConfigMap), obj)
}

// MockConfigMapEventWatcher is a mock of ConfigMapEventWatcher interface.
type MockConfigMapEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMapEventWatcherMockRecorder
}

// MockConfigMapEventWatcherMockRecorder is the mock recorder for MockConfigMapEventWatcher.
type MockConfigMapEventWatcherMockRecorder struct {
	mock *MockConfigMapEventWatcher
}

// NewMockConfigMapEventWatcher creates a new mock instance.
func NewMockConfigMapEventWatcher(ctrl *gomock.Controller) *MockConfigMapEventWatcher {
	mock := &MockConfigMapEventWatcher{ctrl: ctrl}
	mock.recorder = &MockConfigMapEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigMapEventWatcher) EXPECT() *MockConfigMapEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockConfigMapEventWatcher) AddEventHandler(ctx context.Context, h controller.ConfigMapEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockConfigMapEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockConfigMapEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockServiceEventHandler is a mock of ServiceEventHandler interface.
type MockServiceEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEventHandlerMockRecorder
}

// MockServiceEventHandlerMockRecorder is the mock recorder for MockServiceEventHandler.
type MockServiceEventHandlerMockRecorder struct {
	mock *MockServiceEventHandler
}

// NewMockServiceEventHandler creates a new mock instance.
func NewMockServiceEventHandler(ctrl *gomock.Controller) *MockServiceEventHandler {
	mock := &MockServiceEventHandler{ctrl: ctrl}
	mock.recorder = &MockServiceEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEventHandler) EXPECT() *MockServiceEventHandlerMockRecorder {
	return m.recorder
}

// CreateService mocks base method.
func (m *MockServiceEventHandler) CreateService(obj *v1.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService.
func (mr *MockServiceEventHandlerMockRecorder) CreateService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockServiceEventHandler)(nil).CreateService), obj)
}

// UpdateService mocks base method.
func (m *MockServiceEventHandler) UpdateService(old, new *v1.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockServiceEventHandlerMockRecorder) UpdateService(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockServiceEventHandler)(nil).UpdateService), old, new)
}

// DeleteService mocks base method.
func (m *MockServiceEventHandler) DeleteService(obj *v1.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockServiceEventHandlerMockRecorder) DeleteService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockServiceEventHandler)(nil).DeleteService), obj)
}

// GenericService mocks base method.
func (m *MockServiceEventHandler) GenericService(obj *v1.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericService indicates an expected call of GenericService.
func (mr *MockServiceEventHandlerMockRecorder) GenericService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericService", reflect.TypeOf((*MockServiceEventHandler)(nil).GenericService), obj)
}

// MockServiceEventWatcher is a mock of ServiceEventWatcher interface.
type MockServiceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEventWatcherMockRecorder
}

// MockServiceEventWatcherMockRecorder is the mock recorder for MockServiceEventWatcher.
type MockServiceEventWatcherMockRecorder struct {
	mock *MockServiceEventWatcher
}

// NewMockServiceEventWatcher creates a new mock instance.
func NewMockServiceEventWatcher(ctrl *gomock.Controller) *MockServiceEventWatcher {
	mock := &MockServiceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockServiceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEventWatcher) EXPECT() *MockServiceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockServiceEventWatcher) AddEventHandler(ctx context.Context, h controller.ServiceEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockServiceEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockServiceEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockPodEventHandler is a mock of PodEventHandler interface.
type MockPodEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPodEventHandlerMockRecorder
}

// MockPodEventHandlerMockRecorder is the mock recorder for MockPodEventHandler.
type MockPodEventHandlerMockRecorder struct {
	mock *MockPodEventHandler
}

// NewMockPodEventHandler creates a new mock instance.
func NewMockPodEventHandler(ctrl *gomock.Controller) *MockPodEventHandler {
	mock := &MockPodEventHandler{ctrl: ctrl}
	mock.recorder = &MockPodEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodEventHandler) EXPECT() *MockPodEventHandlerMockRecorder {
	return m.recorder
}

// CreatePod mocks base method.
func (m *MockPodEventHandler) CreatePod(obj *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePod", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePod indicates an expected call of CreatePod.
func (mr *MockPodEventHandlerMockRecorder) CreatePod(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePod", reflect.TypeOf((*MockPodEventHandler)(nil).CreatePod), obj)
}

// UpdatePod mocks base method.
func (m *MockPodEventHandler) UpdatePod(old, new *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePod", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePod indicates an expected call of UpdatePod.
func (mr *MockPodEventHandlerMockRecorder) UpdatePod(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePod", reflect.TypeOf((*MockPodEventHandler)(nil).UpdatePod), old, new)
}

// DeletePod mocks base method.
func (m *MockPodEventHandler) DeletePod(obj *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePod", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePod indicates an expected call of DeletePod.
func (mr *MockPodEventHandlerMockRecorder) DeletePod(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePod", reflect.TypeOf((*MockPodEventHandler)(nil).DeletePod), obj)
}

// GenericPod mocks base method.
func (m *MockPodEventHandler) GenericPod(obj *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericPod", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericPod indicates an expected call of GenericPod.
func (mr *MockPodEventHandlerMockRecorder) GenericPod(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericPod", reflect.TypeOf((*MockPodEventHandler)(nil).GenericPod), obj)
}

// MockPodEventWatcher is a mock of PodEventWatcher interface.
type MockPodEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockPodEventWatcherMockRecorder
}

// MockPodEventWatcherMockRecorder is the mock recorder for MockPodEventWatcher.
type MockPodEventWatcherMockRecorder struct {
	mock *MockPodEventWatcher
}

// NewMockPodEventWatcher creates a new mock instance.
func NewMockPodEventWatcher(ctrl *gomock.Controller) *MockPodEventWatcher {
	mock := &MockPodEventWatcher{ctrl: ctrl}
	mock.recorder = &MockPodEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodEventWatcher) EXPECT() *MockPodEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockPodEventWatcher) AddEventHandler(ctx context.Context, h controller.PodEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockPodEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockPodEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockNamespaceEventHandler is a mock of NamespaceEventHandler interface.
type MockNamespaceEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceEventHandlerMockRecorder
}

// MockNamespaceEventHandlerMockRecorder is the mock recorder for MockNamespaceEventHandler.
type MockNamespaceEventHandlerMockRecorder struct {
	mock *MockNamespaceEventHandler
}

// NewMockNamespaceEventHandler creates a new mock instance.
func NewMockNamespaceEventHandler(ctrl *gomock.Controller) *MockNamespaceEventHandler {
	mock := &MockNamespaceEventHandler{ctrl: ctrl}
	mock.recorder = &MockNamespaceEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceEventHandler) EXPECT() *MockNamespaceEventHandlerMockRecorder {
	return m.recorder
}

// CreateNamespace mocks base method.
func (m *MockNamespaceEventHandler) CreateNamespace(obj *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNamespace indicates an expected call of CreateNamespace.
func (mr *MockNamespaceEventHandlerMockRecorder) CreateNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockNamespaceEventHandler)(nil).CreateNamespace), obj)
}

// UpdateNamespace mocks base method.
func (m *MockNamespaceEventHandler) UpdateNamespace(old, new *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNamespace", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNamespace indicates an expected call of UpdateNamespace.
func (mr *MockNamespaceEventHandlerMockRecorder) UpdateNamespace(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespace", reflect.TypeOf((*MockNamespaceEventHandler)(nil).UpdateNamespace), old, new)
}

// DeleteNamespace mocks base method.
func (m *MockNamespaceEventHandler) DeleteNamespace(obj *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespace indicates an expected call of DeleteNamespace.
func (mr *MockNamespaceEventHandlerMockRecorder) DeleteNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockNamespaceEventHandler)(nil).DeleteNamespace), obj)
}

// GenericNamespace mocks base method.
func (m *MockNamespaceEventHandler) GenericNamespace(obj *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericNamespace", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericNamespace indicates an expected call of GenericNamespace.
func (mr *MockNamespaceEventHandlerMockRecorder) GenericNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericNamespace", reflect.TypeOf((*MockNamespaceEventHandler)(nil).GenericNamespace), obj)
}

// MockNamespaceEventWatcher is a mock of NamespaceEventWatcher interface.
type MockNamespaceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceEventWatcherMockRecorder
}

// MockNamespaceEventWatcherMockRecorder is the mock recorder for MockNamespaceEventWatcher.
type MockNamespaceEventWatcherMockRecorder struct {
	mock *MockNamespaceEventWatcher
}

// NewMockNamespaceEventWatcher creates a new mock instance.
func NewMockNamespaceEventWatcher(ctrl *gomock.Controller) *MockNamespaceEventWatcher {
	mock := &MockNamespaceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockNamespaceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceEventWatcher) EXPECT() *MockNamespaceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockNamespaceEventWatcher) AddEventHandler(ctx context.Context, h controller.NamespaceEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockNamespaceEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockNamespaceEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockNodeEventHandler is a mock of NodeEventHandler interface.
type MockNodeEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockNodeEventHandlerMockRecorder
}

// MockNodeEventHandlerMockRecorder is the mock recorder for MockNodeEventHandler.
type MockNodeEventHandlerMockRecorder struct {
	mock *MockNodeEventHandler
}

// NewMockNodeEventHandler creates a new mock instance.
func NewMockNodeEventHandler(ctrl *gomock.Controller) *MockNodeEventHandler {
	mock := &MockNodeEventHandler{ctrl: ctrl}
	mock.recorder = &MockNodeEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeEventHandler) EXPECT() *MockNodeEventHandlerMockRecorder {
	return m.recorder
}

// CreateNode mocks base method.
func (m *MockNodeEventHandler) CreateNode(obj *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNode", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNode indicates an expected call of CreateNode.
func (mr *MockNodeEventHandlerMockRecorder) CreateNode(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNode", reflect.TypeOf((*MockNodeEventHandler)(nil).CreateNode), obj)
}

// UpdateNode mocks base method.
func (m *MockNodeEventHandler) UpdateNode(old, new *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNode", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNode indicates an expected call of UpdateNode.
func (mr *MockNodeEventHandlerMockRecorder) UpdateNode(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNode", reflect.TypeOf((*MockNodeEventHandler)(nil).UpdateNode), old, new)
}

// DeleteNode mocks base method.
func (m *MockNodeEventHandler) DeleteNode(obj *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNode", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNode indicates an expected call of DeleteNode.
func (mr *MockNodeEventHandlerMockRecorder) DeleteNode(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockNodeEventHandler)(nil).DeleteNode), obj)
}

// GenericNode mocks base method.
func (m *MockNodeEventHandler) GenericNode(obj *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericNode", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericNode indicates an expected call of GenericNode.
func (mr *MockNodeEventHandlerMockRecorder) GenericNode(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericNode", reflect.TypeOf((*MockNodeEventHandler)(nil).GenericNode), obj)
}

// MockNodeEventWatcher is a mock of NodeEventWatcher interface.
type MockNodeEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockNodeEventWatcherMockRecorder
}

// MockNodeEventWatcherMockRecorder is the mock recorder for MockNodeEventWatcher.
type MockNodeEventWatcherMockRecorder struct {
	mock *MockNodeEventWatcher
}

// NewMockNodeEventWatcher creates a new mock instance.
func NewMockNodeEventWatcher(ctrl *gomock.Controller) *MockNodeEventWatcher {
	mock := &MockNodeEventWatcher{ctrl: ctrl}
	mock.recorder = &MockNodeEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeEventWatcher) EXPECT() *MockNodeEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockNodeEventWatcher) AddEventHandler(ctx context.Context, h controller.NodeEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockNodeEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockNodeEventWatcher)(nil).AddEventHandler), varargs...)
}
