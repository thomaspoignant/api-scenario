package test

import (
	"github.com/sendgrid/rest"
)

type ClientMock struct {
}

var jsonBody = `{
				"hello":"world",
				"param1":true,
				"param2":123
				}`

var xmlBody =`<root>
				<hello>world</hello>
				<param1>true</param1>
				<param2>123</param2>
		   	  </root>`

func (c *ClientMock) Send(request rest.Request) (*rest.Response, error) {
	testNumber:=request.QueryParams["testNumber"]

	response := &rest.Response{
		StatusCode: 200,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	if testNumber == "1" {
		response.Body = jsonBody
		return response, nil
	} else if testNumber == "2" {
		response.Body = xmlBody
		return response, nil
	}

	return &rest.Response{}, nil
}
