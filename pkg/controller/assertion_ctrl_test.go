package controller_test

import (
	"fmt"
	"github.com/thomaspoignant/api-scenario/pkg/controller"
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/pkg/util"
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
var body, _ = util.StringToJson(`{
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
	 }`)
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
		message:  "'1500' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonIsANumberStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.IsANumber, Property: "id", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id_2' was not a number",
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
		message:  "'' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func TestResponseJsonEmptyStringInvalid(t *testing.T) {
	assertion := model.Assertion{Comparison: model.Empty, Property: "id", Source: model.ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   model.ResponseJson,
		message:  "'id_2' was not empty",
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
		message:  "'[map[primary:true value:indigo.anidter.ykxmid@test.com]]' was not a number",
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
