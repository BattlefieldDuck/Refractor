// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "Refractor/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// InfractionRepo is an autogenerated mock type for the InfractionRepo type
type InfractionRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *InfractionRepo) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *InfractionRepo) GetByID(ctx context.Context, id int64) (*domain.Infraction, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Infraction
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Infraction); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Infraction)
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

// GetByPlayer provides a mock function with given fields: ctx, playerID, platform
func (_m *InfractionRepo) GetByPlayer(ctx context.Context, playerID string, platform string) ([]*domain.Infraction, error) {
	ret := _m.Called(ctx, playerID, platform)

	var r0 []*domain.Infraction
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*domain.Infraction); ok {
		r0 = rf(ctx, playerID, platform)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Infraction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, playerID, platform)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, args, limit, offset
func (_m *InfractionRepo) Search(ctx context.Context, args domain.FindArgs, limit int, offset int) (int, []*domain.Infraction, error) {
	ret := _m.Called(ctx, args, limit, offset)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, domain.FindArgs, int, int) int); ok {
		r0 = rf(ctx, args, limit, offset)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []*domain.Infraction
	if rf, ok := ret.Get(1).(func(context.Context, domain.FindArgs, int, int) []*domain.Infraction); ok {
		r1 = rf(ctx, args, limit, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*domain.Infraction)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, domain.FindArgs, int, int) error); ok {
		r2 = rf(ctx, args, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Store provides a mock function with given fields: ctx, infraction
func (_m *InfractionRepo) Store(ctx context.Context, infraction *domain.Infraction) (*domain.Infraction, error) {
	ret := _m.Called(ctx, infraction)

	var r0 *domain.Infraction
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Infraction) *domain.Infraction); ok {
		r0 = rf(ctx, infraction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Infraction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Infraction) error); ok {
		r1 = rf(ctx, infraction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, args
func (_m *InfractionRepo) Update(ctx context.Context, id int64, args domain.UpdateArgs) (*domain.Infraction, error) {
	ret := _m.Called(ctx, id, args)

	var r0 *domain.Infraction
	if rf, ok := ret.Get(0).(func(context.Context, int64, domain.UpdateArgs) *domain.Infraction); ok {
		r0 = rf(ctx, id, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Infraction)
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
