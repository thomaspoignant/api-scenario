package model_test

import (
	"fmt"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)

func TestPrintResultAssertionSuccess(t *testing.T) {
	test.SetupLog()
	want := "✓\tstatus - '1' was not equal to 20\n"
	res := model.NewResultAssertion(model.NotEqual, true, "1", "20")
	got := test.CaptureOutput(res.Print)
	test.Equals(t, "", want, got)
}

func TestPrintResultAssertionWithError(t *testing.T) {
	test.SetupLog()
	want := "X\tstatus.email - '1' was equal to 1\nrandom error\n"
	res := model.NewResultAssertion(model.NotEqual, false, "1", "1")
	res.Property = "email"
	res.Err = fmt.Errorf("random error")
	got := test.CaptureOutput(res.Print)
	test.Equals(t, "", want, got)
}
