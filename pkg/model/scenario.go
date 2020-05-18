package model

import (
	"github.com/thomaspoignant/api-scenario/pkg/util"
	"log"
)

type Scenario struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	ExportedAt  int    `json:"exported_at"`
	Steps       []Step `json:"steps"`
	Description string `json:"description"`
}

type ScenarioResult struct {
	Name        string       `json:"name"`
	Version     string       `json:"version"`
	Description string       `json:"description"`
	StepResults []ResultStep `json:step_results`
}

func (scenario *Scenario) Run() ScenarioResult {
	result := ScenarioResult{
		Name:        scenario.Name,
		Description: scenario.Description,
		Version:     scenario.Version,
		StepResults: []ResultStep{},
	}
	util.PrintfC(util.Green, "%s\n", scenario.Description)

	for _, step := range scenario.Steps {
		stepRes, err := step.Apply()
		if err != nil {
			log.Fatalf("impossible to execute the step: %v\n%v", err, step)
			break
		}
		result.StepResults = append(result.StepResults, stepRes)
	}

	return result
}
