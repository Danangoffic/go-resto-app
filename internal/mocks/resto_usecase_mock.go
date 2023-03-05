// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	model "resto-app/internal/model"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// CheckSession mocks base method.
func (m *MockUsecase) CheckSession(ctx context.Context, data model.UserSession) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSession", ctx, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession.
func (mr *MockUsecaseMockRecorder) CheckSession(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockUsecase)(nil).CheckSession), ctx, data)
}

// GetMenu mocks base method.
func (m *MockUsecase) GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenu", ctx, orderCode)
	ret0, _ := ret[0].(model.MenuItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu.
func (mr *MockUsecaseMockRecorder) GetMenu(ctx, orderCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockUsecase)(nil).GetMenu), ctx, orderCode)
}

// GetMenuList mocks base method.
func (m *MockUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuList", ctx, menuType)
	ret0, _ := ret[0].([]model.MenuItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuList indicates an expected call of GetMenuList.
func (mr *MockUsecaseMockRecorder) GetMenuList(ctx, menuType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuList", reflect.TypeOf((*MockUsecase)(nil).GetMenuList), ctx, menuType)
}

// GetOrderData mocks base method.
func (m *MockUsecase) GetOrderData(ctx context.Context, request model.GetOrderDataRequest) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderData", ctx, request)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderData indicates an expected call of GetOrderData.
func (mr *MockUsecaseMockRecorder) GetOrderData(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderData", reflect.TypeOf((*MockUsecase)(nil).GetOrderData), ctx, request)
}

// Login mocks base method.
func (m *MockUsecase) Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, request)
	ret0, _ := ret[0].(model.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUsecaseMockRecorder) Login(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsecase)(nil).Login), ctx, request)
}

// Order mocks base method.
func (m *MockUsecase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order", ctx, request)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Order indicates an expected call of Order.
func (mr *MockUsecaseMockRecorder) Order(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockUsecase)(nil).Order), ctx, request)
}

// RegisterUser mocks base method.
func (m *MockUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, request)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUsecaseMockRecorder) RegisterUser(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUsecase)(nil).RegisterUser), ctx, request)
}