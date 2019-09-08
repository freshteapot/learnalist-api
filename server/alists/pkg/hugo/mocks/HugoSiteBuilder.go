// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import alist "github.com/freshteapot/learnalist-api/server/api/alist"

import mock "github.com/stretchr/testify/mock"

// HugoSiteBuilder is an autogenerated mock type for the HugoSiteBuilder type
type HugoSiteBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *HugoSiteBuilder) Build() {
	_m.Called()
}

// MakeContent provides a mock function with given fields:
func (_m *HugoSiteBuilder) MakeContent() {
	_m.Called()
}

// Remove provides a mock function with given fields: uuid
func (_m *HugoSiteBuilder) Remove(uuid string) {
	_m.Called(uuid)
}

// Write provides a mock function with given fields: aList
func (_m *HugoSiteBuilder) Write(aList *alist.Alist) {
	_m.Called(aList)
}
