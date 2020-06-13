// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ManagementStorage is an autogenerated mock type for the ManagementStorage type
type ManagementStorage struct {
	mock.Mock
}

// DeleteList provides a mock function with given fields: listUUID
func (_m *ManagementStorage) DeleteList(listUUID string) error {
	ret := _m.Called(listUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(listUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: userUUID
func (_m *ManagementStorage) DeleteUser(userUUID string) error {
	ret := _m.Called(userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUserUUID provides a mock function with given fields: search
func (_m *ManagementStorage) FindUserUUID(search string) ([]string, error) {
	ret := _m.Called(search)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLists provides a mock function with given fields: userUUID
func (_m *ManagementStorage) GetLists(userUUID string) ([]string, error) {
	ret := _m.Called(userUUID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(userUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}