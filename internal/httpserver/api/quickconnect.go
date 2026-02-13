package api

import (
	"encoding/json"
	"net/http"
)

/*
GET https://demo.jellyfin.org/stable/QuickConnect/Enabled
Response (json):
true
*/

func EndpointQuickConnectEnabled(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(false)
}
