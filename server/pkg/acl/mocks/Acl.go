// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Acl is an autogenerated mock type for the Acl type
type Acl struct {
	mock.Mock
}

// DeleteList provides a mock function with given fields: alistUUID
func (_m *Acl) DeleteList(alistUUID string) error {
	ret := _m.Called(alistUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(alistUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GrantUserListReadAccess provides a mock function with given fields: alistUUID, userUUID
func (_m *Acl) GrantUserListReadAccess(alistUUID string, userUUID string) error {
	ret := _m.Called(alistUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(alistUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GrantUserListWriteAccess provides a mock function with given fields: alistUUID, userUUID
func (_m *Acl) GrantUserListWriteAccess(alistUUID string, userUUID string) error {
	ret := _m.Called(alistUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(alistUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasUserListReadAccess provides a mock function with given fields: alistUUID, userUUID
func (_m *Acl) HasUserListReadAccess(alistUUID string, userUUID string) (bool, error) {
	ret := _m.Called(alistUUID, userUUID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(alistUUID, userUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(alistUUID, userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasUserListWriteAccess provides a mock function with given fields: alistUUID, userUUID
func (_m *Acl) HasUserListWriteAccess(alistUUID string, userUUID string) (bool, error) {
	ret := _m.Called(alistUUID, userUUID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(alistUUID, userUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(alistUUID, userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsListAvailableToFriends provides a mock function with given fields: alistUUID
func (_m *Acl) IsListAvailableToFriends(alistUUID string) (bool, error) {
	ret := _m.Called(alistUUID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(alistUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alistUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsListPrivate provides a mock function with given fields: alistUUID
func (_m *Acl) IsListPrivate(alistUUID string) (bool, error) {
	ret := _m.Called(alistUUID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(alistUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alistUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsListPublic provides a mock function with given fields: alistUUID
func (_m *Acl) IsListPublic(alistUUID string) (bool, error) {
	ret := _m.Called(alistUUID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(alistUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alistUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIsSharedWith provides a mock function with given fields: alistUUID
func (_m *Acl) ListIsSharedWith(alistUUID string) (string, error) {
	ret := _m.Called(alistUUID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(alistUUID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alistUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeListPrivate provides a mock function with given fields: alistUUID, userUUID
func (_m *Acl) MakeListPrivate(alistUUID string, userUUID string) error {
	ret := _m.Called(alistUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(alistUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeUserListReadAccess provides a mock function with given fields: alistUUID, userUUID
func (_m *Acl) RevokeUserListReadAccess(alistUUID string, userUUID string) error {
	ret := _m.Called(alistUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(alistUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeUserListWriteAccess provides a mock function with given fields: alistUUID, userUUID
func (_m *Acl) RevokeUserListWriteAccess(alistUUID string, userUUID string) error {
	ret := _m.Called(alistUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(alistUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShareListWithFriends provides a mock function with given fields: alistUUID
func (_m *Acl) ShareListWithFriends(alistUUID string) error {
	ret := _m.Called(alistUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(alistUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShareListWithPublic provides a mock function with given fields: alistUUID
func (_m *Acl) ShareListWithPublic(alistUUID string) error {
	ret := _m.Called(alistUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(alistUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
