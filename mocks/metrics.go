// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: ./metrics/metrics.go

// Package mock_metrics is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	"github.com/uber-go/dosa/metrics"
)

// MockScope is a mock of Scope interface
type MockScope struct {
	ctrl     *gomock.Controller
	recorder *MockScopeMockRecorder
}

// MockScopeMockRecorder is the mock recorder for MockScope
type MockScopeMockRecorder struct {
	mock *MockScope
}

// NewMockScope creates a new mock instance
func NewMockScope(ctrl *gomock.Controller) *MockScope {
	mock := &MockScope{ctrl: ctrl}
	mock.recorder = &MockScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScope) EXPECT() *MockScopeMockRecorder {
	return m.recorder
}

// Counter mocks base method
func (m *MockScope) Counter(name string) metrics.Counter {
	ret := m.ctrl.Call(m, "Counter", name)
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// Counter indicates an expected call of Counter
func (mr *MockScopeMockRecorder) Counter(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Counter", reflect.TypeOf((*MockScope)(nil).Counter), name)
}

// Tagged mocks base method
func (m *MockScope) Tagged(tags map[string]string) metrics.Scope {
	ret := m.ctrl.Call(m, "Tagged", tags)
	ret0, _ := ret[0].(metrics.Scope)
	return ret0
}

// Tagged indicates an expected call of Tagged
func (mr *MockScopeMockRecorder) Tagged(tags interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tagged", reflect.TypeOf((*MockScope)(nil).Tagged), tags)
}

// SubScope mocks base method
func (m *MockScope) SubScope(name string) metrics.Scope {
	ret := m.ctrl.Call(m, "SubScope", name)
	ret0, _ := ret[0].(metrics.Scope)
	return ret0
}

// SubScope indicates an expected call of SubScope
func (mr *MockScopeMockRecorder) SubScope(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubScope", reflect.TypeOf((*MockScope)(nil).SubScope), name)
}

// Timer mocks base method
func (m *MockScope) Timer(name string) metrics.Timer {
	ret := m.ctrl.Call(m, "Timer", name)
	ret0, _ := ret[0].(metrics.Timer)
	return ret0
}

// Timer indicates an expected call of Timer
func (mr *MockScopeMockRecorder) Timer(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timer", reflect.TypeOf((*MockScope)(nil).Timer), name)
}

// MockCounter is a mock of Counter interface
type MockCounter struct {
	ctrl     *gomock.Controller
	recorder *MockCounterMockRecorder
}

// MockCounterMockRecorder is the mock recorder for MockCounter
type MockCounterMockRecorder struct {
	mock *MockCounter
}

// NewMockCounter creates a new mock instance
func NewMockCounter(ctrl *gomock.Controller) *MockCounter {
	mock := &MockCounter{ctrl: ctrl}
	mock.recorder = &MockCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCounter) EXPECT() *MockCounterMockRecorder {
	return m.recorder
}

// Inc mocks base method
func (m *MockCounter) Inc(delta int64) {
	m.ctrl.Call(m, "Inc", delta)
}

// Inc indicates an expected call of Inc
func (mr *MockCounterMockRecorder) Inc(delta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inc", reflect.TypeOf((*MockCounter)(nil).Inc), delta)
}

// MockTimer is a mock of Timer interface
type MockTimer struct {
	ctrl     *gomock.Controller
	recorder *MockTimerMockRecorder
}

// MockTimerMockRecorder is the mock recorder for MockTimer
type MockTimerMockRecorder struct {
	mock *MockTimer
}

// NewMockTimer creates a new mock instance
func NewMockTimer(ctrl *gomock.Controller) *MockTimer {
	mock := &MockTimer{ctrl: ctrl}
	mock.recorder = &MockTimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimer) EXPECT() *MockTimerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockTimer) Start() time.Time {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockTimerMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTimer)(nil).Start))
}

// Stop mocks base method
func (m *MockTimer) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockTimerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTimer)(nil).Stop))
}