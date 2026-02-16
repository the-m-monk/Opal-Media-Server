package httpserver

import (
	"fmt"
	"net/http"
	"opal/internal/config"
	"opal/internal/httpserver/api"
)

func Start() {
	Addr := config.FetchValue("/server.cfg", "address", true)
	Port := config.FetchValue("/server.cfg", "port", true)

	fileServer := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fileServer))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/web/index.html", http.StatusMovedPermanently)
			return
		}
		http.NotFound(w, r)
	})

	printMux := printRequests(http.DefaultServeMux)

	http.HandleFunc("/System/Info/Public", api.EndpointSystemInfoPublic)
	http.HandleFunc("/System/Info", api.EndpointSystemInfo)
	http.HandleFunc("/System/Endpoint", api.EndpointSystemEndpoint)

	http.HandleFunc("/QuickConnect/Enabled", api.EndpointQuickConnectEnabled)

	http.HandleFunc("/users/public", api.EndpointUsersPublic)
	http.HandleFunc("/Users/authenticatebyname", api.EndpointUsersAuthenticateByName)
	http.HandleFunc("/Users/{id}", api.EndpointUsersById)
	http.HandleFunc("/Users/{id}/Items/Resume", api.EndpointUsersItemsResume)
	http.HandleFunc("/Users/{id}/Items/Latest", api.EndpointUsersItemsLatest)
	http.HandleFunc("/Users/{id}/Items/{itemUuid}", api.EndpointUsersItemsByUuid)
	http.HandleFunc("/Users/{id}/Items", api.EndpointUsersItems)

	http.HandleFunc("/UserViews", api.EndpointUserViews)

	http.HandleFunc("/Branding/Configuration", api.EndpointBrandingConfiguration)

	http.HandleFunc("/Sessions/Capabilities/Full", api.EndpointSessionsCapabilitiesFull)

	http.HandleFunc("/Playback/BitrateTest", api.EndpointPlaybackBitrateTest)

	http.HandleFunc("/DisplayPreferences/usersettings", api.EndpointDisplayPreferencesUsersettings)

	http.HandleFunc("/Items/{uuid}/Images/Primary", api.EndpointItemsByUuidImagesPrimary)

	http.HandleFunc("/Shows/NextUp", api.EndpointShowsNextUp)

	fmt.Printf("Server starting on http://%s:%s\n", Addr, Port)
	http.ListenAndServe(Addr+":"+Port, printMux)
}

func printRequests(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, pattern := mux.Handler(r)

		isUnmapped := pattern == "" || pattern == "/"
		isNotRoot := r.URL.Path != "/"

		if isUnmapped && isNotRoot {
			//if !strings.HasPrefix(r.URL.Path, "/web") {
			fmt.Printf("Unmapped request: %s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
		}

		mux.ServeHTTP(w, r)
	})
}
