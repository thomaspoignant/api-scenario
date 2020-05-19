package model

import (
	"errors"
	"net/http"
	"time"

	"github.com/sendgrid/rest"
	"github.com/thomaspoignant/api-scenario/pkg/util"
)

type Response struct {
	TimeElapsed time.Duration          // e.g 1ms
	StatusCode  int                    // e.g. 200
	Body        map[string]interface{} // e.g. {"result: Success"}
	Header      http.Header            // e.g. map[X-Ratelimit-Limit:[600]]
}

// Create a new responseApi from a rest.Response
func NewResponse(restResponse rest.Response, timeElapsed time.Duration) (Response, error) {
	var body map[string]interface{}
	if len(restResponse.Body) > 0 && !util.IsJson(restResponse.Body) {
		return Response{}, errors.New("there is a result and this is not a valid JSON api Response is not in JSON")
	}

	body, err := util.StringToJson(restResponse.Body)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Body:        body,
		StatusCode:  restResponse.StatusCode,
		TimeElapsed: timeElapsed,
		Header:      restResponse.Headers,
	}, nil
}
