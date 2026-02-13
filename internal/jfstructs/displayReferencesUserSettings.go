package jfstructs

/*
{
    "Id": "3ce5b65d-e116-d731-65d1-efc4a30ec35c",
    "SortBy": "SortName",
    "RememberIndexing": false,
    "PrimaryImageHeight": 250,
    "PrimaryImageWidth": 250,
    "CustomPrefs": {
        "chromecastVersion": "stable",
        "skipForwardLength": "30000",
        "skipBackLength": "10000",
        "enableNextVideoInfoOverlay": "False",
        "tvhome": null,
        "dashboardTheme": null
    },
    "ScrollDirection": "Horizontal",
    "ShowBackdrop": true,
    "RememberSorting": false,
    "SortOrder": "Ascending",
    "ShowSidebar": false,
    "Client": "emby"
}
*/

type ResponseDisplayPreferencesUsersettings struct {
	ID                 string                                            `json:"Id"` //non-stripped uuid allocated to this user's display preferences
	SortBy             string                                            `json:"SortBy"`
	RememberIndexing   bool                                              `json:"RememberIndexing"`
	PrimaryImageHeight int                                               `json:"PrimaryImageHeight"`
	PrimaryImageWidth  int                                               `json:"PrimaryImageWidth"`
	CustomPrefs        ResponseDisplayPreferencesUsersettingsCustomPrefs `json:"CustomPrefs"`
	ScrollDirection    string                                            `json:"ScrollDirection"`
	ShowBackdrop       bool                                              `json:"ShowBackdrop"`
	RememberSorting    bool                                              `json:"RememberSorting"`
	SortOrder          string                                            `json:"SortOrder"`
	ShowSidebar        bool                                              `json:"ShowSidebar"`
	Client             string                                            `json:"Client"` //http query string contains a "client" property, this field should be set to the value of "client"
}

type ResponseDisplayPreferencesUsersettingsCustomPrefs struct {
	ChromecastVersion          string `json:"chromecastVersion"`
	SkipForwardLength          string `json:"skipForwardLength"`
	SkipBackLength             string `json:"skipBackLength"`
	EnableNextVideoInfoOverlay string `json:"enableNextVideoInfoOverlay"`
	Tvhome                     any    `json:"tvhome"`
	DashboardTheme             any    `json:"dashboardTheme"`
}
