// Code generated by MockGen. DO NOT EDIT.
// Source: dependencies.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	"github.com/kudrykv/latex-yearly-planner/internal/core/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSection is a mock of Section interface.
type MockSection struct {
	ctrl     *gomock.Controller
	recorder *MockSectionMockRecorder
}

// MockSectionMockRecorder is the mock recorder for MockSection.
type MockSectionMockRecorder struct {
	mock *MockSection
}

// NewMockSection creates a new mock instance.
func NewMockSection(ctrl *gomock.Controller) *MockSection {
	mock := &MockSection{ctrl: ctrl}
	mock.recorder = &MockSectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSection) EXPECT() *MockSectionMockRecorder {
	return m.recorder
}

// GenerateSection mocks base method.
func (m *MockSection) GenerateSection(arg0 context.Context) (entities.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSection", arg0)
	ret0, _ := ret[0].(entities.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSection indicates an expected call of GenerateSection.
func (mr *MockSectionMockRecorder) GenerateSection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSection", reflect.TypeOf((*MockSection)(nil).GenerateSection), arg0)
}

// IsEnabled mocks base method.
func (m *MockSection) IsEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEnabled indicates an expected call of IsEnabled.
func (mr *MockSectionMockRecorder) IsEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabled", reflect.TypeOf((*MockSection)(nil).IsEnabled))
}

// MockIndexer is a mock of Indexer interface.
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer.
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance.
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// CreateIndex mocks base method.
func (m *MockIndexer) CreateIndex(arg0 context.Context, arg1 entities.Notes) (entities.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndex", arg0, arg1)
	ret0, _ := ret[0].(entities.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIndex indicates an expected call of CreateIndex.
func (mr *MockIndexerMockRecorder) CreateIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndex", reflect.TypeOf((*MockIndexer)(nil).CreateIndex), arg0, arg1)
}
