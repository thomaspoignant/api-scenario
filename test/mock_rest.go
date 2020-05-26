package test

import (
	"github.com/sendgrid/rest"
)

type ClientMock struct {
}

func (c *ClientMock) Send(request rest.Request) (*rest.Response, error) {


	if request.QueryParams["testNumber"] == "1" {
		return &rest.Response{
			Body: `{"hello":"world"}`,
			StatusCode: 200,
			Headers: map[string][]string{
				"Content-Type": {"application/json"},
			},
		}, nil
	}



	return &rest.Response{}, nil
}