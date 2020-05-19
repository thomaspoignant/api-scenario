package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
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
	explanation := ""
	if rv.Type == Created {
		explanation += fmt.Sprintf("%s '%s' is set to '%s'", rv.Type, rv.Key, rv.NewValue)
	} else {
		explanation += fmt.Sprintf("%s '%s' to '%s'", rv.Type, rv.Key, rv.NewValue)
	}

	if rv.Err == nil {
		logrus.Infof(log.SuccessColor.Sprint("\u2713\t") + "%s - %s", explanation)
		return
	}
	logrus.Errorf("X\t%s\n\t- %s",explanation, rv.Err.Error())
}
