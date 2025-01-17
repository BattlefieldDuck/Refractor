// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "Refractor/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ServerRepo is an autogenerated mock type for the ServerRepo type
type ServerRepo struct {
	mock.Mock
}

// Deactivate provides a mock function with given fields: ctx, id
func (_m *ServerRepo) Deactivate(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, args
func (_m *ServerRepo) Exists(ctx context.Context, args domain.FindArgs) (bool, error) {
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

// GetAll provides a mock function with given fields: ctx
func (_m *ServerRepo) GetAll(ctx context.Context) ([]*domain.Server, error) {
	ret := _m.Called(ctx)

	var r0 []*domain.Server
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Server); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ServerRepo) GetByID(ctx context.Context, id int64) (*domain.Server, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Server
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Server); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, server
func (_m *ServerRepo) Store(ctx context.Context, server *domain.Server) error {
	ret := _m.Called(ctx, server)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Server) error); ok {
		r0 = rf(ctx, server)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, id, args
func (_m *ServerRepo) Update(ctx context.Context, id int64, args domain.UpdateArgs) (*domain.Server, error) {
	ret := _m.Called(ctx, id, args)

	var r0 *domain.Server
	if rf, ok := ret.Get(0).(func(context.Context, int64, domain.UpdateArgs) *domain.Server); ok {
		r0 = rf(ctx, id, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, domain.UpdateArgs) error); ok {
		r1 = rf(ctx, id, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
