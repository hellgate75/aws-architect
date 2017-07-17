package action

import (
	"github.com/golang/mock/gomock"
	"fmt"
	"strconv"
)

// MockAction is a mock of Action interface
type MockAction struct {
	ctrl     *gomock.Controller
	recorder *MockActionMockRecorder
}

// MockActionMockRecorder is the mock recorder for MockAction
type MockActionMockRecorder struct {
	mock *MockAction
}

// NewMockAction creates a new mock instance
func NewMockAction(ctrl *gomock.Controller) *MockAction {
	mock := &MockAction{ctrl: ctrl}
	mock.recorder = &MockActionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockAction) EXPECT() *MockActionMockRecorder {
	return _m.recorder
}

// Init mocks base method
func (_m *MockAction) Init() (bool) {
	values := _m.ctrl.Call(_m, "Init")
	return  values[0] == true
}

// Init indicates an expected call of Init
func (_mr *MockActionMockRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init")
}

// Reset mocks base method
func (_m *MockAction) Reset() (bool) {
	values := _m.ctrl.Call(_m, "Reset")
	return  values[0] == true
}

// Reset indicates an expected call of Reset
func (_mr *MockActionMockRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

// Execute mocks base method
func (_m *MockAction) Execute(_param0 chan string) (bool) {
	values := _m.ctrl.Call(_m, "Execute", _param0)
	return  values[0] == true
}

// Execute indicates an expected call of Execute
func (_mr *MockActionMockRecorder) Execute(_param0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Execute", _param0)
}

// IsInProgress mocks base method
func (_m *MockAction) IsInProgress() (bool) {
	values := _m.ctrl.Call(_m, "IsInProgress")
	return  values[0] == true
}

// IsInProgress indicates an expected call of IsInProgress
func (_mr *MockActionMockRecorder) IsInProgress() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsInProgress")
}

// GetCommand mocks base method
func (_m *MockAction) GetCommand() (string) {
	values := _m.ctrl.Call(_m, "GetCommand")
	return  fmt.Sprintf("%s",values[0])
}

// GetCommand indicates an expected call of GetCommand
func (_mr *MockActionMockRecorder) GetCommand() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCommand")
}

// GetName mocks base method
func (_m *MockAction) GetName() (string) {
	values := _m.ctrl.Call(_m, "GetName")
	return  fmt.Sprintf("%s",values[0])
}

// GetName indicates an expected call of GetName
func (_mr *MockActionMockRecorder) GetName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetName")
}

// GetUsage mocks base method
func (_m *MockAction) GetUsage() (string) {
	values := _m.ctrl.Call(_m, "GetUsage")
	return  fmt.Sprintf("%s",values[0])
}

// GetName indicates an expected call of GetUsage
func (_mr *MockActionMockRecorder) GetUsage() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUsage")
}

// AcquireValues mocks base method
func (_m *MockAction) AcquireValues() (bool) {
	values := _m.ctrl.Call(_m, "AcquireValues")
	return  values[0] == true
}

// AcquireValues indicates an expected call of AcquireValues
func (_mr *MockActionMockRecorder) AcquireValues() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcquireValues")
}

// GetExitCode mocks base method
func (_m *MockAction) GetExitCode() (int) {
	values := _m.ctrl.Call(_m, "GetExitCode")
	val,_  := strconv.Atoi(fmt.Sprintf("%v",values[0]))
	return val
}

// GetExitCode indicates an expected call of GetExitCode
func (_mr *MockActionMockRecorder) GetExitCode() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetExitCode")
}

// GetLastMessage mocks base method
func (_m *MockAction) GetLastMessage() (string) {
	values := _m.ctrl.Call(_m, "GetLastMessage")
	return  fmt.Sprintf("%s",values[0])
}

// GetLastMessage indicates an expected call of GetLastMessage
func (_mr *MockActionMockRecorder) GetLastMessage() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLastMessage")
}
