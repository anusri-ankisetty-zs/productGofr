// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anusri-ankisetty-zs/productGofr/services (interfaces: Iservice)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	gofr "developer.zopsmart.com/go/gofr/pkg/gofr"
	models "github.com/anusri-ankisetty-zs/productGofr/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIservice is a mock of Iservice interface.
type MockIservice struct {
	ctrl     *gomock.Controller
	recorder *MockIserviceMockRecorder
}

// MockIserviceMockRecorder is the mock recorder for MockIservice.
type MockIserviceMockRecorder struct {
	mock *MockIservice
}

// NewMockIservice creates a new mock instance.
func NewMockIservice(ctrl *gomock.Controller) *MockIservice {
	mock := &MockIservice{ctrl: ctrl}
	mock.recorder = &MockIserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIservice) EXPECT() *MockIserviceMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockIservice) CreateProduct(arg0 *gofr.Context, arg1 models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockIserviceMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockIservice)(nil).CreateProduct), arg0, arg1)
}

// DeleteById mocks base method.
func (m *MockIservice) DeleteById(arg0 *gofr.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockIserviceMockRecorder) DeleteById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockIservice)(nil).DeleteById), arg0, arg1)
}

// GetAllUsers mocks base method.
func (m *MockIservice) GetAllUsers(arg0 *gofr.Context) ([]*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0)
	ret0, _ := ret[0].([]*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIserviceMockRecorder) GetAllUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIservice)(nil).GetAllUsers), arg0)
}

// GetProductById mocks base method.
func (m *MockIservice) GetProductById(arg0 *gofr.Context, arg1 string) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", arg0, arg1)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockIserviceMockRecorder) GetProductById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockIservice)(nil).GetProductById), arg0, arg1)
}

// UpdateById mocks base method.
func (m *MockIservice) UpdateById(arg0 *gofr.Context, arg1 string, arg2 models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockIserviceMockRecorder) UpdateById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockIservice)(nil).UpdateById), arg0, arg1, arg2)
}
