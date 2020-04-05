// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	user "github.com/freshteapot/learnalist-api/server/pkg/user"
	mock "github.com/stretchr/testify/mock"
)

// Session is an autogenerated mock type for the Session type
type Session struct {
	mock.Mock
}

// Activate provides a mock function with given fields: session
func (_m *Session) Activate(session user.UserSession) error {
	ret := _m.Called(session)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.UserSession) error); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWithChallenge provides a mock function with given fields:
func (_m *Session) CreateWithChallenge() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserUUIDByToken provides a mock function with given fields: token
func (_m *Session) GetUserUUIDByToken(token string) (string, error) {
	ret := _m.Called(token)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsChallengeValid provides a mock function with given fields: challenge
func (_m *Session) IsChallengeValid(challenge string) (bool, error) {
	ret := _m.Called(challenge)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(challenge)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(challenge)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSession provides a mock function with given fields: userUUID
func (_m *Session) NewSession(userUUID string) (user.UserSession, error) {
	ret := _m.Called(userUUID)

	var r0 user.UserSession
	if rf, ok := ret.Get(0).(func(string) user.UserSession); ok {
		r0 = rf(userUUID)
	} else {
		r0 = ret.Get(0).(user.UserSession)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveExpiredChallenges provides a mock function with given fields:
func (_m *Session) RemoveExpiredChallenges() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveSessionForUser provides a mock function with given fields: userUUID, token
func (_m *Session) RemoveSessionForUser(userUUID string, token string) error {
	ret := _m.Called(userUUID, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userUUID, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveSessionsForUser provides a mock function with given fields: userUUID
func (_m *Session) RemoveSessionsForUser(userUUID string) error {
	ret := _m.Called(userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}