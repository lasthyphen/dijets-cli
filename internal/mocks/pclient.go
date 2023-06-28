// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/lasthyphen/dijetsnode/api"

	crypto "github.com/lasthyphen/dijetsnode/utils/crypto"

	ids "github.com/lasthyphen/dijetsnode/ids"

	mock "github.com/stretchr/testify/mock"

	platformvm "github.com/lasthyphen/dijetsnode/vms/platformvm"

	rpc "github.com/lasthyphen/dijetsnode/utils/rpc"

	status "github.com/lasthyphen/dijetsnode/vms/platformvm/status"

	time "time"
)

// PClient is an autogenerated mock type for the Client type
type PClient struct {
	mock.Mock
}

// AddDelegator provides a mock function with given fields: ctx, user, from, changeAddr, rewardAddress, nodeID, stakeAmount, startTime, endTime, options
func (_m *PClient) AddDelegator(ctx context.Context, user api.UserPass, from []ids.ShortID, changeAddr ids.ShortID, rewardAddress ids.ShortID, nodeID ids.NodeID, stakeAmount uint64, startTime uint64, endTime uint64, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, from, changeAddr, rewardAddress, nodeID, stakeAmount, startTime, endTime)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ShortID, ids.NodeID, uint64, uint64, uint64, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, user, from, changeAddr, rewardAddress, nodeID, stakeAmount, startTime, endTime, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ShortID, ids.NodeID, uint64, uint64, uint64, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, from, changeAddr, rewardAddress, nodeID, stakeAmount, startTime, endTime, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddSubnetValidator provides a mock function with given fields: ctx, user, from, changeAddr, subnetID, nodeID, stakeAmount, startTime, endTime, options
func (_m *PClient) AddSubnetValidator(ctx context.Context, user api.UserPass, from []ids.ShortID, changeAddr ids.ShortID, subnetID ids.ID, nodeID ids.NodeID, stakeAmount uint64, startTime uint64, endTime uint64, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, from, changeAddr, subnetID, nodeID, stakeAmount, startTime, endTime)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ID, ids.NodeID, uint64, uint64, uint64, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, user, from, changeAddr, subnetID, nodeID, stakeAmount, startTime, endTime, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ID, ids.NodeID, uint64, uint64, uint64, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, from, changeAddr, subnetID, nodeID, stakeAmount, startTime, endTime, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddValidator provides a mock function with given fields: ctx, user, from, changeAddr, rewardAddress, nodeID, stakeAmount, startTime, endTime, delegationFeeRate, options
func (_m *PClient) AddValidator(ctx context.Context, user api.UserPass, from []ids.ShortID, changeAddr ids.ShortID, rewardAddress ids.ShortID, nodeID ids.NodeID, stakeAmount uint64, startTime uint64, endTime uint64, delegationFeeRate float32, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, from, changeAddr, rewardAddress, nodeID, stakeAmount, startTime, endTime, delegationFeeRate)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ShortID, ids.NodeID, uint64, uint64, uint64, float32, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, user, from, changeAddr, rewardAddress, nodeID, stakeAmount, startTime, endTime, delegationFeeRate, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ShortID, ids.NodeID, uint64, uint64, uint64, float32, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, from, changeAddr, rewardAddress, nodeID, stakeAmount, startTime, endTime, delegationFeeRate, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AwaitTxDecided provides a mock function with given fields: ctx, txID, freq, options
func (_m *PClient) AwaitTxDecided(ctx context.Context, txID ids.ID, freq time.Duration, options ...rpc.Option) (*platformvm.GetTxStatusResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, txID, freq)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *platformvm.GetTxStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, time.Duration, ...rpc.Option) *platformvm.GetTxStatusResponse); ok {
		r0 = rf(ctx, txID, freq, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformvm.GetTxStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, time.Duration, ...rpc.Option) error); ok {
		r1 = rf(ctx, txID, freq, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAddress provides a mock function with given fields: ctx, user, options
func (_m *PClient) CreateAddress(ctx context.Context, user api.UserPass, options ...rpc.Option) (ids.ShortID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ShortID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, ...rpc.Option) ids.ShortID); ok {
		r0 = rf(ctx, user, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ShortID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBlockchain provides a mock function with given fields: ctx, user, from, changeAddr, subnetID, vmID, fxIDs, name, genesisData, options
func (_m *PClient) CreateBlockchain(ctx context.Context, user api.UserPass, from []ids.ShortID, changeAddr ids.ShortID, subnetID ids.ID, vmID string, fxIDs []string, name string, genesisData []byte, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, from, changeAddr, subnetID, vmID, fxIDs, name, genesisData)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ID, string, []string, string, []byte, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, user, from, changeAddr, subnetID, vmID, fxIDs, name, genesisData, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ID, string, []string, string, []byte, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, from, changeAddr, subnetID, vmID, fxIDs, name, genesisData, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSubnet provides a mock function with given fields: ctx, user, from, changeAddr, controlKeys, threshold, options
func (_m *PClient) CreateSubnet(ctx context.Context, user api.UserPass, from []ids.ShortID, changeAddr ids.ShortID, controlKeys []ids.ShortID, threshold uint32, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, from, changeAddr, controlKeys, threshold)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, []ids.ShortID, uint32, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, user, from, changeAddr, controlKeys, threshold, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, []ids.ShortID, uint32, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, from, changeAddr, controlKeys, threshold, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportAVAX provides a mock function with given fields: ctx, user, from, changeAddr, to, toChainIDAlias, amount, options
func (_m *PClient) ExportAVAX(ctx context.Context, user api.UserPass, from []ids.ShortID, changeAddr ids.ShortID, to ids.ShortID, toChainIDAlias string, amount uint64, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, from, changeAddr, to, toChainIDAlias, amount)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ShortID, string, uint64, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, user, from, changeAddr, to, toChainIDAlias, amount, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ShortID, string, uint64, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, from, changeAddr, to, toChainIDAlias, amount, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportKey provides a mock function with given fields: ctx, user, address, options
func (_m *PClient) ExportKey(ctx context.Context, user api.UserPass, address ids.ShortID, options ...rpc.Option) (*crypto.PrivateKeySECP256K1R, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *crypto.PrivateKeySECP256K1R
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, ids.ShortID, ...rpc.Option) *crypto.PrivateKeySECP256K1R); ok {
		r0 = rf(ctx, user, address, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*crypto.PrivateKeySECP256K1R)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, ids.ShortID, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, address, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAtomicUTXOs provides a mock function with given fields: ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options
func (_m *PClient) GetAtomicUTXOs(ctx context.Context, addrs []ids.ShortID, sourceChain string, limit uint32, startAddress ids.ShortID, startUTXOID ids.ID, options ...rpc.Option) ([][]byte, ids.ShortID, ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, addrs, sourceChain, limit, startAddress, startUTXOID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) [][]byte); ok {
		r0 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 ids.ShortID
	if rf, ok := ret.Get(1).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) ids.ShortID); ok {
		r1 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ids.ShortID)
		}
	}

	var r2 ids.ID
	if rf, ok := ret.Get(2).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) ids.ID); ok {
		r2 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(ids.ID)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, []ids.ShortID, string, uint32, ids.ShortID, ids.ID, ...rpc.Option) error); ok {
		r3 = rf(ctx, addrs, sourceChain, limit, startAddress, startUTXOID, options...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetBalance provides a mock function with given fields: ctx, addrs, options
func (_m *PClient) GetBalance(ctx context.Context, addrs []ids.ShortID, options ...rpc.Option) (*platformvm.GetBalanceResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, addrs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *platformvm.GetBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, ...rpc.Option) *platformvm.GetBalanceResponse); ok {
		r0 = rf(ctx, addrs, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformvm.GetBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []ids.ShortID, ...rpc.Option) error); ok {
		r1 = rf(ctx, addrs, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlock provides a mock function with given fields: ctx, blockID, options
func (_m *PClient) GetBlock(ctx context.Context, blockID ids.ID, options ...rpc.Option) ([]byte, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, blockID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) []byte); ok {
		r0 = rf(ctx, blockID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, blockID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockchainStatus provides a mock function with given fields: ctx, blockchainID, options
func (_m *PClient) GetBlockchainStatus(ctx context.Context, blockchainID string, options ...rpc.Option) (status.BlockchainStatus, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, blockchainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 status.BlockchainStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, ...rpc.Option) status.BlockchainStatus); ok {
		r0 = rf(ctx, blockchainID, options...)
	} else {
		r0 = ret.Get(0).(status.BlockchainStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...rpc.Option) error); ok {
		r1 = rf(ctx, blockchainID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockchains provides a mock function with given fields: ctx, options
func (_m *PClient) GetBlockchains(ctx context.Context, options ...rpc.Option) ([]platformvm.APIBlockchain, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []platformvm.APIBlockchain
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) []platformvm.APIBlockchain); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]platformvm.APIBlockchain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentSupply provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetCurrentSupply(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentValidators provides a mock function with given fields: ctx, subnetID, nodeIDs, options
func (_m *PClient) GetCurrentValidators(ctx context.Context, subnetID ids.ID, nodeIDs []ids.NodeID, options ...rpc.Option) ([]platformvm.ClientPermissionlessValidator, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID, nodeIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []platformvm.ClientPermissionlessValidator
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, []ids.NodeID, ...rpc.Option) []platformvm.ClientPermissionlessValidator); ok {
		r0 = rf(ctx, subnetID, nodeIDs, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]platformvm.ClientPermissionlessValidator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, []ids.NodeID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, nodeIDs, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeight provides a mock function with given fields: ctx, options
func (_m *PClient) GetHeight(ctx context.Context, options ...rpc.Option) (uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxStakeAmount provides a mock function with given fields: ctx, subnetID, nodeID, startTime, endTime, options
func (_m *PClient) GetMaxStakeAmount(ctx context.Context, subnetID ids.ID, nodeID ids.NodeID, startTime uint64, endTime uint64, options ...rpc.Option) (uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID, nodeID, startTime, endTime)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ids.NodeID, uint64, uint64, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, subnetID, nodeID, startTime, endTime, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ids.NodeID, uint64, uint64, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, nodeID, startTime, endTime, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMinStake provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetMinStake(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r2 = rf(ctx, subnetID, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPendingValidators provides a mock function with given fields: ctx, subnetID, nodeIDs, options
func (_m *PClient) GetPendingValidators(ctx context.Context, subnetID ids.ID, nodeIDs []ids.NodeID, options ...rpc.Option) ([]interface{}, []interface{}, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID, nodeIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, []ids.NodeID, ...rpc.Option) []interface{}); ok {
		r0 = rf(ctx, subnetID, nodeIDs, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 []interface{}
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, []ids.NodeID, ...rpc.Option) []interface{}); ok {
		r1 = rf(ctx, subnetID, nodeIDs, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ids.ID, []ids.NodeID, ...rpc.Option) error); ok {
		r2 = rf(ctx, subnetID, nodeIDs, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRewardUTXOs provides a mock function with given fields: _a0, _a1, _a2
func (_m *PClient) GetRewardUTXOs(_a0 context.Context, _a1 *api.GetTxArgs, _a2 ...rpc.Option) ([][]byte, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetTxArgs, ...rpc.Option) [][]byte); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.GetTxArgs, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStake provides a mock function with given fields: ctx, addrs, options
func (_m *PClient) GetStake(ctx context.Context, addrs []ids.ShortID, options ...rpc.Option) (map[ids.ID]uint64, [][]byte, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, addrs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[ids.ID]uint64
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, ...rpc.Option) map[ids.ID]uint64); ok {
		r0 = rf(ctx, addrs, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ids.ID]uint64)
		}
	}

	var r1 [][]byte
	if rf, ok := ret.Get(1).(func(context.Context, []ids.ShortID, ...rpc.Option) [][]byte); ok {
		r1 = rf(ctx, addrs, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []ids.ShortID, ...rpc.Option) error); ok {
		r2 = rf(ctx, addrs, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetStakingAssetID provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetStakingAssetID(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubnets provides a mock function with given fields: ctx, subnetIDs, options
func (_m *PClient) GetSubnets(ctx context.Context, subnetIDs []ids.ID, options ...rpc.Option) ([]platformvm.ClientSubnet, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetIDs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []platformvm.ClientSubnet
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ID, ...rpc.Option) []platformvm.ClientSubnet); ok {
		r0 = rf(ctx, subnetIDs, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]platformvm.ClientSubnet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetIDs, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTimestamp provides a mock function with given fields: ctx, options
func (_m *PClient) GetTimestamp(ctx context.Context, options ...rpc.Option) (time.Time, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) time.Time); ok {
		r0 = rf(ctx, options...)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalStake provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) GetTotalStake(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) uint64); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTx provides a mock function with given fields: ctx, txID, options
func (_m *PClient) GetTx(ctx context.Context, txID ids.ID, options ...rpc.Option) ([]byte, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, txID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) []byte); ok {
		r0 = rf(ctx, txID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, txID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxStatus provides a mock function with given fields: ctx, txID, options
func (_m *PClient) GetTxStatus(ctx context.Context, txID ids.ID, options ...rpc.Option) (*platformvm.GetTxStatusResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, txID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *platformvm.GetTxStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) *platformvm.GetTxStatusResponse); ok {
		r0 = rf(ctx, txID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*platformvm.GetTxStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, txID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUTXOs provides a mock function with given fields: ctx, addrs, limit, startAddress, startUTXOID, options
func (_m *PClient) GetUTXOs(ctx context.Context, addrs []ids.ShortID, limit uint32, startAddress ids.ShortID, startUTXOID ids.ID, options ...rpc.Option) ([][]byte, ids.ShortID, ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, addrs, limit, startAddress, startUTXOID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) [][]byte); ok {
		r0 = rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 ids.ShortID
	if rf, ok := ret.Get(1).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) ids.ShortID); ok {
		r1 = rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ids.ShortID)
		}
	}

	var r2 ids.ID
	if rf, ok := ret.Get(2).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) ids.ID); ok {
		r2 = rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(ids.ID)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, []ids.ShortID, uint32, ids.ShortID, ids.ID, ...rpc.Option) error); ok {
		r3 = rf(ctx, addrs, limit, startAddress, startUTXOID, options...)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetValidatorsAt provides a mock function with given fields: ctx, subnetID, height, options
func (_m *PClient) GetValidatorsAt(ctx context.Context, subnetID ids.ID, height uint64, options ...rpc.Option) (map[ids.NodeID]uint64, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID, height)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[ids.NodeID]uint64
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, uint64, ...rpc.Option) map[ids.NodeID]uint64); ok {
		r0 = rf(ctx, subnetID, height, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ids.NodeID]uint64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, uint64, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, height, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportAVAX provides a mock function with given fields: ctx, user, from, changeAddr, to, sourceChain, options
func (_m *PClient) ImportAVAX(ctx context.Context, user api.UserPass, from []ids.ShortID, changeAddr ids.ShortID, to ids.ShortID, sourceChain string, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, from, changeAddr, to, sourceChain)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ShortID, string, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, user, from, changeAddr, to, sourceChain, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, []ids.ShortID, ids.ShortID, ids.ShortID, string, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, from, changeAddr, to, sourceChain, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportKey provides a mock function with given fields: ctx, user, privateKey, options
func (_m *PClient) ImportKey(ctx context.Context, user api.UserPass, privateKey *crypto.PrivateKeySECP256K1R, options ...rpc.Option) (ids.ShortID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user, privateKey)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ShortID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, *crypto.PrivateKeySECP256K1R, ...rpc.Option) ids.ShortID); ok {
		r0 = rf(ctx, user, privateKey, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ShortID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, *crypto.PrivateKeySECP256K1R, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, privateKey, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IssueTx provides a mock function with given fields: ctx, tx, options
func (_m *PClient) IssueTx(ctx context.Context, tx []byte, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, tx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, []byte, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, tx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, ...rpc.Option) error); ok {
		r1 = rf(ctx, tx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAddresses provides a mock function with given fields: ctx, user, options
func (_m *PClient) ListAddresses(ctx context.Context, user api.UserPass, options ...rpc.Option) ([]ids.ShortID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, user)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []ids.ShortID
	if rf, ok := ret.Get(0).(func(context.Context, api.UserPass, ...rpc.Option) []ids.ShortID); ok {
		r0 = rf(ctx, user, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ids.ShortID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.UserPass, ...rpc.Option) error); ok {
		r1 = rf(ctx, user, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SampleValidators provides a mock function with given fields: ctx, subnetID, sampleSize, options
func (_m *PClient) SampleValidators(ctx context.Context, subnetID ids.ID, sampleSize uint16, options ...rpc.Option) ([]ids.NodeID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID, sampleSize)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []ids.NodeID
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, uint16, ...rpc.Option) []ids.NodeID); ok {
		r0 = rf(ctx, subnetID, sampleSize, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ids.NodeID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, uint16, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, sampleSize, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatedBy provides a mock function with given fields: ctx, blockchainID, options
func (_m *PClient) ValidatedBy(ctx context.Context, blockchainID ids.ID, options ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, blockchainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) ids.ID); ok {
		r0 = rf(ctx, blockchainID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, blockchainID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validates provides a mock function with given fields: ctx, subnetID, options
func (_m *PClient) Validates(ctx context.Context, subnetID ids.ID, options ...rpc.Option) ([]ids.ID, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subnetID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, ids.ID, ...rpc.Option) []ids.ID); ok {
		r0 = rf(ctx, subnetID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ids.ID, ...rpc.Option) error); ok {
		r1 = rf(ctx, subnetID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
