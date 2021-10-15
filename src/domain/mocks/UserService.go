// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "Refractor/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// DeactivateUser provides a mock function with given fields: c, userID
func (_m *UserService) DeactivateUser(c context.Context, userID string) error {
	ret := _m.Called(c, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields: c
func (_m *UserService) GetAllUsers(c context.Context) ([]*domain.User, error) {
	ret := _m.Called(c)

	var r0 []*domain.User
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.User); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: c, userID
func (_m *UserService) GetByID(c context.Context, userID string) (*domain.User, error) {
	ret := _m.Called(c, userID)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLinkedPlayers provides a mock function with given fields: ctx, userID
func (_m *UserService) GetLinkedPlayers(ctx context.Context, userID string) ([]*domain.Player, error) {
	ret := _m.Called(ctx, userID)

	var r0 []*domain.Player
	if rf, ok := ret.Get(0).(func(context.Context, string) []*domain.Player); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Player)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkPlayer provides a mock function with given fields: c, userID, platform, playerID
func (_m *UserService) LinkPlayer(c context.Context, userID string, platform string, playerID string) error {
	ret := _m.Called(c, userID, platform, playerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(c, userID, platform, playerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReactivateUser provides a mock function with given fields: c, userID
func (_m *UserService) ReactivateUser(c context.Context, userID string) error {
	ret := _m.Called(c, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnlinkPlayer provides a mock function with given fields: c, userID, platform, playerID
func (_m *UserService) UnlinkPlayer(c context.Context, userID string, platform string, playerID string) error {
	ret := _m.Called(c, userID, platform, playerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(c, userID, platform, playerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}