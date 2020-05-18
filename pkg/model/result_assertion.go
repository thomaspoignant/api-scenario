package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thomaspoignant/api-scenario/pkg/log"
)

type resultAssertion struct {
	Success  bool
	Source   Source
	Property string
	Err      error
	Message  string
}

func NewResultAssertion(comparison Comparison, success bool, v ...interface{}) resultAssertion {
	msg := comparison.GetMessage()
	var message string
	if success {
		message = fmt.Sprintf(msg.Success, v...)
	} else {
		message = fmt.Sprintf(msg.Failure, v...)
	}
	return resultAssertion{
		Success: success,
		Message: message,
		Err:     nil,
	}
}

var sourceDisplayName = map[Source]string{
	ResponseJson:   "body",
	ResponseTime:   "Response time",
	ResponseStatus: "status",
}

func (ar *resultAssertion) Print() {
	output := ""
	source := sourceDisplayName[ar.Source]
	if len(ar.Property) > 0 {
		source += "." + ar.Property
	}

	if ar.Success {
		output += log.SuccessColor.Sprintf("\u2713\t")
		output += fmt.Sprintf("%s - %s", source, ar.Message)
	} else {
		output += log.ErrorColor.Sprintf("X\t%s - %s", source, ar.Message)
	}

	// Add error if we have it
	if ar.Err != nil {
		output += log.ErrorColor.Sprintf(" - %s", ar.Err)
	}

	logrus.Info(output)
}
