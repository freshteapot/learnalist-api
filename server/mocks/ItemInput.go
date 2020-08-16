// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ItemInput is an autogenerated mock type for the ItemInput type
type ItemInput struct {
	mock.Mock
}

// Created provides a mock function with given fields:
func (_m *ItemInput) Created() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// DecrThreshold provides a mock function with given fields:
func (_m *ItemInput) DecrThreshold() {
	_m.Called()
}

// IncrThreshold provides a mock function with given fields:
func (_m *ItemInput) IncrThreshold() {
	_m.Called()
}

// String provides a mock function with given fields:
func (_m *ItemInput) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UUID provides a mock function with given fields:
func (_m *ItemInput) UUID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WhenNext provides a mock function with given fields:
func (_m *ItemInput) WhenNext() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}
