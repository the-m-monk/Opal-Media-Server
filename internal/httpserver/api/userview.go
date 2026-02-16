package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"opal/internal/jfstructs"
	"opal/internal/librarymgmt"
	"opal/internal/usermgmt"
	"strings"
)

//Despite the name, /UserViews actually gets a list of all libraries

func EndpointUserViews(w http.ResponseWriter, r *http.Request) {
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
		if t.TokenString == getHeaderAuthToken(r) && query.Get("userId") == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
	}

	res := jfstructs.ResponseUserViews{
		StartIndex:       0,
		TotalRecordCount: len(librarymgmt.AllLibraries), //TODO: add user id based library restrictions
	}

	//TODO: add user id based library restrictions

	for _, library := range librarymgmt.LibTree {
		newItem := jfstructs.ResponseUserViewsItem{
			Name:     library.Name,
			ServerID: "STUBBED",
			//This is technically incorrect behaviour, standard api responses from a jellyfin server returns stripped uuids,
			//however not stripping uuids appears to actually have no effect on the client
			ID:          library.RootUuid,
			CanDownload: true,
			//TODO: might be incorrect behaviour, investigate further with special characters
			SortName: strings.ToLower(library.Name),
			//This does not match jellyfin's api, this field is supposed to contain the actual path to the library
			//However exposing the real path to the client is a vulnerability (probably)
			Path:                     fmt.Sprintf("/Libraries/%s", library.Uuid),
			EnableMediaSourceDisplay: true, //I have no idea what this does, sounds important though
			PlayAccess:               "Full",
			IsFolder:                 true,
			ParentID:                 "STUBBED", //TODO: this feels important as all items reference this same random uuid
			Type:                     "CollectionFolder",
			ChildCount:               len(library.Children),
			//TODO: probably need to set:
			//	UserData,
			//	DisplayPreferencesId,
			//	ImageBlurHashes,
			//	PrimaryImageAspectRatio
			//	CollectionType
			LocationType: "FileSystem",
			MediaType:    "Unknown",
		}

		newItem.ImageTags.Primary = "library.jpg"
		res.Items = append(res.Items, newItem)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
