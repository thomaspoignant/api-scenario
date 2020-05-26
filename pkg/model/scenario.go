package model

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"strings"
)

type Scenario struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	ExportedAt  int    `json:"exported_at"`
	Steps       []Step `json:"steps"`
	Description string `json:"description"`
}

// InitScenarioFromFile creates a scenario from the input file.
func InitScenarioFromFile(inputFile string) (Scenario, error) {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return Scenario{}, fmt.Errorf("Impossible to locate the file: %s\n Error: %v", inputFile, err)
	}

	// Unmarshall file to launch the scenario.
	data := Scenario{}

	// Check if we have a yaml or json scenario
	if strings.HasSuffix(inputFile, ".yml") || strings.HasSuffix(inputFile, ".yaml") {
		err = yaml.Unmarshal([]byte(file), &data)
	} else {
		err = json.Unmarshal([]byte(file), &data)
	}

	if err != nil {
		return Scenario{}, fmt.Errorf("Impossible to read file: %s\n%v", inputFile, err)
	}
	return data, nil
}
