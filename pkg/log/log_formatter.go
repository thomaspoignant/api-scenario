package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var SuccessColor = color.New(color.FgGreen)
var ErrorColor = color.New(color.FgRed)

type OutputFormatter struct {
	DisableColors bool
}

func (f *OutputFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.DisableColors {
		SuccessColor.DisableColor()
		ErrorColor.DisableColor()
	}
	return []byte(fmt.Sprintln(entry.Message)), nil
}