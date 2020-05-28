package model

import (
	"net/url"
)

type Step struct {
	StepType      StepType      `json:"step_type"`
	URL           string        `json:"Url,omitempty"`
	Variables     []Variable    `json:"variables,omitempty"`
	MultipartForm []interface{} `json:"multipart_form,omitempty"`

	Auth struct {
	} `json:"auth,omitempty"`
	Note       string              `json:"note,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	Assertions []Assertion         `json:"assertions,omitempty"`
	Method     string              `json:"Method,omitempty"`
	Duration   int                 `json:"duration,omitempty"`
	Body       string              `json:"body,omitempty"`
	Form       struct {
	} `json:"form,omitempty"`
}

func (step *Step) ExtractUrl() (string, map[string]string, error) {
	u, err := url.Parse(step.URL)
	if err != nil {
		return "", nil, err
	}

	baseUrl := u.Scheme + "://" + u.Host + u.Path
	convertedQueryParams := make(map[string]string)
	for key, value := range u.Query() {
		if len(value) > 0 {
			convertedQueryParams[key] = value[0]
		}
	}
	return baseUrl, convertedQueryParams, nil
}
