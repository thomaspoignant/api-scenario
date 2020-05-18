package model

import (
	"errors"
	"fmt"
	"github.com/jmoiron/jsonq"
	"github.com/sendgrid/rest"
	"github.com/thomaspoignant/api-scenario/pkg/log"
	"github.com/thomaspoignant/api-scenario/pkg/model/context"
	"github.com/thomaspoignant/api-scenario/pkg/util"
	"net/url"
	"reflect"
	"strconv"
	"time"
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

func (step *Step) Run() (ResultStep, error) {
	log.Logger.Info("------------------------")
	switch step.StepType {
	case Pause:
		return step.pause(step.Duration)
	case RequestStep:
		return step.request()
	default:
		return ResultStep{}, fmt.Errorf("%s is an invalid step_type", step.StepType)
	}
}

func (step *Step) pause(numberOfSecond int) (ResultStep, error) {
	start := time.Now()
	log.Logger.Infof("Waiting for %ds", numberOfSecond)
	// compute pause time and wait
	duration := time.Duration(numberOfSecond) * time.Second
	time.Sleep(duration)

	result := ResultStep{
		StepType: Pause,
		StepTime: time.Now().Sub(start),
	}
	return result, nil
}

func (step *Step) request() (ResultStep, error) {

	// convert step to api apiReq
	apiReq, err := step.convertToRestRequest()
	if err != nil {
		return ResultStep{}, errors.New("impossible to convert the ")
	}

	// init the result
	result := ResultStep{request: apiReq}

	// apply variable on the request
	result.VariableApplied = apiReq.PatchWithContext()
	apiReq.AddHeadersFromFlags()

	// Display results
	log.Logger.Infof("%s %s", apiReq.Method, apiReq.displayUrl())
	if len(apiReq.Body) > 0 {
		log.Logger.Debugf("Body: %v", string(apiReq.Body))
	}
	if len(apiReq.Headers) > 0{
		log.Logger.Debug("Headers:")
		for key, value := range apiReq.Headers {
			log.Logger.Debugf("\t%s: %s", key, value)
		}
	}
	if len(result.VariableApplied) > 0 {
		log.Logger.Info("Variables Used:")
		printVariables(result.VariableApplied)
	}
	log.Logger.Infof("---")

	// call the API
	start := time.Now()
	res, err := rest.Send(*apiReq.Request)
	elapsed := time.Now().Sub(start)
	result.StepTime = elapsed
	if err != nil {
		return result, err
	}

	// Create a response
	response, err := NewResponse(*res, elapsed)
	if err != nil {
		return result, err
	}
	result.response = response

	// Check the assertions
	result.Assertion = assertResponse(response, step.Assertions)

	// Add variables to context
	result.VariableCreated = attachVariablesToContext(response, step.Variables)
	if len(result.VariableCreated) > 0 {
		log.Logger.Info("Variables Created:")
		printVariables(result.VariableCreated)
	}
	return result, nil
}

func (step *Step) convertToRestRequest() (Request, error) {
	baseUrl, queryParams, err := step.extractUrl()
	if err != nil {
		return Request{}, err
	}

	// Convert headers format to the api.ApiRequest format
	headers := make(map[string]string)
	for key, value := range step.Headers {
		if len(value) > 0 {
			headers[key] = value[0]
		}
	}

	return Request{
		Request: &rest.Request{
			Method:      rest.Method(step.Method),
			Headers:     headers,
			QueryParams: queryParams,
			Body:        []byte(step.Body),
			BaseURL:     baseUrl,
		},
	}, nil
}

func (step *Step) extractUrl() (string, map[string]string, error) {
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

func assertResponse(response Response, assertions []Assertion) []resultAssertion {

	if len(assertions)>0 {
		log.Logger.Info("Assertions:")
	}

	var result []resultAssertion
	for _, assertion := range assertions {
		assertionResult := assertion.Assert(response)
		result = append(result, assertionResult)
		assertionResult.Print()
	}
	return result
}

func attachVariablesToContext(response Response, vars []Variable) []ResultVariable {
	var result []ResultVariable

	for _, variable := range vars {
		if len(variable.Name) == 0 {
			break
		}

		switch variable.Source {
		case ResponseTime:
			value := strconv.FormatInt(int64(response.TimeElapsed.Round(time.Millisecond)/time.Millisecond), 10)
			context.GetContext().Add(variable.Name, value)
			result = append(result, ResultVariable{Key: variable.Name, NewValue: value, Type: Created})
			break

		case ResponseStatus:
			value := fmt.Sprintf("%v", response.StatusCode)
			context.GetContext().Add(variable.Name, value)
			result = append(result, ResultVariable{Key: variable.Name, NewValue: value, Type: Created})
			break

		case ResponseJson:
			// Convert key name
			jqPath := util.JsonConvertKeyName(variable.Property)
			jq := jsonq.NewQuery(response.Body)
			extractedKey, err := jq.Interface(jqPath[:]...)
			if err != nil {
				result = append(result, ResultVariable{Key: variable.Name, Err: err, Type: Created})
			}

			switch value := extractedKey.(type) {
			case string:
				context.GetContext().Add(variable.Name, value)
				result = append(result, ResultVariable{Key: variable.Name, NewValue: value, Type: Created})
				break
			case bool:
				castValue := strconv.FormatBool(value)
				context.GetContext().Add(variable.Name, castValue)
				result = append(result, ResultVariable{Key: variable.Name, NewValue: castValue, Type: Created})
				break
			case float64:
				castValue := fmt.Sprintf("%g", value)
				context.GetContext().Add(variable.Name, castValue)
				result = append(result, ResultVariable{Key: variable.Name, NewValue: castValue, Type: Created})
				break
			default:
				result = append(result, ResultVariable{
					Key:  variable.Name,
					Err:  fmt.Errorf("type %s not valid type to export as a variable", reflect.TypeOf(extractedKey)),
					Type: Created,
				})
				break
			}
		}
	}
	return result
}

func printVariables(variables []ResultVariable) {
	for _, currentVar := range variables {
		currentVar.Print()
	}
}
