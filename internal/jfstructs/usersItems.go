package jfstructs

import "time"

/*
Example:
GET /Users/ea2c447b4f6640ff83d0b65b0fe4b92a/Items
Params:
StartIndex=0
Limit=100
Fields=PrimaryImageAspectRatio,SortName,Path,ChildCount,MediaSourceCount,PrimaryImageAspectRatio
ImageTypeLimit=1  //limit to only one blurhash image?
ParentId=ba1e191a58b0a4ab34c09e33b231b6eb //Id of folder
SortBy=IsFolder,SortName //Put folders first, sort by name amongst folders and shows seperately
SortOrder=Ascending //Alphabetical

{
    "Items": [
        {
            "Name": "Folder name",
            "ServerId": "85e7e9a4fdd743bea5510c6401540719",
            "Id": "c0be5a952643a215a615d99962d5a3d8",
            "SortName": "folder name",
            "Path": "/path/on/disk", //Dont use real path
            "ChannelId": null,
            "IsFolder": true,
            "Type": "Folder",
            "UserData": {
                "UnplayedItemCount": 0,
                "PlaybackPositionTicks": 0,
                "PlayCount": 0,
                "IsFavorite": false,
                "Played": true,
                "Key": "c0be5a95-2643-a215-a615-d99962d5a3d8",
                "ItemId": "c0be5a952643a215a615d99962d5a3d8"
            },
            "ChildCount": 1, //Number of files in folder
            "ImageTags": {},
            "BackdropImageTags": [],
            "ImageBlurHashes": {},
            "LocationType": "FileSystem",
            "MediaType": "Unknown"
        },
    ],
    "TotalRecordCount": 1,
    "StartIndex": 0
}
*/

type ResponseUsersItems struct {
	Items            []ResponseUsersItemsItem `json:"Items"`
	TotalRecordCount int                      `json:"TotalRecordCount"`
	StartIndex       int                      `json:"StartIndex"`
}

type ResponseUsersItemsItem struct {
	Name                    string                       `json:"Name"`
	ServerID                string                       `json:"ServerId"`
	ID                      string                       `json:"Id"`
	SortName                string                       `json:"SortName"`
	Path                    string                       `json:"Path"`
	ChannelID               any                          `json:"ChannelId"`
	IsFolder                bool                         `json:"IsFolder"`
	Type                    string                       `json:"Type"`
	UserData                ResponseUsersItemsUserData   `json:"UserData,omitempty"`
	ChildCount              int                          `json:"ChildCount,omitempty"`
	ImageTags               ResponseUsersItemsImageTags  `json:"ImageTags,omitempty"`
	BackdropImageTags       []any                        `json:"BackdropImageTags"`
	ImageBlurHashes         ResponseUsersItemsBlurHashes `json:"ImageBlurHashes,omitempty"`
	LocationType            string                       `json:"LocationType"`
	MediaType               string                       `json:"MediaType"`
	PrimaryImageAspectRatio float64                      `json:"PrimaryImageAspectRatio,omitempty"`
	HasSubtitles            bool                         `json:"HasSubtitles,omitempty"`
	Container               string                       `json:"Container,omitempty"`
	PremiereDate            time.Time                    `json:"PremiereDate,omitempty"`
	CriticRating            int                          `json:"CriticRating,omitempty"`
	OfficialRating          string                       `json:"OfficialRating,omitempty"`
	CommunityRating         float64                      `json:"CommunityRating,omitempty"`
	RunTimeTicks            int64                        `json:"RunTimeTicks,omitempty"`
	ProductionYear          int                          `json:"ProductionYear,omitempty"`
	VideoType               string                       `json:"VideoType,omitempty"`
}

type ResponseUsersItemsUserData struct {
	UnplayedItemCount     int       `json:"UnplayedItemCount,omitempty"`
	PlaybackPositionTicks int64     `json:"PlaybackPositionTicks"`
	PlayCount             int       `json:"PlayCount"`
	IsFavorite            bool      `json:"IsFavorite"`
	Played                bool      `json:"Played"`
	Key                   string    `json:"Key"`
	ItemID                string    `json:"ItemId"`
	LastPlayedDate        time.Time `json:"LastPlayedDate,omitempty"`
	PlayedPercentage      float64   `json:"PlayedPercentage,omitempty"`
}

type ResponseUsersItemsImageTags struct {
	Primary string `json:"Primary,omitempty"`
	Logo    string `json:"Logo,omitempty"`
	Thumb   string `json:"Thumb,omitempty"`
}

type ResponseUsersItemsBlurHashes struct {
	Primary  map[string]string `json:"Primary,omitempty"`
	Backdrop map[string]string `json:"Backdrop,omitempty"`
	Logo     map[string]string `json:"Logo,omitempty"`
	Thumb    map[string]string `json:"Thumb,omitempty"`
}
