package log_test

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thomaspoignant/api-scenario/pkg/log"
	"testing"
	"time"
)

func Test_OutputFormatter(t *testing.T) {
	tf := &log.OutputFormatter{DisableColors: true}
	testCases := []struct {
		value    logrus.Entry
		expected string
	}{
		{logrus.Entry{Level: logrus.ErrorLevel, Time: time.Now(), Message: "foo"}, "foo\n"},
		{logrus.Entry{Level: logrus.InfoLevel, Time: time.Now(), Message: "foo"}, "foo\n"},
		{logrus.Entry{Level: logrus.DebugLevel, Time: time.Now(), Message: "foo"}, "foo\n"},
	}

	for _, tc := range testCases {
		b, _ := tf.Format(&tc.value)
		fmt.Print(string(b))
		if string(b) != tc.expected {
			t.Errorf("formatting expected for %v (result was %q instead of %q)", tc.value, string(b), tc.expected)
		}
	}
}