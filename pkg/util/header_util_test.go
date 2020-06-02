package util_test

import (
	"github.com/thomaspoignant/api-scenario/pkg/util"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)

func TestTokenWithBearer(t *testing.T) {
	want := "Bearer Token123"
	got := util.AddBearerPrefix("Bearer Token123")
	test.Equals(t, "Should be equals", want, got)
}

func TestTokenWithoutBearer(t *testing.T) {
	want := "Bearer Token123"
	got := util.AddBearerPrefix("Token123")
	test.Equals(t, "Should add bearer to the token", want, got)
}

func TestShouldTrimToken(t *testing.T) {
	want := "Bearer Token123"
	got := util.AddBearerPrefix("   Token123    ")
	test.Equals(t, "Should trim around token", want, got)
}

func TestShouldTrimTokenWithBearer(t *testing.T) {
	want := "Bearer Token123"
	got := util.AddBearerPrefix("   Bearer Token123    ")
	test.Equals(t, "Should trim before checking if started with Bearer", want, got)
}
