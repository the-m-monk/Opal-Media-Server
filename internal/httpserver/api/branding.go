package api

import (
	"encoding/json"
	"log/slog"
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
			slog.Warn("failed to laod custom css", "path", cssPath, "reason", err)
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
	enc.Encode(branding)
}
