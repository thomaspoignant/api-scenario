package config

import (
	"errors"
	"flag"
	"strings"
)

var Configuration Config

type Config struct {
	accept        string
	acceptCharset string
	authorization string
	baseUrl       string
	contentType   string
	userAgent     string
	pauseDuration int
}

func InitConfig() (Config, error) {
	accept := flag.String("Accept", "application/scim+json", "Accept-Charset header to override.")
	acceptCharset := flag.String("Accept-Charset", "utf-8", "Accept-Charset header to override.")
	contentType := flag.String("Content-Type", "application/scim+json; charset=utf-8", "Content-Type header to override.")
	authorization := flag.String("Authorization", "", "[MANDATORY] Authorization header to pass to all the request.")
	userAgent := flag.String("User-Agent", "SCIM Integration (rest-scenario)", "User-Agent header to override.")
	pauseDuration := flag.Int("Pause-Duration", 5.0, "Pause between each API call.")
	baseUrl := flag.String("Base-Url", "", "[MANDATORY] Base URL of your SCIM API.")
	flag.Parse()

	// check mandatory parameters
	// TODO: uncomment this
	//err := mandatoryStrings(*authorization, *baseUrl)
	//if err != nil {
	//	return Config{}, err
	//}

	// format authorization header
	authorizationHeader, err := formatAuthorization(*authorization)
	if err != nil {
		return Config{}, err
	}

	Configuration = Config{
		accept:        *accept,
		acceptCharset: *acceptCharset,
		contentType:   *contentType,
		authorization: authorizationHeader,
		userAgent:     *userAgent,
		pauseDuration: *pauseDuration,
		baseUrl:       *baseUrl,
	}

	return Configuration, nil
}

func formatAuthorization(param string) (string, error) {

	// Check if param starts with "Bearer"
	const bearer = "bearer "
	if strings.HasPrefix(strings.ToLower(param), bearer) {
		return param, nil
	}

	return bearer + param, nil
}

func mandatoryStrings(params ...string) error {
	for _, param := range params {
		if len(strings.TrimSpace(param)) == 0 {
			return errors.New("missing mandatory parameter")
		}
	}
	return nil
}
