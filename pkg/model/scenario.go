package model

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type Scenario struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	ExportedAt  int    `json:"exported_at"`
	Steps       []Step `json:"steps"`
	Description string `json:"description"`
}
func (scenario *Scenario) Run() ScenarioResult {
	result := ScenarioResult{
		Name:        scenario.Name,
		Description: scenario.Description,
		Version:     scenario.Version,
		StepResults: []ResultStep{},
	}


	logrus.Infof("Running api-scenario: %s (%s)", scenario.Name, scenario.Version)
	logrus.Infof("%s\n", scenario.Description)

	for _, step := range scenario.Steps {
		stepRes, err := step.Run()
		if err != nil {
			logrus.Fatalf("impossible to execute the step: %v\n%v", err, step)
			break
		}
		result.StepResults = append(result.StepResults, stepRes)
	}

	return result
}

func InitScenarioFromFile(inputFile string) (Scenario, error){
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return Scenario{}, fmt.Errorf("Impossible to locate the file: %s\n Error: %v", inputFile, err)
	}

	// Unmarshall file to launch the scenario.
	data := Scenario{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return Scenario{}, fmt.Errorf("Impossible to read file: %s\n%v", inputFile, err)
	}

	return data, nil
}

type ScenarioResult struct {
	Name        string       `json:"name"`
	Version     string       `json:"version"`
	Description string       `json:"description"`
	StepResults []ResultStep `json:step_results`
}

func (scenario *ScenarioResult) IsSuccess() bool {
	for _, stepResult := range scenario.StepResults {
		if !stepResult.IsSuccess() {
			return false
		}
	}
	return true
}
