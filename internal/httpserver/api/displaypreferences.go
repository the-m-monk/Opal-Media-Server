package api

import (
	"encoding/json"
	"net/http"
	"opal/internal/jfstructs"
	"opal/internal/usermgmt"
)

func EndpointDisplayPreferencesUsersettings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	query := r.URL.Query()

	aTokenString := getHeaderAuthToken(r)
	reqUserId := query.Get("userId")

	isReqValid := false

	//TODO: token lock halts any other goroutines from reading currentTokens, should change in future
	usermgmt.CurrentTokensMutex.Lock()

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == aTokenString && reqUserId == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
	}

	res := jfstructs.ResponseDisplayPreferencesUsersettings{
		ID:                 "STUBBED",
		SortBy:             "SortName",
		RememberIndexing:   false,
		PrimaryImageHeight: 250,
		PrimaryImageWidth:  250,
		ScrollDirection:    "Horizontal",
		ShowBackdrop:       true,
		RememberSorting:    false,
		SortOrder:          "Ascending",
		ShowSidebar:        false,
		Client:             query.Get("client"),
		CustomPrefs: jfstructs.ResponseDisplayPreferencesUsersettingsCustomPrefs{
			ChromecastVersion:          "stable",
			SkipForwardLength:          "30000",
			SkipBackLength:             "10000",
			EnableNextVideoInfoOverlay: "false",
			Tvhome:                     nil,
			DashboardTheme:             nil,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
