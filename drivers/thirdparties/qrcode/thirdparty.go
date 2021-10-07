package qrcode

import (
	"booking-ticket/business/qrcode"
	"context"
	"encoding/json"
	"net/http"
)

type QrcodeAPI struct {
	httpClient http.Client
	ApiKey     string
}

func NewQrCodeAPI(qrcodeApi QrcodeAPI) *QrcodeAPI {
	return &QrcodeAPI{
		httpClient: http.Client{},
		ApiKey:     qrcodeApi.ApiKey,
	}
}

func (qrcodeApi *QrcodeAPI) GetQrCode(ctx context.Context, dataBooking string) (qrcode.Domain, error) {
	req, _ := http.NewRequest("GET", "https://api.happi.dev/v1/qrcode?data="+dataBooking+"&apikey="+qrcodeApi.ApiKey, nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := qrcodeApi.httpClient.Do(req)
	if err != nil {
		return qrcode.Domain{}, err
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return qrcode.Domain{}, err
	}

	return data.ToDomain(), nil
}
