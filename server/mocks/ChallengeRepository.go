// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	challenge "github.com/freshteapot/learnalist-api/server/pkg/challenge"
	mock "github.com/stretchr/testify/mock"
)

// ChallengeRepository is an autogenerated mock type for the ChallengeRepository type
type ChallengeRepository struct {
	mock.Mock
}

// AddRecord provides a mock function with given fields: UUID, extUUID, userUUID
func (_m *ChallengeRepository) AddRecord(UUID string, extUUID string, userUUID string) (int, error) {
	ret := _m.Called(UUID, extUUID, userUUID)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string) int); ok {
		r0 = rf(UUID, extUUID, userUUID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(UUID, extUUID, userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: userUUID, _a1
func (_m *ChallengeRepository) Create(userUUID string, _a1 challenge.ChallengeInfo) error {
	ret := _m.Called(userUUID, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, challenge.ChallengeInfo) error); ok {
		r0 = rf(userUUID, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: UUID
func (_m *ChallengeRepository) Delete(UUID string) error {
	ret := _m.Called(UUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(UUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRecord provides a mock function with given fields: extUUID, userUUID
func (_m *ChallengeRepository) DeleteRecord(extUUID string, userUUID string) error {
	ret := _m.Called(extUUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(extUUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: userUUID
func (_m *ChallengeRepository) DeleteUser(userUUID string) error {
	ret := _m.Called(userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: UUID
func (_m *ChallengeRepository) Get(UUID string) (challenge.ChallengeInfo, error) {
	ret := _m.Called(UUID)

	var r0 challenge.ChallengeInfo
	if rf, ok := ret.Get(0).(func(string) challenge.ChallengeInfo); ok {
		r0 = rf(UUID)
	} else {
		r0 = ret.Get(0).(challenge.ChallengeInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChallengesByUser provides a mock function with given fields: userUUID
func (_m *ChallengeRepository) GetChallengesByUser(userUUID string) ([]challenge.ChallengeShortInfo, error) {
	ret := _m.Called(userUUID)

	var r0 []challenge.ChallengeShortInfo
	if rf, ok := ret.Get(0).(func(string) []challenge.ChallengeShortInfo); ok {
		r0 = rf(userUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]challenge.ChallengeShortInfo)
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

// Join provides a mock function with given fields: UUID, userUUID
func (_m *ChallengeRepository) Join(UUID string, userUUID string) error {
	ret := _m.Called(UUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(UUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Leave provides a mock function with given fields: UUID, userUUID
func (_m *ChallengeRepository) Leave(UUID string, userUUID string) error {
	ret := _m.Called(UUID, userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(UUID, userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
