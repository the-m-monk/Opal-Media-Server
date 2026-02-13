package jfstructs

import "time"

type CommonUsersConfiguration struct {
	AudioLanguagePreference    string   `json:"AudioLanguagePreference,omitempty"` //usersPublic does not implement but usersAuthenticateByName does
	PlayDefaultAudioTrack      bool     `json:"PlayDefaultAudioTrack"`
	SubtitleLanguagePreference string   `json:"SubtitleLanguagePreference"`
	DisplayMissingEpisodes     bool     `json:"DisplayMissingEpisodes"`
	GroupedFolders             []any    `json:"GroupedFolders"`
	SubtitleMode               string   `json:"SubtitleMode"`
	DisplayCollectionsView     bool     `json:"DisplayCollectionsView"`
	EnableLocalPassword        bool     `json:"EnableLocalPassword"`
	OrderedViews               []string `json:"OrderedViews"`
	LatestItemsExcludes        []any    `json:"LatestItemsExcludes"`
	MyMediaExcludes            []any    `json:"MyMediaExcludes"`
	HidePlayedInLatest         bool     `json:"HidePlayedInLatest"`
	RememberAudioSelections    bool     `json:"RememberAudioSelections"`
	RememberSubtitleSelections bool     `json:"RememberSubtitleSelections"`
	EnableNextEpisodeAutoPlay  bool     `json:"EnableNextEpisodeAutoPlay"`
	CastReceiverID             string   `json:"CastReceiverId"`
}

type CommonUsersPolicy struct {
	IsAdministrator                  bool   `json:"IsAdministrator"`
	IsHidden                         bool   `json:"IsHidden"`
	EnableCollectionManagement       bool   `json:"EnableCollectionManagement"`
	EnableSubtitleManagement         bool   `json:"EnableSubtitleManagement"`
	EnableLyricManagement            bool   `json:"EnableLyricManagement"`
	IsDisabled                       bool   `json:"IsDisabled"`
	BlockedTags                      []any  `json:"BlockedTags"`
	AllowedTags                      []any  `json:"AllowedTags"`
	EnableUserPreferenceAccess       bool   `json:"EnableUserPreferenceAccess"`
	AccessSchedules                  []any  `json:"AccessSchedules"`
	BlockUnratedItems                []any  `json:"BlockUnratedItems"`
	EnableRemoteControlOfOtherUsers  bool   `json:"EnableRemoteControlOfOtherUsers"`
	EnableSharedDeviceControl        bool   `json:"EnableSharedDeviceControl"`
	EnableRemoteAccess               bool   `json:"EnableRemoteAccess"`
	EnableLiveTvManagement           bool   `json:"EnableLiveTvManagement"`
	EnableLiveTvAccess               bool   `json:"EnableLiveTvAccess"`
	EnableMediaPlayback              bool   `json:"EnableMediaPlayback"`
	EnableAudioPlaybackTranscoding   bool   `json:"EnableAudioPlaybackTranscoding"`
	EnableVideoPlaybackTranscoding   bool   `json:"EnableVideoPlaybackTranscoding"`
	EnablePlaybackRemuxing           bool   `json:"EnablePlaybackRemuxing"`
	ForceRemoteSourceTranscoding     bool   `json:"ForceRemoteSourceTranscoding"`
	EnableContentDeletion            bool   `json:"EnableContentDeletion"`
	EnableContentDeletionFromFolders []any  `json:"EnableContentDeletionFromFolders"`
	EnableContentDownloading         bool   `json:"EnableContentDownloading"`
	EnableSyncTranscoding            bool   `json:"EnableSyncTranscoding"`
	EnableMediaConversion            bool   `json:"EnableMediaConversion"`
	EnabledDevices                   []any  `json:"EnabledDevices"`
	EnableAllDevices                 bool   `json:"EnableAllDevices"`
	EnabledChannels                  []any  `json:"EnabledChannels"`
	EnableAllChannels                bool   `json:"EnableAllChannels"`
	EnabledFolders                   []any  `json:"EnabledFolders"`
	EnableAllFolders                 bool   `json:"EnableAllFolders"`
	InvalidLoginAttemptCount         int    `json:"InvalidLoginAttemptCount"`
	LoginAttemptsBeforeLockout       int    `json:"LoginAttemptsBeforeLockout"`
	MaxActiveSessions                int    `json:"MaxActiveSessions"`
	EnablePublicSharing              bool   `json:"EnablePublicSharing"`
	BlockedMediaFolders              []any  `json:"BlockedMediaFolders"`
	BlockedChannels                  []any  `json:"BlockedChannels"`
	RemoteClientBitrateLimit         int    `json:"RemoteClientBitrateLimit"`
	AuthenticationProviderID         string `json:"AuthenticationProviderId"`
	PasswordResetProviderID          string `json:"PasswordResetProviderId"`
	SyncPlayAccess                   string `json:"SyncPlayAccess"`
}

type CommonUser struct {
	Name                      string                   `json:"Name"`
	ServerID                  string                   `json:"ServerId"`
	ID                        string                   `json:"Id"`
	PrimaryImageTag           string                   `json:"PrimaryImageTag,omitempty"`
	HasPassword               bool                     `json:"HasPassword"`
	HasConfiguredPassword     bool                     `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword bool                     `json:"HasConfiguredEasyPassword"`
	EnableAutoLogin           bool                     `json:"EnableAutoLogin"`
	LastLoginDate             time.Time                `json:"LastLoginDate"`
	LastActivityDate          time.Time                `json:"LastActivityDate"`
	Configuration             CommonUsersConfiguration `json:"Configuration"`
	Policy                    CommonUsersPolicy        `json:"Policy"`
}

type CommonItemUserData struct {
	PlayedPercentage      float64   `json:"PlayedPercentage"`
	PlaybackPositionTicks int64     `json:"PlaybackPositionTicks"`
	PlayCount             int       `json:"PlayCount"`
	IsFavorite            bool      `json:"IsFavorite"`
	LastPlayedDate        time.Time `json:"LastPlayedDate"`
	Played                bool      `json:"Played"`
	Key                   string    `json:"Key"`
	ItemID                string    `json:"ItemId"`
}

type CommonItem struct {
	Name                    string             `json:"Name"`
	ServerID                string             `json:"ServerId"`
	ID                      string             `json:"Id"`
	Path                    string             `json:"Path"`
	HasSubtitles            bool               `json:"HasSubtitles"`
	Container               string             `json:"Container"`
	PremiereDate            time.Time          `json:"PremiereDate"`
	ChannelID               any                `json:"ChannelId"`
	RunTimeTicks            int64              `json:"RunTimeTicks"`
	ProductionYear          int                `json:"ProductionYear"`
	IsFolder                bool               `json:"IsFolder"`
	Type                    string             `json:"Type"`
	UserData                CommonItemUserData `json:"UserData"`
	PrimaryImageAspectRatio float64            `json:"PrimaryImageAspectRatio"`
	VideoType               string             `json:"VideoType"`
	ImageTags               struct {
		Primary string `json:"Primary"`
	} `json:"ImageTags"`
	/*

		BackdropImageTags []any `json:"BackdropImageTags"`
		ImageBlurHashes   struct {
			Primary struct {
				Four1B26B069Aaad042A8293C6Fe41D9738 string `json:"41b26b069aaad042a8293c6fe41d9738"`
			} `json:"Primary"`
		} `json:"ImageBlurHashes"`
	*/
	LocationType string `json:"LocationType"`
	MediaType    string `json:"MediaType"`
}

type CommonItemList struct {
	Items            []CommonItem `json:"Items"`
	TotalRecordCount int          `json:"TotalRecordCount"`
	StartIndex       int          `json:"StartIndex"`
}
