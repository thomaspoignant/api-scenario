package model

import (
	"fmt"
	"github.com/fatih/color"
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
	source := sourceDisplayName[ar.Source]
	if len(ar.Property) > 0 {
		source += "." + ar.Property
	}
	if ar.Success {
		color.New(color.FgGreen).Print("\u2713\t")
	} else {
		color.New(color.FgRed).Print("X\t")
	}
	fmt.Printf("%s", source)
	if ar.Err != nil {
		fmt.Printf(" - %s \n", ar.Err)
	} else {
		fmt.Printf(" - %s \n", ar.Message)
	}
}
