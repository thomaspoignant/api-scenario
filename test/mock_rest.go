package test

import (
	"github.com/sendgrid/rest"
)

type ClientMock struct {
}

func (c *ClientMock) Send(request rest.Request) (*rest.Response, error) {
	testNumber:=request.QueryParams["testNumber"]
	if testNumber == "1" {
		return &rest.Response{
			Body: `{
					"hello":"world",
					"param1":true,
					"param2":123
					}`,
			StatusCode: 200,
			Headers: map[string][]string{
				"Content-Type": {"application/json"},
			},
		}, nil
	} else if testNumber == "2" {
		return &rest.Response{
			Body: `<root>
						<hello>world</hello>
						<param1>true</param1>
						<param2>123</param2>
				   </root>
					`,
			StatusCode: 200,
			Headers: map[string][]string{
				"Content-Type": {"application/json"},
			},
		}, nil
	}

	return &rest.Response{}, nil
}
