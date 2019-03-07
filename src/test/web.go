package test

import (
	mockhttp "github.com/karupanerura/go-mock-http-response"
	"net/http"
)

func MockHTTPResponse(statusCode int, headers map[string]string, body []byte) {
	http.DefaultClient = mockhttp.NewResponseMock(statusCode, headers, body).MakeClient()
}
