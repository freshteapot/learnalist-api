// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	alist "github.com/freshteapot/learnalist-api/server/api/alist"
	mock "github.com/stretchr/testify/mock"

	models "github.com/freshteapot/learnalist-api/server/api/models"

	oauth "github.com/freshteapot/learnalist-api/server/pkg/oauth"

	user "github.com/freshteapot/learnalist-api/server/pkg/user"
)

// Datastore is an autogenerated mock type for the Datastore type
type Datastore struct {
	mock.Mock
}

// GetAlist provides a mock function with given fields: uuid
func (_m *Datastore) GetAlist(uuid string) (*alist.Alist, error) {
	ret := _m.Called(uuid)

	var r0 *alist.Alist
	if rf, ok := ret.Get(0).(func(string) *alist.Alist); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*alist.Alist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListsByUserWithFilters provides a mock function with given fields: uuid, labels, listType
func (_m *Datastore) GetListsByUserWithFilters(uuid string, labels string, listType string) []*alist.Alist {
	ret := _m.Called(uuid, labels, listType)

	var r0 []*alist.Alist
	if rf, ok := ret.Get(0).(func(string, string, string) []*alist.Alist); ok {
		r0 = rf(uuid, labels, listType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*alist.Alist)
		}
	}

	return r0
}

// GetUserLabels provides a mock function with given fields: uuid
func (_m *Datastore) GetUserLabels(uuid string) ([]string, error) {
	ret := _m.Called(uuid)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OAuthHandler provides a mock function with given fields:
func (_m *Datastore) OAuthHandler() oauth.OAuthReadWriter {
	ret := _m.Called()

	var r0 oauth.OAuthReadWriter
	if rf, ok := ret.Get(0).(func() oauth.OAuthReadWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oauth.OAuthReadWriter)
		}
	}

	return r0
}

// PostAlistLabel provides a mock function with given fields: label
func (_m *Datastore) PostAlistLabel(label *models.AlistLabel) (int, error) {
	ret := _m.Called(label)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.AlistLabel) int); ok {
		r0 = rf(label)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.AlistLabel) error); ok {
		r1 = rf(label)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostUserLabel provides a mock function with given fields: label
func (_m *Datastore) PostUserLabel(label *models.UserLabel) (int, error) {
	ret := _m.Called(label)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.UserLabel) int); ok {
		r0 = rf(label)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.UserLabel) error); ok {
		r1 = rf(label)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveAlist provides a mock function with given fields: alist_uuid, user_uuid
func (_m *Datastore) RemoveAlist(alist_uuid string, user_uuid string) error {
	ret := _m.Called(alist_uuid, user_uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(alist_uuid, user_uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveUserLabel provides a mock function with given fields: label, uuid
func (_m *Datastore) RemoveUserLabel(label string, uuid string) error {
	ret := _m.Called(label, uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(label, uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveAlist provides a mock function with given fields: method, aList
func (_m *Datastore) SaveAlist(method string, aList alist.Alist) (*alist.Alist, error) {
	ret := _m.Called(method, aList)

	var r0 *alist.Alist
	if rf, ok := ret.Get(0).(func(string, alist.Alist) *alist.Alist); ok {
		r0 = rf(method, aList)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*alist.Alist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, alist.Alist) error); ok {
		r1 = rf(method, aList)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserExists provides a mock function with given fields: userUUID
func (_m *Datastore) UserExists(userUUID string) bool {
	ret := _m.Called(userUUID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(userUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UserFromIDP provides a mock function with given fields:
func (_m *Datastore) UserFromIDP() user.UserFromIDP {
	ret := _m.Called()

	var r0 user.UserFromIDP
	if rf, ok := ret.Get(0).(func() user.UserFromIDP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(user.UserFromIDP)
		}
	}

	return r0
}

// UserSession provides a mock function with given fields:
func (_m *Datastore) UserSession() user.Session {
	ret := _m.Called()

	var r0 user.Session
	if rf, ok := ret.Get(0).(func() user.Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(user.Session)
		}
	}

	return r0
}

// UserWithUsernameAndPassword provides a mock function with given fields:
func (_m *Datastore) UserWithUsernameAndPassword() user.UserWithUsernameAndPassword {
	ret := _m.Called()

	var r0 user.UserWithUsernameAndPassword
	if rf, ok := ret.Get(0).(func() user.UserWithUsernameAndPassword); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(user.UserWithUsernameAndPassword)
		}
	}

	return r0
}
