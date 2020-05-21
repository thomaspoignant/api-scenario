package test

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/thomaspoignant/api-scenario/pkg/log"
	"os"
)

func SetupLog() {
	logrus.SetFormatter(&log.OutputFormatter{DisableColors: true})
	logrus.SetLevel(logrus.TraceLevel)
}

func CaptureOutput(f func()) string {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)
	f()
	logrus.SetOutput(os.Stdout)
	return buf.String()
}
