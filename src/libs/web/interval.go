package web

import (
	"time"
)

type IntervalRequest interface {
	GetJsonData(url string, v interface{}, headers map[string]string) (err error)
	PostJsonData(url string, request interface{}, v interface{}, headers map[string]string) (err error)
}

type intervalRequest struct {
	ticker *time.Ticker
}

func NewIntervalRequest(interval time.Duration) IntervalRequest {
	return &intervalRequest{ticker: time.NewTicker(interval)}
}

func (i intervalRequest) GetJsonData(url string, v interface{}, headers map[string]string) (err error) {
	<-i.ticker.C
	return GetJsonData(url, v, headers)
}

func (i intervalRequest) PostJsonData(url string, request interface{}, v interface{}, headers map[string]string) (err error) {
	<-i.ticker.C
	return PostJsonData(url, request, v, headers)
}
