// Code generated by mockery v2.23.2. DO NOT EDIT.

package mocks

import (
	context "context"

	logr "github.com/go-logr/logr"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// DatasourceChanged provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Client) DatasourceChanged(_a0 context.Context, _a1 logr.Logger, _a2 string, _a3 string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, logr.Logger, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
