package model

import (
	"fmt"
	"github.com/thomaspoignant/rest-scenario/pkg/util"
	"github.com/thomaspoignant/rest-scenario/test"
	"net/http"
	"testing"
	"time"
)

type expectedResult struct {
	source   Source
	message  string
	property string
	success  bool
	err      bool
}

func te(t *testing.T, assertion Assertion, response Response, expected expectedResult) {
	got := assertion.Assert(response)

	if expected.err {
		test.Ko(t, got.Err)
	} else {
		test.Ok(t, got.Err)
	}

	test.Equals(t, "invalid source", expected.source, got.Source)
	test.Equals(t, "invalid message", expected.message, got.Message)
	test.Equals(t, "should not have property for response_status source", expected.property, got.Property)
	test.Equals(t, "wrong Assertion result", expected.success, got.Success)
}

// response_status
func Test_ResponseStatus_isANumber_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Value: "200", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'200' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_EqualNumber_valid(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Value: "200", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'200' was a number equal to 200",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_EqualNumber_notAnNumber(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Value: "qwerty", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'qwerty' should be a number to compare with equal_number",
		property: assertion.Property,
		success:  false,
		err:      true,
	})

}
func Test_ResponseStatus_EqualNumber_notExpected(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Value: "200", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusAccepted}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'202' was not a number equal to 200",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseStatus_Equal_valid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "200", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'200' was equal to 200",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_Equal_invalid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "200", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusAccepted}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'202' was not equal to 200",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseStatus_IsLessThan_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'303' was less than 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_IsLessThan_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Value: "qwerty", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'qwerty' should be a number to compare with is_less_than",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}
func Test_ResponseStatus_IsLessThan_invalidWhenValueEquals(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'400' was not less than 400",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseStatus_IsLessThan_invalidWhenValueOver(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Value: "200", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'400' was not less than 200",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseStatus_IsLessThanOrEqual_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThanOrEqual, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'303' was less than or equal to 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_IsLessThanOrEqual_equal(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThanOrEqual, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'400' was less than or equal to 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_IsLessThanOrEqual_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThanOrEqual, Value: "qwerty", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'qwerty' should be a number to compare with is_less_than_or_equals",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}
func Test_ResponseStatus_IsLessThanOrEqual_invalidWhenValueOver(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThanOrEqual, Value: "200", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'400' was not less than or equal to 200",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseStatus_IsGreaterThan_invalidWhenValueEquals(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'400' was not greater than 400",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseStatus_IsGreaterThan_invalidWhenValueOver(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'200' was not greater than 400",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseStatus_IsGreaterThan_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusBadGateway}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'502' was greater than 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_IsGreaterThan_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Value: "qwerty", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'qwerty' should be a number to compare with is_greater_than",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}
func Test_ResponseStatus_IsGreaterThanOrEqual_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThanOrEqual, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusBadGateway}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'502' was greater than or equal to 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_IsGreaterThanOrEqual_equal(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThanOrEqual, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusBadRequest}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'400' was greater than or equal to 400",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseStatus_IsGreaterThanOrEqual_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThanOrEqual, Value: "qwerty", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusSeeOther}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'qwerty' should be a number to compare with is_greater_than_or_equal",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}
func Test_ResponseStatus_IsGreaterThanOrEqual_invalidWhenValueOver(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThanOrEqual, Value: "400", Source: ResponseStatus}
	response := Response{StatusCode: http.StatusOK}
	te(t, assertion, response, expectedResult{
		source:   ResponseStatus,
		message:  "'200' was not greater than or equal to 400",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseStatusNotSupportedComparison(t *testing.T) {
	unexpectedComparison := []Comparison{
		Contains, DoesNotContain, NotEmpty, Empty, IsNull, HasValue, HasKey,
	}

	for _, comparison := range unexpectedComparison {
		assertion := Assertion{Comparison: comparison, Value: "400", Source: ResponseStatus}
		response := Response{StatusCode: http.StatusOK}
		te(t, assertion, response, expectedResult{
			source:   ResponseStatus,
			message:  fmt.Sprintf("the comparison %s was not supported for the Source", comparison.String()),
			property: assertion.Property,
			success:  false,
			err:      true,
		})
	}
}

// response_time
func Test_ResponseTime_EqualNumber_valid(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Value: "1", Source: ResponseTime}
	response := Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'1' was a number equal to 1",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseTime_EqualNumber_compareAsFloat(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Value: "1.1000", Source: ResponseTime}
	response := Response{TimeElapsed: time.Duration(1.1 * 1000 * 1000 * 1000)}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'1.1' was a number equal to 1.1000",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseTime_EqualNumber_invalid(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Value: "1", Source: ResponseTime}
	response := Response{TimeElapsed: 2 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'2' was not a number equal to 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseTime_Equal_valid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "1", Source: ResponseTime}
	response := Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'1' was equal to 1",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseTime_Equal_invalid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "1", Source: ResponseTime}
	response := Response{TimeElapsed: 2 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'2' was not equal to 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseTime_Equal_compareAsString(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "1.1000", Source: ResponseTime}
	response := Response{TimeElapsed: time.Duration(1.1 * 1000 * 1000 * 1000)}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'1.1' was not equal to 1.1000",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseTime_IsLessThan_invalidWhenValueEquals(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Value: "1", Source: ResponseTime}
	response := Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'1' was not less than 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseTime_IsLessThan_invalidWhenValueOver(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Value: "1", Source: ResponseTime}
	response := Response{TimeElapsed: 2 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'2' was not less than 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseTime_IsLessThan_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Value: "2", Source: ResponseTime}
	response := Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'1' was less than 2",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseTime_IsMoreThan_invalidWhenValueEquals(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Value: "1", Source: ResponseTime}
	response := Response{TimeElapsed: 1 * time.Second}

	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'1' was not greater than 1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseTime_IsMoreThan_invalidWhenValueOver(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Value: "2", Source: ResponseTime}
	response := Response{TimeElapsed: 1 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'1' was not greater than 2",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseTime_IsMoreThan_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Value: "2", Source: ResponseTime}
	response := Response{TimeElapsed: 3 * time.Second}
	te(t, assertion, response, expectedResult{
		source:   ResponseTime,
		message:  "'3' was greater than 2",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseTime_NotSupportedComparison(t *testing.T) {
	unexpectedComparison := []Comparison{
		Contains, DoesNotContain, NotEmpty, Empty, IsNull, HasValue, HasKey,
	}

	for _, comparison := range unexpectedComparison {
		assertion := Assertion{Comparison: comparison, Value: "1", Source: ResponseTime}
		response := Response{StatusCode: http.StatusOK}
		te(t, assertion, response, expectedResult{
			source:   ResponseTime,
			message:  fmt.Sprintf("the comparison %s was not supported for the Source", comparison.String()),
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
var response = Response{
	Body: body,
}

func Test_ResponseJson_equals_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "Anidter", Property: "name.familyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'Anidter' was equal to Anidter",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_equals_string_complexPath(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "indigo.anidter.ykxmid@test.com", Property: "emails[0].value", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'indigo.anidter.ykxmid@test.com' was equal to indigo.anidter.ykxmid@test.com",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_equals_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "Anidter1", Property: "name.familyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'Anidter' was not equal to Anidter1",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_equals_number_valid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "1500", Property: "point", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was equal to 1500",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_notEquals_number_valid(t *testing.T) {
	assertion := Assertion{Comparison: NotEqual, Value: "1501", Property: "point", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was not equal to 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_notEquals_number_invalid(t *testing.T) {
	assertion := Assertion{Comparison: NotEqual, Value: "1500", Property: "point", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was equal to 1500",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_notEquals_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: NotEqual, Value: "not valid name", Property: "name.familyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'Anidter' was not equal to not valid name",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_notEquals_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: NotEqual, Value: "Anidter", Property: "name.familyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'Anidter' was equal to Anidter",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_contains_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: Contains, Value: "not", Property: "name.familyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'Anidter' does not contains not",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_contains_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: Contains, Value: "idt", Property: "name.familyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'Anidter' does contains idt",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func Test_ResponseJson_doesNotContain_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: DoesNotContain, Value: "not", Property: "name.familyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'Anidter' does not contains not",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_doesNotContain_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: DoesNotContain, Value: "idt", Property: "name.familyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'Anidter' does contains idt",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_isANumber_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Property: "pointStr", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isANumber_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Property: "id", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'id_2' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_equalNumber_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Property: "pointStr", Value: "1500.00", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was a number equal to 1500.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_equalNumber_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Property: "pointStr", Value: "1501", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was not a number equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_equalNumber_string_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Property: "id", Value: "1501", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'id_2' was not a number impossible to use equal_number",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}
func Test_ResponseJson_equalNumber_string_invalidAssertion(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Property: "pointStr", Value: "toto", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'toto' should be a number to compare with equal_number",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func Test_ResponseJson_isLessThan_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Property: "pointStr", Value: "1501.00", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was less than 1501.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isLessThan_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Property: "pointStr", Value: "1499", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was not less than 1499",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_isLessThan_string_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Property: "id", Value: "1501", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'id_2' was not less than 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_isLessThan_string_invalidAssertion(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThan, Property: "pointStr", Value: "toto", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was less than toto",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func Test_ResponseJson_isGreaterThan_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Property: "pointStr", Value: "1499.00", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was greater than 1499.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isGreaterThan_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Property: "pointStr", Value: "1500", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was not greater than 1500",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_isGreaterThan_string_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Property: "id", Value: "1501", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'id_2' was greater than 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isGreaterThan_string_invalidAssertion(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThan, Property: "pointStr", Value: "toto", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was not greater than toto",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_isLessThanOrEqual_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThanOrEqual, Property: "pointStr", Value: "1501.00", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was less than or equal to 1501.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isLessThanOrEqual_string_equal(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThanOrEqual, Property: "pointStr", Value: "1500.00", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was less than or equal to 1500.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isLessThanOrEqual_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThanOrEqual, Property: "pointStr", Value: "1499", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was not less than or equal to 1499",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_isLessThanOrEqual_string_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: IsLessThanOrEqual, Property: "id", Value: "1501", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'id_2' was not less than or equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_isGreaterThanOrEqual_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThanOrEqual, Property: "pointStr", Value: "1499.00", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was greater than or equal to 1499.00",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isGreaterThanOrEqual_string_notANumber(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThanOrEqual, Property: "id", Value: "1501", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'id_2' was greater than or equal to 1501",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isGreaterThanOrEqual_string_invalidAssertion(t *testing.T) {
	assertion := Assertion{Comparison: IsGreaterThanOrEqual, Property: "pointStr", Value: "toto", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was not greater than or equal to toto",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_empty_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: Empty, Property: "companyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_empty_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: Empty, Property: "id", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'id_2' was not empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_notEmpty_string_valid(t *testing.T) {
	assertion := Assertion{Comparison: NotEmpty, Property: "id", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'id' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_notEmpty_string_invalid(t *testing.T) {
	assertion := Assertion{Comparison: NotEmpty, Property: "companyName", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'companyName' was empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_equalsNumber_number_valid(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Value: "1500", Property: "point", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was a number equal to 1500",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_equalsNumber_number_invalid(t *testing.T) {
	assertion := Assertion{Comparison: EqualNumber, Value: "1501", Property: "point", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was not a number equal to 1501",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_notEquals_bool_invalid(t *testing.T) {
	assertion := Assertion{Comparison: NotEqual, Value: "false", Property: "active", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'true' was not equal to false",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_notEquals_bool_valid(t *testing.T) {
	assertion := Assertion{Comparison: NotEqual, Value: "true", Property: "active", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'true' was equal to true",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_equals_bool_valid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "true", Property: "active", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'true' was equal to true",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isANumber_bool(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Value: "true", Property: "active", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'active' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_equals_bool_invalid(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "false", Property: "active", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'true' was not equal to false",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_notEmpty_valid(t *testing.T) {
	assertion := Assertion{Comparison: NotEmpty, Property: "emails", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'emails' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_notEmpty_invalid(t *testing.T) {
	assertion := Assertion{Comparison: NotEmpty, Property: "roles", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'roles' was empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_notEmpty_notEmptyObject(t *testing.T) {
	assertion := Assertion{Comparison: NotEmpty, Property: "meta", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'meta' was not empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func Test_ResponseJson_empty_valid(t *testing.T) {
	assertion := Assertion{Comparison: Empty, Property: "roles", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'roles' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_empty_invalid(t *testing.T) {
	assertion := Assertion{Comparison: Empty, Property: "emails", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'emails' was not empty",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_empty_emptyObject(t *testing.T) {
	assertion := Assertion{Comparison: Empty, Property: "company", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'company' was empty",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}

func Test_ResponseJson_HasValue_valid(t *testing.T) {
	assertion := Assertion{Comparison: HasValue, Value: "urn:ietf:params:scim:schemas:core:2.0:User", Property: "schemas", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'schemas' had value",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_HasValue_invalid(t *testing.T) {
	assertion := Assertion{Comparison: HasValue, Value: "test", Property: "schemas", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'schemas' had no value",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_isNumber(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Property: "name", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'name' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}
func Test_ResponseJson_isNull_nullObject(t *testing.T) {
	assertion := Assertion{Comparison: IsNull, Property: "building", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'building' was null",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isNull_invalid(t *testing.T) {
	assertion := Assertion{Comparison: IsNull, Property: "emails", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'emails' was not null",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_hasKey_valid(t *testing.T) {
	assertion := Assertion{Comparison: HasKey, Property: "meta", Value: "resourceType", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'meta' key does exist",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_hasKey_invalid(t *testing.T) {
	assertion := Assertion{Comparison: HasKey, Property: "meta", Value: "invalidKey", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'meta' key does not exist",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_hasValue_valid(t *testing.T) {
	assertion := Assertion{Comparison: HasValue, Property: "meta", Value: "User", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'meta' had value",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_hasValue_invalid(t *testing.T) {
	assertion := Assertion{Comparison: HasValue, Property: "meta", Value: "invalidValue", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'meta' had no value",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_isANumber_valid(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Property: "point", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'1500' was a number",
		property: assertion.Property,
		success:  true,
		err:      false,
	})
}
func Test_ResponseJson_isANumber_invalid(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Property: "emails", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'[map[primary:true value:indigo.anidter.ykxmid@test.com]]' was not a number",
		property: assertion.Property,
		success:  false,
		err:      false,
	})
}

func Test_ResponseJson_inexistantKey(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Property: "inexistant", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "Unable to locate inexistant property in path 'inexistant' in JSON",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func Test_ResponseJson_emptyJson(t *testing.T) {
	assertion := Assertion{Comparison: IsANumber, Property: "inexistant", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "Unable to locate inexistant property in path 'inexistant' in JSON",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}

func Test_ResponseJson_equals_bool_compareWithSomethingElse(t *testing.T) {
	assertion := Assertion{Comparison: Equal, Value: "toto", Property: "active", Source: ResponseJson}
	te(t, assertion, response, expectedResult{
		source:   ResponseJson,
		message:  "'toto' was not comparable with a boolean value true",
		property: assertion.Property,
		success:  false,
		err:      true,
	})
}
