package model

import (
	"fmt"
	"github.com/thomaspoignant/api-scenario/pkg/log"
)

type ResultVariableType int

//go:generate enumer -type=ResultVariableType -json -linecomment -output resultvariabletype_generated.go
const (
	Created ResultVariableType = iota //Variable
	Used                              //Set
)

type ResultVariable struct {
	Key      string
	NewValue string
	Err      error
	Type     ResultVariableType
}

func (rv *ResultVariable) Print() {
	output := ""
	if rv.Err == nil {
		output += log.SuccessColor.Sprint("\u2713\t")
	} else {
		output += log.ErrorColor.Sprint("X\t")
	}

	if rv.Type == Created {
		output += fmt.Sprintf("%s '%s' is set to '%s'", rv.Type, rv.Key, rv.NewValue)
	} else {
		output += fmt.Sprintf("%s '%s' to '%s'", rv.Type, rv.Key, rv.NewValue)
	}

	if rv.Err != nil {
		output += log.ErrorColor.Sprintf(" - %s", rv.Err.Error())
	}

	log.Logger.Info(output)
}
