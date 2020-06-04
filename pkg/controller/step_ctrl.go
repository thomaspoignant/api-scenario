package controller

import (
	"errors"
	"fmt"
	"github.com/jmoiron/jsonq"
	"github.com/thomaspoignant/api-scenario/pkg/util"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/sendgrid/rest"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thomaspoignant/api-scenario/pkg/context"
	"github.com/thomaspoignant/api-scenario/pkg/model"
)

type StepController interface {
	Run(step model.Step) (model.ResultStep, error)
}

type stepControllerImpl struct {
	client        RestClient
	assertionCtrl AssertionController
}

func NewStepController(client RestClient, assertionCtrl AssertionController) StepController {
	return &stepControllerImpl{
		client:        client,
		assertionCtrl: assertionCtrl,
	}
}

// Run is running the step and assert it.
func (sc *stepControllerImpl) Run(step model.Step) (model.ResultStep, error) {

	switch step.StepType {
	case model.Pause:
		return sc.pause(step.Duration)

	case model.RequestStep:
		return sc.request(step)

	default:
		// Cannot happen, all value tested
		return model.ResultStep{}, fmt.Errorf("%s is an invalid step_type", step.StepType)
	}
}

// pause is stopping the thread during numberOfSecond seconds.
func (sc *stepControllerImpl) pause(numberOfSecond int) (model.ResultStep, error) {
	start := time.Now()
	logrus.Info("------------------------")
	logrus.Infof("Waiting for %ds", numberOfSecond)
	// compute pause time and wait
	duration := time.Duration(numberOfSecond) * time.Second
	time.Sleep(duration)

	result := model.ResultStep{
		StepType: model.Pause,
		StepTime: time.Since(start),
	}
	return result, nil
}

// request is calling a Rest HTTP endpoint and assert the response.
func (sc *stepControllerImpl) request(step model.Step) (model.ResultStep, error) {
	// convert step to api req

	req, variables, err := convertAndPatchToHttpRequest(step)
	if err != nil {
		return model.ResultStep{}, errors.New("impossible to convert the request")
	}

	// init the result
	result := model.ResultStep{
		StepType: model.RequestStep,
		Request:  req,
	}

	// apply variable on the request
	result.VariablesApplied = variables

	// Display request
	printRestRequest(req, result.VariablesApplied)

	// call the API
	start := time.Now()
	res, err := sc.client.Send(req)
	elapsed := time.Since(start)
	result.StepTime = elapsed
	if err != nil {
		return result, err
	}
	logrus.Infof("Time elapsed: %v", elapsed)

	// Create a response
	response, err := model.NewResponse(*res, elapsed)
	if err != nil {
		return result, err
	}
	result.Response = response

	// Check the assertions
	result.Assertions = sc.assertResponse(response, step.Assertions)

	// Add variables to context
	result.VariablesCreated = attachVariablesToContext(response, step.Variables)

	if len(result.VariablesCreated) > 0 {
		logrus.Info("Variables  created:")
		for _, currentVar := range result.VariablesCreated {
			currentVar.Print()
		}
	}
	return result, nil
}

// assertResponse assert the response of a REST Call.
func (sc *stepControllerImpl) assertResponse(response model.Response, assertions []model.Assertion) []model.ResultAssertion {
	if len(assertions) > 0 {
		logrus.Info("Assertions:")
	}

	var result []model.ResultAssertion
	for _, assertion := range assertions {
		assertionResult := sc.assertionCtrl.Assert(assertion, response)
		result = append(result, assertionResult)
		assertionResult.Print()
	}
	return result
}

// attachVariablesToContext extract variable from the response and add it to the context.
func attachVariablesToContext(response model.Response, vars []model.Variable) []model.ResultVariable {
	var result []model.ResultVariable

	for _, variable := range vars {
		if len(variable.Name) == 0 {
			continue
		}

		switch variable.Source {
		case model.ResponseTime:
			value := strconv.FormatInt(int64(response.TimeElapsed.Round(time.Millisecond)/time.Millisecond), 10)
			context.GetContext().Add(variable.Name, value)
			result = append(result, model.ResultVariable{Key: variable.Name, NewValue: value, Type: model.Created})

		case model.ResponseStatus:
			value := fmt.Sprintf("%v", response.StatusCode)
			context.GetContext().Add(variable.Name, value)
			result = append(result, model.ResultVariable{Key: variable.Name, NewValue: value, Type: model.Created})

		case model.ResponseHeader:
			header := response.Header[variable.Property]
			if len(header) > 0 {
				// TODO: Works fine if we have only one value for the header
				context.GetContext().Add(variable.Name, header[0])
				result = append(result, model.ResultVariable{Key: variable.Name, NewValue: header[0], Type: model.Created})
			}

		case model.ResponseText:
			context.GetContext().Add(variable.Name, response.Body)
			result = append(result, model.ResultVariable{Key: variable.Name, NewValue: response.Body, Type: model.Created})

		case model.ResponseJson:
			result = append(result, attachVariablesFromResponseJson(variable, response))
		}
	}
	return result
}

// attachVariablesFromResponseJson extract variable from the JSON response and add it to the context.
func attachVariablesFromResponseJson(variable model.Variable, response model.Response) model.ResultVariable {
	// Convert key name
	jqPath := util.JsonConvertKeyName(variable.Property)

	// Convert body to map[string]interface{}
	body, err := util.StringToJson(response.Body)
	if err != nil {
		return model.ResultVariable{Key: variable.Name, Err: err, Type: model.Created}
	}

	jq := jsonq.NewQuery(body)
	extractedKey, err := jq.Interface(jqPath[:]...)
	if err != nil {
		return model.ResultVariable{Key: variable.Name, Err: err, Type: model.Created}
	}

	switch value := extractedKey.(type) {
	case string:
		context.GetContext().Add(variable.Name, value)
		return model.ResultVariable{Key: variable.Name, NewValue: value, Type: model.Created}

	case bool:
		castValue := strconv.FormatBool(value)
		context.GetContext().Add(variable.Name, castValue)
		return model.ResultVariable{Key: variable.Name, NewValue: castValue, Type: model.Created}

	case float64:
		castValue := fmt.Sprintf("%g", value)
		context.GetContext().Add(variable.Name, castValue)
		return model.ResultVariable{Key: variable.Name, NewValue: castValue, Type: model.Created}

	default:
		return model.ResultVariable{
			Key:  variable.Name,
			Err:  fmt.Errorf("type %s not valid type to export as a variable", reflect.TypeOf(extractedKey)),
			Type: model.Created,
		}
	}
}

// convertAndPatchToHttpRequest create the HTTP request to call.
func convertAndPatchToHttpRequest(step model.Step) (rest.Request, []model.ResultVariable, error) {

	var result []model.ResultVariable
	urlPatched := context.GetContext().Patch(step.URL)
	baseUrl, queryParams, err := extractUrl(urlPatched)
	if err != nil {
		return rest.Request{}, result, err
	}

	// Convert headers format to the api.ApiRequest format
	headers := make(map[string]string)
	for key, value := range step.Headers {
		if len(value) > 0 {
			headers[key] = value[0]
		}
	}

	// Add headers from command line.
	// It can override existing headers.
	for key, value := range viper.GetStringMapString("headers") {
		headers[key] = context.GetContext().Patch(value)
	}

	// Patches
	bodyPatched := patchVariable(step.Body, "body", &result)
	for key, value := range headers {
		headers[key] = patchVariable(value, "headers."+key, &result)
	}

	return rest.Request{
		Method:      rest.Method(step.Method),
		Headers:     headers,
		QueryParams: queryParams,
		Body:        []byte(bodyPatched),
		BaseURL:     baseUrl,
	}, result, nil
}

// extractUrl extract URL and params to have URL and params in 2 variables.
func extractUrl(urlToParse string) (string, map[string]string, error) {
	u, err := url.Parse(urlToParse)
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

// patchVariable is applying a patch with the context on the "initial" string and also
// update the slice of 'variables"
func patchVariable(initial string, name string, variables *[]model.ResultVariable) string {
	initialValue := string(initial)
	patchedValue := context.GetContext().Patch(initial)

	if initialValue != patchedValue {
		*variables = append(*variables, model.ResultVariable{
			Type:     model.Used,
			NewValue: patchedValue,
			Key:      name,
		})
	}

	return patchedValue
}

// printRestRequest is logging a user friendly description of the request.
func printRestRequest(req rest.Request, appliedVar []model.ResultVariable) {
	logrus.Info("------------------------")
	// Compose URL
	params := ""
	for key, value := range req.QueryParams {
		if len(params) == 0 {
			params += "?"
		} else {
			params += "&"
		}
		params += fmt.Sprintf("%s=%s", key, value)
	}
	url := req.BaseURL + params

	logrus.Infof("%s %s", req.Method, url)
	if len(req.Body) > 0 {
		logrus.Debugf("Body: %v", string(req.Body))
	}
	if len(req.Headers) > 0 {
		logrus.Debug("Headers:")
		for key, value := range req.Headers {
			logrus.Debugf("\t%s: %s", key, value)
		}
	}
	if len(appliedVar) > 0 {
		logrus.Info("Variables Used:")
		for _, currentVar := range appliedVar {
			currentVar.Print()
		}
	}
	logrus.Infof("---")
}
