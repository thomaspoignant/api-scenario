package model_test

import (
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)


func Test_Success(t *testing.T) {
	test.SetupLog()
	want := "✓\tstatus - '1' was not equal to 20\n"
	res := model.NewResultAssertion(model.NotEqual, true, "1", "20")
	got := test.CaptureOutput(res.Print)
	test.Equals(t, "", want, got)
}

func Test_Error(t *testing.T) {
	test.SetupLog()
	want := "X\tstatus - '1' was equal to 1\n"
	res := model.NewResultAssertion(model.NotEqual, false, "1", "1")
	got := test.CaptureOutput(res.Print)
	test.Equals(t, "", want, got)
}
