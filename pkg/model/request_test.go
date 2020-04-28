package model_test

import (
	"github.com/sendgrid/rest"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)

func Test_PatchWithContext_valid(t *testing.T) {
	req := model.Request{
		Request: &rest.Request{
			Body:    []byte(`{"hello":"world_{{random_int(1,1)}}"}`),
			BaseURL: "http://perdu.com/{{random_int(1,1)}}",
			Headers: map[string]string{
				"Content-Type": "other_test_{{random_int(1,1)}}",
			},
			QueryParams: map[string]string{
				"param1": "param1_{{random_int(1,1)}}",
				"param2": "param2_{{random_int(1,1)}}",
			},
			Method: rest.Get,
		},
	}

	want := []model.ResultVariable{
		{Key: "body", NewValue: "{\"hello\":\"world_1\"}", Type: model.Used},
		{Key: "url", NewValue: "http://perdu.com/1", Type: model.Used},
		{Key: "params[param1]", NewValue: "param1_1", Type: model.Used},
		{Key: "params[param2]", NewValue: "param2_1", Type: model.Used},
		{Key: "headers.Content-Type", NewValue: "other_test_1", Type: model.Used},
	}
	got := req.PatchWithContext()

	test.Equals(t, "", want, got)

}
