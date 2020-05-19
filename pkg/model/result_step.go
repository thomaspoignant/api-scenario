package model

import (
	"time"
)

type ResultStep struct {
	// Common result for every step types
	StepType StepType
	StepTime time.Duration
	Warning  string

	// Specific for type request
	request         Request
	response        Response
	Assertion       []ResultAssertion
	VariableApplied []ResultVariable
	VariableCreated []ResultVariable
}

func (step *ResultStep) IsSuccess() bool {
	for _, assert := range step.Assertion {
		if !assert.Success {
			return false
		}
	}

	variables := append(step.VariableApplied, step.VariableCreated...)
	for _, variable := range variables {
		if variable.Err != nil {
			return false
		}
	}
	return true
}