package util

import (
	"strings"
)

// AddBearerPrefix is formatting a token to be sure to have the Bearer prefix
func AddBearerPrefix(param string) string {
	const bearer = "Bearer "
	param = strings.TrimSpace(param)
	if strings.HasPrefix(strings.ToLower(param), strings.ToLower(bearer)) {
		return param
	}
	return bearer + strings.TrimSpace(param)
}
