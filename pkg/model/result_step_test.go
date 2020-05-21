package model_test

import (
	"fmt"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)

func Test_ResultStepIsSuccess_Pause(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.Pause,
	}

	got := resStep.IsSuccess()
	test.Equals(t, "Pause should always be success", true, got)
}

func Test_ResultStepIsSuccess_Request_SuccessEmpty(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
	}

	got := resStep.IsSuccess()
	test.Equals(t, "Empty request should always be success", true, got)
}

func Test_ResultStepIsSuccess_Request_Success(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
		VariableCreated: []model.ResultVariable{
			{Err: nil},
		},
		VariableApplied: []model.ResultVariable{
			{Err: nil},
		},
		Assertion: []model.ResultAssertion{
			{Success: true},
		},
	}

	got := resStep.IsSuccess()
	test.Equals(t, "No error should be success", true, got)
}

func Test_ResultStepIsSuccess_Request_AssertionFailed(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
		VariableCreated: []model.ResultVariable{
			{Err: nil},
		},
		VariableApplied: []model.ResultVariable{
			{Err: nil},
		},
		Assertion: []model.ResultAssertion{
			{Success: false},
		},
	}

	got := resStep.IsSuccess()
	test.Equals(t, "No error should be success", false, got)
}

func Test_ResultStepIsSuccess_Request_VariableCreatedFailed(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
		VariableCreated: []model.ResultVariable{
			{Err: fmt.Errorf("random error")},
		},
		VariableApplied: []model.ResultVariable{
			{Err: nil},
		},
		Assertion: []model.ResultAssertion{
			{Success: true},
		},
	}

	got := resStep.IsSuccess()
	test.Equals(t, "VariableCreated error should be on error", false, got)
}

func Test_ResultStepIsSuccess_Request_VariableAppliedFailed(t *testing.T) {
	resStep := model.ResultStep{
		StepType: model.RequestStep,
		VariableCreated: []model.ResultVariable{
			{Err: nil},
		},
		VariableApplied: []model.ResultVariable{

			{Err: fmt.Errorf("random error")},
		},
		Assertion: []model.ResultAssertion{
			{Success: true},
		},
	}

	got := resStep.IsSuccess()
	test.Equals(t, "VariableApplied error should be on error", false, got)
}
