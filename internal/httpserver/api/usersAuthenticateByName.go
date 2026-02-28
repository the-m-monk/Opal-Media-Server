package api

import (
	"encoding/json"
	"net/http"
	"opal/internal/jfstructs"
	"opal/internal/usermgmt"
	"strings"
)

func EndpointUsersAuthenticateByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var authReq jfstructs.RequestUsersAuthenticateByName
	err := json.NewDecoder(r.Body).Decode(&authReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var userRecord usermgmt.UserRecord

	//TODO: add lookup map
	for _, user := range usermgmt.UserDB {
		if user.Username == authReq.Username {
			userRecord = user
		}
	}

	if userRecord.Username == "" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	var aT usermgmt.AccessToken

	if usermgmt.Authenticate(userRecord, authReq.Pw, &aT) != "SUCCESS" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	aT.UserId = usermgmt.GetId(authReq.Username)
	aT.UserName = authReq.Username

	//TODO: replace with rwmutex, also inforce usage for endpoints
	usermgmt.CurrentTokensMutex.Lock()
	usermgmt.CurrentTokens = append(usermgmt.CurrentTokens, aT)
	usermgmt.CurrentTokensMutex.Unlock()

	h := r.Header.Get("Authorization")
	h = strings.TrimPrefix(h, "MediaBrowser ")

	authMap := make(map[string]string)

	authPairs := strings.Split(h, ",")

	for _, pair := range authPairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			val := strings.Trim(strings.TrimSpace(kv[1]), "\"")
			authMap[key] = val
		}
	}

	response := jfstructs.ResponseUsersAuthenticateByName{
		AccessToken: aT.TokenString,
		ServerID:    "STUBBED",
		User: jfstructs.CommonUser{
			Name:                      aT.UserName,
			ServerID:                  "STUBBED",
			ID:                        aT.UserId,
			HasPassword:               true,
			HasConfiguredPassword:     true,
			HasConfiguredEasyPassword: false, //TODO: unsure of what this option does
			EnableAutoLogin:           true,
			//"LastLoginDate": "2026-02-26T00:44:12.1394014Z",
			//"LastActivityDate": "2026-02-26T00:44:12.1394014Z",
			Configuration: jfstructs.CommonUsersConfiguration{
				PlayDefaultAudioTrack:      true,
				SubtitleLanguagePreference: "",
				DisplayMissingEpisodes:     false,
				GroupedFolders:             []any{},
				SubtitleMode:               "Default",
				DisplayCollectionsView:     false,
				EnableLocalPassword:        false,
				OrderedViews:               []string{},
				LatestItemsExcludes:        []any{},
				MyMediaExcludes:            []any{},
				HidePlayedInLatest:         true,
				RememberAudioSelections:    true,
				RememberSubtitleSelections: true,
				EnableNextEpisodeAutoPlay:  true,
				CastReceiverID:             "STUBBED",
			},
			Policy: jfstructs.CommonUsersPolicy{
				IsAdministrator:                  false,
				IsHidden:                         true,
				EnableCollectionManagement:       false,
				EnableSubtitleManagement:         false,
				EnableLyricManagement:            false,
				IsDisabled:                       false,
				BlockedTags:                      []any{},
				AllowedTags:                      []any{},
				EnableUserPreferenceAccess:       true,
				AccessSchedules:                  []any{},
				BlockUnratedItems:                []any{},
				EnableRemoteControlOfOtherUsers:  false,
				EnableSharedDeviceControl:        false,
				EnableRemoteAccess:               true,
				EnableLiveTvManagement:           false, //NOTE: live tv will probably never be supported, seems like alot of work for nothing
				EnableLiveTvAccess:               false,
				EnableMediaPlayback:              true,
				EnableAudioPlaybackTranscoding:   true,
				EnableVideoPlaybackTranscoding:   true,
				EnablePlaybackRemuxing:           true,
				ForceRemoteSourceTranscoding:     false,
				EnableContentDeletion:            false,
				EnableContentDeletionFromFolders: []any{},
				EnableContentDownloading:         true,
				EnableSyncTranscoding:            true,
				EnableMediaConversion:            true,
				EnabledDevices:                   []any{},
				EnableAllDevices:                 true,
				EnabledChannels:                  []any{},
				EnableAllChannels:                true,
				EnabledFolders:                   []any{},
				EnableAllFolders:                 true,
				InvalidLoginAttemptCount:         0,
				LoginAttemptsBeforeLockout:       -1,
				MaxActiveSessions:                0,
				EnablePublicSharing:              true,
				BlockedMediaFolders:              []any{},
				BlockedChannels:                  []any{},
				RemoteClientBitrateLimit:         0,
				AuthenticationProviderID:         "Jellyfin.Server.Implementations.Users.DefaultAuthenticationProvider",
				PasswordResetProviderID:          "Jellyfin.Server.Implementations.Users.DefaultPasswordResetProvider",
				SyncPlayAccess:                   "CreateAndJoinGroups",
			},
		},
		SessionInfo: jfstructs.ResponseUsersAuthenticateByNameSessionInfo{
			PlayState: jfstructs.ResponseUsersAuthenticateByNameSessionInfoPlayState{
				CanSeek:       false,
				IsPaused:      false,
				IsMuted:       false,
				RepeatMode:    "RepeatNone",
				PlaybackOrder: "Default",
			},
			AdditionalUsers: []any{},
			Capabilities: jfstructs.ResponseUsersAuthenticateByNameSessionInfoCapabilities{
				PlayableMediaTypes: []string{"Video", "Audio"},
				SupportedCommands: []string{"MoveUp",
					"MoveDown",
					"MoveLeft",
					"MoveRight",
					"PageUp",
					"PageDown",
					"PreviousLetter",
					"NextLetter",
					"ToggleOsd",
					"ToggleContextMenu",
					"Select",
					"Back",
					"SendKey",
					"SendString",
					"GoHome",
					"GoToSettings",
					"VolumeUp",
					"VolumeDown",
					"Mute",
					"Unmute",
					"ToggleMute",
					"SetVolume",
					"SetAudioStreamIndex",
					"SetSubtitleStreamIndex",
					"DisplayContent",
					"GoToSearch",
					"DisplayMessage",
					"SetRepeatMode",
					"SetShuffleQueue",
					"ChannelUp",
					"ChannelDown",
					"PlayMediaSource",
					"PlayTrailers",
				},
				SupportsMediaControl: true,
				//Opal uses v5 uuids that are persistent so this could be set to true but jellyfin sets this to false and so clients probably expect this to be false
				SupportsPersistentIdentifier: false,
			},
			//"RemoteEndPoint": "127.0.0.1",
			PlayableMediaTypes: []string{"Video", "Audio"},
			//I think this is the uuid of the SessionInfo object, Opal is mostly stateless and creates SessionInfo upon request instead of storing it as an object so we'll just hardcode it
			ID:       "STUBBED_ID_EndpointUsersAuthenticateByName",
			UserID:   aT.UserId,
			UserName: aT.UserName,
			Client:   authMap["Client"],
			/*
			   "LastActivityDate": "2026-02-26T00:44:12.1405613Z",
			   "LastPlaybackCheckIn": "0001-01-01T00:00:00.0000000Z",
			*/
			DeviceName:               authMap["Device"],
			DeviceID:                 authMap["DeviceId"],
			ApplicationVersion:       authMap["Version"],
			IsActive:                 true,
			SupportsMediaControl:     false,
			SupportsRemoteControl:    false,
			NowPlayingQueue:          []any{},
			NowPlayingQueueFullItems: []any{},
			HasCustomDeviceName:      false,
			ServerID:                 "STUBBED",
			SupportedCommands: []string{"MoveUp",
				"MoveDown",
				"MoveLeft",
				"MoveRight",
				"PageUp",
				"PageDown",
				"PreviousLetter",
				"NextLetter",
				"ToggleOsd",
				"ToggleContextMenu",
				"Select",
				"Back",
				"SendKey",
				"SendString",
				"GoHome",
				"GoToSettings",
				"VolumeUp",
				"VolumeDown",
				"Mute",
				"Unmute",
				"ToggleMute",
				"SetVolume",
				"SetAudioStreamIndex",
				"SetSubtitleStreamIndex",
				"DisplayContent",
				"GoToSearch",
				"DisplayMessage",
				"SetRepeatMode",
				"SetShuffleQueue",
				"ChannelUp",
				"ChannelDown",
				"PlayMediaSource",
				"PlayTrailers",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
