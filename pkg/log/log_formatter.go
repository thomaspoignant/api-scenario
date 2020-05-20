package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var SuccessColor = color.New(color.FgGreen)
var errorColor = color.New(color.FgRed)

type OutputFormatter struct {
	DisableColors bool
}

// Format is used by logrus to print logs in stdout. This formatter is basic and just print the message.
func (f *OutputFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.DisableColors {
		SuccessColor.DisableColor()
		errorColor.DisableColor()
	}

	if entry.Level <= logrus.ErrorLevel {
		return []byte(errorColor.Sprintln(entry.Message)), nil
	}
	return []byte(fmt.Sprintln(entry.Message)), nil
}