package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from
// The headers of an HTTP request
// Example:
// Authorization: ApiKey {insert ApiKey here}
func GetAPIKey(headers http.Header) (string, error){
	value := headers.Get("Authorization")

	if value == "" {
		return "", errors.New("no authentication found")
	}

	vals := strings.Split(value, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}