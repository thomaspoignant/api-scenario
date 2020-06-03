package model

// An Assertions defines one expected result
type Assertion struct {
	Comparison Comparison `json:"comparison"`
	Value      string     `json:"value"`
	Source     Source     `json:"Source"`
	Property   string     `json:"property,omitempty"`
}
