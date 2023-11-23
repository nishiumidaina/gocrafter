// Code generated by MockGen. DO NOT EDIT.
// Source: ./loginReward_service.go

// Package loginReward is a generated GoMock package.
package loginReward

import (
	reflect "reflect"
	time "time"

	loginReward "github.com/game-core/gocrafter/api/presentation/request/loginReward"
	loginReward0 "github.com/game-core/gocrafter/api/presentation/response/loginReward"
	gomock "github.com/golang/mock/gomock"
)

// MockLoginRewardService is a mock of LoginRewardService interface.
type MockLoginRewardService struct {
	ctrl     *gomock.Controller
	recorder *MockLoginRewardServiceMockRecorder
}

// MockLoginRewardServiceMockRecorder is the mock recorder for MockLoginRewardService.
type MockLoginRewardServiceMockRecorder struct {
	mock *MockLoginRewardService
}

// NewMockLoginRewardService creates a new mock instance.
func NewMockLoginRewardService(ctrl *gomock.Controller) *MockLoginRewardService {
	mock := &MockLoginRewardService{ctrl: ctrl}
	mock.recorder = &MockLoginRewardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginRewardService) EXPECT() *MockLoginRewardServiceMockRecorder {
	return m.recorder
}

// GetLoginRewardModel mocks base method.
func (m *MockLoginRewardService) GetLoginRewardModel(req *loginReward.GetLoginRewardModel, now time.Time) (*loginReward0.GetLoginRewardModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoginRewardModel", req, now)
	ret0, _ := ret[0].(*loginReward0.GetLoginRewardModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoginRewardModel indicates an expected call of GetLoginRewardModel.
func (mr *MockLoginRewardServiceMockRecorder) GetLoginRewardModel(req, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoginRewardModel", reflect.TypeOf((*MockLoginRewardService)(nil).GetLoginRewardModel), req, now)
}

// ReceiveLoginReward mocks base method.
func (m *MockLoginRewardService) ReceiveLoginReward(req *loginReward.ReceiveLoginReward, now time.Time) (*loginReward0.ReceiveLoginReward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveLoginReward", req, now)
	ret0, _ := ret[0].(*loginReward0.ReceiveLoginReward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveLoginReward indicates an expected call of ReceiveLoginReward.
func (mr *MockLoginRewardServiceMockRecorder) ReceiveLoginReward(req, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveLoginReward", reflect.TypeOf((*MockLoginRewardService)(nil).ReceiveLoginReward), req, now)
}
