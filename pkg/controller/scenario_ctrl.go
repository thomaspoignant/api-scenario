package controller

import (
	"github.com/sirupsen/logrus"
	"github.com/thomaspoignant/api-scenario/pkg/model"
)

type ScenarioController interface {
	Run(scenario model.Scenario) model.ScenarioResult
}

func NewScenarioController(stepCtrl StepController) ScenarioController {
	return &scenarioControllerImpl{
		stepController: stepCtrl,
	}
}

type scenarioControllerImpl struct {
	stepController StepController
}

func (s *scenarioControllerImpl) Run(scenario model.Scenario) model.ScenarioResult {
	result := model.ScenarioResult{
		Name:        scenario.Name,
		Description: scenario.Description,
		Version:     scenario.Version,
		StepResults: []model.ResultStep{},
	}

	logrus.Infof("Running api-scenario: %s (%s)", scenario.Name, scenario.Version)
	logrus.Infof("%s\n", scenario.Description)

	for _, step := range scenario.Steps {
		stepRes, err := s.stepController.Run(step)
		if err != nil {
			logrus.Errorf("impossible to execute the step: %v\n%v", err, step)
			continue
		}
		result.StepResults = append(result.StepResults, stepRes)
	}

	return result
}
