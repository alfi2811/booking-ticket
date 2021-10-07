package qrcode

import "booking-ticket/business/qrcode"

type Response struct {
	QrCode string `json:"qrcode"`
}

func (resp *Response) ToDomain() qrcode.Domain {
	return qrcode.Domain{
		QrCode: resp.QrCode,
	}
}
