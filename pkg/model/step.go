package model

type Step struct {
	StepType  StepType   `json:"step_type"`
	URL       string     `json:"Url,omitempty"`
	Variables []Variable `json:"variables,omitempty"`

	Headers    map[string][]string `json:"headers,omitempty"`
	Assertions []Assertion         `json:"assertions,omitempty"`
	Method     string              `json:"Method,omitempty"`
	Duration   int                 `json:"duration,omitempty"`
	Body       string              `json:"body,omitempty"`
}
