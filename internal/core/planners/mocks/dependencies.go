// Code generated by MockGen. DO NOT EDIT.
// Source: dependencies.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	planners "github.com/kudrykv/latex-yearly-planner/internal/core/planners"
	entities "github.com/kudrykv/latex-yearly-planner/internal/core/planners/entities"
)

// MockPlannerBuilder is a mock of PlannerBuilder interface.
type MockPlannerBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockPlannerBuilderMockRecorder
}

// MockPlannerBuilderMockRecorder is the mock recorder for MockPlannerBuilder.
type MockPlannerBuilderMockRecorder struct {
	mock *MockPlannerBuilder
}

// NewMockPlannerBuilder creates a new mock instance.
func NewMockPlannerBuilder(ctrl *gomock.Controller) *MockPlannerBuilder {
	mock := &MockPlannerBuilder{ctrl: ctrl}
	mock.recorder = &MockPlannerBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlannerBuilder) EXPECT() *MockPlannerBuilderMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockPlannerBuilder) Generate(arg0 context.Context) (entities.NoteStructure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0)
	ret0, _ := ret[0].(entities.NoteStructure)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockPlannerBuilderMockRecorder) Generate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockPlannerBuilder)(nil).Generate), arg0)
}

// MockNoteWriter is a mock of NoteWriter interface.
type MockNoteWriter struct {
	ctrl     *gomock.Controller
	recorder *MockNoteWriterMockRecorder
}

// MockNoteWriterMockRecorder is the mock recorder for MockNoteWriter.
type MockNoteWriterMockRecorder struct {
	mock *MockNoteWriter
}

// NewMockNoteWriter creates a new mock instance.
func NewMockNoteWriter(ctrl *gomock.Controller) *MockNoteWriter {
	mock := &MockNoteWriter{ctrl: ctrl}
	mock.recorder = &MockNoteWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNoteWriter) EXPECT() *MockNoteWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockNoteWriter) Write(arg0 context.Context, arg1 entities.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockNoteWriterMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockNoteWriter)(nil).Write), arg0, arg1)
}

// MockCommander is a mock of Commander interface.
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderMockRecorder
}

// MockCommanderMockRecorder is the mock recorder for MockCommander.
type MockCommanderMockRecorder struct {
	mock *MockCommander
}

// NewMockCommander creates a new mock instance.
func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &MockCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommander) EXPECT() *MockCommanderMockRecorder {
	return m.recorder
}

// CreateCommand mocks base method.
func (m *MockCommander) CreateCommand(arg0 planners.CommandName, arg1 ...planners.StringArg) planners.Command {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCommand", varargs...)
	ret0, _ := ret[0].(planners.Command)
	return ret0
}

// CreateCommand indicates an expected call of CreateCommand.
func (mr *MockCommanderMockRecorder) CreateCommand(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommand", reflect.TypeOf((*MockCommander)(nil).CreateCommand), varargs...)
}

// MockCommand is a mock of Command interface.
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand.
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance.
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockCommand) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockCommandMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCommand)(nil).Run), arg0)
}

// SetBasePath mocks base method.
func (m *MockCommand) SetBasePath(path planners.BasePath) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBasePath", path)
}

// SetBasePath indicates an expected call of SetBasePath.
func (mr *MockCommandMockRecorder) SetBasePath(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBasePath", reflect.TypeOf((*MockCommand)(nil).SetBasePath), path)
}
