// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/atlaskerr/titan/http/ready (interfaces: Readiness)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Readiness is a mock of Readiness interface
type Readiness struct {
	ctrl     *gomock.Controller
	recorder *ReadinessMockRecorder
}

// ReadinessMockRecorder is the mock recorder for Readiness
type ReadinessMockRecorder struct {
	mock *Readiness
}

// NewReadiness creates a new mock instance
func NewReadiness(ctrl *gomock.Controller) *Readiness {
	mock := &Readiness{ctrl: ctrl}
	mock.recorder = &ReadinessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Readiness) EXPECT() *ReadinessMockRecorder {
	return m.recorder
}

// Ready mocks base method
func (m *Readiness) Ready() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ready indicates an expected call of Ready
func (mr *ReadinessMockRecorder) Ready() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*Readiness)(nil).Ready))
}
