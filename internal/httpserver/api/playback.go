package api

import "net/http"

func EndpointPlaybackBitrateTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	http.Error(w, "STUBBED", http.StatusNotImplemented)
}
