package api

import (
	"encoding/json"
	"net/http"
	"opal/internal/jfstructs"
	"opal/internal/usermgmt"
)

func EndpointShowsNextUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	//TODO: implement user id based library restrictions
	query := r.URL.Query()

	isReqValid := false

	//TODO: token lock halts any other goroutines from reading currentTokens, should change in future
	usermgmt.CurrentTokensMutex.Lock()

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == getHeaderAuthToken(r) && query.Get("UserId") == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
	}

	res := jfstructs.CommonItemList{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
