package context_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/thomaspoignant/api-scenario/pkg/context"
	"github.com/thomaspoignant/api-scenario/test"
	"strconv"
	"strings"
	"testing"
	"time"
)

func resetContext() {
	context.GetContext().ResetContext()
}

func TestPatchOneElement(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	ctx.Add("SCIMBaseURL", "http://google.fr")

	input := "{{SCIMBaseURL}}/Users"
	want := "http://google.fr/Users"
	got := ctx.Patch(input)

	test.Equals(t, "Should patch on variables", want, got)
}

func TestNothingToPatch(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	ctx.Add("SCIMBaseURL", "http://google.fr")

	input := "http://yahoo.fr/Users"
	want := "http://yahoo.fr/Users"
	got := ctx.Patch(input)

	test.Equals(t, "Should not changed anything", want, got)
}

func TestMultipleElementsWithSameKey(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	ctx.Add("SCIMBaseURL", "http://google.fr")

	input := "{{SCIMBaseURL}}/Users?param={{SCIMBaseURL}}"
	want := "http://google.fr/Users?param=http://google.fr"
	got := ctx.Patch(input)

	test.Equals(t, "Should replace multiple time the same variables", want, got)
}

func TestMultipleElementsWithMultipleKeys(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	ctx.Add("SCIMBaseURL", "http://google.fr")
	ctx.Add("name", "John")

	input := "{{SCIMBaseURL}}/Users?param={{name}}"
	want := "http://google.fr/Users?param=John"
	got := ctx.Patch(input)

	test.Equals(t, "Should replace multiple variables", want, got)
}

func TestContextNotInit(t *testing.T) {
	resetContext()
	ctx := context.GetContext()

	input := "{{SCIMBaseURL}}/Users?param={{name}}"
	want := "{{SCIMBaseURL}}/Users?param={{name}}"
	got := ctx.Patch(input)

	test.Equals(t, "Should replace multiple variables", want, got)
}

func TestBuiltinTimestamp(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{timestamp}}"
	got := ctx.Patch(input)

	_, err := strconv.Atoi(got)
	if err != nil {
		t.Error("timestamp should be a number", got)
	}
}

func TestBuiltinUTC(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{utc_datetime}}"
	got := ctx.Patch(input)

	if got == input {
		t.Error("patch not applied", got)
	}
}

func TestBuiltinRandomInt(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{random_int}}"
	got := ctx.Patch(input)

	_, err := strconv.Atoi(got)
	if err != nil {
		t.Error("randomIntWithRange should be a int", got)
	}
}

func TestBuiltinUUID(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{uuid}}"
	got := ctx.Patch(input)

	_, err := uuid.Parse(got)
	if err != nil {
		t.Error("Patched value is not a uuid", got)
	}
}

func TestBuiltinRandomIntWithRange(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{random_int(0,100)}};{{random_int(123,11)}};{{random_int(0,100)}}"
	got := ctx.Patch(input)

	splitedRes := strings.Split(got, ";")

	test.Equals(t, "we should have 3 elements as results", len(splitedRes), 3)

	for _, item := range splitedRes {
		_, err := strconv.Atoi(item)
		if err != nil {
			t.Error("randomIntWithRange should be a int", item)
		}
	}
}

func TestBuiltinRandomIntWithRangeNegative(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{random_int(-100,-1)}}"
	got := ctx.Patch(input)
	value, err := strconv.Atoi(got)
	if err != nil || value > -1 {
		t.Error("randomIntWithRange should be a negative int", got)
	}
}

func TestBuiltinRandomString(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{random_string(10)}};{{random_string(123,11)}};{{random_string()}}"
	got := ctx.Patch(input)
	res := strings.Split(got, ";")

	test.Equals(t, "we should have 3 elements as results", len(res), 3)
	test.Equals(t, "should have replace the first pattern", len(res[0]), 10)
	test.Assert(t, res[0] != "{{random_string(10)}}", "Should have been patched by a random string, %q", res[0])
	test.Equals(t, "should have replace the first pattern", res[1], "{{random_string(123,11)}}")
	test.Equals(t, "should have replace the first pattern", res[2], "{{random_string()}}")
}

func TestBuiltinMd5(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{md5(TOTO)}}"
	got := ctx.Patch(input)
	want := "04c1d7cd203ef496f200ee5a096b5764"

	test.Equals(t, "MD5 should be equals", want, got)
}

func TestBuiltinEncodeBase64(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{encode_base64(TOTO)}}"
	got := ctx.Patch(input)
	want := "VE9UTw=="

	test.Equals(t, "Base64 should be equals", want, got)
}

func TestBuiltinSha1(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{sha1(TOTO)}}"
	got := ctx.Patch(input)
	want := "eefaf6bedac8f0f58af507ce3fde2a1b77b1cd89"

	test.Equals(t, "SHA1 should be equals", want, got)
}

func TestBuiltinSha256(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{sha256(TOTO)}}"
	got := ctx.Patch(input)
	want := "f2efb991e19f0edff35aa412b47e49be6f4f694028fe15598951619de915d54a"

	test.Equals(t, "SHA256 should be equals", want, got)
}

func TestBuiltinUrlEncode(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{url_encode(TOTO?TITI)}}"
	got := ctx.Patch(input)
	want := "TOTO%3FTITI"

	test.Equals(t, "SHA256 should be equals", want, got)
}

func TestBuiltinHmacSha1(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{hmac_sha1(TOTO, TITI)}}"
	got := ctx.Patch(input)
	want := "ded3645c5095344face07f545a7e8a62c9f971a3"

	test.Equals(t, "HMAC_SHA1 should be equals", want, got)
}

func TestBuiltinHmacSha256(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{hmac_sha256(TOTO, TITI)}}"
	got := ctx.Patch(input)
	want := "7e74526556d94dfb746a65cf298d8d000dff61bef5a8466b15d11f7256c01001"

	test.Equals(t, "HMAC_SHA256 should be equals", want, got)
}

func TestBuiltinTimestampFormat(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	t1 := time.Date(2020, time.April, 16, 21, 8, 17, 0, time.Local)
	timestamp := t1.Unix()

	input := fmt.Sprintf("{{format_timestamp(%d, YYYY-YY-MM-DD-HH-hh-mm-ss)}}", timestamp)
	got := ctx.Patch(input)
	want := "2020-20-04-16-21-09-08-17"

	test.Equals(t, "timestamp format should equals", want, got)
}

func TestBuiltinTimestampOffsetPositive(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{timestamp_offset(5)}}"
	got := ctx.Patch(input)
	now := time.Now()
	i, err := strconv.Atoi(got)
	if err != nil {
		t.Errorf("this is not a number %q", got)
	}

	timeGot := time.Unix(int64(i), 0)
	if now.After(timeGot) {
		t.Errorf("got %q should be after %q", got, now)
	}
}

func TestBuiltinTimestampOffsetNegative(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{timestamp_offset(-5)}}"
	got := ctx.Patch(input)
	now := time.Now()
	i, err := strconv.Atoi(got)
	if err != nil {
		t.Errorf("this is not a number %q", got)
	}

	timeGot := time.Unix(int64(i), 0)
	if now.Before(timeGot) {
		t.Errorf("got %d should be before %d", timeGot.Unix(), now.Unix())
	}
}

func TestBuiltinTimestampOffsetWrongType(t *testing.T) {
	resetContext()
	ctx := context.GetContext()
	input := "{{timestamp_offset(T)}}"
	got := ctx.Patch(input)
	test.Equals(t, "we should not have patched the value", input, got)
}
