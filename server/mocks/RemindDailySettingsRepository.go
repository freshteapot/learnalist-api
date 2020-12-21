// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	openapi "github.com/freshteapot/learnalist-api/server/pkg/openapi"
	remind "github.com/freshteapot/learnalist-api/server/pkg/remind"
	mock "github.com/stretchr/testify/mock"
)

// RemindDailySettingsRepository is an autogenerated mock type for the RemindDailySettingsRepository type
type RemindDailySettingsRepository struct {
	mock.Mock
}

// ActivityHappened provides a mock function with given fields: userUUID, appIdentifier
func (_m *RemindDailySettingsRepository) ActivityHappened(userUUID string, appIdentifier string) error {
	ret := _m.Called(userUUID, appIdentifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userUUID, appIdentifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByApp provides a mock function with given fields: userUUID, appIdentifier
func (_m *RemindDailySettingsRepository) DeleteByApp(userUUID string, appIdentifier string) error {
	ret := _m.Called(userUUID, appIdentifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userUUID, appIdentifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByUser provides a mock function with given fields: userUUID
func (_m *RemindDailySettingsRepository) DeleteByUser(userUUID string) error {
	ret := _m.Called(userUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: userUUID, settings, whenNext
func (_m *RemindDailySettingsRepository) Save(userUUID string, settings openapi.RemindDailySettings, whenNext string) error {
	ret := _m.Called(userUUID, settings, whenNext)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, openapi.RemindDailySettings, string) error); ok {
		r0 = rf(userUUID, settings, whenNext)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WhoToRemind provides a mock function with given fields:
func (_m *RemindDailySettingsRepository) WhoToRemind() []remind.RemindMe {
	ret := _m.Called()

	var r0 []remind.RemindMe
	if rf, ok := ret.Get(0).(func() []remind.RemindMe); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]remind.RemindMe)
		}
	}

	return r0
}
