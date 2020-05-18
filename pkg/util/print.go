package util

import (
	"github.com/fatih/color"
)

var Red = color.New(color.FgRed)
var Green = color.New(color.FgGreen)
var noColor = color.New(color.Reset)
var isQuiet = false

func DisableColor() {
	Red.DisableColor()
	Green.DisableColor()
	noColor.DisableColor()
}

func DisableOutput() {
	isQuiet = true
}

func PrintC(c *color.Color, a ...interface{}) (n int, err error) {
	if isQuiet {
		return 0, nil
	}
	return c.Print(a...)
}

func PrintfC(c *color.Color, format string, a ...interface{}) (n int, err error) {
	if isQuiet {
		return 0, nil
	}
	return c.Printf(format, a...)
}

func PrintlnC(c *color.Color, a ...interface{}) (n int, err error) {
	if isQuiet {
		return 0, nil
	}
	return c.Println(a...)
}

func Print(a ...interface{}) (n int, err error) {
	return PrintC(Red, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return PrintlnC(noColor, a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return PrintfC(noColor, format, a...)
}
