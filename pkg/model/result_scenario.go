package model

type ScenarioResult struct {
	Name        string       `json:"name"`
	Version     string       `json:"version"`
	Description string       `json:"description"`
	StepResults []ResultStep `json:"stepResults"`
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
