package model_test

import (
	"github.com/sendgrid/rest"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"

	"github.com/spf13/viper"
)

var req model.Request

func initReq() {
	req = model.Request{
		Request: &rest.Request{
			Body:    []byte(`{"hello":"world_{{random_int(1,1)}}"}`),
			BaseURL: "http://perdu.com/{{random_int(1,1)}}",
			Headers: map[string]string{
				"Content-Type": "other_test_{{random_int(1,1)}}",
			},
			QueryParams: map[string]string{
				"param1": "param1_{{random_int(1,1)}}",
			},
			Method: rest.Get,
		},
	}
}

func Test_PatchWithContext_valid(t *testing.T) {
	initReq()
	want := []model.ResultVariable{
		{Key: "body", NewValue: "{\"hello\":\"world_1\"}", Type: model.Used},
		{Key: "url", NewValue: "http://perdu.com/1", Type: model.Used},
		{Key: "params[param1]", NewValue: "param1_1", Type: model.Used},
		{Key: "headers.Content-Type", NewValue: "other_test_1", Type: model.Used},
	}
	got := req.PatchWithContext()
	test.Equals(t, "", want, got)
}

func Test_OverrideHeadersFromViper(t *testing.T) {
	initReq()
	want := map[string]string{
		"Content-Type": "application/JSON",
		"Accept":       "application/JSON",
	}
	viper.Set("headers", want)
	req.AddHeadersFromFlags()
	got := req.Headers

	test.Equals(t, "Header should be override", want, got)
}

func Test_AddHeadersFromViper(t *testing.T) {
	initReq()
	want := map[string]string{
		"Accept":       "application/JSON",
		"Content-Type": "other_test_{{random_int(1,1)}}",
	}
	viper.Set("headers", map[string]string{
		"Accept": "application/JSON",
	})
	req.AddHeadersFromFlags()
	got := req.Headers

	test.Equals(t, "Header should be override", want, got)
}
