package config_helper_test

import (
	"github.com/thomaspoignant/api-scenario/pkg/config_helper"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)

func Test_TokenWithBearer(t *testing.T) {
	want := "Bearer Token123"
	got := config_helper.FormatAuthorization("Bearer Token123")
	test.Equals(t, "Should be equals", want, got)
}

func Test_TokenWithoutBearer(t *testing.T) {
	want := "Bearer Token123"
	got := config_helper.FormatAuthorization("Token123")
	test.Equals(t, "Should add bearer to the token", want, got)
}

func Test_ShouldTrimToken(t *testing.T) {
	want := "Bearer Token123"
	got := config_helper.FormatAuthorization("   Token123    ")
	test.Equals(t, "Should trim around token", want, got)
}

func Test_ShouldTrimTokenWithBearer(t *testing.T) {
	want := "Bearer Token123"
	got := config_helper.FormatAuthorization("   Bearer Token123    ")
	test.Equals(t, "Should trim before checking if started with Bearer", want, got)
}

