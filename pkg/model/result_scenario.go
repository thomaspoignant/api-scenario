package model

type ScenarioResult struct {
	Name        string       `json:"name,omitempty"`
	Version     string       `json:"version,omitempty"`
	Description string       `json:"description,omitempty"`
	StepResults []ResultStep `json:"step_results,omitempty"`
}

// IsSuccess check if the scenario was success.
func (scenario *ScenarioResult) IsSuccess() bool {
	for _, stepResult := range scenario.StepResults {
		if !stepResult.IsSuccess() {
			return false
		}
	}
	return true
}
