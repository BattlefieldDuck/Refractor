// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "Refractor/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AttachmentRepo is an autogenerated mock type for the AttachmentRepo type
type AttachmentRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *AttachmentRepo) Delete(ctx context.Context, id int64) error {
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
func (_m *AttachmentRepo) GetByID(ctx context.Context, id int64) (*domain.Attachment, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Attachment
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Attachment); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Attachment)
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

// GetByInfraction provides a mock function with given fields: ctx, infractionID
func (_m *AttachmentRepo) GetByInfraction(ctx context.Context, infractionID int64) ([]*domain.Attachment, error) {
	ret := _m.Called(ctx, infractionID)

	var r0 []*domain.Attachment
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*domain.Attachment); ok {
		r0 = rf(ctx, infractionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, infractionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, attachment
func (_m *AttachmentRepo) Store(ctx context.Context, attachment *domain.Attachment) error {
	ret := _m.Called(ctx, attachment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Attachment) error); ok {
		r0 = rf(ctx, attachment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}