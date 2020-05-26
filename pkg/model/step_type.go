package model

type StepType int

//go:generate enumer -type=StepType -json -linecomment  -output steptype_gen.go
const (
	Pause       StepType = iota //pause
	RequestStep                 //request
)
