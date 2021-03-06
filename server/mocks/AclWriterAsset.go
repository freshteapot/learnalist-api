// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AclWriterAsset is an autogenerated mock type for the AclWriterAsset type
type AclWriterAsset struct {
	mock.Mock
}

// DeleteAsset provides a mock function with given fields: extUUID
func (_m *AclWriterAsset) DeleteAsset(extUUID string) error {
	ret := _m.Called(extUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(extUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GrantUserAssetReadAccess provides a mock function with given fields: extUUID, userUUID
func (_m *AclWriterAsset) GrantUserAssetReadAccess(extUUID string, userUUID string) error {
	ret := _m.Called(extUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(extUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MakeAssetPrivate provides a mock function with given fields: extUUID, userUUID
func (_m *AclWriterAsset) MakeAssetPrivate(extUUID string, userUUID string) error {
	ret := _m.Called(extUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(extUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeUserAssetReadAccess provides a mock function with given fields: extUUID, userUUID
func (_m *AclWriterAsset) RevokeUserAssetReadAccess(extUUID string, userUUID string) error {
	ret := _m.Called(extUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(extUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShareAssetWithPublic provides a mock function with given fields: extUUID
func (_m *AclWriterAsset) ShareAssetWithPublic(extUUID string) error {
	ret := _m.Called(extUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(extUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
