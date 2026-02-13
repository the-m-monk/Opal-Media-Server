package api

import (
	"encoding/json"
	"net/http"
	"opal/internal/jfstructs"
)

func EndpointSystemInfoPublic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	info := jfstructs.ResponseSystemInfoPublic{
		LocalAddress:           r.Host, //TODO: this is probably incorrect behaviour
		ServerName:             "STUBBED",
		Version:                "10.11.6",
		ProductName:            "Jellyfin Server",
		OperatingSystem:        "",
		Id:                     "STUBBED",
		StartupWizardCompleted: true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func EndpointSystemInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	info := jfstructs.ResponseSystemInfo{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

/*
GET https://demo.jellyfin.org/stable/System/Endpoint
{"IsLocal":false,"IsInNetwork":false}
*/

func EndpointSystemEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//TODO: implement
	w.Write([]byte(`{"IsLocal":false,"IsInNetwork":false}`))
}
