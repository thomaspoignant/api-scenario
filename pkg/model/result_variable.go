package model

import (
	"fmt"
	"github.com/fatih/color"
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
		color.New(color.FgGreen).Print("\u2713\t")
	} else {
		color.New(color.FgRed).Print("X\t")
	}

	fmt.Printf("%s '%s' set to '%s'", rv.Type, rv.Key, rv.NewValue)
	if rv.Err != nil {
		fmt.Printf(" - %s", rv.Err.Error())
	} else {
		fmt.Print("\n")
	}

}
