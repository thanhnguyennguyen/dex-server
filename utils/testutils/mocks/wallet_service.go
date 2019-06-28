// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import common "github.com/ethereum/go-ethereum/common"
import mock "github.com/stretchr/testify/mock"

import types "github.com/tomochain/tomox-sdk/types"

// WalletService is an autogenerated mock type for the WalletService type
type WalletService struct {
	mock.Mock
}

// CreateAdminWallet provides a mock function with given fields: a
func (_m *WalletService) CreateAdminWallet(a common.Address) (*types.Wallet, error) {
	ret := _m.Called(a)

	var r0 *types.Wallet
	if rf, ok := ret.Get(0).(func(common.Address) *types.Wallet); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *WalletService) GetAll() ([]types.Wallet, error) {
	ret := _m.Called()

	var r0 []types.Wallet
	if rf, ok := ret.Get(0).(func() []types.Wallet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAddress provides a mock function with given fields: a
func (_m *WalletService) GetByAddress(a common.Address) (*types.Wallet, error) {
	ret := _m.Called(a)

	var r0 *types.Wallet
	if rf, ok := ret.Get(0).(func(common.Address) *types.Wallet); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultAdminWallet provides a mock function with given fields:
func (_m *WalletService) GetDefaultAdminWallet() (*types.Wallet, error) {
	ret := _m.Called()

	var r0 *types.Wallet
	if rf, ok := ret.Get(0).(func() *types.Wallet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperatorWallets provides a mock function with given fields:
func (_m *WalletService) GetOperatorWallets() ([]*types.Wallet, error) {
	ret := _m.Called()

	var r0 []*types.Wallet
	if rf, ok := ret.Get(0).(func() []*types.Wallet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
