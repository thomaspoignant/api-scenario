package model

import (
	"github.com/thomaspoignant/api-scenario/pkg/util"
)

type ResultVariableType int

//go:generate enumer -type=ResultVariableType -json -linecomment -output resultvariabletype_generated.go
const (
	Created ResultVariableType = iota //Request
	Used                              //Set
)

type ResultVariable struct {
	Key      string
	NewValue string
	Err      error
	Type     ResultVariableType
}

func (rv *ResultVariable) Print() {
	if rv.Err == nil {
		util.PrintC(util.Green, "\u2713\t")
	} else {
		util.PrintC(util.Red, "X\t")
	}

	util.Printf("%s '%s' set to '%s'", rv.Type, rv.Key, rv.NewValue)
	if rv.Err != nil {
		util.Printf(" - %s", rv.Err.Error())
	} else {
		util.Print("\n")
	}

}
