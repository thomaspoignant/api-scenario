package controller_test

import (
	"github.com/spf13/viper"
	"github.com/thomaspoignant/api-scenario/pkg/context"
	"testing"
	"time"

	"github.com/sendgrid/rest"
	"github.com/thomaspoignant/api-scenario/pkg/controller"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
)

// Pause
func TestStepPause(t *testing.T) {
	sc := controller.NewStepController(&test.ClientMock{}, controller.NewAssertionController())
	step := model.Step{
		StepType: model.Pause,
		Duration: 1,
	}

	start := time.Now()
	if _, err := sc.Run(step); err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	end := time.Now()

	if start.Add(1 * time.Second).After(end) {
		t.Errorf("We should wait for at least 1 seconds, start:%v - end:%v", start, end)
	}
}

func TestOutputPause(t *testing.T) {
	test.SetupLog()
	sc := controller.NewStepController(&test.ClientMock{}, controller.NewAssertionController())
	step := model.Step{
		StepType: model.Pause,
		Duration: 1,
	}

	want := "------------------------\nWaiting for 1s\n"
	got := test.CaptureOutput(func() {
		if _, err := sc.Run(step); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	test.Equals(t, "Output messages are different", want, got)
}

// Request
func TestRequestValid(t *testing.T) {
	test.SetupLog()
	testNumber := "1"
	sc := controller.NewStepController(&test.ClientMock{}, controller.NewAssertionController())

	context.GetContext().Add("baseUrl", "test.com")
	viper.Set("headers", map[string]string{
		"Accept": "application/json",
	})
	step := model.Step{
		Body: `{"hello":"world_{{random_int(1,1)}}"}`,
		URL:  "http://{{baseUrl}}/1/{{random_int(1,1)}}?param1=param1_{{random_int(1,1)}}&testNumber=" + testNumber,
		Headers: map[string][]string{
			"Content-Type": {"other_test_{{random_int(1,1)}}"},
		},
		Method:   "GET",
		StepType: model.RequestStep,
		Variables: []model.Variable{
			{
				Source:   model.ResponseJson,
				Property: "hello",
				Name:     "hello",
			},
			{
				Source:   model.ResponseJson,
				Property: "param1",
				Name:     "param1",
			},
			{
				Source:   model.ResponseJson,
				Property: "param2",
				Name:     "param2",
			},
			{
				Source: model.ResponseStatus,
				Name:   "response_status",
			},
			{
				Source: model.ResponseTime,
				Name:   "response_time",
			},
			{
				Source: model.ResponseTime,
				// if variable has no name we ignore it
				Name: "",
			},
			{
				Source:   model.ResponseHeader,
				Name:     "response_header",
				Property: "Content-Type",
			},
			{
				Source: model.ResponseText,
				Name:   "response_text",
			},
		},
		Assertions: []model.Assertion{
			{
				Comparison: model.Equal,
				Value:      "200",
				Source:     model.ResponseStatus,
			},
		},
	}
	got, _ := sc.Run(step)

	test.Equals(t, "StepType should be request", model.RequestStep, got.StepType)
	test.Assert(t, got.StepTime > 0, "StepTime should be positive")

	// Check patch on request
	test.Equals(t, "Should have patch URL", "http://test.com/1/1", got.Request.BaseURL)
	test.Equals(t, "Should return method", rest.Get, got.Request.Method)
	wantHeaders := map[string]string{
		"Content-Type": "other_test_1",
		"Accept":       "application/json",
	}
	test.Equals(t, "Should have patch headers", wantHeaders, got.Request.Headers)
	wantParams := map[string]string{
		"param1":     "param1_1",
		"testNumber": "1",
	}
	test.Equals(t, "Should have patch params", wantParams, got.Request.QueryParams)
	test.Equals(t, "Should have patch body", `{"hello":"world_1"}`, string(got.Request.Body))

	// Check response
	test.Assert(t, got.Response.TimeElapsed > 0, "TimeElapsed should be positive")
	test.Equals(t, "Should have response status = 200", 200, got.Response.StatusCode)
	test.Equals(t, "Should have 1 assertion", 1, len(got.Assertions))
	test.Equals(t, "Should have valid assertion", true, got.Assertions[0].Success)
	test.Equals(t, "Should have patch 2 elements of the request", 2, len(got.VariablesApplied))
	test.Equals(t, "Should have create 7 variables", 7, len(got.VariablesCreated))
}
