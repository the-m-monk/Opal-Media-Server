package jfstructs

import "time"

//TODO: userview returns an extended version of CommonItem that is easier to keep as a seperate defintion rather than extend CommonItem

/*
GET /UserViews?userId=ea2c447b4f6640ff83d0b65b0fe4b92a

{
    "Items": [
        {
            "Name": "Movies",
            "ServerId": "85e7e9a4fdd743bea5510c6401540719",
            "Id": "f137a2dd21bbc1b99aa5c0f6bf02a805",
            "Etag": "a90d43b3f3a28363e447622e61b93b01",
            "DateCreated": "2026-01-24T12:39:13.6612996Z",
            "DateLastMediaAdded": "0001-01-01T00:00:00.0000000Z",
            "CanDelete": false,
            "CanDownload": false,
            "SortName": "movies",
            "ExternalUrls": [],
            "Path": "/home/mmonk/.local/share/jellyfin/root/default/Movies",
            "EnableMediaSourceDisplay": true,
            "ChannelId": null,
            "Taglines": [],
            "Genres": [],
            "PlayAccess": "Full",
            "RemoteTrailers": [],
            "ProviderIds": {},
            "IsFolder": true,
            "ParentId": "e9d5075a555c1cbc394eec4cef295274",
            "Type": "CollectionFolder",
            "People": [],
            "Studios": [],
            "GenreItems": [],
            "LocalTrailerCount": 0,
            "UserData": {
                "PlaybackPositionTicks": 0,
                "PlayCount": 0,
                "IsFavorite": false,
                "Played": false,
                "Key": "f137a2dd-21bb-c1b9-9aa5-c0f6bf02a805",
                "ItemId": "f137a2dd21bbc1b99aa5c0f6bf02a805"
            },
            "ChildCount": 2,
            "SpecialFeatureCount": 0,
            "DisplayPreferencesId": "f137a2dd21bbc1b99aa5c0f6bf02a805",
            "Tags": [],
            "PrimaryImageAspectRatio": 1.7777777777777777,
            "CollectionType": "movies",
            "ImageTags": {
                "Primary": "ec22d3b46cbb21c6b1c9d0f041123c1c"
            },
            "BackdropImageTags": [],
            "ImageBlurHashes": {
                "Primary": {
                    "ec22d3b46cbb21c6b1c9d0f041123c1c": "WF5$ufjEDNbus8kCU^kCkqaKkWf6Dia#%gofRjayb]aejGkBaea#"
                }
            },
            "LocationType": "FileSystem",
            "MediaType": "Unknown",
            "LockedFields": [],
            "LockData": false
        },
        {
            "Name": "Movies plus shows",
            "ServerId": "85e7e9a4fdd743bea5510c6401540719",
            "Id": "ba1e191a58b0a4ab34c09e33b231b6eb",
            "Etag": "cc06e6c65455e4e87b54215d7686da91",
            "DateCreated": "2026-01-25T02:03:48.5200841Z",
            "DateLastMediaAdded": "0001-01-01T00:00:00.0000000Z",
            "CanDelete": false,
            "CanDownload": false,
            "SortName": "movies plus shows",
            "ExternalUrls": [],
            "Path": "/home/mmonk/.local/share/jellyfin/root/default/Movies plus shows",
            "EnableMediaSourceDisplay": true,
            "ChannelId": null,
            "Taglines": [],
            "Genres": [],
            "PlayAccess": "Full",
            "RemoteTrailers": [],
            "ProviderIds": {},
            "IsFolder": true,
            "ParentId": "e9d5075a555c1cbc394eec4cef295274",
            "Type": "CollectionFolder",
            "People": [],
            "Studios": [],
            "GenreItems": [],
            "LocalTrailerCount": 0,
            "UserData": {
                "PlaybackPositionTicks": 0,
                "PlayCount": 0,
                "IsFavorite": false,
                "Played": false,
                "Key": "ba1e191a-58b0-a4ab-34c0-9e33b231b6eb",
                "ItemId": "ba1e191a58b0a4ab34c09e33b231b6eb"
            },
            "ChildCount": 4,
            "SpecialFeatureCount": 0,
            "DisplayPreferencesId": "ba1e191a58b0a4ab34c09e33b231b6eb",
            "Tags": [],
            "PrimaryImageAspectRatio": 1.7777777777777777,
            "ImageTags": {
                "Primary": "27739ba3164e102573071433d9c803bf"
            },
            "BackdropImageTags": [],
            "ImageBlurHashes": {
                "Primary": {
                    "27739ba3164e102573071433d9c803bf": "WC7LM*ozMJaKbvofQnaetloyV@bH00ayofWBofV@W:ayV[ayoff6"
                }
            },
            "LocationType": "FileSystem",
            "MediaType": "Unknown",
            "LockedFields": [],
            "LockData": false
        },
        {
            "Name": "Shows",
            "ServerId": "85e7e9a4fdd743bea5510c6401540719",
            "Id": "a656b907eb3a73532e40e44b968d0225",
            "Etag": "23012a6d9bf348629870c4b73efbe34d",
            "DateCreated": "2026-01-25T02:03:22.981531Z",
            "DateLastMediaAdded": "0001-01-01T00:00:00.0000000Z",
            "CanDelete": false,
            "CanDownload": false,
            "SortName": "shows",
            "ExternalUrls": [],
            "Path": "/home/mmonk/.local/share/jellyfin/root/default/Shows",
            "EnableMediaSourceDisplay": true,
            "ChannelId": null,
            "Taglines": [],
            "Genres": [],
            "PlayAccess": "Full",
            "RemoteTrailers": [],
            "ProviderIds": {},
            "IsFolder": true,
            "ParentId": "e9d5075a555c1cbc394eec4cef295274",
            "Type": "CollectionFolder",
            "People": [],
            "Studios": [],
            "GenreItems": [],
            "LocalTrailerCount": 0,
            "UserData": {
                "PlaybackPositionTicks": 0,
                "PlayCount": 0,
                "IsFavorite": false,
                "Played": false,
                "Key": "a656b907-eb3a-7353-2e40-e44b968d0225",
                "ItemId": "a656b907eb3a73532e40e44b968d0225"
            },
            "ChildCount": 4,
            "SpecialFeatureCount": 0,
            "DisplayPreferencesId": "a656b907eb3a73532e40e44b968d0225",
            "Tags": [],
            "CollectionType": "tvshows",
            "ImageTags": {},
            "BackdropImageTags": [],
            "ImageBlurHashes": {},
            "LocationType": "FileSystem",
            "MediaType": "Unknown",
            "LockedFields": [],
            "LockData": false
        }
    ],
    "TotalRecordCount": 3,
    "StartIndex": 0
}

{
    "Name": "Movies",
    "ServerId": "85e7e9a4fdd743bea5510c6401540719",
    "Id": "f137a2dd21bbc1b99aa5c0f6bf02a805",
    "Etag": "a90d43b3f3a28363e447622e61b93b01",
    "DateCreated": "2026-01-24T12:39:13.6612996Z",
    "DateLastMediaAdded": "0001-01-01T00:00:00.0000000Z",
    "CanDelete": false,
    "CanDownload": false,
    "SortName": "movies",
    "ExternalUrls": [],
    "Path": "/home/mmonk/.local/share/jellyfin/root/default/Movies",
    "EnableMediaSourceDisplay": true,
    "ChannelId": null,
    "Taglines": [],
    "Genres": [],
    "PlayAccess": "Full",
    "RemoteTrailers": [],
    "ProviderIds": {},
    "IsFolder": true,
    "ParentId": "e9d5075a555c1cbc394eec4cef295274",
    "Type": "CollectionFolder",
    "People": [],
    "Studios": [],
    "GenreItems": [],
    "LocalTrailerCount": 0,
    "UserData": {
        "PlaybackPositionTicks": 0,
        "PlayCount": 0,
        "IsFavorite": false,
        "Played": false,
        "Key": "f137a2dd-21bb-c1b9-9aa5-c0f6bf02a805",
        "ItemId": "f137a2dd21bbc1b99aa5c0f6bf02a805"
    },
    "ChildCount": 2,
    "SpecialFeatureCount": 0,
    "DisplayPreferencesId": "f137a2dd21bbc1b99aa5c0f6bf02a805",
    "Tags": [],
    "PrimaryImageAspectRatio": 1.7777777777777777,
    "CollectionType": "movies",
    "ImageTags": {
        "Primary": "ec22d3b46cbb21c6b1c9d0f041123c1c"
    },
    "BackdropImageTags": [],
    "ImageBlurHashes": {
        "Primary": {
            "ec22d3b46cbb21c6b1c9d0f041123c1c": "WF5$ufjEDNbus8kCU^kCkqaKkWf6Dia#%gofRjayb]aejGkBaea#"
        }
    },
    "LocationType": "FileSystem",
    "MediaType": "Unknown",
    "LockedFields": [],
    "LockData": false
},
*/

type ResponseUserViewsItemUserData struct {
	PlaybackPositionTicks int    `json:"PlaybackPositionTicks"`
	PlayCount             int    `json:"PlayCount"`
	IsFavorite            bool   `json:"IsFavorite"`
	Played                bool   `json:"Played"`
	Key                   string `json:"Key"`
	ItemID                string `json:"ItemId"`
}

type ResponseUserViewsItem struct {
	Name                     string                        `json:"Name"`
	ServerID                 string                        `json:"ServerId"`
	ID                       string                        `json:"Id"`
	Etag                     string                        `json:"Etag"`
	DateCreated              time.Time                     `json:"DateCreated"`
	DateLastMediaAdded       time.Time                     `json:"DateLastMediaAdded"`
	CanDelete                bool                          `json:"CanDelete"`
	CanDownload              bool                          `json:"CanDownload"`
	SortName                 string                        `json:"SortName"`
	ExternalUrls             []any                         `json:"ExternalUrls"`
	Path                     string                        `json:"Path"`
	EnableMediaSourceDisplay bool                          `json:"EnableMediaSourceDisplay"`
	ChannelID                any                           `json:"ChannelId"`
	Taglines                 []any                         `json:"Taglines"`
	Genres                   []any                         `json:"Genres"`
	PlayAccess               string                        `json:"PlayAccess"`
	RemoteTrailers           []any                         `json:"RemoteTrailers"`
	ProviderIds              struct{}                      `json:"ProviderIds"`
	IsFolder                 bool                          `json:"IsFolder"`
	ParentID                 string                        `json:"ParentId"`
	Type                     string                        `json:"Type"`
	People                   []any                         `json:"People"`
	Studios                  []any                         `json:"Studios"`
	GenreItems               []any                         `json:"GenreItems"`
	LocalTrailerCount        int                           `json:"LocalTrailerCount"`
	UserData                 ResponseUserViewsItemUserData `json:"UserData"`
	ChildCount               int                           `json:"ChildCount"`
	SpecialFeatureCount      int                           `json:"SpecialFeatureCount"`
	DisplayPreferencesID     string                        `json:"DisplayPreferencesId"`
	Tags                     []any                         `json:"Tags"`
	PrimaryImageAspectRatio  float64                       `json:"PrimaryImageAspectRatio,omitempty"`
	// Currently we make no distinction between various types of media collections, this may be added in future releases if needed
	//	CollectionType           string                        `json:"CollectionType,omitempty"`
	ImageTags struct {
		Primary string `json:"Primary"`
	} `json:"ImageTags,omitempty"`
	BackdropImageTags []any `json:"BackdropImageTags"`
	ImageBlurHashes   struct {
		Primary struct {
			//Ec22D3B46Cbb21C6B1C9D0F041123C1C string `json:"ec22d3b46cbb21c6b1c9d0f041123c1c"`
		} `json:"Primary"`
	} `json:"ImageBlurHashes,omitempty"`
	LocationType string `json:"LocationType"`
	MediaType    string `json:"MediaType"`
	LockedFields []any  `json:"LockedFields"`
	LockData     bool   `json:"LockData"`
}

type ResponseUserViews struct {
	Items            []ResponseUserViewsItem `json:"Items"`
	TotalRecordCount int                     `json:"TotalRecordCount"`
	StartIndex       int                     `json:"StartIndex"`
}
