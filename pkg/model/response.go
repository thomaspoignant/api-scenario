package model

import (
	"net/http"
	"time"

	"github.com/sendgrid/rest"
)

type Response struct {
	TimeElapsed time.Duration `json:"time_elapsed,omitempty"` // e.g 1ms
	StatusCode  int           `json:"status_code,omitempty"`  // e.g. 200
	Body        string        `json:"body,omitempty"`         // e.g. {"result: Success"}
	Header      http.Header   `json:"header,omitempty"`       // e.g. map[X-Ratelimit-Limit:[600]]
}

// Create a new responseApi from a rest.Response
func NewResponse(restResponse rest.Response, timeElapsed time.Duration) (Response, error) {
	return Response{
		Body:        restResponse.Body,
		StatusCode:  restResponse.StatusCode,
		TimeElapsed: timeElapsed,
		Header:      restResponse.Headers,
	}, nil
}
