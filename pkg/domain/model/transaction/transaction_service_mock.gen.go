// Code generated by MockGen. DO NOT EDIT.
// Source: ./transaction_service.go

// Package transaction is a generated GoMock package.
package transaction

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// CommonBegin mocks base method.
func (m *MockTransactionService) CommonBegin(ctx context.Context) (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommonBegin", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommonBegin indicates an expected call of CommonBegin.
func (mr *MockTransactionServiceMockRecorder) CommonBegin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommonBegin", reflect.TypeOf((*MockTransactionService)(nil).CommonBegin), ctx)
}

// CommonEnd mocks base method.
func (m *MockTransactionService) CommonEnd(ctx context.Context, tx *gorm.DB, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CommonEnd", ctx, tx, err)
}

// CommonEnd indicates an expected call of CommonEnd.
func (mr *MockTransactionServiceMockRecorder) CommonEnd(ctx, tx, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommonEnd", reflect.TypeOf((*MockTransactionService)(nil).CommonEnd), ctx, tx, err)
}

// MasterBegin mocks base method.
func (m *MockTransactionService) MasterBegin(ctx context.Context) (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MasterBegin", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MasterBegin indicates an expected call of MasterBegin.
func (mr *MockTransactionServiceMockRecorder) MasterBegin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MasterBegin", reflect.TypeOf((*MockTransactionService)(nil).MasterBegin), ctx)
}

// MasterEnd mocks base method.
func (m *MockTransactionService) MasterEnd(ctx context.Context, tx *gorm.DB, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MasterEnd", ctx, tx, err)
}

// MasterEnd indicates an expected call of MasterEnd.
func (mr *MockTransactionServiceMockRecorder) MasterEnd(ctx, tx, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MasterEnd", reflect.TypeOf((*MockTransactionService)(nil).MasterEnd), ctx, tx, err)
}

// MultiUserBegin mocks base method.
func (m *MockTransactionService) MultiUserBegin(ctx context.Context, userIds []string) (map[string]*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiUserBegin", ctx, userIds)
	ret0, _ := ret[0].(map[string]*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiUserBegin indicates an expected call of MultiUserBegin.
func (mr *MockTransactionServiceMockRecorder) MultiUserBegin(ctx, userIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiUserBegin", reflect.TypeOf((*MockTransactionService)(nil).MultiUserBegin), ctx, userIds)
}

// MultiUserEnd mocks base method.
func (m *MockTransactionService) MultiUserEnd(ctx context.Context, txs map[string]*gorm.DB, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MultiUserEnd", ctx, txs, err)
}

// MultiUserEnd indicates an expected call of MultiUserEnd.
func (mr *MockTransactionServiceMockRecorder) MultiUserEnd(ctx, txs, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiUserEnd", reflect.TypeOf((*MockTransactionService)(nil).MultiUserEnd), ctx, txs, err)
}

// UserBegin mocks base method.
func (m *MockTransactionService) UserBegin(ctx context.Context, userId string) (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserBegin", ctx, userId)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserBegin indicates an expected call of UserBegin.
func (mr *MockTransactionServiceMockRecorder) UserBegin(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserBegin", reflect.TypeOf((*MockTransactionService)(nil).UserBegin), ctx, userId)
}

// UserEnd mocks base method.
func (m *MockTransactionService) UserEnd(ctx context.Context, tx *gorm.DB, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserEnd", ctx, tx, err)
}

// UserEnd indicates an expected call of UserEnd.
func (mr *MockTransactionServiceMockRecorder) UserEnd(ctx, tx, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserEnd", reflect.TypeOf((*MockTransactionService)(nil).UserEnd), ctx, tx, err)
}
