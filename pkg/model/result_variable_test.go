package model_test

import (
	"fmt"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)

func Test_ResultVariable_Print(t *testing.T) {
	test.SetupLog()
	testCases := []struct {
		value    model.ResultVariable
		expected string
	}{
		{
			model.ResultVariable{Key: "test", NewValue: "newValue", Type: model.Used},
			"✓\tSet 'test' to 'newValue'\n",
		},
		{
			model.ResultVariable{Key: "test", NewValue: "newValue", Type: model.Created},
			"✓\tVariable 'test' is set to 'newValue'\n",
		},
		{
			model.ResultVariable{Key: "test", NewValue: "newValue", Err: fmt.Errorf("random error"), Type: model.Created},
			"X\tVariable 'test' is set to 'newValue'\n\t- random error\n",
		},
		{
			model.ResultVariable{Key: "test", NewValue: "newValue", Err: fmt.Errorf("random error"), Type: model.Used},
			"X\tSet 'test' to 'newValue'\n\t- random error\n",
		},
	}

	for _, tc := range testCases {
		got := test.CaptureOutput(tc.value.Print)
		test.Equals(t, "Output should be equals", tc.expected, got)
	}
}
