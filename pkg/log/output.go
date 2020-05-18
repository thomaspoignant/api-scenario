package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = &logrus.Logger{
	Out:   os.Stderr,
	Level: logrus.InfoLevel,
	Formatter: &OutputFormatter{
		DisableColors: false,
	},
}

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