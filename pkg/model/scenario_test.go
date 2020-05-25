package model

import (
	"reflect"
	"testing"
)

func TestInitScenarioFromFile(t *testing.T) {
	type args struct {
		inputFile string
	}

	expected := Scenario{
		Name:        "Simple API Test Example",
		Description: "A full description ...",
		Version:     "1.0",
		Steps: []Step{
			{
				StepType: Pause,
				Duration: 1,
			},
			{
				StepType: RequestStep,
				Method:   "GET",
				URL:      "https://reqres.in/api/users",
				Variables: []Variable{
					{
						Source:   ResponseJson,
						Property: "data[0].id",
						Name:     "user_id",
					},
				},
				Headers: map[string][]string{
					"Content-Type": {"application/json"},
				},
				Assertions: []Assertion{
					{
						Comparison: EqualNumber,
						Value:      "200",
						Source:     ResponseStatus,
					},
				},
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    Scenario
		wantErr bool
	}{
		{"Yaml valid test", args{inputFile: "../../testdata/scenario_valid_yaml.yml"}, expected, false},
		{"JSON valid test", args{inputFile: "../../testdata/scenario_valid_json.json"}, expected, false},
		{"File does not exist", args{inputFile: "../../testdata/does_not_exist.json"}, Scenario{}, true},
		{"Yaml invalid test", args{inputFile: "../../testdata/scenario_invalid_yaml.yml"}, Scenario{}, true},
		{"JSON invalid test", args{inputFile: "../../testdata/scenario_invalid_json.json"}, Scenario{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitScenarioFromFile(tt.args.inputFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitScenarioFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitScenarioFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
