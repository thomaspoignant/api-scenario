package model

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
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
	color.Green("%s", scenario.Description)

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
