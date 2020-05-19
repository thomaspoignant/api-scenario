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

// An Assertion defines one expected result
type Assertion struct {
	Comparison Comparison `json:"comparison"`
	Value      string     `json:"value"`
	Source     Source     `json:"Source"`
	Property   string     `json:"property,omitempty"`
}

const ComparisonNotSupportedMessage = "the comparison %s was not supported for the Source"

// Assert if the Assertion is valid for the current response,
func (a *Assertion) Assert(resp Response) ResultAssertion {

	switch a.Source {
	case ResponseStatus:
		res := a.assertNumber(float64(resp.StatusCode))
		res.Source = a.Source
		return res

	case ResponseTime:
		apiTime := float64(resp.TimeElapsed) / float64(time.Second)
		res := a.assertNumber(apiTime)
		res.Source = a.Source
		return res

	case ResponseJson:
		res := a.assertResponseJson(resp.Body)
		res.Source = a.Source
		return res

	case ResponseHeader:
		res := a.assertResponseHeader(resp.Header)
		res.Source = a.Source
		return res

	default:
		message := fmt.Sprintf("the Source %s is not valid", a.Source)
		return ResultAssertion{Success: false, Err: errors.New(message), Message: message, Source: a.Source}
	}
}

func (a *Assertion) assertResponseHeader(h http.Header) ResultAssertion {
	//search for header using canonical key format
	values := h.Values(a.Property)
	if values == nil {
		if a.Comparison == IsNull {
			result := NewResultAssertion(IsNull, true, a.Property)
			result.Property = a.Property
			return result
		}
		message := fmt.Sprintf("Header %q not found.", a.Property)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message), Property: a.Property}
	}

	//Compare fisrt value of the given header key
	//TODO handle many values for same key
	result := a.assertValue(h.Get(a.Property))
	result.Property = a.Property
	return result
}

func (a *Assertion) assertResponseJson(body map[string]interface{}) ResultAssertion {
	// Convert property from Json syntax to an array of fields
	jqPath := util.JsonConvertKeyName(a.Property)

	// Init jsonq library with the response data
	jq := jsonq.NewQuery(body)

	// Get the key in the data
	extractedKey, err := jq.Interface(jqPath[:]...)
	if err != nil {
		//If the key if not found and we are looking for is_null this is a Success
		if a.Comparison == IsNull {
			result := NewResultAssertion(IsNull, true, a.Property)
			result.Property = a.Property
			return result
		}
		message := fmt.Sprintf("Unable to locate %s property in path '%s' in JSON", assertionProperty, assertionProperty)
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message), Property: assertionProperty}
	}

	result := a.assertValue(extractedKey)
	result.Property = a.Property
	return result
}

func (a *Assertion) assertValue(got interface{}) ResultAssertion {
	switch got := got.(type) {
	case string:
		return a.assertString(got)
	case bool:
		return a.assertBool(got)
	case float64:
		return a.assertNumber(got)
	case []interface{}:
		return a.assertArray(got)
	case map[string]interface{}:
		return a.assertMap(got)
	default:
		// Not supposed to happen
		message := fmt.Sprintf("%s comparison is not available for element of type %s", a.Comparison, reflect.TypeOf(got))
		return ResultAssertion{Success: false, Message: message, Err: errors.New(message)}
	}
}

func (a *Assertion) assertNumber(apiValue float64) ResultAssertion {
	comparison := a.Comparison
	assertionValue := a.Value

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

func (a *Assertion) assertString(apiValue string) ResultAssertion {
	comparison := a.Comparison
	assertionValue := a.Value
	propertyName := a.Property
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
			return a.assertNumber(apiValueAsNumber)
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

func (a *Assertion) assertBool(apiValue bool) ResultAssertion {

	comparison := a.Comparison
	assertionValue := a.Value
	propertyName := a.Property

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

func (a *Assertion) assertArray(apiValue []interface{}) ResultAssertion {
	comparison := a.Comparison
	propertyName := a.Property

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
			newAssertion := Assertion{
				Comparison: Equal,
				Value:      a.Value,
				Property:   a.Property,
				Source:     a.Source,
			}
			assert := newAssertion.assertValue(value)
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

func (a *Assertion) assertMap(apiValue map[string]interface{}) ResultAssertion {
	comparison := a.Comparison
	assertionValue := a.Value
	propertyName := a.Property

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
			newAssertion := Assertion{
				Comparison: Equal,
				Value:      a.Value,
				Property:   a.Property,
				Source:     a.Source,
			}
			assert := newAssertion.assertValue(value)
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
