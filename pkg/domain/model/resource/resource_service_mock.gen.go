// Code generated by MockGen. DO NOT EDIT.
// Source: ./resource_service.go

// Package resource is a generated GoMock package.
package resource

import (
	gomock "github.com/golang/mock/gomock"
)

// MockResourceService is a mock of ResourceService interface.
type MockResourceService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceServiceMockRecorder
}

// MockResourceServiceMockRecorder is the mock recorder for MockResourceService.
type MockResourceServiceMockRecorder struct {
	mock *MockResourceService
}

// NewMockResourceService creates a new mock instance.
func NewMockResourceService(ctrl *gomock.Controller) *MockResourceService {
	mock := &MockResourceService{ctrl: ctrl}
	mock.recorder = &MockResourceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceService) EXPECT() *MockResourceServiceMockRecorder {
	return m.recorder
}
