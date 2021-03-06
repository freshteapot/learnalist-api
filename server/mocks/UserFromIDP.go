// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UserFromIDP is an autogenerated mock type for the UserFromIDP type
type UserFromIDP struct {
	mock.Mock
}

// Lookup provides a mock function with given fields: idp, kind, identifier
func (_m *UserFromIDP) Lookup(idp string, kind string, identifier string) (string, error) {
	ret := _m.Called(idp, kind, identifier)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(idp, kind, identifier)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(idp, kind, identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: idp, kind, identifier, info
func (_m *UserFromIDP) Register(idp string, kind string, identifier string, info []byte) (string, error) {
	ret := _m.Called(idp, kind, identifier, info)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, []byte) string); ok {
		r0 = rf(idp, kind, identifier, info)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, []byte) error); ok {
		r1 = rf(idp, kind, identifier, info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
