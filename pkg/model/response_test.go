package model_test

import (
	"github.com/sendgrid/rest"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
	"time"
)

func TestNewResponseCreateAValidResponse(t *testing.T) {
	restResp := rest.Response{
		StatusCode: 200,
		Body:       `{ "hello": "world"}`,
		Headers: map[string][]string{
			"Accept": {"toto"},
		},
	}

	expectedDuration := time.Duration(1 * time.Second)
	response, err := model.NewResponse(restResp, expectedDuration)
	test.Ok(t, err)
	test.Equals(t, "Invalid duration", expectedDuration, response.TimeElapsed)
	test.Equals(t, "Invalid status code", restResp.StatusCode, response.StatusCode)
	expectedBody := make(map[string]interface{})
	expectedBody["hello"] = "world"
	test.Equals(t, "Invalid body", expectedBody, response.Body)
}

func TestNewResponseInvalidBody(t *testing.T) {
	restResp := rest.Response{
		StatusCode: 200,
		Body:       `{ "hello": "world"`,
		Headers: map[string][]string{
			"Accept": {"toto"},
		},
	}

	expectedDuration := time.Duration(1 * time.Second)
	_, err := model.NewResponse(restResp, expectedDuration)
	test.Ko(t, err)
}

func TestNewResponseEmptyBody(t *testing.T) {
	restResp := rest.Response{
		StatusCode: 200,
		Body:       "",
		Headers: map[string][]string{
			"Accept": {"toto"},
		},
	}

	expectedDuration := time.Duration(1 * time.Second)
	response, err := model.NewResponse(restResp, expectedDuration)
	test.Ok(t, err)
	test.Equals(t, "Invalid duration", expectedDuration, response.TimeElapsed)
	test.Equals(t, "Invalid status code", restResp.StatusCode, response.StatusCode)
	test.Equals(t, "Body should be an empty map", 0, len(response.Body))
}
