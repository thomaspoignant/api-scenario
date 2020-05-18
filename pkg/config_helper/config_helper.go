package config_helper

import (
	"strings"
)

func FormatAuthorization(param string) string {
	// Check if param starts with "Bearer"
	const bearer = "bearer "
	if strings.HasPrefix(strings.ToLower(param), bearer) {
		return param
	}

	return bearer + param
}
