package api

import (
	"encoding/json"
	"net/http"
	"opal/internal/jfstructs"
)

func EndpointBrandingConfiguration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	branding := jfstructs.ResponseBrandingConfiguration{
		LoginDisclaimer:     "",
		CustomCss:           "",
		SplashscreenEnabled: false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(branding)
}
