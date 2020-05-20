package model

type Comparison int

//go:generate enumer -type=Comparison -json -linecomment -output comparison_generated.go
const (
	EqualNumber          Comparison = iota //equal_number
	Equal                                  //equal
	NotEqual                               //not_equal
	IsANumber                              //is_a_number
	IsLessThan                             //is_less_than
	IsLessThanOrEqual                      //is_less_than_or_equals
	IsGreaterThan                          //is_greater_than
	IsGreaterThanOrEqual                   //is_greater_than_or_equal
	Contains                               //contains
	DoesNotContain                         //does_not_contains
	NotEmpty                               //not_empty
	Empty                                  //empty
	IsNull                                 //is_null
	HasValue                               //has_value
	HasKey                                 //has_key
)

type comparisonMessage struct {
	Success string
	Failure string
}

var messages = map[Comparison]comparisonMessage{
	Equal: {
		Success: "'%v' was equal to %s",
		Failure: "'%v' was not equal to %s"},
	NotEqual: {
		Success: "'%v' was not equal to %s",
		Failure: "'%v' was equal to %s"},
	EqualNumber: {
		Success: "'%v' was a number equal to %s",
		Failure: "'%v' was not a number equal to %s"},
	IsANumber: {
		Success: "'%v' was a number",
		Failure: "'%v' was not a number"},
	IsLessThan: {
		Success: "'%v' was less than %s",
		Failure: "'%v' was not less than %s"},
	IsLessThanOrEqual: {
		Success: "'%v' was less than or equal to %s",
		Failure: "'%v' was not less than or equal to %s"},
	IsGreaterThan: {
		Success: "'%v' was greater than %s",
		Failure: "'%v' was not greater than %s"},
	IsGreaterThanOrEqual: {
		Success: "'%v' was greater than or equal to %s",
		Failure: "'%v' was not greater than or equal to %s"},
	Contains: {
		Success: "'%v' does contains %s",
		Failure: "'%v' does not contains %s"},
	DoesNotContain: {
		Success: "'%v' does not contains %s",
		Failure: "'%v' does contains %s"},
	NotEmpty: {
		Success: "'%v' was not empty",
		Failure: "'%v' was empty"},
	Empty: {
		Success: "'%v' was empty",
		Failure: "'%v' was not empty"},
	IsNull: {Success: "'%v' was null",
		Failure: "'%v' was not null"},
	HasValue: {
		Success: "'%v' had value",
		Failure: "'%v' had no value"},
	HasKey: {
		Success: "'%v' key does exist",
		Failure: "'%v' key does not exist"},
}

func (comparison Comparison) GetMessage() comparisonMessage {
	return messages[comparison]
}
