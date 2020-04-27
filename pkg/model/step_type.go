package model

type StepType int

//go:generate enumer -type=StepType -json -linecomment  -output steptype_generated.go
const (
	Pause       StepType = iota //pause
	RequestStep                 //request
)
