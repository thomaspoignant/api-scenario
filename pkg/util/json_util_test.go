package util_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/thomaspoignant/api-scenario/pkg/util"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
)

func TestIsJsonEmptyDocument(t *testing.T) {
	want := true
	got := util.IsJson("{}")
	test.Equals(t, "wrong assertion result", want, got)
}

func TestIsJsonValid(t *testing.T) {
	want := true
	got := util.IsJson(`{
		"hello": "world"
	}`)
	test.Equals(t, "wrong assertion result", want, got)
}

func TestIsJsonEmptyString(t *testing.T) {
	want := false
	got := util.IsJson("")
	test.Equals(t, "wrong assertion result", want, got)
}

func TestIsJsonInvalidJson(t *testing.T) {
	want := false
	got := util.IsJson(`{
		"hello": "world"
		"world": "hello"
	}`)
	test.Equals(t, "wrong assertion result", want, got)
}

func TestJsonConvertKeyNameValid(t *testing.T) {
	want := []string{"emails", "0", "value"}
	got := util.JsonConvertKeyName("emails[0].value")

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStringToJsonValid(t *testing.T) {
	input := `{
		"hello": "world",
		"world": ["hello"]
	}`

	got, err := util.StringToJson(input)
	test.Ok(t, err)

	world := make([]interface{}, 1)
	world[0] = "hello"

	want := make(map[string]interface{})
	want["hello"] = "world"
	want["world"] = world

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestStringToJsonEmpty(t *testing.T) {
	input := ""
	got, err := util.StringToJson(input)
	test.Ok(t, err)

	want := make(map[string]interface{})
	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
