package model

import (
	"github.com/sendgrid/rest"
	"github.com/thomaspoignant/api-scenario/pkg/model/context"
)

type Request struct {
	*rest.Request
}

func (req *Request) PatchWithContext() []ResultVariable {
	var result []ResultVariable
	// Patch body
	bodyAsStr := string(req.Body)
	bodyPatched := context.GetContext().Patch(bodyAsStr)
	req.Body = []byte(bodyPatched)
	if bodyAsStr != bodyPatched {
		result = append(result, ResultVariable{
			Key:      "body",
			NewValue: bodyPatched,
			Type:     Used,
		})
	}

	// Patch URL
	urlBefore := req.BaseURL
	req.BaseURL = context.GetContext().Patch(req.BaseURL)
	if urlBefore != req.BaseURL {
		result = append(result, ResultVariable{
			Key:      "url",
			NewValue: req.BaseURL,
			Type:     Used,
		})
	}

	// Patch query params
	for key, value := range req.QueryParams {
		paramValue := value
		req.QueryParams[key] = context.GetContext().Patch(value)
		if paramValue != req.QueryParams[key] {
			result = append(result, ResultVariable{
				Key:      "params[" + key + "]",
				NewValue: req.QueryParams[key],
				Type:     Used,
			})
		}
	}

	// Patch headers
	for key, value := range req.Headers {
		headerValue := value
		req.Headers[key] = context.GetContext().Patch(value)
		if headerValue != req.Headers[key] {
			result = append(result, ResultVariable{
				Key:      "headers." + key,
				NewValue: req.Headers[key],
				Type:     Used,
			})
		}
	}
	return result
}
