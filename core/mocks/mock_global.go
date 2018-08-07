// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/Go-IOS-Protocol/core/global (interfaces: Global)

// Package core_mock is a generated GoMock package.
package core_mock

import (
	gomock "github.com/golang/mock/gomock"
	common "github.com/iost-official/Go-IOS-Protocol/common"
	block "github.com/iost-official/Go-IOS-Protocol/core/block"
	global "github.com/iost-official/Go-IOS-Protocol/core/global"
	state "github.com/iost-official/Go-IOS-Protocol/core/state"
	tx "github.com/iost-official/Go-IOS-Protocol/core/tx"
	reflect "reflect"
)

// MockGlobal is a mock of Global interface
type MockGlobal struct {
	ctrl     *gomock.Controller
	recorder *MockGlobalMockRecorder
}

// MockGlobalMockRecorder is the mock recorder for MockGlobal
type MockGlobalMockRecorder struct {
	mock *MockGlobal
}

// NewMockGlobal creates a new mock instance
func NewMockGlobal(ctrl *gomock.Controller) *MockGlobal {
	mock := &MockGlobal{ctrl: ctrl}
	mock.recorder = &MockGlobalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGlobal) EXPECT() *MockGlobalMockRecorder {
	return m.recorder
}

// BlockChain mocks base method
func (m *MockGlobal) BlockChain() block.Chain {
	ret := m.ctrl.Call(m, "BlockChain")
	ret0, _ := ret[0].(block.Chain)
	return ret0
}

// BlockChain indicates an expected call of BlockChain
func (mr *MockGlobalMockRecorder) BlockChain() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockChain", reflect.TypeOf((*MockGlobal)(nil).BlockChain))
}

// Config mocks base method
func (m *MockGlobal) Config() *common.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*common.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockGlobalMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockGlobal)(nil).Config))
}

// Mode mocks base method
func (m *MockGlobal) Mode() global.Mode {
	ret := m.ctrl.Call(m, "Mode")
	ret0, _ := ret[0].(global.Mode)
	return ret0
}

// Mode indicates an expected call of Mode
func (mr *MockGlobalMockRecorder) Mode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mode", reflect.TypeOf((*MockGlobal)(nil).Mode))
}

// SetMode mocks base method
func (m *MockGlobal) SetMode(arg0 global.Mode) bool {
	ret := m.ctrl.Call(m, "SetMode", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetMode indicates an expected call of SetMode
func (mr *MockGlobalMockRecorder) SetMode(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMode", reflect.TypeOf((*MockGlobal)(nil).SetMode), arg0)
}

// StdPool mocks base method
func (m *MockGlobal) StdPool() state.Pool {
	ret := m.ctrl.Call(m, "StdPool")
	ret0, _ := ret[0].(state.Pool)
	return ret0
}

// StdPool indicates an expected call of StdPool
func (mr *MockGlobalMockRecorder) StdPool() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdPool", reflect.TypeOf((*MockGlobal)(nil).StdPool))
}

// TxDB mocks base method
func (m *MockGlobal) TxDB() tx.TxPool {
	ret := m.ctrl.Call(m, "TxDB")
	ret0, _ := ret[0].(tx.TxPool)
	return ret0
}

// TxDB indicates an expected call of TxDB
func (mr *MockGlobalMockRecorder) TxDB() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxDB", reflect.TypeOf((*MockGlobal)(nil).TxDB))
}