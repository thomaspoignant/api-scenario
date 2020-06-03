package controller_test

import (
	"errors"
	"github.com/thomaspoignant/api-scenario/pkg/controller"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"strings"
	"testing"
)

type MockStepController struct {
	wantedResponse int
}

func (m MockStepController) Run(step model.Step) (model.ResultStep, error) {

	switch m.wantedResponse {
	case 1:
		return model.ResultStep{
			StepType: model.Pause,
			StepTime: 5,
		}, nil
	case 2:
		return model.ResultStep{
			StepType: model.RequestStep,
			StepTime: 5,
			Assertions: []model.ResultAssertion{
				{Success: false},
			},
		}, nil
	case 3:
		return model.ResultStep{}, errors.New("RequestXXX is an invalid step_type")
	}
	return model.ResultStep{}, nil
}

func TestScenarioError(t *testing.T) {

	test.SetupLog()
	scenario := model.Scenario{
		Name:    "Test Scenario",
		Version: "1.0",
		Steps: []model.Step{
			{StepType: model.Pause, Duration: 5},
		},
		Description: "This is a test scenario",
	}

	ctrl := controller.NewScenarioController(MockStepController{2})
	var got model.ScenarioResult
	output := test.CaptureOutput(func() {
		got = ctrl.Run(scenario)
	})

	test.Equals(t, "Name should be the same", scenario.Name, got.Name)
	test.Equals(t, "Description should be the same", scenario.Description, got.Description)
	test.Equals(t, "Version should be the same", scenario.Version, got.Version)
	test.Equals(t, "There is no error scenario should be a success", false, got.IsSuccess())
	test.Equals(t, "Should have the same number of step", len(scenario.Steps), len(got.StepResults))

	wantedOutput := "Running api-scenario: Test Scenario (1.0)\nThis is a test scenario\n\n"
	test.Equals(t, "Output should be equals", wantedOutput, output)
}

func TestScenarioSuccess(t *testing.T) {

	test.SetupLog()
	scenario := model.Scenario{
		Name:    "Test Scenario",
		Version: "1.0",
		Steps: []model.Step{
			{StepType: model.Pause, Duration: 5},
		},
		Description: "This is a test scenario",
	}

	ctrl := controller.NewScenarioController(MockStepController{1})
	var got model.ScenarioResult
	output := test.CaptureOutput(func() {
		got = ctrl.Run(scenario)
	})

	test.Equals(t, "Name should be the same", scenario.Name, got.Name)
	test.Equals(t, "Description should be the same", scenario.Description, got.Description)
	test.Equals(t, "Version should be the same", scenario.Version, got.Version)
	test.Equals(t, "There is no error scenario should be a success", true, got.IsSuccess())
	test.Equals(t, "Should have the same number of step", len(scenario.Steps), len(got.StepResults))

	wantedOutput := "Running api-scenario: Test Scenario (1.0)\nThis is a test scenario\n\n"
	test.Equals(t, "Output should be equals", wantedOutput, output)
}

func TestScenarioImpossible(t *testing.T) {

	test.SetupLog()
	scenario := model.Scenario{
		Name:    "Test Scenario",
		Version: "1.0",
		Steps: []model.Step{
			{StepType: model.Pause, Duration: 5},
		},
		Description: "This is a test scenario",
	}

	ctrl := controller.NewScenarioController(MockStepController{3})
	var got model.ScenarioResult
	output := test.CaptureOutput(func() {
		got = ctrl.Run(scenario)
	})

	test.Equals(t, "Name should be the same", scenario.Name, got.Name)
	test.Equals(t, "Description should be the same", scenario.Description, got.Description)
	test.Equals(t, "Version should be the same", scenario.Version, got.Version)
	test.Equals(t, "There is no error scenario should be a success", true, got.IsSuccess())
	test.Equals(t, "Should not have step results cause it is ignored", 0, len(got.StepResults))

	wantedPrefix := "Running api-scenario: Test Scenario (1.0)\nThis is a test scenario\n\nimpossible to execute the step: RequestXXX is an invalid step_type"
	test.Assert(t, strings.HasPrefix(output, wantedPrefix), "Output should starts with %v and got %v", wantedPrefix, output)
}
