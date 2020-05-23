// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"
)

// OAuth2ConfigInterface is an autogenerated mock type for the OAuth2ConfigInterface type
type OAuth2ConfigInterface struct {
	mock.Mock
}

// AuthCodeURL provides a mock function with given fields: state, opts
func (_m *OAuth2ConfigInterface) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, state)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...oauth2.AuthCodeOption) string); ok {
		r0 = rf(state, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Client provides a mock function with given fields: ctx, t
func (_m *OAuth2ConfigInterface) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	ret := _m.Called(ctx, t)

	var r0 *http.Client
	if rf, ok := ret.Get(0).(func(context.Context, *oauth2.Token) *http.Client); ok {
		r0 = rf(ctx, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Client)
		}
	}

	return r0
}

// Exchange provides a mock function with given fields: ctx, code, opts
func (_m *OAuth2ConfigInterface) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, code)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *oauth2.Token
	if rf, ok := ret.Get(0).(func(context.Context, string, ...oauth2.AuthCodeOption) *oauth2.Token); ok {
		r0 = rf(ctx, code, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...oauth2.AuthCodeOption) error); ok {
		r1 = rf(ctx, code, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenSource provides a mock function with given fields: ctx, token
func (_m *OAuth2ConfigInterface) TokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource {
	ret := _m.Called(ctx, token)

	var r0 oauth2.TokenSource
	if rf, ok := ret.Get(0).(func(context.Context, *oauth2.Token) oauth2.TokenSource); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oauth2.TokenSource)
		}
	}

	return r0
}