// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	context "context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"

	mock "github.com/stretchr/testify/mock"
)

// mockRegistryListAllPager is an autogenerated mock type for the registryListAllPager type
type mockRegistryListAllPager struct {
	mock.Mock
}

// Err provides a mock function with given fields:
func (_m *mockRegistryListAllPager) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NextPage provides a mock function with given fields: ctx
func (_m *mockRegistryListAllPager) NextPage(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PageResponse provides a mock function with given fields:
func (_m *mockRegistryListAllPager) PageResponse() armcontainerregistry.RegistriesListResponse {
	ret := _m.Called()

	var r0 armcontainerregistry.RegistriesListResponse
	if rf, ok := ret.Get(0).(func() armcontainerregistry.RegistriesListResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(armcontainerregistry.RegistriesListResponse)
	}

	return r0
}
