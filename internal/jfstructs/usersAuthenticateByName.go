package jfstructs

import "time"

/*
Example response:
POST https://demo.jellyfin.org/stable/Users/authenticatebyname
{
	"User":{
		"Name":"demo",
		"ServerId":"f0b3381645f04afb9a0e392e74b6a1b0",
		"Id":"4ed1b8b42a7c4ea682f0fe5d08d4e278",
		"PrimaryImageTag":"6f04f86b8bc21df4b638617ab76e0ec1",
		"HasPassword":false,
		"HasConfiguredPassword":false,
		"HasConfiguredEasyPassword":false,
		"EnableAutoLogin":false,
		"LastLoginDate":"2026-01-23T00:02:44.4648009Z",
		"LastActivityDate":"2026-01-23T00:02:44.4648009Z",
		"Configuration":{
			"AudioLanguagePreference":"",
			"PlayDefaultAudioTrack":true,
			"SubtitleLanguagePreference":"",
			"DisplayMissingEpisodes":true,
			"GroupedFolders":[],
			"SubtitleMode":"Default",
			"DisplayCollectionsView":false,
			"EnableLocalPassword":false,
			"OrderedViews":["f137a2dd21bbc1b99aa5c0f6bf02a805","7e64e319657a9516ec78490da03edccb","accbf8ec18a80e054486551fbf1306b9","a656b907eb3a73532e40e44b968d0225"],
			"LatestItemsExcludes":[],
			"MyMediaExcludes":[],
			"HidePlayedInLatest":false,
			"RememberAudioSelections":true,
			"RememberSubtitleSelections":true,
			"EnableNextEpisodeAutoPlay":true,
			"CastReceiverId":"F007D354"
		},
		"Policy":{
			"IsAdministrator":false,
			"IsHidden":false,
			"EnableCollectionManagement":false,
			"EnableSubtitleManagement":false,
			"EnableLyricManagement":false,
			"IsDisabled":false,
			"BlockedTags":[],
			"AllowedTags":[],
			"EnableUserPreferenceAccess":true,
			"AccessSchedules":[],
			"BlockUnratedItems":[],
			"EnableRemoteControlOfOtherUsers":false,
			"EnableSharedDeviceControl":false,
			"EnableRemoteAccess":true,
			"EnableLiveTvManagement":true,
			"EnableLiveTvAccess":true,
			"EnableMediaPlayback":true,
			"EnableAudioPlaybackTranscoding":true,
			"EnableVideoPlaybackTranscoding":true,
			"EnablePlaybackRemuxing":true,
			"ForceRemoteSourceTranscoding":false,
			"EnableContentDeletion":false,
			"EnableContentDeletionFromFolders":[],
			"EnableContentDownloading":false,
			"EnableSyncTranscoding":true,
			"EnableMediaConversion":true,
			"EnabledDevices":[],
			"EnableAllDevices":true,
			"EnabledChannels":[],
			"EnableAllChannels":false,
			"EnabledFolders":[],
			"EnableAllFolders":true,
			"InvalidLoginAttemptCount":0,
			"LoginAttemptsBeforeLockout":-1,
			"MaxActiveSessions":0,
			"EnablePublicSharing":true,
			"BlockedMediaFolders":[],
			"BlockedChannels":[],
			"RemoteClientBitrateLimit":0,
			"AuthenticationProviderId":"Jellyfin.Server.Implementations.Users.DefaultAuthenticationProvider",
			"PasswordResetProviderId":"Jellyfin.Server.Implementations.Users.DefaultPasswordResetProvider",
			"SyncPlayAccess":"CreateAndJoinGroups"
		}
	},
	"SessionInfo":{
		"PlayState":{
			"CanSeek":false,
			"IsPaused":false,
			"IsMuted":false,
			"RepeatMode":"RepeatNone",
			"PlaybackOrder":"Default"},
			"AdditionalUsers":[],
			"Capabilities":{
				"PlayableMediaTypes":["Audio","Video"],
				"SupportedCommands":["MoveUp","MoveDown","MoveLeft","MoveRight","PageUp","PageDown","PreviousLetter","NextLetter","ToggleOsd","ToggleContextMenu","Select","Back","SendKey","SendString","GoHome","GoToSettings","VolumeUp","VolumeDown","Mute","Unmute","ToggleMute","SetVolume","SetAudioStreamIndex","SetSubtitleStreamIndex","DisplayContent","GoToSearch","DisplayMessage","SetRepeatMode","SetShuffleQueue","ChannelUp","ChannelDown","PlayMediaSource","PlayTrailers"],
				"SupportsMediaControl":true,
				"SupportsPersistentIdentifier":false
			},
			"RemoteEndPoint":"110.32.176.54",
			"PlayableMediaTypes":["Audio","Video"],
			"Id":"69d07c2fc8c761cc1afb64d10bd559f7",
			"UserId":"4ed1b8b42a7c4ea682f0fe5d08d4e278",
			"UserName":"demo",
			"Client":"Jellyfin Web",
			"LastActivityDate":"2026-01-23T00:02:44.4786794Z",
			"LastPlaybackCheckIn":"0001-01-01T00:00:00.0000000Z",
			"DeviceName":"Firefox",
			"DeviceId":"TW96aWxsYS81LjAgKFgxMTsgTGludXggeDg2XzY0OyBydjoxNDAuMCkgR2Vja28vMjAxMDAxMDEgRmlyZWZveC8xNDAuMHwxNzY5MDgyMTgzMTY2",
			"ApplicationVersion":"10.11.6",
			"IsActive":true,
			"SupportsMediaControl":false,
			"SupportsRemoteControl":false,
			"NowPlayingQueue":[],
			"NowPlayingQueueFullItems":[],
			"HasCustomDeviceName":false,
			"ServerId":"f0b3381645f04afb9a0e392e74b6a1b0",
			"UserPrimaryImageTag":"6f04f86b8bc21df4b638617ab76e0ec1",
			"SupportedCommands":["MoveUp","MoveDown","MoveLeft","MoveRight","PageUp","PageDown","PreviousLetter","NextLetter","ToggleOsd","ToggleContextMenu","Select","Back","SendKey","SendString","GoHome","GoToSettings","VolumeUp","VolumeDown","Mute","Unmute","ToggleMute","SetVolume","SetAudioStreamIndex","SetSubtitleStreamIndex","DisplayContent","GoToSearch","DisplayMessage","SetRepeatMode","SetShuffleQueue","ChannelUp","ChannelDown","PlayMediaSource","PlayTrailers"]
		},
	"AccessToken":"30ca66c1f93e4281838d4a05107eb1be",
	"ServerId":"f0b3381645f04afb9a0e392e74b6a1b0"
}
*/

type RequestUsersAuthenticateByName struct {
	Username string `json:"Username"`
	Pw       string `json:"Pw"`
}

type ResponseUsersAuthenticateByName struct {
	User        CommonUser                                 `json:"User"`
	SessionInfo ResponseUsersAuthenticateByNameSessionInfo `json:"SessionInfo"`
	AccessToken string                                     `json:"AccessToken"`
	ServerID    string                                     `json:"ServerId"`
}

type ResponseUsersAuthenticateByNameSessionInfo struct {
	PlayState                ResponseUsersAuthenticateByNameSessionInfoPlayState    `json:"PlayState"`
	AdditionalUsers          []any                                                  `json:"AdditionalUsers"`
	Capabilities             ResponseUsersAuthenticateByNameSessionInfoCapabilities `json:"Capabilities"`
	RemoteEndPoint           string                                                 `json:"RemoteEndPoint"`
	PlayableMediaTypes       []string                                               `json:"PlayableMediaTypes"`
	ID                       string                                                 `json:"Id"`
	UserID                   string                                                 `json:"UserId"`
	UserName                 string                                                 `json:"UserName"`
	Client                   string                                                 `json:"Client"`
	LastActivityDate         time.Time                                              `json:"LastActivityDate"`
	LastPlaybackCheckIn      time.Time                                              `json:"LastPlaybackCheckIn"`
	DeviceName               string                                                 `json:"DeviceName"`
	DeviceID                 string                                                 `json:"DeviceId"`
	ApplicationVersion       string                                                 `json:"ApplicationVersion"`
	IsActive                 bool                                                   `json:"IsActive"`
	SupportsMediaControl     bool                                                   `json:"SupportsMediaControl"`
	SupportsRemoteControl    bool                                                   `json:"SupportsRemoteControl"`
	NowPlayingQueue          []any                                                  `json:"NowPlayingQueue"`
	NowPlayingQueueFullItems []any                                                  `json:"NowPlayingQueueFullItems"`
	HasCustomDeviceName      bool                                                   `json:"HasCustomDeviceName"`
	ServerID                 string                                                 `json:"ServerId"`
	UserPrimaryImageTag      string                                                 `json:"UserPrimaryImageTag"`
	SupportedCommands        []string                                               `json:"SupportedCommands"`
}

type ResponseUsersAuthenticateByNameSessionInfoPlayState struct {
	CanSeek       bool   `json:"CanSeek"`
	IsPaused      bool   `json:"IsPaused"`
	IsMuted       bool   `json:"IsMuted"`
	RepeatMode    string `json:"RepeatMode"`
	PlaybackOrder string `json:"PlaybackOrder"`
}

type ResponseUsersAuthenticateByNameSessionInfoCapabilities struct {
	PlayableMediaTypes           []string `json:"PlayableMediaTypes"`
	SupportedCommands            []string `json:"SupportedCommands"`
	SupportsMediaControl         bool     `json:"SupportsMediaControl"`
	SupportsPersistentIdentifier bool     `json:"SupportsPersistentIdentifier"`
}
