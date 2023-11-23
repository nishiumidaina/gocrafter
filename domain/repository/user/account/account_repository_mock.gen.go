// Code generated by MockGen. DO NOT EDIT.
// Source: ./account_repository.gen.go

// Package account is a generated GoMock package.
package account

import (
	reflect "reflect"

	account "github.com/game-core/gocrafter/domain/entity/user/account"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountRepository) Create(entity *account.Account, shardKey int, tx *gorm.DB) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", entity, shardKey, tx)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountRepositoryMockRecorder) Create(entity, shardKey, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepository)(nil).Create), entity, shardKey, tx)
}

// Delete mocks base method.
func (m *MockAccountRepository) Delete(entity *account.Account, shardKey int, tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", entity, shardKey, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccountRepositoryMockRecorder) Delete(entity, shardKey, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccountRepository)(nil).Delete), entity, shardKey, tx)
}

// FindByID mocks base method.
func (m *MockAccountRepository) FindByID(ID int64, shardKey int) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ID, shardKey)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockAccountRepositoryMockRecorder) FindByID(ID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockAccountRepository)(nil).FindByID), ID, shardKey)
}

// FindOrNilByID mocks base method.
func (m *MockAccountRepository) FindOrNilByID(ID int64, shardKey int) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByID", ID, shardKey)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByID indicates an expected call of FindOrNilByID.
func (mr *MockAccountRepositoryMockRecorder) FindOrNilByID(ID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByID", reflect.TypeOf((*MockAccountRepository)(nil).FindOrNilByID), ID, shardKey)
}

// FindOrNilByIDAndUUID mocks base method.
func (m *MockAccountRepository) FindOrNilByIDAndUUID(ID int64, UUID string, shardKey int) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByIDAndUUID", ID, UUID, shardKey)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByIDAndUUID indicates an expected call of FindOrNilByIDAndUUID.
func (mr *MockAccountRepositoryMockRecorder) FindOrNilByIDAndUUID(ID, UUID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByIDAndUUID", reflect.TypeOf((*MockAccountRepository)(nil).FindOrNilByIDAndUUID), ID, UUID, shardKey)
}

// FindOrNilByUUID mocks base method.
func (m *MockAccountRepository) FindOrNilByUUID(UUID string, shardKey int) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByUUID", UUID, shardKey)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByUUID indicates an expected call of FindOrNilByUUID.
func (mr *MockAccountRepositoryMockRecorder) FindOrNilByUUID(UUID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByUUID", reflect.TypeOf((*MockAccountRepository)(nil).FindOrNilByUUID), UUID, shardKey)
}

// List mocks base method.
func (m *MockAccountRepository) List(limit int64, shardKey int) (*account.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", limit, shardKey)
	ret0, _ := ret[0].(*account.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAccountRepositoryMockRecorder) List(limit, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAccountRepository)(nil).List), limit, shardKey)
}

// ListByIDAndUUID mocks base method.
func (m *MockAccountRepository) ListByIDAndUUID(ID int64, UUID string, shardKey int) (*account.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIDAndUUID", ID, UUID, shardKey)
	ret0, _ := ret[0].(*account.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIDAndUUID indicates an expected call of ListByIDAndUUID.
func (mr *MockAccountRepositoryMockRecorder) ListByIDAndUUID(ID, UUID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIDAndUUID", reflect.TypeOf((*MockAccountRepository)(nil).ListByIDAndUUID), ID, UUID, shardKey)
}

// ListByUUID mocks base method.
func (m *MockAccountRepository) ListByUUID(UUID string, shardKey int) (*account.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUUID", UUID, shardKey)
	ret0, _ := ret[0].(*account.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByUUID indicates an expected call of ListByUUID.
func (mr *MockAccountRepositoryMockRecorder) ListByUUID(UUID, shardKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUUID", reflect.TypeOf((*MockAccountRepository)(nil).ListByUUID), UUID, shardKey)
}

// Update mocks base method.
func (m *MockAccountRepository) Update(entity *account.Account, shardKey int, tx *gorm.DB) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", entity, shardKey, tx)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAccountRepositoryMockRecorder) Update(entity, shardKey, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountRepository)(nil).Update), entity, shardKey, tx)
}
