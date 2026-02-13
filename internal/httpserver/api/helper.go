package api

import (
	"net/http"
	"opal/internal/usermgmt"
	"strings"
)

func getHeaderAuthToken(r *http.Request) string {
	h := r.Header.Get("Authorization")
	authKeyVals := strings.Split(h, " ")

	for _, keyval := range authKeyVals {
		if strings.HasPrefix(keyval, `Token="`) {
			return strings.TrimSuffix(strings.TrimPrefix(keyval, `Token="`), `"`)
		}
	}

	return ""
}

func isHeaderAuthTokenValid(r *http.Request) bool {
	token := getHeaderAuthToken(r)
	if token == "" {
		return false
	}

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == token {
			return true
		}
	}

	return false
}
