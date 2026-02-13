package jfstructs

import (
	"time"
)

type ResponseUsersItemsByUuidFolder struct {
	Name                     string             `json:"Name"`
	ServerID                 string             `json:"ServerId"`
	Id                       string             `json:"Id"`
	Etag                     string             `json:"Etag"`
	DateCreated              time.Time          `json:"DateCreated"`
	DateLastMediaAdded       time.Time          `json:"DateLastMediaAdded"`
	CanDelete                bool               `json:"CanDelete"`
	CanDownload              bool               `json:"CanDownload"`
	SortName                 string             `json:"SortName"`
	ExternalUrls             []any              `json:"ExternalUrls"`
	Path                     string             `json:"Path"`
	EnableMediaSourceDisplay bool               `json:"EnableMediaSourceDisplay"`
	ChannelID                any                `json:"ChannelId"`
	Taglines                 []any              `json:"Taglines"`
	Genres                   []any              `json:"Genres"`
	PlayAccess               string             `json:"PlayAccess"`
	RemoteTrailers           []any              `json:"RemoteTrailers"`
	ProviderIds              struct{}           `json:"ProviderIds"`
	IsFolder                 bool               `json:"IsFolder"`
	ParentId                 string             `json:"ParentId"`
	Type                     string             `json:"Type"`
	People                   []any              `json:"People"`
	Studios                  []any              `json:"Studios"`
	GenreItems               []any              `json:"GenreItems"`
	LocalTrailerCount        int                `json:"LocalTrailerCount"`
	UserData                 CommonItemUserData `json:"UserData"`
	ChildCount               int                `json:"ChildCount"`
	SpecialFeatureCount      int                `json:"SpecialFeatureCount"`
	DisplayPreferencesID     string             `json:"DisplayPreferencesId"`
	Tags                     []any              `json:"Tags"`
	PrimaryImageAspectRatio  float64            `json:"PrimaryImageAspectRatio"`
	/*
		ImageTags               struct {
			Primary string `json:"Primary"`
		} `json:"ImageTags"`
		BackdropImageTags []any `json:"BackdropImageTags"`

		ImageBlurHashes   struct {
			Primary struct {
				Two7739Ba3164E102573071433D9C803Bf string `json:"27739ba3164e102573071433d9c803bf"`
			} `json:"Primary"`
		} `json:"ImageBlurHashes"`
	*/
	LocationType string `json:"LocationType"`
	MediaType    string `json:"MediaType"`
	LockedFields []any  `json:"LockedFields"`
	LockData     bool   `json:"LockData"`
}
