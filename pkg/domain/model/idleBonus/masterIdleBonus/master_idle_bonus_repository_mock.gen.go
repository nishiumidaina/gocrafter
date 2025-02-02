// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_idle_bonus_repository.gen.go

// Package masterIdleBonus is a generated GoMock package.
package masterIdleBonus

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterIdleBonusRepository is a mock of MasterIdleBonusRepository interface.
type MockMasterIdleBonusRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterIdleBonusRepositoryMockRecorder
}

// MockMasterIdleBonusRepositoryMockRecorder is the mock recorder for MockMasterIdleBonusRepository.
type MockMasterIdleBonusRepositoryMockRecorder struct {
	mock *MockMasterIdleBonusRepository
}

// NewMockMasterIdleBonusRepository creates a new mock instance.
func NewMockMasterIdleBonusRepository(ctrl *gomock.Controller) *MockMasterIdleBonusRepository {
	mock := &MockMasterIdleBonusRepository{ctrl: ctrl}
	mock.recorder = &MockMasterIdleBonusRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterIdleBonusRepository) EXPECT() *MockMasterIdleBonusRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterIdleBonusRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterIdleBonus) (*MasterIdleBonus, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterIdleBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterIdleBonusRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterIdleBonuses) (MasterIdleBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterIdleBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterIdleBonusRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterIdleBonus) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterIdleBonusRepository) Find(ctx context.Context, id int64) (*MasterIdleBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterIdleBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).Find), ctx, id)
}

// FindByMasterIdleBonusEventId mocks base method.
func (m *MockMasterIdleBonusRepository) FindByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (*MasterIdleBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterIdleBonusEventId", ctx, masterIdleBonusEventId)
	ret0, _ := ret[0].(*MasterIdleBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterIdleBonusEventId indicates an expected call of FindByMasterIdleBonusEventId.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) FindByMasterIdleBonusEventId(ctx, masterIdleBonusEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterIdleBonusEventId", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).FindByMasterIdleBonusEventId), ctx, masterIdleBonusEventId)
}

// FindList mocks base method.
func (m *MockMasterIdleBonusRepository) FindList(ctx context.Context) (MasterIdleBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterIdleBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).FindList), ctx)
}

// FindListByMasterIdleBonusEventId mocks base method.
func (m *MockMasterIdleBonusRepository) FindListByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (MasterIdleBonuses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterIdleBonusEventId", ctx, masterIdleBonusEventId)
	ret0, _ := ret[0].(MasterIdleBonuses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterIdleBonusEventId indicates an expected call of FindListByMasterIdleBonusEventId.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) FindListByMasterIdleBonusEventId(ctx, masterIdleBonusEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterIdleBonusEventId", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).FindListByMasterIdleBonusEventId), ctx, masterIdleBonusEventId)
}

// FindOrNil mocks base method.
func (m *MockMasterIdleBonusRepository) FindOrNil(ctx context.Context, id int64) (*MasterIdleBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterIdleBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).FindOrNil), ctx, id)
}

// FindOrNilByMasterIdleBonusEventId mocks base method.
func (m *MockMasterIdleBonusRepository) FindOrNilByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (*MasterIdleBonus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterIdleBonusEventId", ctx, masterIdleBonusEventId)
	ret0, _ := ret[0].(*MasterIdleBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterIdleBonusEventId indicates an expected call of FindOrNilByMasterIdleBonusEventId.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) FindOrNilByMasterIdleBonusEventId(ctx, masterIdleBonusEventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterIdleBonusEventId", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).FindOrNilByMasterIdleBonusEventId), ctx, masterIdleBonusEventId)
}

// Update mocks base method.
func (m_2 *MockMasterIdleBonusRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterIdleBonus) (*MasterIdleBonus, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterIdleBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterIdleBonusRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterIdleBonusRepository)(nil).Update), ctx, tx, m)
}
