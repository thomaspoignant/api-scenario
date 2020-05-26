//+build wireinject

package controller

import "github.com/google/wire"

func InitializeScenarioController() (ScenarioController, error) {
	wire.Build(NewScenarioController, NewRestClient, NewStepController, NewAssertionController)
	return &scenarioControllerImpl{}, nil
}