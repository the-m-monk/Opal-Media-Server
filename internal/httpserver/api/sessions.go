package api

import (
	"net/http"
)

/*
POST https://demo.jellyfin.org/stable/Sessions/Capabilities/Full
Client sends playback and other capabilities
*/

func EndpointSessionsCapabilitiesFull(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
