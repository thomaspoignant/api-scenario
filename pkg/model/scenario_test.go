package model_test

import (
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"reflect"
	"testing"
)

func TestInitScenarioFromFile(t *testing.T) {
	type args struct {
		inputFile string
	}

	expected := model.Scenario{
		Name:        "Simple API Test Example",
		Description: "A full description ...",
		Version:     "1.0",
		Steps: []model.Step{
			{
				StepType: model.Pause,
				Duration: 1,
			},
			{
				StepType: model.RequestStep,
				Method:   "GET",
				URL:      "https://reqres.in/api/users",
				Variables: []model.Variable{
					{
						Source:   model.ResponseJson,
						Property: "data[0].id",
						Name:     "user_id",
					},
				},
				Headers: map[string][]string{
					"Content-Type": {"application/json"},
				},
				Assertions: []model.Assertion{
					{
						Comparison: model.EqualNumber,
						Value:      "200",
						Source:     model.ResponseStatus,
					},
				},
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    model.Scenario
		wantErr bool
	}{
		{"Yaml valid test", args{inputFile: "../../testdata/scenario_valid_yaml.yml"}, expected, false},
		{"JSON valid test", args{inputFile: "../../testdata/scenario_valid_json.json"}, expected, false},
		{"File does not exist", args{inputFile: "../../testdata/does_not_exist.json"}, model.Scenario{}, true},
		{"Yaml invalid test", args{inputFile: "../../testdata/scenario_invalid_yaml.yml"}, model.Scenario{}, true},
		{"JSON invalid test", args{inputFile: "../../testdata/scenario_invalid_json.json"}, model.Scenario{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.InitScenarioFromFile(tt.args.inputFile)
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
