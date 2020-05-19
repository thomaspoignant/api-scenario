package model

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/jsonq"
	"github.com/thomaspoignant/api-scenario/pkg/util"
)

type Assertion struct {
	Comparison Comparison `json:"comparison"`
	Value      string     `json:"value"`
	Source     Source     `json:"Source"`
	Property   string     `json:"property,omitempty"`
}

const ComparisonNotSupportedMessage = "the comparison %s was not supported for the Source"

// Assert if the Assertion is valid for the current response,
// return true/false if test applied or an error.
func (assertion *Assertion) Assert(response Response) ResultAssertion {
	switch assertion.Source {
	case ResponseStatus:
		res := assertNumber(assertion.Comparison, assertion.Value, float64(response.StatusCode))
		res.Source = assertion.Source
		return res

	case ResponseTime:
		apiTime := float64(response.TimeElapsed) / float64(time.Second)
		res := assertNumber(assertion.Comparison, assertion.Value, apiTime)
		res.Source = assertion.Source
		return res

	case ResponseJson:
		res := assertResponseJson(assertion.Comparison, assertion.Value, assertion.Property, response.Body)
		res.Source = assertion.Source
		return res

	case ResponseHeader:
		res := assertResponseHeader(assertion.Comparison, assertion.Value, assertion.Property, response.Header)
		res.Source = assertion.Source
		return res

	default:
		message := fmt.Sprintf("the Source %s is not valid", assertion.Source)
		return ResultAssertion{Success: false, Err: errors.New(message), Message: message, Source: assertion.Source}
	}
}

func assertResponseHeader(assertionComparison Comparison,
	assertionValue string,
	assertionProperty string,
	headers http.Header) resultAssertion {

	//search for header using canonical key format
	values := headers.Values(assertionProperty)
	if values == nil {
		if assertionComparison == IsNull {
			result := NewResultAssertion(IsNull, true, assertionProperty)
			result.Property = assertionProperty
			return result
		}
		message := fmt.Sprintf("Header %q not found.", assertionProperty, assertionProperty)
		return resultAssertion{Success: false, Message: message, Err: errors.New(message), Property: assertionProperty}
	}

	//Compare fisrt value of the given header key
	//TODO handle many values for same key
	result := assertValue(assertionComparison, assertionValue, headers.Get(assertionProperty), assertionProperty)
	result.Property = assertionProperty
	return result

}

func assertResponseJson(
	assertionComparison Comparison,
	assertionValue string,
	assertionProperty string,
	body map[string]interface{}) ResultAssertion {
	// Convert property from Json syntax to an array of fields
	jqPath := util.JsonConvertKeyName(assertionProperty)

	// Init jsonq library with the response data
	jq := jsonq.NewQuery(body)

	// Get the key in the data
	extractedKey, err := jq.Interface(jqPath[:]...)
	if err != nil {
		//If the key if not found and we are looking for is_null this is a Success
		if assertionComparison == IsNull {
			result := NewResultAssertion(IsNull, true, assertionProperty)
			result.Property = assertionProperty
			return result
		}
		message := fmt.Sprintf("Unable to locate %s property in path '%s' in JSON", assertionProperty, assertionProperty)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message), Property: assertionProperty}
	}

	result := assertValue(assertionComparison, assertionValue, extractedKey, assertionProperty)
	result.Property = assertionProperty
	return result
}

func assertValue(comparison Comparison, assertionValue string, res interface{}, propertyName string) ResultAssertion {
	switch value := res.(type) {
	case string:
		return assertString(comparison, assertionValue, value, propertyName)
	case bool:
		return assertBool(comparison, assertionValue, value, propertyName)
	case float64:
		return assertNumber(comparison, assertionValue, value)
	case []interface{}:
		return assertArray(comparison, assertionValue, value, propertyName)
	case map[string]interface{}:
		return assertInterface(comparison, assertionValue, value, propertyName)
	default:
		// Not supposed to happen
		message := fmt.Sprintf("%s comparison is not available for element of type %s", comparison, reflect.TypeOf(res))
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

func assertNumber(comparison Comparison, assertionValue string, apiValue float64) ResultAssertion {
	switch comparison {
	case IsANumber:
		return NewResultAssertion(comparison, true, apiValue)

	case Equal:
		apiValueAsString := fmt.Sprintf("%g", apiValue)
		success := assertionValue == apiValueAsString
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case NotEqual:
		apiValueAsString := fmt.Sprintf("%g", apiValue)
		success := assertionValue != apiValueAsString
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case EqualNumber:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := testValue == apiValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case IsLessThan:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := apiValue < testValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case IsGreaterThan:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := apiValue > testValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case IsLessThanOrEqual:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := apiValue <= testValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case IsGreaterThanOrEqual:
		testValue, err := strconv.ParseFloat(assertionValue, 64)
		if err != nil {
			message := fmt.Sprintf("'%s' should be a number to compare with %s", assertionValue, comparison)
			return ResultAssertion{Err: err, Success: false, Message: message}
		}
		success := apiValue >= testValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

func assertString(comparison Comparison, assertionValue string, apiValue string, propertyName string) ResultAssertion {
	switch comparison {
	case Equal:
		success := assertionValue == apiValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case NotEqual:
		success := assertionValue != apiValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case Contains:
		success := strings.Contains(apiValue, assertionValue)
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case DoesNotContain:
		success := !strings.Contains(apiValue, assertionValue)
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case IsANumber:
		success := isNumeric(apiValue)
		return NewResultAssertion(comparison, success, apiValue)

	case EqualNumber:
		if isNumeric(apiValue) {
			apiValueAsNumber, _ := strconv.ParseFloat(apiValue, 64)
			return assertNumber(comparison, assertionValue, apiValueAsNumber)
		}
		message := fmt.Sprintf("'%s' was not a number impossible to use %s", apiValue, comparison)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}

	case IsLessThan:
		success := apiValue < assertionValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case IsLessThanOrEqual:
		success := apiValue <= assertionValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case IsGreaterThan:
		success := apiValue > assertionValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case IsGreaterThanOrEqual:
		success := apiValue >= assertionValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case NotEmpty:
		success := strings.TrimSpace(apiValue) != ""
		return NewResultAssertion(comparison, success, propertyName)

	case Empty:
		success := strings.TrimSpace(apiValue) == ""
		return NewResultAssertion(comparison, success, apiValue)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

func assertBool(comparison Comparison, assertionValue string, apiValue bool, propertyName string) ResultAssertion {
	// Parse the Assertion value to have the bool value
	testValue, err := strconv.ParseBool(assertionValue)
	if err != nil {
		message := fmt.Sprintf("'%s' was not comparable with a boolean value %t", assertionValue, apiValue)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}

	switch comparison {
	case IsANumber:
		return NewResultAssertion(comparison, false, propertyName)

	case Equal:
		success := testValue == apiValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	case NotEqual:
		success := testValue != apiValue
		return NewResultAssertion(comparison, success, apiValue, assertionValue)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

func assertArray(comparison Comparison, assertionValue string, apiValue []interface{}, propertyName string) ResultAssertion {
	switch comparison {
	case IsANumber:
		return NewResultAssertion(comparison, false, apiValue)

	case IsNull:
		return NewResultAssertion(comparison, false, propertyName)

	case NotEmpty:
		success := len(apiValue) > 0
		return NewResultAssertion(comparison, success, propertyName)

	case Empty:
		success := len(apiValue) == 0
		return NewResultAssertion(comparison, success, propertyName)

	case HasValue:
		for _, value := range apiValue {
			assert := assertValue(Equal, assertionValue, value, propertyName)
			if assert.Success {
				return NewResultAssertion(comparison, true, propertyName)
			}
		}
		return NewResultAssertion(comparison, false, propertyName)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

func assertInterface(comparison Comparison, assertionValue string, apiValue map[string]interface{}, propertyName string) ResultAssertion {
	switch comparison {
	case IsANumber:
		return NewResultAssertion(comparison, false, propertyName)

	case Empty:
		success := len(apiValue) == 0
		return NewResultAssertion(comparison, success, propertyName)

	case NotEmpty:
		success := len(apiValue) >= 0
		return NewResultAssertion(comparison, success, propertyName)

	case HasKey:
		_, success := apiValue[assertionValue]
		return NewResultAssertion(comparison, success, propertyName)

	case HasValue:
		for _, value := range apiValue {
			assert := assertValue(Equal, assertionValue, value, propertyName)
			if assert.Success {
				return NewResultAssertion(comparison, true, propertyName)
			}
		}
		return NewResultAssertion(comparison, false, propertyName)

	default:
		message := fmt.Sprintf(ComparisonNotSupportedMessage, comparison)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
