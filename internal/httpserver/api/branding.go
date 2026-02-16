package api

import (
	"encoding/json"
	"log"
	"net/http"
	"opal/internal/config"
	"opal/internal/jfstructs"
	"os"
	"path/filepath"
)

func EndpointBrandingConfiguration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cssPath := filepath.Join(config.ConfigDirectoryPath, "theme.css")
	customCss, err := os.ReadFile(cssPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("[WARN] Failed to load custom css from %s: %v\n", cssPath, err)
		}
		customCss = []byte("")
	}

	branding := jfstructs.ResponseBrandingConfiguration{
		LoginDisclaimer:     "",
		CustomCss:           string(customCss),
		SplashscreenEnabled: false,
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false) //Needed to prevent the json encoder from mangling branding.CustomCss
	if err := enc.Encode(branding); err != nil {
		log.Printf("[ERROR] EndpointBrandingConfiguration: %v", err)
	}
}
