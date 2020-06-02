package util

import (
	"encoding/json"
	"io"
	"regexp"
	"strings"
)

func IsJson(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func JsonConvertKeyName(property string) []string {
	regex, _ := regexp.Compile(`([^\.\[.*\]]+)`)
	return regex.FindAllString(property, -1)
}

func StringToJson(s string) (map[string]interface{}, error) {
	// Parse the document in Json
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(s))
	if err := dec.Decode(&data); err != io.EOF && err != nil {
		return data, err
	}

	return data, nil
}
