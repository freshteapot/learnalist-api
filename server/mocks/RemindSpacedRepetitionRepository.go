// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	time "time"

	remind "github.com/freshteapot/learnalist-api/server/pkg/remind"
	mock "github.com/stretchr/testify/mock"
)

// RemindSpacedRepetitionRepository is an autogenerated mock type for the RemindSpacedRepetitionRepository type
type RemindSpacedRepetitionRepository struct {
	mock.Mock
}

// DeleteByUser provides a mock function with given fields: userUUID
func (_m *RemindSpacedRepetitionRepository) DeleteByUser(userUUID string) error {
	ret := _m.Called(userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetReminders provides a mock function with given fields:
func (_m *RemindSpacedRepetitionRepository) GetReminders() []remind.SpacedRepetitionReminder {
	ret := _m.Called()

	var r0 []remind.SpacedRepetitionReminder
	if rf, ok := ret.Get(0).(func() []remind.SpacedRepetitionReminder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]remind.SpacedRepetitionReminder)
		}
	}

	return r0
}

// SetReminder provides a mock function with given fields: userUUID, whenNext, lastActive
func (_m *RemindSpacedRepetitionRepository) SetReminder(userUUID string, whenNext time.Time, lastActive time.Time) error {
	ret := _m.Called(userUUID, whenNext, lastActive)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) error); ok {
		r0 = rf(userUUID, whenNext, lastActive)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSent provides a mock function with given fields: userUUID, sent
func (_m *RemindSpacedRepetitionRepository) UpdateSent(userUUID string, sent int) error {
	ret := _m.Called(userUUID, sent)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(userUUID, sent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
