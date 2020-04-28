package context

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/metakeule/fmtdate"
	"github.com/thanhpk/randstr"
	hash2 "hash"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type context struct {
	variables map[string]string
}

var instance *context
var once sync.Once

func GetContext() *context {
	once.Do(func() {
		instance = &context{
			variables: make(map[string]string),
		}
	})
	return instance
}

func (context *context) AddToContext(key string, value string) {
	context.variables[key] = value
}

func (context *context) ResetContext() {
	context.variables = map[string]string{}
}

func (context *context) Patch(str string) string {
	result := str
	for key, value := range context.variables {
		if strings.Contains(result, "{{"+key+"}}") {
			result = strings.ReplaceAll(result, "{{"+key+"}}", value)
			//TODO : add to patched values
		}

	}
	result = context.patchBuiltin(result)
	return result
}

func (context *context) patchBuiltin(s string) string {
	s = timestamp(s)
	s = utcDatetime(s)
	s = randomInt(s)
	s = randUuid(s)
	s = randomIntWithRange(s)
	s = randomString(s)
	s = hash(s)
	s = hmacSha(s)
	s = formatTimestamp(s)
	s = timestampOffset(s)
	return s
}

func randomInt(s string) string {
	if key := "{{random_int}}"; strings.Contains(s, key) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		s = strings.ReplaceAll(s, key, strconv.Itoa(r1.Int()))
	}
	return s
}

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

func timestamp(s string) string {
	if key := "{{timestamp}}"; strings.Contains(s, key) {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		s = strings.ReplaceAll(s, key, timestamp)
	}
	return s
}

func utcDatetime(s string) string {
	if key := "{{utc_datetime}}"; strings.Contains(s, key) {
		loc, _ := time.LoadLocation("UTC")
		s = strings.ReplaceAll(s, key, time.Now().In(loc).Format(time.RFC3339Nano))
	}
	return s
}

func randUuid(s string) string {
	if key := "{{uuid}}"; strings.Contains(s, key) {
		s = strings.ReplaceAll(s, key, uuid.New().String())
	}
	return s
}

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
			break

		case "encode_base64":
			replaceValue = base64.StdEncoding.EncodeToString([]byte(itemToHash))
			break

		case "sha1":
			replaceValue = fmt.Sprintf("%x", sha1.Sum([]byte(itemToHash)))
			break

		case "sha256":
			replaceValue = fmt.Sprintf("%x", sha256.Sum256([]byte(itemToHash)))
			break

		case "url_encode":
			replaceValue = url.QueryEscape(itemToHash)
			break
		}
		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}

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
			break

		case "hmac_sha256":
			mac = hmac.New(sha256.New, []byte(key))
			break
		}
		mac.Write([]byte(value))
		expectedMAC := mac.Sum(nil)
		replaceValue := hex.EncodeToString(expectedMAC)
		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}

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
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		replaceValue := fmtdate.Format(format, time.Unix(timestamp, 0))


		fmt.Println("--------")
		fmt.Println(format)
		fmt.Println(timestamp)
		fmt.Println(err)
		fmt.Println(replaceValue)
		fmt.Println("--------")

		s = strings.Replace(s, variable, replaceValue, 1)
	}
	return s
}

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
