package model_test

import (
	"fmt"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)

func TestResultStepIsSuccessPause(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.Pause,
	}

	got := resStep.IsSuccess()
	test.Equals(t, "Pause should always be success", true, got)
}

func TestResultStepIsSuccessRequestSuccessEmpty(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
	}

	got := resStep.IsSuccess()
	test.Equals(t, "Empty request should always be success", true, got)
}

func TestResultStepIsSuccessRequestSuccess(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
		VariablesCreated: []model.ResultVariable{
			{Err: nil},
		},
		VariablesApplied: []model.ResultVariable{
			{Err: nil},
		},
		Assertions: []model.ResultAssertion{
			{Success: true},
		},
	}

	got := resStep.IsSuccess()
	test.Equals(t, "No error should be success", true, got)
}

func TestResultStepIsSuccessRequestAssertionFailed(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
		VariablesCreated: []model.ResultVariable{
			{Err: nil},
		},
		VariablesApplied: []model.ResultVariable{
			{Err: nil},
		},
		Assertions: []model.ResultAssertion{
			{Success: false},
		},
	}

	got := resStep.IsSuccess()
	test.Equals(t, "No error should be success", false, got)
}

func TestResultStepIsSuccessRequestVariableCreatedFailed(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
		VariablesCreated: []model.ResultVariable{
			{Err: fmt.Errorf("random error")},
		},
		VariablesApplied: []model.ResultVariable{
			{Err: nil},
		},
		Assertions: []model.ResultAssertion{
			{Success: true},
		},
	}

	got := resStep.IsSuccess()
	test.Equals(t, "VariablesCreated error should be on error", false, got)
}

func TestResultStepIsSuccessRequestVariableAppliedFailed(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
		VariablesCreated: []model.ResultVariable{
			{Err: nil},
		},
		VariablesApplied: []model.ResultVariable{

			{Err: fmt.Errorf("random error")},
		},
		Assertions: []model.ResultAssertion{
			{Success: true},
		},
	}

	got := resStep.IsSuccess()
	test.Equals(t, "VariablesApplied error should be on error", false, got)
}
