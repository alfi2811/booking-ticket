// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	qrcode "booking-ticket/business/qrcode"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetQrCode provides a mock function with given fields: ctx, dataBooking
func (_m *Repository) GetQrCode(ctx context.Context, dataBooking string) (qrcode.Domain, error) {
	ret := _m.Called(ctx, dataBooking)

	var r0 qrcode.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) qrcode.Domain); ok {
		r0 = rf(ctx, dataBooking)
	} else {
		r0 = ret.Get(0).(qrcode.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, dataBooking)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}