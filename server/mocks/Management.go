// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Management is an autogenerated mock type for the Management type
type Management struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: userUUID
func (_m *Management) DeleteUser(userUUID string) error {
	ret := _m.Called(userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUser provides a mock function with given fields: search
func (_m *Management) FindUser(search string) ([]string, error) {
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

// GetInfo provides a mock function with given fields: userUUID
func (_m *Management) GetInfo(userUUID string) ([]byte, error) {
	ret := _m.Called(userUUID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(userUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
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

// SaveInfo provides a mock function with given fields: userUUID, info
func (_m *Management) SaveInfo(userUUID string, info []byte) error {
	ret := _m.Called(userUUID, info)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(userUUID, info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
