package config_helper

import (
	"strings"
)

func FormatAuthorization(param string) string {
	// Check if param starts with "Bearer"
	const bearer = "Bearer "
	param = strings.TrimSpace(param)
	if strings.HasPrefix(strings.ToLower(param), strings.ToLower(bearer)) {
		return param
	}
	return bearer + strings.TrimSpace(param)
}
