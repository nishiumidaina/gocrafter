// Code generated by MockGen. DO NOT EDIT.
// Source: ./loginRewardReward_repository.gen.go

// Package loginReward is a generated GoMock package.
package loginReward

import (
	reflect "reflect"

	loginReward "github.com/game-core/gocrafter/domain/entity/master/loginReward"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockLoginRewardRewardRepository is a mock of LoginRewardRewardRepository interface.
type MockLoginRewardRewardRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLoginRewardRewardRepositoryMockRecorder
}

// MockLoginRewardRewardRepositoryMockRecorder is the mock recorder for MockLoginRewardRewardRepository.
type MockLoginRewardRewardRepositoryMockRecorder struct {
	mock *MockLoginRewardRewardRepository
}

// NewMockLoginRewardRewardRepository creates a new mock instance.
func NewMockLoginRewardRewardRepository(ctrl *gomock.Controller) *MockLoginRewardRewardRepository {
	mock := &MockLoginRewardRewardRepository{ctrl: ctrl}
	mock.recorder = &MockLoginRewardRewardRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginRewardRewardRepository) EXPECT() *MockLoginRewardRewardRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLoginRewardRewardRepository) Create(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", entity, tx)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) Create(entity, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).Create), entity, tx)
}

// Delete mocks base method.
func (m *MockLoginRewardRewardRepository) Delete(entity *loginReward.LoginRewardReward, tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", entity, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) Delete(entity, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).Delete), entity, tx)
}

// FindByID mocks base method.
func (m *MockLoginRewardRewardRepository) FindByID(ID int64) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ID)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindByID), ID)
}

// FindByItemID mocks base method.
func (m *MockLoginRewardRewardRepository) FindByItemID(ItemID int64) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByItemID", ItemID)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByItemID indicates an expected call of FindByItemID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindByItemID(ItemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByItemID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindByItemID), ItemID)
}

// FindByLoginRewardID mocks base method.
func (m *MockLoginRewardRewardRepository) FindByLoginRewardID(LoginRewardID int64) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLoginRewardID", LoginRewardID)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLoginRewardID indicates an expected call of FindByLoginRewardID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindByLoginRewardID(LoginRewardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLoginRewardID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindByLoginRewardID), LoginRewardID)
}

// FindByLoginRewardIDAndItemID mocks base method.
func (m *MockLoginRewardRewardRepository) FindByLoginRewardIDAndItemID(LoginRewardID, ItemID int64) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLoginRewardIDAndItemID", LoginRewardID, ItemID)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLoginRewardIDAndItemID indicates an expected call of FindByLoginRewardIDAndItemID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindByLoginRewardIDAndItemID(LoginRewardID, ItemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLoginRewardIDAndItemID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindByLoginRewardIDAndItemID), LoginRewardID, ItemID)
}

// FindByName mocks base method.
func (m *MockLoginRewardRewardRepository) FindByName(Name string) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", Name)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindByName(Name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindByName), Name)
}

// FindOrNilByID mocks base method.
func (m *MockLoginRewardRewardRepository) FindOrNilByID(ID int64) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByID", ID)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByID indicates an expected call of FindOrNilByID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindOrNilByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindOrNilByID), ID)
}

// FindOrNilByItemID mocks base method.
func (m *MockLoginRewardRewardRepository) FindOrNilByItemID(ItemID int64) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByItemID", ItemID)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByItemID indicates an expected call of FindOrNilByItemID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindOrNilByItemID(ItemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByItemID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindOrNilByItemID), ItemID)
}

// FindOrNilByLoginRewardID mocks base method.
func (m *MockLoginRewardRewardRepository) FindOrNilByLoginRewardID(LoginRewardID int64) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByLoginRewardID", LoginRewardID)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByLoginRewardID indicates an expected call of FindOrNilByLoginRewardID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindOrNilByLoginRewardID(LoginRewardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByLoginRewardID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindOrNilByLoginRewardID), LoginRewardID)
}

// FindOrNilByLoginRewardIDAndItemID mocks base method.
func (m *MockLoginRewardRewardRepository) FindOrNilByLoginRewardIDAndItemID(LoginRewardID, ItemID int64) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByLoginRewardIDAndItemID", LoginRewardID, ItemID)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByLoginRewardIDAndItemID indicates an expected call of FindOrNilByLoginRewardIDAndItemID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindOrNilByLoginRewardIDAndItemID(LoginRewardID, ItemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByLoginRewardIDAndItemID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindOrNilByLoginRewardIDAndItemID), LoginRewardID, ItemID)
}

// FindOrNilByName mocks base method.
func (m *MockLoginRewardRewardRepository) FindOrNilByName(Name string) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByName", Name)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByName indicates an expected call of FindOrNilByName.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) FindOrNilByName(Name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByName", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).FindOrNilByName), Name)
}

// List mocks base method.
func (m *MockLoginRewardRewardRepository) List(limit int64) (*loginReward.LoginRewardRewards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", limit)
	ret0, _ := ret[0].(*loginReward.LoginRewardRewards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) List(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).List), limit)
}

// ListByItemID mocks base method.
func (m *MockLoginRewardRewardRepository) ListByItemID(ItemID int64) (*loginReward.LoginRewardRewards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByItemID", ItemID)
	ret0, _ := ret[0].(*loginReward.LoginRewardRewards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByItemID indicates an expected call of ListByItemID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) ListByItemID(ItemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByItemID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).ListByItemID), ItemID)
}

// ListByLoginRewardID mocks base method.
func (m *MockLoginRewardRewardRepository) ListByLoginRewardID(LoginRewardID int64) (*loginReward.LoginRewardRewards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByLoginRewardID", LoginRewardID)
	ret0, _ := ret[0].(*loginReward.LoginRewardRewards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByLoginRewardID indicates an expected call of ListByLoginRewardID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) ListByLoginRewardID(LoginRewardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByLoginRewardID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).ListByLoginRewardID), LoginRewardID)
}

// ListByLoginRewardIDAndItemID mocks base method.
func (m *MockLoginRewardRewardRepository) ListByLoginRewardIDAndItemID(LoginRewardID, ItemID int64) (*loginReward.LoginRewardRewards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByLoginRewardIDAndItemID", LoginRewardID, ItemID)
	ret0, _ := ret[0].(*loginReward.LoginRewardRewards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByLoginRewardIDAndItemID indicates an expected call of ListByLoginRewardIDAndItemID.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) ListByLoginRewardIDAndItemID(LoginRewardID, ItemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByLoginRewardIDAndItemID", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).ListByLoginRewardIDAndItemID), LoginRewardID, ItemID)
}

// ListByName mocks base method.
func (m *MockLoginRewardRewardRepository) ListByName(Name string) (*loginReward.LoginRewardRewards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByName", Name)
	ret0, _ := ret[0].(*loginReward.LoginRewardRewards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByName indicates an expected call of ListByName.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) ListByName(Name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByName", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).ListByName), Name)
}

// Update mocks base method.
func (m *MockLoginRewardRewardRepository) Update(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", entity, tx)
	ret0, _ := ret[0].(*loginReward.LoginRewardReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockLoginRewardRewardRepositoryMockRecorder) Update(entity, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLoginRewardRewardRepository)(nil).Update), entity, tx)
}
