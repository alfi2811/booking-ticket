package qrcode

import (
	"context"
)

type Domain struct {
	QrCode string `json:"qrcode"`
}

type Repository interface {
	GetQrCode(ctx context.Context, dataBooking string) (Domain, error)
}
