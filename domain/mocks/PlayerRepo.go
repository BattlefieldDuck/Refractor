// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "Refractor/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PlayerRepo is an autogenerated mock type for the PlayerRepo type
type PlayerRepo struct {
	mock.Mock
}

// Exists provides a mock function with given fields: ctx, args
func (_m *PlayerRepo) Exists(ctx context.Context, args domain.FindArgs) (bool, error) {
	ret := _m.Called(ctx, args)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, domain.FindArgs) bool); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.FindArgs) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, platform, id
func (_m *PlayerRepo) GetByID(ctx context.Context, platform string, id string) (*domain.Player, error) {
	ret := _m.Called(ctx, platform, id)

	var r0 *domain.Player
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.Player); ok {
		r0 = rf(ctx, platform, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Player)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, platform, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchByName provides a mock function with given fields: ctx, name, limit, offset
func (_m *PlayerRepo) SearchByName(ctx context.Context, name string, limit int, offset int) (int, []*domain.Player, error) {
	ret := _m.Called(ctx, name, limit, offset)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) int); ok {
		r0 = rf(ctx, name, limit, offset)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []*domain.Player
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) []*domain.Player); ok {
		r1 = rf(ctx, name, limit, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*domain.Player)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int, int) error); ok {
		r2 = rf(ctx, name, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Store provides a mock function with given fields: ctx, player
func (_m *PlayerRepo) Store(ctx context.Context, player *domain.Player) error {
	ret := _m.Called(ctx, player)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Player) error); ok {
		r0 = rf(ctx, player)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, platform, id, args
func (_m *PlayerRepo) Update(ctx context.Context, platform string, id string, args domain.UpdateArgs) (*domain.Player, error) {
	ret := _m.Called(ctx, platform, id, args)

	var r0 *domain.Player
	if rf, ok := ret.Get(0).(func(context.Context, string, string, domain.UpdateArgs) *domain.Player); ok {
		r0 = rf(ctx, platform, id, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Player)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, domain.UpdateArgs) error); ok {
		r1 = rf(ctx, platform, id, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
