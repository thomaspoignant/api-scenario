package model

import (
	"github.com/sendgrid/rest"
	"time"
)

type ResultStep struct {
	// Common result for every step types
	StepType StepType      `json:"step_type"`
	StepTime time.Duration `json:"step_time,omitempty"`

	// Specific for type request
	Request          rest.Request      `json:"request,omitempty"`
	Response         Response          `json:"response,omitempty"`
	Assertions       []ResultAssertion `json:"assertions,omitempty"`
	VariablesApplied []ResultVariable  `json:"variables_applied,omitempty"`
	VariablesCreated []ResultVariable  `json:"variables_created,omitempty"`
}

// IsSuccess check if the step was a success or not.
func (step *ResultStep) IsSuccess() bool {
	for _, assert := range step.Assertions {
		if !assert.Success {
			return false
		}
	}

	variables := append(step.VariablesApplied, step.VariablesCreated...)
	for _, variable := range variables {
		if variable.Err != nil {
			return false
		}
	}
	return true
}
