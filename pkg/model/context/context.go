package context

import (
	"strings"
	"sync"
)

// context is a key value store where we keep all the variable name and replacement string.
type context struct {
	variables map[string]string
}

var instance *context
var once sync.Once

// GetContext allows to get a singleton of the context.
func GetContext() *context {
	once.Do(func() {
		instance = &context{
			variables: make(map[string]string),
		}
	})
	return instance
}

// Add a new variable to the context.
func (context *context) Add(key string, value string) {
	context.variables[key] = value
}

// ResetContext remove all the variable in the context.
func (context *context) ResetContext() {
	context.variables = map[string]string{}
}

// Patch is taking a string and patch all the variable he found in the string
// a variable is something inside {{variable}}
func (context *context) Patch(str string) string {
	result := str
	for key, value := range context.variables {
		if strings.Contains(result, "{{"+key+"}}") {
			result = strings.ReplaceAll(result, "{{"+key+"}}", value)
		}
	}
	result = context.patchBuiltin(result)
	return result
}

// patchBuiltin is applying builtin patches.
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
