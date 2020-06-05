package controller_test

import (
	"fmt"
	"github.com/thomaspoignant/api-scenario/pkg/controller"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"net/http"
	"testing"
	"time"
)

type expectedResult struct {
	source   model.Source
	message  string
	property string
	success  bool
	err      bool
}

func te(t *testing.T, assertion model.Assertion, response model.Response, expected expectedResult) {
	ctrl := controller.NewAssertionController()
	got := ctrl.Assert(assertion, response)

	if expected.err {
		test.Ko(t, got.Err)
	} else {
		test.Ok(t, got.Err)
	}

	test.Equals(t, "invalid source", expected.source, got.Source)
	test.Equals(t, "invalid message", expected.message, got.Message)
	test.Equals(t, "should not have property for response_status source", expected.property, got.Property)
	test.Equals(t, "wrong model.Assertions result", expected.success, got.Success)
}

// response_status

func TestResponseStatusIsANumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Value: "200", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'200' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusEqualNumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "200", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'200' was a number equal to 200",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusEqualNumberNotAnNumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "qwerty", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'qwerty' should be a number to compare with equal_number",
		property: assertion.Property,
		success:  false,
		err:      true,
	})

}

func TestResponseStatusEqualNumberNotExpected(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "200", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusAccepted}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'202' was not a number equal to 200",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseStatusEqualValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "200", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'200' was equal to 200",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusEqualInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "200", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusAccepted}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'202' was not equal to 200",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseStatusIsLessThanValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'303' was less than 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusIsLessThanNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "qwerty", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'qwerty' should be a number to compare with is_less_than",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseStatusIsLessThanInvalidWhenValueEquals(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'400' was not less than 400",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseStatusIsLessThanInvalidWhenValueOver(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "200", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'400' was not less than 200",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseStatusIsLessThanOrEqualValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'303' was less than or equal to 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusIsLessThanOrEqualEquals(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'400' was less than or equal to 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusIsLessThanOrEqualNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Value: "qwerty", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'qwerty' should be a number to compare with is_less_than_or_equals",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseStatusIsLessThanOrEqualInvalidWhenValueOver(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Value: "200", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'400' was not less than or equal to 200",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseStatusIsGreaterThanInvalidWhenValueEquals(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'400' was not greater than 400",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseStatusIsGreaterThanInvalidWhenValueOver(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'200' was not greater than 400",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseStatusIsGreaterThanValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusBadGateway}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'502' was greater than 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusIsGreaterThanNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "qwerty", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'qwerty' should be a number to compare with is_greater_than",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseStatusIsGreaterThanOrEqualValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusBadGateway}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'502' was greater than or equal to 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusIsGreaterThanOrEqualEquals(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'400' was greater than or equal to 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseStatusIsGreaterThanOrEqualNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Value: "qwerty", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'qwerty' should be a number to compare with is_greater_than_or_equal",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseStatusIsGreaterThanOrEqualInvalidWhenValueOver(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Value: "400", Source: model.ResponseStatus}
	response := model.Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseStatus,
		message:  "'200' was not greater than or equal to 400",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseStatusNotSupportedComparison(t *testing.T) {
	unexpectedComparison := []model.Comparison{
		model.Contains, model.DoesNotContain, model.NotEmpty, model.Empty, model.IsNull, model.HasValue, model.HasKey,
	}

	for _, comparison := range unexpectedComparison {
		assertion := model.Assertion{Comparison: comparison, Value: "400", Source: model.ResponseStatus}
		response := model.Response{StatusCode: http.StatusOK}
		te(t, assertion, response, expectedResult{
			source:   model.ResponseStatus,
			message:  fmt.Sprintf("the comparison %s was not supported for the source", comparison.String()),
			property: assertion.Property,
			success:  false,
			err:      true,
		})
	}
}

// response_time

func TestResponseTimeEqualNumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "1", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'1' was a number equal to 1",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseTimeEqualNumberCompareAsFloat(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "1.1000", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: time.Duration(1.1 * 1000 * 1000 * 1000)}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'1.1' was a number equal to 1.1000",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseTimeEqualNumberInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "1", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 2 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'2' was not a number equal to 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseTimeEqualValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "1", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'1' was equal to 1",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseTimeEqualInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "1", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 2 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'2' was not equal to 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseTimeEqualCompareAsString(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "1.1000", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: time.Duration(1.1 * 1000 * 1000 * 1000)}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'1.1' was not equal to 1.1000",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseTimeIsLessThanInvalidWhenValueEquals(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "1", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'1' was not less than 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseTimeIsLessThanInvalidWhenValueOver(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "1", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 2 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'2' was not less than 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseTimeIsLessThanValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "2", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'1' was less than 2",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseTimeIsMoreThanInvalidWhenValueEquals(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "1", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 1 * time.Second}

	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'1' was not greater than 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseTimeIsMoreThanInvalidWhenValueOver(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "2", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'1' was not greater than 2",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseTimeIsMoreThanValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "2", Source: model.ResponseTime}
	response := model.Response{TimeElapsed: 3 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseTime,
		message:  "'3' was greater than 2",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseTimeNotSupportedComparison(t *testing.T) {
	unexpectedComparison := []model.Comparison{
		model.Contains, model.DoesNotContain, model.NotEmpty, model.Empty, model.IsNull, model.HasValue, model.HasKey,
	}

	for _, comparison := range unexpectedComparison {
		assertion := model.Assertion{Comparison: comparison, Value: "1", Source: model.ResponseTime}
		response := model.Response{StatusCode: http.StatusOK}
		te(t, assertion, response, expectedResult{
			source:   model.ResponseTime,
			message:  fmt.Sprintf("the comparison %s was not supported for the source", comparison.String()),
			property: assertion.Property,
			success:  false,
			err:      true,
		})
	}
}

// response_json
var body = `{
		"schemas": [
		  "urn:ietf:params:scim:schemas:core:2.0:User"
		],
		"id": "id_2",
		"point": 1500,
		"pointStr": "1500",
		"userName": "indigo.anidter_ykxmid",
		"name": {
		  "familyName": "Anidter",
		  "givenName": "Indigo"
		},
		"active": true,
		"emails": [
		  {
			"value": "indigo.anidter.ykxmid@test.com",
			"primary": true
		  }
		],
		"roles": [],
		"meta": {
		  "resourceType": "User",
		  "created": "2020-01-09T09:04:34.588Z",
		  "lastModified": "2020-01-09T09:05:55.943Z",
		  "location": "**REQUIRED**/Users/id_2"
		},
		"company": {},
		"building": null,
		"companyName": ""
	 }`

var response = model.Response{
	Body: body,
}

func TestResponseJsonEqualsStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "Anidter", Property: "name.familyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'Anidter' was equal to Anidter",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonEqualsStringComplexPath(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "indigo.anidter.ykxmid@test.com", Property: "emails[0].value", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'indigo.anidter.ykxmid@test.com' was equal to indigo.anidter.ykxmid@test.com",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonEqualsStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "Anidter1", Property: "name.familyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'Anidter' was not equal to Anidter1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonEqualsNumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "1500", Property: "point", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was equal to 1500",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonNotEqualsNumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "1501", Property: "point", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was not equal to 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonNotEqualsNumberInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "1500", Property: "point", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was equal to 1500",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonNotEqualsStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "not valid name", Property: "name.familyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'Anidter' was not equal to not valid name",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonNotEqualsStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "Anidter", Property: "name.familyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'Anidter' was equal to Anidter",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonContainsStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Contains, Value: "not", Property: "name.familyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'Anidter' does not contains not",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonContainsStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Contains, Value: "idt", Property: "name.familyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'Anidter' does contains idt",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonDoesNotContainStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.DoesNotContain, Value: "not", Property: "name.familyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'Anidter' does not contains not",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonDoesNotContainStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.DoesNotContain, Value: "idt", Property: "name.familyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'Anidter' does contains idt",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsANumberStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "pointStr", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'pointStr' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsANumberStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "id", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonEqualNumberStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Property: "pointStr", Value: "1500.00", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was a number equal to 1500.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonEqualNumberStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Property: "pointStr", Value: "1501", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was not a number equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonEqualNumberStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Property: "id", Value: "1501", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id_2' was not a number impossible to use equal_number",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseJsonEqualNumberStringInvalidAssertion(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Property: "pointStr", Value: "toto", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'toto' should be a number to compare with equal_number",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseJsonIsLessThanStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Property: "pointStr", Value: "1501.00", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was less than 1501.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsLessThanStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Property: "pointStr", Value: "1499", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was not less than 1499",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsLessThanStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Property: "id", Value: "1501", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id_2' was not less than 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsLessThanStringInvalidAssertion(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Property: "pointStr", Value: "toto", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was less than toto",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsGreaterThanStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Property: "pointStr", Value: "1499.00", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was greater than 1499.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsGreaterThanStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Property: "pointStr", Value: "1500", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was not greater than 1500",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsGreaterThanStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Property: "id", Value: "1501", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id_2' was greater than 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsGreaterThanStringInvalidAssertion(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Property: "pointStr", Value: "toto", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was not greater than toto",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsLessThanOrEqualStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Property: "pointStr", Value: "1501.00", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was less than or equal to 1501.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsLessThanOrEqualStringEqual(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Property: "pointStr", Value: "1500.00", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was less than or equal to 1500.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsLessThanOrEqualStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Property: "pointStr", Value: "1499", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was not less than or equal to 1499",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsLessThanOrEqualStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Property: "id", Value: "1501", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id_2' was not less than or equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsGreaterThanOrEqualStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Property: "pointStr", Value: "1499.00", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was greater than or equal to 1499.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsGreaterThanOrEqualStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Property: "id", Value: "1501", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id_2' was greater than or equal to 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsGreaterThanOrEqualStringInvalidAssertion(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Property: "pointStr", Value: "toto", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was not greater than or equal to toto",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonEmptyStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "companyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'companyName' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonEmptyStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "id", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id' was not empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonNotEmptyStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "id", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonNotEmptyStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "companyName", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'companyName' was empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonEqualsNumberNumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "1500", Property: "point", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was a number equal to 1500",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonEqualsNumberNumberInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "1501", Property: "point", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was not a number equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonNotEqualsBoolInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "false", Property: "active", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'true' was not equal to false",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonNotEqualsBoolValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "true", Property: "active", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'true' was equal to true",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonEqualsBoolValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "true", Property: "active", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'true' was equal to true",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsANumberBool(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Value: "true", Property: "active", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'active' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonEqualsBoolInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "false", Property: "active", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'true' was not equal to false",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonNotEmptyValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "emails", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'emails' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonNotEmptyInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "roles", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'roles' was empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonNotEmptyNotEmptyObject(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "meta", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'meta' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonEmptyValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "roles", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'roles' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonEmptyInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "emails", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'emails' was not empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonEmptyEmptyObject(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "company", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'company' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonHasSchemaValueValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasValue, Value: "urn:ietf:params:scim:schemas:core:2.0:User", Property: "schemas", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'schemas' had value",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonHasSchemaValueInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasValue, Value: "test", Property: "schemas", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'schemas' had no value",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsNumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "name", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'name' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsNullNullObject(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsNull, Property: "building", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'building' was null",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsNullInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsNull, Property: "emails", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'emails' was not null",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonHasKeyValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasKey, Property: "meta", Value: "resourceType", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'meta' key does exist",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonHasKeyInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasKey, Property: "meta", Value: "invalidKey", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'meta' key does not exist",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonHasValueValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasValue, Property: "meta", Value: "User", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'meta' had value",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonHasValueInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasValue, Property: "meta", Value: "invalidValue", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'meta' had no value",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonIsANumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "point", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'1500' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsANumberInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "emails", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'emails' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseJsonNoKey(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "inexistant", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "Unable to locate inexistant property in path 'inexistant' in JSON",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseJsonEqualsBoolCompareWithSomethingElse(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "toto", Property: "active", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'toto' was not comparable with a boolean value true",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseJsonInvalidJson(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "toto", Property: "active", Source: model.ResponseJson}
	te(t, assertion, model.Response{
		Body: `{"hello":"world"`,
	}, expectedResult{
		source:   model.ResponseJson,
		message:  "there is a result and this is not a valid JSON api Response is not in JSON",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

// response_header
var header http.Header = map[string][]string{
	"Content-Type": []string{"application/json; charset=utf-8"},
}

var responseH = model.Response{
	Header: header,
}

func TestResponseHeaderEquals(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "application/json; charset=utf-8", Property: "Content-Type", Source: model.ResponseHeader}
	te(t, assertion, responseH, expectedResult{
		source:   model.ResponseHeader,
		message:  "'application/json; charset=utf-8' was equal to application/json; charset=utf-8",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseHeaderEqualsNonCanonicalFormat(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "application/json; charset=utf-8", Property: "content-type", Source: model.ResponseHeader}
	te(t, assertion, responseH, expectedResult{
		source:   model.ResponseHeader,
		message:  "'application/json; charset=utf-8' was equal to application/json; charset=utf-8",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseHeaderHasKey(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasKey, Value: "", Property: "Content-Type", Source: model.ResponseHeader}
	te(t, assertion, responseH, expectedResult{
		source:   model.ResponseHeader,
		message:  "'Content-Type' key does exist",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseHeaderIsNull(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsNull, Value: "", Property: "unkown-key", Source: model.ResponseHeader}
	te(t, assertion, responseH, expectedResult{
		source:   model.ResponseHeader,
		message:  "'unkown-key' was null",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseHeaderErrorNotFound(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "toto", Property: "unkown-key", Source: model.ResponseHeader}
	te(t, assertion, responseH, expectedResult{
		source:   model.ResponseHeader,
		message:  "Header \"unkown-key\" not found.",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

// response_text
func TestResponseTextEqualValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.Equal, Value: "result", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' was equal to result",
		success: true,
		err:     false,
	})
}

func TestResponseTextEqualInValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.Equal, Value: "not result", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' was not equal to not result",
		success: false,
		err:     false,
	})
}

func TestResponseTextDoesNotEqualInValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "result", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' was equal to result",
		success: false,
		err:     false,
	})
}

func TestResponseTextDoesNotEqualValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "not result", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' was not equal to not result",
		success: true,
		err:     false,
	})
}

func TestResponseTextIsEmptyValid(t *testing.T) {
	responseText := model.Response{Body: ""}
	assertion := model.Assertion{Comparison: model.Empty, Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'body' was empty",
		success: true,
		err:     false,
	})
}

func TestResponseTextIsEmptyInValid(t *testing.T) {
	responseText := model.Response{Body: "res"}
	assertion := model.Assertion{Comparison: model.Empty, Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'body' was not empty",
		success: false,
		err:     false,
	})
}

func TestResponseTextIsNotEmptyValid(t *testing.T) {
	responseText := model.Response{Body: "res"}
	assertion := model.Assertion{Comparison: model.NotEmpty, Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'body' was not empty",
		success: true,
		err:     false,
	})
}

func TestResponseTextIsNotEmptyInValid(t *testing.T) {
	responseText := model.Response{Body: ""}
	assertion := model.Assertion{Comparison: model.NotEmpty, Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'body' was empty",
		success: false,
		err:     false,
	})
}

func TestResponseTextContainsValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.Contains, Value: "sul", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' does contains sul",
		success: true,
		err:     false,
	})
}

func TestResponseTextContainsInValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.Contains, Value: "suls", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' does not contains suls",
		success: false,
		err:     false,
	})
}

func TestResponseTextDoesNotContainsValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.DoesNotContain, Value: "suls", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' does not contains suls",
		success: true,
		err:     false,
	})
}

func TestResponseTextDoesNotContainsInValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.DoesNotContain, Value: "sul", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' does contains sul",
		success: false,
		err:     false,
	})
}

func TestResponseTextIsANumberValid(t *testing.T) {
	responseText := model.Response{Body: "10"}
	assertion := model.Assertion{Comparison: model.IsANumber, Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'10' was a number",
		success: true,
		err:     false,
	})
}

func TestResponseTextIsANumberInValid(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.IsANumber, Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'body' was not a number",
		success: false,
		err:     false,
	})
}

func TestResponseTextEqualNumberValid(t *testing.T) {
	responseText := model.Response{Body: "10"}
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "10", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'10' was a number equal to 10",
		success: true,
		err:     false,
	})
}

func TestResponseTextEqualNumberInValid(t *testing.T) {
	responseText := model.Response{Body: "11"}
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "10", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'11' was not a number equal to 10",
		success: false,
		err:     false,
	})
}

func TestResponseTextEqualNumberInValidString(t *testing.T) {
	responseText := model.Response{Body: "result"}
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "10", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'result' was not a number impossible to use equal_number",
		success: false,
		err:     true,
	})
}

func TestResponseTextLessThanValid(t *testing.T) {
	responseText := model.Response{Body: "10"}
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "11", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'10' was less than 11",
		success: true,
		err:     false,
	})
}

func TestResponseTextLessThanInValid(t *testing.T) {
	responseText := model.Response{Body: "10"}
	assertion := model.Assertion{Comparison: model.IsLessThan, Value: "9", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'10' was not less than 9",
		success: false,
		err:     false,
	})
}

func TestResponseTextLessThanOrEqualsValid(t *testing.T) {
	responseText := model.Response{Body: "11"}
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Value: "11", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'11' was less than or equal to 11",
		success: true,
		err:     false,
	})
}

func TestResponseTextLessThanOrEqualInValid(t *testing.T) {
	responseText := model.Response{Body: "10"}
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Value: "9", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'10' was not less than or equal to 9",
		success: false,
		err:     false,
	})
}

func TestResponseTextGreaterThanValid(t *testing.T) {
	responseText := model.Response{Body: "11"}
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "10", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'11' was greater than 10",
		success: true,
		err:     false,
	})
}

func TestResponseTextGreaterThanInValid(t *testing.T) {
	responseText := model.Response{Body: "9"}
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Value: "10", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'9' was not greater than 10",
		success: false,
		err:     false,
	})
}

func TestResponseTextGreaterThanOrEqualsValid(t *testing.T) {
	responseText := model.Response{Body: "11"}
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Value: "11", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'11' was greater than or equal to 11",
		success: true,
		err:     false,
	})
}

func TestResponseTextGreaterThanOrEqualInValid(t *testing.T) {
	responseText := model.Response{Body: "9"}
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Value: "10", Source: model.ResponseText}
	te(t, assertion, responseText, expectedResult{
		source:  model.ResponseText,
		message: "'9' was not greater than or equal to 10",
		success: false,
		err:     false,
	})
}

// response_xml
var bodyXml = `<?xml version="1.0" encoding="UTF-8"?>
<root>
   <active>true</active>
   <building null="true" />
   <company />
   <companyName />
   <emails>
      <email>
         <primary>true</primary>
         <value>indigo.anidter.ykxmid@test.com</value>
      </email>
	  <email>
         <primary>false</primary>
         <value>indigo.anidter.ykxmid2@test.com</value>
      </email>
   </emails>
   <id>id_2</id>
   <meta>
      <created>2020-01-09T09:04:34.588Z</created>
      <lastModified>2020-01-09T09:05:55.943Z</lastModified>
      <location>**REQUIRED**/Users/id_2</location>
      <resourceType>User</resourceType>
   </meta>
   <name>
      <familyName>Anidter</familyName>
      <givenName>Indigo</givenName>
   </name>
   <point>1500</point>
   <pointStr>1500</pointStr>
   <roles />
   <schemas>
      <element>urn:ietf:params:scim:schemas:core:2.0:User</element>
   </schemas>
   <userName>indigo.anidter_ykxmid</userName>
</root>`

var responseXml = model.Response{
	Body: bodyXml,
}

func TestResponseXmlEqualsStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "Anidter", Property: "root.name.familyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'Anidter' was equal to Anidter",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlEqualsStringComplexPath(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "indigo.anidter.ykxmid@test.com", Property: "root.emails.email[0].value", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'indigo.anidter.ykxmid@test.com' was equal to indigo.anidter.ykxmid@test.com",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlEqualsStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "Anidter1", Property: "root.name.familyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'Anidter' was not equal to Anidter1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlEqualsNumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "1500", Property: "root.point", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was equal to 1500",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlNotEqualsNumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "1501", Property: "root.point", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was not equal to 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlNotEqualsNumberInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "1500", Property: "root.point", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was equal to 1500",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlNotEqualsStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "not valid name", Property: "root.name.familyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'Anidter' was not equal to not valid name",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlNotEqualsStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "Anidter", Property: "root.name.familyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'Anidter' was equal to Anidter",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlContainsStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Contains, Value: "not", Property: "root.name.familyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'Anidter' does not contains not",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlContainsStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Contains, Value: "idt", Property: "root.name.familyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'Anidter' does contains idt",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlDoesNotContainStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.DoesNotContain, Value: "not", Property: "root.name.familyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'Anidter' does not contains not",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlDoesNotContainStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.DoesNotContain, Value: "idt", Property: "root.name.familyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'Anidter' does contains idt",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsANumberStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "root.pointStr", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.pointStr' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsANumberStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "root.id", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.id' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlEqualNumberStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Property: "root.pointStr", Value: "1500.00", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was a number equal to 1500.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlEqualNumberStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Property: "root.pointStr", Value: "1501", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was not a number equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlEqualNumberStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Property: "root.id", Value: "1501", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'id_2' was not a number impossible to use equal_number",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseXmlEqualNumberStringInvalidAssertion(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Property: "root.pointStr", Value: "toto", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'toto' should be a number to compare with equal_number",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseXmlIsLessThanStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Property: "root.pointStr", Value: "1501.00", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was less than 1501.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsLessThanStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Property: "root.pointStr", Value: "1499", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was not less than 1499",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsLessThanStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Property: "root.id", Value: "1501", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'id_2' was not less than 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsLessThanStringInvalidAssertion(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThan, Property: "root.pointStr", Value: "toto", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was less than toto",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsGreaterThanStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Property: "root.pointStr", Value: "1499.00", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was greater than 1499.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsGreaterThanStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Property: "root.pointStr", Value: "1500", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was not greater than 1500",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsGreaterThanStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Property: "root.id", Value: "1501", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'id_2' was greater than 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsGreaterThanStringInvalidAssertion(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThan, Property: "root.pointStr", Value: "toto", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was not greater than toto",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsLessThanOrEqualStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Property: "root.pointStr", Value: "1501.00", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was less than or equal to 1501.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsLessThanOrEqualStringEqual(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Property: "root.pointStr", Value: "1500.00", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was less than or equal to 1500.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsLessThanOrEqualStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Property: "root.pointStr", Value: "1499", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was not less than or equal to 1499",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsLessThanOrEqualStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsLessThanOrEqual, Property: "root.id", Value: "1501", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'id_2' was not less than or equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsGreaterThanOrEqualStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Property: "root.pointStr", Value: "1499.00", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was greater than or equal to 1499.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsGreaterThanOrEqualStringNotANumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Property: "root.id", Value: "1501", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'id_2' was greater than or equal to 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsGreaterThanOrEqualStringInvalidAssertion(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsGreaterThanOrEqual, Property: "root.pointStr", Value: "toto", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was not greater than or equal to toto",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlEmptyStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "root.companyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.companyName' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlEmptyStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "root.id", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.id' was not empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlNotEmptyStringValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "root.id", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.id' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlNotEmptyStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "root.companyName", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.companyName' was empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlEqualsNumberNumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "1500", Property: "root.point", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was a number equal to 1500",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlEqualsNumberNumberInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.EqualNumber, Value: "1501", Property: "root.point", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'1500' was not a number equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlNotEqualsBoolInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "false", Property: "root.active", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'true' was not equal to false",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlNotEqualsBoolValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEqual, Value: "true", Property: "root.active", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'true' was equal to true",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlEqualsBoolValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "true", Property: "root.active", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'true' was equal to true",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsANumberBool(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Value: "true", Property: "root.active", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.active' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlEqualsBoolInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "false", Property: "root.active", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'true' was not equal to false",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlNotEmptyValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "root.emails", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.emails' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlNotEmptyInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "root.roles", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.roles' was empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlNotEmptyNotEmptyObject(t *testing.T) {
	assertion := model.Assertion{Comparison: model.NotEmpty, Property: "root.meta", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.meta' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlEmptyValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "root.roles", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.roles' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlEmptyInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "root.emails", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.emails' was not empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlEmptyEmptyObject(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "root.company", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.company' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlHasSchemaValueValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasValue, Value: "urn:ietf:params:scim:schemas:core:2.0:User", Property: "root.schemas", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.schemas' had value",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlHasSchemaValueInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasValue, Value: "test", Property: "root.schemas", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.schemas' had no value",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsNumber(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "root.name", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.name' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsNullNullObject(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsNull, Property: "root.building", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.building' was null",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsNullInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsNull, Property: "root.emails", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.emails' was not null",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlHasKeyValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasKey, Property: "root.meta", Value: "resourceType", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.meta' key does exist",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlHasKeyInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasKey, Property: "root.meta", Value: "invalidKey", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.meta' key does not exist",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlHasValueValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasValue, Property: "root.meta", Value: "User", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.meta' had value",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlHasValueInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.HasValue, Property: "root.meta", Value: "invalidValue", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.meta' had no value",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlIsANumberValid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "root.point", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.point' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseXmlIsANumberInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "root.emails", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'root.emails' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlNoKey(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "root.inexistant", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "Unable to locate root.inexistant property in path 'root.inexistant' in JSON",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseXmlEqualsBoolCompareWithSomethingElse(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "toto", Property: "root.active", Source: model.ResponseXml}
	te(t, assertion, responseXml, expectedResult{
		source:   model.ResponseXml,
		message:  "'true' was not equal to toto",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func TestResponseXmlInvalidXml(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "toto", Property: "root.active", Source: model.ResponseXml}
	te(t, assertion, model.Response{
		Body: `<root><hello>world`,
	}, expectedResult{
		source:   model.ResponseXml,
		message:  "xml.Decoder.Token() - XML syntax error on line 1: unexpected EOF",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func TestResponseXmlEmptyBody(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Equal, Value: "toto", Property: "root.active", Source: model.ResponseXml}
	te(t, assertion, model.Response{
		Body: ``,
	}, expectedResult{
		source:   model.ResponseXml,
		message:  "there is a result and this is not a valid XML api Response",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}
