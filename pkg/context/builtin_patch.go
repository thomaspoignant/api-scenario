package context

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	hash2 "hash"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/metakeule/fmtdate"
	"github.com/thanhpk/randstr"
)

// randomInt replace {{random_int}} by a Random integer between 0 and 18446744073709551615
func randomInt(s string) string {
	if key := "{{random_int}}"; strings.Contains(s, key) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		s = strings.ReplaceAll(s, key, strconv.Itoa(r1.Int()))
	}
	return s
}

// randomIntWithRange replace {{random_int(a,b)}} by a Random integer value between a and b, inclusive.
func randomIntWithRange(s string) string {
	r := regexp.MustCompile(`{{random_int\(([-]?[0-9]+),([-]?[0-9]+)\)}}`)
	for _, subMatch := range r.FindAllStringSubmatch(s, -1) {
		variable := subMatch[0]
		start := subMatch[1]
		end := subMatch[2]

		// We don't check the error because regex ensure we have int as value
		min, _ := strconv.Atoi(start)
		max, _ := strconv.Atoi(end)

		if min > max {
			oldMin := min
			min = max
			max = oldMin
		}

		rand.Seed(time.Now().UnixNano())
		replaceValue := strconv.Itoa(rand.Intn(max-min+1) + min)
		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}

// randomString replace {{random_string(length)}} by a Random alphanumeric string of the specified length
// (max 1000 characters).
func randomString(s string) string {
	r := regexp.MustCompile(`{{random_string\(([0-9]+)\)}}`)
	for _, subMatch := range r.FindAllStringSubmatch(s, -1) {
		variable := subMatch[0]
		extractedLength := subMatch[1]

		// We are sure that Atoi will work cause extractedLength had to be a int in the regex
		length, _ := strconv.Atoi(extractedLength)
		replaceValue := randstr.String(length)
		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}

// timestamp replace {{timestamp}} by the current Integer Unix timestamp (seconds elapsed since January 1, 1970 00:00 UTC).
func timestamp(s string) string {
	if key := "{{timestamp}}"; strings.Contains(s, key) {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		s = strings.ReplaceAll(s, key, timestamp)
	}
	return s
}

// utcDatetime replace {{utc_datetime}} by the current UTC datetime string in ISO 8601 format.
func utcDatetime(s string) string {
	if key := "{{utc_datetime}}"; strings.Contains(s, key) {
		loc, _ := time.LoadLocation("UTC")
		s = strings.ReplaceAll(s, key, time.Now().In(loc).Format(time.RFC3339Nano))
	}
	return s
}

// randUuid replace {{uuid}} by a Random universally unique identifier (UUID)..
func randUuid(s string) string {
	if key := "{{uuid}}"; strings.Contains(s, key) {
		s = strings.ReplaceAll(s, key, uuid.New().String())
	}
	return s
}

// hash a value and replace the variable using different hashing methods. Values should be passed without surrounding quotes.
func hash(s string) string {
	r := regexp.MustCompile(`{{(md5|encode_base64|sha1|sha256|url_encode)\((.+)\)}}`)
	for _, subMatch := range r.FindAllStringSubmatch(s, -1) {
		variable := subMatch[0]
		algorithm := subMatch[1]
		itemToHash := subMatch[2]

		var replaceValue string = variable
		switch algorithm {
		case "md5":
			replaceValue = fmt.Sprintf("%x", md5.Sum([]byte(itemToHash)))
		case "encode_base64":
			replaceValue = base64.StdEncoding.EncodeToString([]byte(itemToHash))
		case "sha1":
			replaceValue = fmt.Sprintf("%x", sha1.Sum([]byte(itemToHash)))
		case "sha256":
			replaceValue = fmt.Sprintf("%x", sha256.Sum256([]byte(itemToHash)))
		case "url_encode":
			replaceValue = url.QueryEscape(itemToHash)
		}
		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}

// hmacSha a value and replace the variable using different hmac methods. Values should be passed without surrounding quotes.
func hmacSha(s string) string {
	r := regexp.MustCompile(`{{(hmac_sha1|hmac_sha256)\((.+),(.+)\)}}`)
	for _, subMatch := range r.FindAllStringSubmatch(s, -1) {
		variable := subMatch[0]
		algorithm := subMatch[1]
		value := strings.TrimSpace(subMatch[2])
		key := strings.TrimSpace(subMatch[3])

		var mac hash2.Hash
		switch algorithm {
		case "hmac_sha1":
			mac = hmac.New(sha1.New, []byte(key))
		case "hmac_sha256":
			mac = hmac.New(sha256.New, []byte(key))
		}
		if _, err := mac.Write([]byte(value)); err != nil {
			panic(fmt.Sprintf("hmac of %q: %v", s, err))
		}
		expectedMAC := mac.Sum(nil)
		replaceValue := hex.EncodeToString(expectedMAC)
		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}

/*
formatTimestamp replace {{format_timestamp(value, format)}} by a formatted value of the timestamp.
Timestamp of the specified value in the specified format.
Any delimiters (e.g. -, /, ., *, etc.) can be used in the format with a combination of any of the following date/time
format options. Also accepts variables. E.g. {{format_timestamp({{timestamp}}, YYYY-MM-DD)}}

    YYYY - 4 digit year (e.g. 2019)
    YYYY - 4 digit year (e.g. 2016)
    YY - 2 digit year (e.g. 16)
    MM - month
    DD - day
    HH - 24 hour (e.g. 13 == 1pm)
    hh - 12 hour (e.g. 01 == 1pm)
    mm - minutes
    ss - seconds
*/
func formatTimestamp(s string) string {
	r := regexp.MustCompile(`{{format_timestamp\(([0-9]+),(.+)\)}}`)
	for _, subMatch := range r.FindAllStringSubmatch(s, -1) {
		variable := subMatch[0]
		timestampStr := strings.TrimSpace(subMatch[1])
		format := strings.TrimSpace(subMatch[2])

		// replace pattern to match
		format = strings.ReplaceAll(format, "hh", "h")
		format = strings.ReplaceAll(format, "HH", "hh")

		// ParseInt will work because regex extract an int
		timestamp, _ := strconv.ParseInt(timestampStr, 10, 64)
		replaceValue := fmtdate.Format(format, time.Unix(timestamp, 0))
		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}

// timestampOffset replace {{timestamp_offset(value)}} by Integer Unix timestamp offset by the specified value in seconds (going back in time would
// be a negative offset value). Values should be passed without surrounding quotes.
func timestampOffset(s string) string {
	r := regexp.MustCompile(`{{timestamp_offset\(([-]?[0-9]+)\)}}`)
	for _, subMatch := range r.FindAllStringSubmatch(s, -1) {
		variable := subMatch[0]
		value := strings.TrimSpace(subMatch[1])

		// ParseInt will work because regex extract an int
		valueInSeconds, _ := strconv.ParseInt(value, 10, 64)
		timeIn := time.Now().Local().Add(time.Second * time.Duration(valueInSeconds))
		replaceValue := strconv.FormatInt(timeIn.Unix(), 10)
		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}
