// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BinaryChecker is an autogenerated mock type for the BinaryChecker type
type BinaryChecker struct {
	mock.Mock
}

// ExistsWithVersion provides a mock function with given fields: name, binaryPrefix, version
func (_m *BinaryChecker) ExistsWithVersion(name string, binaryPrefix string, version string) (bool, error) {
	ret := _m.Called(name, binaryPrefix, version)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(name, binaryPrefix, version)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(name, binaryPrefix, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBinaryChecker interface {
	mock.TestingT
	Cleanup(func())
}

// NewBinaryChecker creates a new instance of BinaryChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBinaryChecker(t mockConstructorTestingTNewBinaryChecker) *BinaryChecker {
	mock := &BinaryChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}