package controller

import (
	"errors"
	"fmt"
	"github.com/jmoiron/jsonq"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/pkg/util"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type AssertionController interface {
	Assert(assertion model.Assertion, resp model.Response) model.ResultAssertion
}

type assertionControllerImpl struct {
}

func NewAssertionController() AssertionController {
	return &assertionControllerImpl{}
}

const ComparisonNotSupportedMessage = "the comparison %s was not supported for the source"

// Assert is testing an assertion on a API response.
func (ctrl *assertionControllerImpl) Assert(assertion model.Assertion, resp model.Response) model.ResultAssertion {

	switch assertion.Source {
	case model.ResponseStatus:
		res := ctrl.assertNumber(assertion, float64(resp.StatusCode))
		res.Source = assertion.Source
		return res

	case model.ResponseTime:
		apiTime := float64(resp.TimeElapsed) / float64(time.Second)
		res := ctrl.assertNumber(assertion, apiTime)
		res.Source = assertion.Source
		return res

	case model.ResponseJson:
		res := ctrl.assertResponseJson(assertion, resp.Body)
		res.Source = assertion.Source
		return res

	case model.ResponseHeader:
		res := ctrl.assertResponseHeader(assertion, resp.Header)
		res.Source = assertion.Source
		return res

	case model.ResponseText:
		var res model.ResultAssertion

		if util.IsNumeric(resp.Body) {
			bodyAsFloat, _ := strconv.ParseFloat(resp.Body, 64)
			res = ctrl.assertNumber(assertion, bodyAsFloat)
		} else {
			res = ctrl.assertString(assertion, resp.Body)
		}

		res.Source = assertion.Source
		return res

	default:
		message := fmt.Sprintf("the Source %s is not valid", assertion.Source)
		return model.ResultAssertion{Success: false, Err: errors.New(message), Message: message, Source: assertion.Source}
	}
}

// assertResponseHeader is testing an assertion on the HTTP headers.
func (ctrl *assertionControllerImpl) assertResponseHeader(assertion model.Assertion, h http.Header) model.ResultAssertion {

	//search for header using canonical key format
	values := h.Values(assertion.Property)
	if values == nil {
		if assertion.Comparison == model.IsNull {
			result := model.NewResultAssertion(model.IsNull, true, assertion.Property)
			result.Property = assertion.Property
			return result
		}
		message := fmt.Sprintf("Header %q not found.", assertion.Property)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message), Property: assertion.Property}
	}

	if assertion.Comparison == model.HasKey {
		result := model.NewResultAssertion(model.HasKey, true, assertion.Property)
		result.Property = assertion.Property
		return result
	}

	// Compare first value of the given header key
	// TODO handle many values for same key
	result := ctrl.assertValue(assertion, h.Get(assertion.Property))
	result.Property = assertion.Property
	return result
}

// assertResponseHeader is testing an assertion on the JSON body of the response.
func (ctrl *assertionControllerImpl) assertResponseJson(assertion model.Assertion, bodyAsString string) model.ResultAssertion {

	if len(bodyAsString) > 0 && !util.IsJson(bodyAsString) {
		message := "there is a result and this is not a valid JSON api Response is not in JSON"
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message), Property: assertion.Property}
	}

	body, err := util.StringToJson(bodyAsString)
	if err != nil {
		return model.ResultAssertion{Success: false, Message: err.Error(), Err: err, Property: assertion.Property}
	}

	// Convert property from Json syntax to an array of fields
	jqPath := util.JsonConvertKeyName(assertion.Property)

	// Init jsonq library with the response data
	jq := jsonq.NewQuery(body)

	// Get the key in the data
	extractedKey, err := jq.Interface(jqPath[:]...)
	if err != nil {
		//If the key if not found and we are looking for is_null this is assertion Success
		if assertion.Comparison == model.IsNull {
			result := model.NewResultAssertion(model.IsNull, true, assertion.Property)
			result.Property = assertion.Property
			return result
		}
		message := fmt.Sprintf("Unable to locate %s property in path '%s' in JSON", assertion.Property, assertion.Property)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message), Property: assertion.Property}
	}

	result := ctrl.assertValue(assertion, extractedKey)
	result.Property = assertion.Property
	return result
}

// assertValue determine the type of the result and use the correct assert method.
func (ctrl *assertionControllerImpl) assertValue(assertion model.Assertion, apiValue interface{}) model.ResultAssertion {
	switch apiValue := apiValue.(type) {
	case string:
		return ctrl.assertString(assertion, apiValue)
	case bool:
		return ctrl.assertBool(assertion, apiValue)
	case float64:
		return ctrl.assertNumber(assertion, apiValue)
	case []interface{}:
		return ctrl.assertArray(assertion, apiValue)
	case map[string]interface{}:
		return ctrl.assertMap(assertion, apiValue)
	default:
		// Not supposed to happen
		message := fmt.Sprintf("%s comparison is not available for element of type %s", assertion.Comparison, reflect.TypeOf(apiValue))
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

// assertNumber is testing the assertion "assertion" with the value "apiValue"
func (ctrl *assertionControllerImpl) assertNumber(assertion model.Assertion, apiValue float64) model.ResultAssertion {
	comparison := assertion.Comparison
	assertionValue := assertion.Value

	switch comparison {
	case model.IsANumber:
		return model.NewResultAssertion(comparison, true, apiValue)

	case model.Equal:
		apiValueAsString := fmt.Sprintf("%g", apiValue)
		success := assertionValue == apiValueAsString
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.NotEqual:
		apiValueAsString := fmt.Sprintf("%g", apiValue)
		success := assertionValue != apiValueAsString
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.EqualNumber:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return model.ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := testValue == apiValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.IsLessThan:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return model.ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := apiValue < testValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.IsGreaterThan:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return model.ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := apiValue > testValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.IsLessThanOrEqual:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return model.ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := apiValue <= testValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.IsGreaterThanOrEqual:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return model.ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := apiValue >= testValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

// assertString is testing the value "apiValue" with the comparison on strings.
func (ctrl *assertionControllerImpl) assertString(assertion model.Assertion, apiValue string) model.ResultAssertion {
	comparison := assertion.Comparison
	assertionValue := assertion.Value
	propertyName := assertion.Property
	switch comparison {
	case model.Equal:
		success := assertionValue == apiValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.NotEqual:
		success := assertionValue != apiValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.Contains:
		success := strings.Contains(apiValue, assertionValue)
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.DoesNotContain:
		success := !strings.Contains(apiValue, assertionValue)
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.IsANumber:
		success := util.IsNumeric(apiValue)
		return model.NewResultAssertion(comparison, success, apiValue)

	case model.EqualNumber:
		if util.IsNumeric(apiValue) {
			apiValueAsNumber, _ := strconv.ParseFloat(apiValue, 64)
			return ctrl.assertNumber(assertion, apiValueAsNumber)
		}
		message := fmt.Sprintf("'%s' was not a number impossible to use %s", apiValue, comparison)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message)}

	case model.IsLessThan:
		success := apiValue < assertionValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.IsLessThanOrEqual:
		success := apiValue <= assertionValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.IsGreaterThan:
		success := apiValue > assertionValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.IsGreaterThanOrEqual:
		success := apiValue >= assertionValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.NotEmpty:
		success := strings.TrimSpace(apiValue) != ""
		return model.NewResultAssertion(comparison, success, propertyName)

	case model.Empty:
		success := strings.TrimSpace(apiValue) == ""
		return model.NewResultAssertion(comparison, success, apiValue)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

// assertBool is testing the value "apiValue" with the comparison on boolean.
func (ctrl *assertionControllerImpl) assertBool(assertion model.Assertion, apiValue bool) model.ResultAssertion {

	comparison := assertion.Comparison
	assertionValue := assertion.Value
	propertyName := assertion.Property

	// Parse the Assertions value to have the bool value
	testValue, err := strconv.ParseBool(assertionValue)
	if err != nil {
		message := fmt.Sprintf("'%s' was not comparable with a boolean value %t", assertionValue, apiValue)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}

	switch comparison {
	case model.IsANumber:
		return model.NewResultAssertion(comparison, false, propertyName)

	case model.Equal:
		success := testValue == apiValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	case model.NotEqual:
		success := testValue != apiValue
		return model.NewResultAssertion(comparison, success, apiValue, assertionValue)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

// assertArray is testing the value "apiValue" with the comparison on array/JSON Array.
func (ctrl *assertionControllerImpl) assertArray(assertion model.Assertion, apiValue []interface{}) model.ResultAssertion {
	comparison := assertion.Comparison
	propertyName := assertion.Property

	switch comparison {
	case model.IsANumber:
		return model.NewResultAssertion(comparison, false, apiValue)

	case model.IsNull:
		return model.NewResultAssertion(comparison, false, propertyName)

	case model.NotEmpty:
		success := len(apiValue) > 0
		return model.NewResultAssertion(comparison, success, propertyName)

	case model.Empty:
		success := len(apiValue) == 0
		return model.NewResultAssertion(comparison, success, propertyName)

	case model.Contains:
		return model.NewResultAssertion(comparison, sliceContainsValue(apiValue, assertion.Value), propertyName, assertion.Value)

	case model.DoesNotContain:
		return model.NewResultAssertion(comparison, !sliceContainsValue(apiValue, assertion.Value), propertyName, assertion.Value)

	case model.HasValue:
		for _, value := range apiValue {
			newAssertion := model.Assertion{
				Comparison: model.Equal,
				Value:      assertion.Value,
				Property:   assertion.Property,
				Source:     assertion.Source,
			}
			assert := ctrl.assertValue(newAssertion, value)
			if assert.Success {
				return model.NewResultAssertion(comparison, true, propertyName)
			}
		}
		return model.NewResultAssertion(comparison, false, propertyName)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

// assertArray is testing the value "apiValue" with the comparison on map/JSON Object.
func (ctrl *assertionControllerImpl) assertMap(assertion model.Assertion, apiValue map[string]interface{}) model.ResultAssertion {
	comparison := assertion.Comparison
	assertionValue := assertion.Value
	propertyName := assertion.Property

	switch comparison {
	case model.IsANumber:
		return model.NewResultAssertion(comparison, false, propertyName)

	case model.Empty:
		success := len(apiValue) == 0
		return model.NewResultAssertion(comparison, success, propertyName)

	case model.NotEmpty:
		success := len(apiValue) >= 0
		return model.NewResultAssertion(comparison, success, propertyName)

	case model.HasKey:
		_, success := apiValue[assertionValue]
		return model.NewResultAssertion(comparison, success, propertyName)

	case model.HasValue:
		for _, value := range apiValue {
			newAssertion := model.Assertion{
				Comparison: model.Equal,
				Value:      assertion.Value,
				Property:   assertion.Property,
				Source:     assertion.Source,
			}
			assert := ctrl.assertValue(newAssertion, value)
			if assert.Success {
				return model.NewResultAssertion(comparison, true, propertyName)
			}
		}
		return model.NewResultAssertion(comparison, false, propertyName)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return model.ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

// sliceContainsValue checks if a value is contains in a slice.
func sliceContainsValue(arr []interface{}, value string) bool {
	for _, v := range arr {
		switch item := v.(type) {
		case string:
			if item == value {
				return true
			}
			break

		case bool:
			if strconv.FormatBool(item) == value {
				return true
			}

		case float64:
			s64 := strconv.FormatFloat(item, 'E', -1, 64)
			if s64 == value {
				return true
			}
		}
	}

	return false
}
