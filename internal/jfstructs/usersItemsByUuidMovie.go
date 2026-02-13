package jfstructs

import "time"

/*
"MediaStreams": [
                {
                    "Codec": "hevc",
                    "TimeBase": "1/1000",
                    "VideoRange": "SDR",
                    "VideoRangeType": "SDR",
                    "AudioSpatialFormat": "None",
                    "DisplayTitle": "1080p HEVC SDR",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "BitRate": 2480084,
                    "BitDepth": 10,
                    "RefFrames": 1,
                    "IsDefault": true,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 1040,
                    "Width": 1920,
                    "AverageFrameRate": 23.976025,
                    "RealFrameRate": 23.976025,
                    "ReferenceFrameRate": 23.976025,
                    "Profile": "Main 10",
                    "Type": "Video",
                    "AspectRatio": "1.85:1",
                    "Index": 0,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": false,
                    "PixelFormat": "yuv420p10le",
                    "Level": 120,
                    "IsAnamorphic": false
                },
                {
                    "Codec": "aac",
                    "Language": "eng",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedDefault": "Default",
                    "LocalizedExternal": "External",
                    "DisplayTitle": "English - AAC - Stereo",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "ChannelLayout": "stereo",
                    "BitRate": 184134,
                    "Channels": 2,
                    "SampleRate": 48000,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Profile": "LC",
                    "Type": "Audio",
                    "Index": 1,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": false,
                    "Level": 0
                },
                {
                    "Codec": "aac",
                    "Language": "jpn",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedDefault": "Default",
                    "LocalizedExternal": "External",
                    "DisplayTitle": "Japanese - AAC - Stereo",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "ChannelLayout": "stereo",
                    "BitRate": 174185,
                    "Channels": 2,
                    "SampleRate": 48000,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Profile": "LC",
                    "Type": "Audio",
                    "Index": 2,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": false,
                    "Level": 0
                },
                {
                    "Codec": "PGSSUB",
                    "Language": "eng",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "English - PGSSUB",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 1080,
                    "Width": 1920,
                    "Type": "Subtitle",
                    "Index": 3,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "PGSSUB",
                    "Language": "jpn",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Japanese - PGSSUB",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 1080,
                    "Width": 1920,
                    "Type": "Subtitle",
                    "Index": 4,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "ara",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Arabic - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 5,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "dan",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Danish - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 6,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "nld",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Dutch - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 7,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "fas",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Persian - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 8,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "fra",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "French - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 9,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "deu",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "German - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 10,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "vie",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Vietnamese - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 11,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "hun",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Hungarian - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 12,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "ita",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Italian - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 13,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "pol",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Polish - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 14,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "spa",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Spanish - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 15,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "swe",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Swedish - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 16,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "tha",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Thai - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 17,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "mjpeg",
                    "ColorSpace": "bt470bg",
                    "TimeBase": "1/90000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "BitDepth": 8,
                    "RefFrames": 1,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 314,
                    "Width": 222,
                    "RealFrameRate": 90000,
                    "ReferenceFrameRate": 90000,
                    "Profile": "Baseline",
                    "Type": "EmbeddedImage",
                    "AspectRatio": "111:157",
                    "Index": 18,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": false,
                    "PixelFormat": "yuvj420p",
                    "Level": -99,
                    "IsAnamorphic": false
                }
            ],
*/

type ResponseUsersItemsByUuidMovieMediaStreamsStream struct {
	Codec                    string  `json:"Codec"`
	CodecTag                 string  `json:"CodecTag,omitempty"`
	TimeBase                 string  `json:"TimeBase"`
	VideoRange               string  `json:"VideoRange"`
	VideoRangeType           string  `json:"VideoRangeType"`
	AudioSpatialFormat       string  `json:"AudioSpatialFormat"`
	DisplayTitle             string  `json:"DisplayTitle,omitempty"`
	IsInterlaced             bool    `json:"IsInterlaced"`
	IsAVC                    bool    `json:"IsAVC"`
	BitRate                  int     `json:"BitRate,omitempty"`
	BitDepth                 int     `json:"BitDepth,omitempty"`
	RefFrames                int     `json:"RefFrames,omitempty"`
	IsDefault                bool    `json:"IsDefault"`
	IsForced                 bool    `json:"IsForced"`
	IsHearingImpaired        bool    `json:"IsHearingImpaired"`
	Height                   int     `json:"Height,omitempty"`
	Width                    int     `json:"Width,omitempty"`
	AverageFrameRate         float64 `json:"AverageFrameRate,omitempty"`
	RealFrameRate            float64 `json:"RealFrameRate,omitempty"`
	ReferenceFrameRate       float64 `json:"ReferenceFrameRate,omitempty"`
	Profile                  string  `json:"Profile,omitempty"`
	Type                     string  `json:"Type"`
	AspectRatio              string  `json:"AspectRatio,omitempty"`
	Index                    int     `json:"Index"`
	IsExternal               bool    `json:"IsExternal"`
	IsTextSubtitleStream     bool    `json:"IsTextSubtitleStream"`
	SupportsExternalStream   bool    `json:"SupportsExternalStream"`
	PixelFormat              string  `json:"PixelFormat,omitempty"`
	Level                    int     `json:"Level"`
	IsAnamorphic             bool    `json:"IsAnamorphic,omitempty"`
	Language                 string  `json:"Language,omitempty"`
	LocalizedDefault         string  `json:"LocalizedDefault,omitempty"`
	LocalizedExternal        string  `json:"LocalizedExternal,omitempty"`
	ChannelLayout            string  `json:"ChannelLayout,omitempty"`
	Channels                 int     `json:"Channels,omitempty"`
	SampleRate               int     `json:"SampleRate,omitempty"`
	LocalizedUndefined       string  `json:"LocalizedUndefined,omitempty"`
	LocalizedForced          string  `json:"LocalizedForced,omitempty"`
	LocalizedHearingImpaired string  `json:"LocalizedHearingImpaired,omitempty"`
	ColorSpace               string  `json:"ColorSpace,omitempty"`
}

type ResponseUsersItemsByUuidMovieMediaSources struct {
	Protocol                            string                                            `json:"Protocol"`
	ID                                  string                                            `json:"Id"`
	Path                                string                                            `json:"Path"`
	Type                                string                                            `json:"Type"`
	Container                           string                                            `json:"Container"`
	Size                                int                                               `json:"Size"`
	Name                                string                                            `json:"Name"`
	IsRemote                            bool                                              `json:"IsRemote"`
	ETag                                string                                            `json:"ETag"`
	RunTimeTicks                        int64                                             `json:"RunTimeTicks"`
	ReadAtNativeFramerate               bool                                              `json:"ReadAtNativeFramerate"`
	IgnoreDts                           bool                                              `json:"IgnoreDts"`
	IgnoreIndex                         bool                                              `json:"IgnoreIndex"`
	GenPtsInput                         bool                                              `json:"GenPtsInput"`
	SupportsTranscoding                 bool                                              `json:"SupportsTranscoding"`
	SupportsDirectStream                bool                                              `json:"SupportsDirectStream"`
	SupportsDirectPlay                  bool                                              `json:"SupportsDirectPlay"`
	IsInfiniteStream                    bool                                              `json:"IsInfiniteStream"`
	UseMostCompatibleTranscodingProfile bool                                              `json:"UseMostCompatibleTranscodingProfile"`
	RequiresOpening                     bool                                              `json:"RequiresOpening"`
	RequiresClosing                     bool                                              `json:"RequiresClosing"`
	RequiresLooping                     bool                                              `json:"RequiresLooping"`
	SupportsProbing                     bool                                              `json:"SupportsProbing"`
	VideoType                           string                                            `json:"VideoType"`
	MediaStreams                        []ResponseUsersItemsByUuidMovieMediaStreamsStream `json:"MediaStreams"`
	MediaAttachments                    []any                                             `json:"MediaAttachments"`
	Formats                             []any                                             `json:"Formats"`
	Bitrate                             int                                               `json:"Bitrate"`
	RequiredHTTPHeaders                 struct {
	} `json:"RequiredHttpHeaders"`
	TranscodingSubProtocol  string `json:"TranscodingSubProtocol"`
	DefaultAudioStreamIndex int    `json:"DefaultAudioStreamIndex"`
	HasSegments             bool   `json:"HasSegments"`
}

type ResponseUsersItemsByUuidMovie struct {
	Name          string                                            `json:"Name"`
	OriginalTitle string                                            `json:"OriginalTitle"`
	ServerID      string                                            `json:"ServerId"`
	ID            string                                            `json:"Id"`
	Container     string                                            `json:"Container"`
	PremiereDate  time.Time                                         `json:"PremiereDate"`
	Overview      string                                            `json:"Overview"`
	Path          string                                            `json:"Path"`
	RunTimeTicks  int64                                             `json:"RunTimeTicks"`
	PlayAccess    string                                            `json:"PlayAccess"`
	MediaType     string                                            `json:"MediaType"`
	LocationType  string                                            `json:"LocationType"`
	MediaStreams  []ResponseUsersItemsByUuidMovieMediaStreamsStream `json:"MediaStreams"`
	MediaSources  []ResponseUsersItemsByUuidMovieMediaSources       `json:"MediaSources"`
	VideoType     string                                            `json:"VideoType"`
	CanDownload   bool                                              `json:"CanDownload"`
	//Etag          string    `json:"Etag"`
	//DateCreated  time.Time `json:"DateCreated"`
	//CanDelete    bool      `json:"CanDelete"`

	//HasSubtitles bool      `json:"HasSubtitles"`

	//SortName     string    `json:"SortName"`

	//
	//LockedFields []any     `json:"LockedFields"`
	//LockData     bool      `json:"LockData"`
	//Width        int       `json:"Width"`
	//Height       int       `json:"Height"`

	/*
	    ExternalUrls  []struct {
	   		Name string `json:"Name"`
	   		URL  string `json:"Url"`
	   	} `json:"ExternalUrls"`
	*/
	/*

			MediaAttachments []struct {
				Codec    string `json:"Codec"`
				CodecTag string `json:"CodecTag"`
				Index    int    `json:"Index"`
				FileName string `json:"FileName"`
				MimeType string `json:"MimeType"`
			} `json:"MediaAttachments"`
			Formats             []any `json:"Formats"`
			Bitrate             int   `json:"Bitrate"`
			RequiredHTTPHeaders struct {
			} `json:"RequiredHttpHeaders"`
			TranscodingSubProtocol     string `json:"TranscodingSubProtocol"`
			DefaultAudioStreamIndex    int    `json:"DefaultAudioStreamIndex"`
			DefaultSubtitleStreamIndex int    `json:"DefaultSubtitleStreamIndex"`
			HasSegments                bool   `json:"HasSegments"`
		} `json:"MediaSources"`
	*/
	/*
		CriticRating             int      `json:"CriticRating"`
		ProductionLocations      []string `json:"ProductionLocations"`

		EnableMediaSourceDisplay bool     `json:"EnableMediaSourceDisplay"`
		OfficialRating           string   `json:"OfficialRating"`
		ChannelID                any      `json:"ChannelId"`

		Taglines                 []string `json:"Taglines"`
		Genres                   []string `json:"Genres"`
		CommunityRating          float64  `json:"CommunityRating"`

		ProductionYear           int      `json:"ProductionYear"`
		RemoteTrailers           []struct {
			URL  string `json:"Url"`
			Name string `json:"Name"`
		} `json:"RemoteTrailers"`
		ProviderIds struct {
			Imdb           string `json:"Imdb"`
			Tmdb           string `json:"Tmdb"`
			TmdbCollection string `json:"TmdbCollection"`
		} `json:"ProviderIds"`
		IsHD     bool   `json:"IsHD"`
		IsFolder bool   `json:"IsFolder"`
		ParentID string `json:"ParentId"`
		Type     string `json:"Type"`
	*/
	/*
		People   []struct {
			Name            string `json:"Name"`
			ID              string `json:"Id"`
			Role            string `json:"Role"`
			Type            string `json:"Type"`
			PrimaryImageTag string `json:"PrimaryImageTag,omitempty"`
			ImageBlurHashes struct {
				Primary struct {
					Ef76Eb6226B3921B504Ed8E8Bb1Bbab5 string `json:"ef76eb6226b3921b504ed8e8bb1bbab5"`
				} `json:"Primary"`
			} `json:"ImageBlurHashes"`
		} `json:"People"`
	*/
	/*
		Studios []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"Studios"`
		GenreItems []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"GenreItems"`
		LocalTrailerCount int `json:"LocalTrailerCount"`
	*/
	/*
		UserData struct {
			PlayedPercentage      float64   `json:"PlayedPercentage"`
			PlaybackPositionTicks int64     `json:"PlaybackPositionTicks"`
			PlayCount             int       `json:"PlayCount"`
			IsFavorite            bool      `json:"IsFavorite"`
			LastPlayedDate        time.Time `json:"LastPlayedDate"`
			Played                bool      `json:"Played"`
			Key                   string    `json:"Key"`
			ItemID                string    `json:"ItemId"`
		} `json:"UserData"`
		SpecialFeatureCount     int      `json:"SpecialFeatureCount"`
		DisplayPreferencesID    string   `json:"DisplayPreferencesId"`
		Tags                    []string `json:"Tags"`
		PrimaryImageAspectRatio float64  `json:"PrimaryImageAspectRatio"`
		MediaStreams            []struct {
			Codec                    string  `json:"Codec"`
			TimeBase                 string  `json:"TimeBase"`
			VideoRange               string  `json:"VideoRange"`
			VideoRangeType           string  `json:"VideoRangeType"`
			AudioSpatialFormat       string  `json:"AudioSpatialFormat"`
			DisplayTitle             string  `json:"DisplayTitle,omitempty"`
			IsInterlaced             bool    `json:"IsInterlaced"`
			IsAVC                    bool    `json:"IsAVC"`
			BitRate                  int     `json:"BitRate,omitempty"`
			BitDepth                 int     `json:"BitDepth,omitempty"`
			RefFrames                int     `json:"RefFrames,omitempty"`
			IsDefault                bool    `json:"IsDefault"`
			IsForced                 bool    `json:"IsForced"`
			IsHearingImpaired        bool    `json:"IsHearingImpaired"`
			Height                   int     `json:"Height,omitempty"`
			Width                    int     `json:"Width,omitempty"`
			AverageFrameRate         float64 `json:"AverageFrameRate,omitempty"`
			RealFrameRate            float64 `json:"RealFrameRate,omitempty"`
			ReferenceFrameRate       float64 `json:"ReferenceFrameRate,omitempty"`
			Profile                  string  `json:"Profile,omitempty"`
			Type                     string  `json:"Type"`
			AspectRatio              string  `json:"AspectRatio,omitempty"`
			Index                    int     `json:"Index"`
			IsExternal               bool    `json:"IsExternal"`
			IsTextSubtitleStream     bool    `json:"IsTextSubtitleStream"`
			SupportsExternalStream   bool    `json:"SupportsExternalStream"`
			PixelFormat              string  `json:"PixelFormat,omitempty"`
			Level                    int     `json:"Level"`
			IsAnamorphic             bool    `json:"IsAnamorphic,omitempty"`
			Language                 string  `json:"Language,omitempty"`
			LocalizedDefault         string  `json:"LocalizedDefault,omitempty"`
			LocalizedExternal        string  `json:"LocalizedExternal,omitempty"`
			ChannelLayout            string  `json:"ChannelLayout,omitempty"`
			Channels                 int     `json:"Channels,omitempty"`
			SampleRate               int     `json:"SampleRate,omitempty"`
			LocalizedUndefined       string  `json:"LocalizedUndefined,omitempty"`
			LocalizedForced          string  `json:"LocalizedForced,omitempty"`
			LocalizedHearingImpaired string  `json:"LocalizedHearingImpaired,omitempty"`
			ColorSpace               string  `json:"ColorSpace,omitempty"`
		} `json:"MediaStreams"`

		ImageTags struct {
			Logo    string `json:"Logo"`
			Thumb   string `json:"Thumb"`
			Primary string `json:"Primary"`
		} `json:"ImageTags"`
		BackdropImageTags []string `json:"BackdropImageTags"`
		ImageBlurHashes   struct {
			Backdrop struct {
				SixD10Dfc7Cea2613F76324108A2518334 string `json:"6d10dfc7cea2613f76324108a2518334"`
			} `json:"Backdrop"`
			Logo struct {
				Ecf86C54317Ef8676E98Ca49E688Df29 string `json:"ecf86c54317ef8676e98ca49e688df29"`
			} `json:"Logo"`
			Thumb struct {
				D12457F1F322Ff757Bdfcfae0E839Bdf string `json:"d12457f1f322ff757bdfcfae0e839bdf"`
			} `json:"Thumb"`
			Primary struct {
				NineC3152Facacd7C88A6Db5E4C5B041051 string `json:"9c3152facacd7c88a6db5e4c5b041051"`
			} `json:"Primary"`
		} `json:"ImageBlurHashes"`
		Chapters []struct {
			StartPositionTicks int       `json:"StartPositionTicks"`
			Name               string    `json:"Name"`
			ImageDateModified  time.Time `json:"ImageDateModified"`
		} `json:"Chapters"`
		Trickplay struct {
		} `json:"Trickplay"`
	*/
}

/*
type AutoGenerated struct {
	Name          string    `json:"Name"`
	OriginalTitle string    `json:"OriginalTitle"`
	ServerID      string    `json:"ServerId"`
	ID            string    `json:"Id"`
	Etag          string    `json:"Etag"`
	DateCreated   time.Time `json:"DateCreated"`
	CanDelete     bool      `json:"CanDelete"`
	CanDownload   bool      `json:"CanDownload"`
	Container     string    `json:"Container"`
	SortName      string    `json:"SortName"`
	PremiereDate  time.Time `json:"PremiereDate"`
	ExternalUrls  []struct {
		Name string `json:"Name"`
		URL  string `json:"Url"`
	} `json:"ExternalUrls"`
	MediaSources []struct {
		Protocol                            string `json:"Protocol"`
		ID                                  string `json:"Id"`
		Path                                string `json:"Path"`
		Type                                string `json:"Type"`
		Container                           string `json:"Container"`
		Size                                int    `json:"Size"`
		Name                                string `json:"Name"`
		IsRemote                            bool   `json:"IsRemote"`
		ETag                                string `json:"ETag"`
		RunTimeTicks                        int64  `json:"RunTimeTicks"`
		ReadAtNativeFramerate               bool   `json:"ReadAtNativeFramerate"`
		IgnoreDts                           bool   `json:"IgnoreDts"`
		IgnoreIndex                         bool   `json:"IgnoreIndex"`
		GenPtsInput                         bool   `json:"GenPtsInput"`
		SupportsTranscoding                 bool   `json:"SupportsTranscoding"`
		SupportsDirectStream                bool   `json:"SupportsDirectStream"`
		SupportsDirectPlay                  bool   `json:"SupportsDirectPlay"`
		IsInfiniteStream                    bool   `json:"IsInfiniteStream"`
		UseMostCompatibleTranscodingProfile bool   `json:"UseMostCompatibleTranscodingProfile"`
		RequiresOpening                     bool   `json:"RequiresOpening"`
		RequiresClosing                     bool   `json:"RequiresClosing"`
		RequiresLooping                     bool   `json:"RequiresLooping"`
		SupportsProbing                     bool   `json:"SupportsProbing"`
		VideoType                           string `json:"VideoType"`
		MediaStreams                        []struct {
			Codec                  string `json:"Codec"`
			CodecTag               string `json:"CodecTag"`
			Language               string `json:"Language"`
			TimeBase               string `json:"TimeBase"`
			VideoRange             string `json:"VideoRange"`
			VideoRangeType         string `json:"VideoRangeType"`
			AudioSpatialFormat     string `json:"AudioSpatialFormat"`
			DisplayTitle           string `json:"DisplayTitle"`
			NalLengthSize          string `json:"NalLengthSize,omitempty"`
			IsInterlaced           bool   `json:"IsInterlaced"`
			IsAVC                  bool   `json:"IsAVC"`
			BitRate                int    `json:"BitRate"`
			BitDepth               int    `json:"BitDepth,omitempty"`
			RefFrames              int    `json:"RefFrames,omitempty"`
			IsDefault              bool   `json:"IsDefault"`
			IsForced               bool   `json:"IsForced"`
			IsHearingImpaired      bool   `json:"IsHearingImpaired"`
			Height                 int    `json:"Height,omitempty"`
			Width                  int    `json:"Width,omitempty"`
			AverageFrameRate       int    `json:"AverageFrameRate,omitempty"`
			RealFrameRate          int    `json:"RealFrameRate,omitempty"`
			ReferenceFrameRate     int    `json:"ReferenceFrameRate,omitempty"`
			Profile                string `json:"Profile"`
			Type                   string `json:"Type"`
			AspectRatio            string `json:"AspectRatio,omitempty"`
			Index                  int    `json:"Index"`
			IsExternal             bool   `json:"IsExternal"`
			IsTextSubtitleStream   bool   `json:"IsTextSubtitleStream"`
			SupportsExternalStream bool   `json:"SupportsExternalStream"`
			PixelFormat            string `json:"PixelFormat,omitempty"`
			Level                  int    `json:"Level"`
			IsAnamorphic           bool   `json:"IsAnamorphic,omitempty"`
			Title                  string `json:"Title,omitempty"`
			LocalizedDefault       string `json:"LocalizedDefault,omitempty"`
			LocalizedExternal      string `json:"LocalizedExternal,omitempty"`
			ChannelLayout          string `json:"ChannelLayout,omitempty"`
			Channels               int    `json:"Channels,omitempty"`
			SampleRate             int    `json:"SampleRate,omitempty"`
		} `json:"MediaStreams"`
		MediaAttachments    []any `json:"MediaAttachments"`
		Formats             []any `json:"Formats"`
		Bitrate             int   `json:"Bitrate"`
		RequiredHTTPHeaders struct {
		} `json:"RequiredHttpHeaders"`
		TranscodingSubProtocol  string `json:"TranscodingSubProtocol"`
		DefaultAudioStreamIndex int    `json:"DefaultAudioStreamIndex"`
		HasSegments             bool   `json:"HasSegments"`
	} `json:"MediaSources"`
	CriticRating             int      `json:"CriticRating"`
	ProductionLocations      []string `json:"ProductionLocations"`
	Path                     string   `json:"Path"`
	EnableMediaSourceDisplay bool     `json:"EnableMediaSourceDisplay"`
	OfficialRating           string   `json:"OfficialRating"`
	ChannelID                any      `json:"ChannelId"`
	Overview                 string   `json:"Overview"`
	Taglines                 []string `json:"Taglines"`
	Genres                   []string `json:"Genres"`
	CommunityRating          float64  `json:"CommunityRating"`
	RunTimeTicks             int64    `json:"RunTimeTicks"`
	PlayAccess               string   `json:"PlayAccess"`
	ProductionYear           int      `json:"ProductionYear"`
	RemoteTrailers           []struct {
		URL  string `json:"Url"`
		Name string `json:"Name"`
	} `json:"RemoteTrailers"`
	ProviderIds struct {
		Imdb string `json:"Imdb"`
		Tmdb string `json:"Tmdb"`
	} `json:"ProviderIds"`
	IsHD     bool   `json:"IsHD"`
	IsFolder bool   `json:"IsFolder"`
	ParentID string `json:"ParentId"`
	Type     string `json:"Type"`
	People   []struct {
		Name            string `json:"Name"`
		ID              string `json:"Id"`
		Role            string `json:"Role"`
		Type            string `json:"Type"`
		PrimaryImageTag string `json:"PrimaryImageTag,omitempty"`
		ImageBlurHashes struct {
			Primary struct {
				ThreeE2Fa050E881C894D264B770E1579E52 string `json:"3e2fa050e881c894d264b770e1579e52"`
			} `json:"Primary"`
		} `json:"ImageBlurHashes"`
	} `json:"People"`
	Studios []struct {
		Name string `json:"Name"`
		ID   string `json:"Id"`
	} `json:"Studios"`
	GenreItems []struct {
		Name string `json:"Name"`
		ID   string `json:"Id"`
	} `json:"GenreItems"`
	LocalTrailerCount int `json:"LocalTrailerCount"`
	UserData          struct {
		PlaybackPositionTicks int    `json:"PlaybackPositionTicks"`
		PlayCount             int    `json:"PlayCount"`
		IsFavorite            bool   `json:"IsFavorite"`
		Played                bool   `json:"Played"`
		Key                   string `json:"Key"`
		ItemID                string `json:"ItemId"`
	} `json:"UserData"`
	SpecialFeatureCount     int      `json:"SpecialFeatureCount"`
	DisplayPreferencesID    string   `json:"DisplayPreferencesId"`
	Tags                    []string `json:"Tags"`
	PrimaryImageAspectRatio float64  `json:"PrimaryImageAspectRatio"`
	MediaStreams            []struct {
		Codec                  string `json:"Codec"`
		CodecTag               string `json:"CodecTag"`
		Language               string `json:"Language"`
		TimeBase               string `json:"TimeBase"`
		VideoRange             string `json:"VideoRange"`
		VideoRangeType         string `json:"VideoRangeType"`
		AudioSpatialFormat     string `json:"AudioSpatialFormat"`
		DisplayTitle           string `json:"DisplayTitle"`
		NalLengthSize          string `json:"NalLengthSize,omitempty"`
		IsInterlaced           bool   `json:"IsInterlaced"`
		IsAVC                  bool   `json:"IsAVC"`
		BitRate                int    `json:"BitRate"`
		BitDepth               int    `json:"BitDepth,omitempty"`
		RefFrames              int    `json:"RefFrames,omitempty"`
		IsDefault              bool   `json:"IsDefault"`
		IsForced               bool   `json:"IsForced"`
		IsHearingImpaired      bool   `json:"IsHearingImpaired"`
		Height                 int    `json:"Height,omitempty"`
		Width                  int    `json:"Width,omitempty"`
		AverageFrameRate       int    `json:"AverageFrameRate,omitempty"`
		RealFrameRate          int    `json:"RealFrameRate,omitempty"`
		ReferenceFrameRate     int    `json:"ReferenceFrameRate,omitempty"`
		Profile                string `json:"Profile"`
		Type                   string `json:"Type"`
		AspectRatio            string `json:"AspectRatio,omitempty"`
		Index                  int    `json:"Index"`
		IsExternal             bool   `json:"IsExternal"`
		IsTextSubtitleStream   bool   `json:"IsTextSubtitleStream"`
		SupportsExternalStream bool   `json:"SupportsExternalStream"`
		PixelFormat            string `json:"PixelFormat,omitempty"`
		Level                  int    `json:"Level"`
		IsAnamorphic           bool   `json:"IsAnamorphic,omitempty"`
		Title                  string `json:"Title,omitempty"`
		LocalizedDefault       string `json:"LocalizedDefault,omitempty"`
		LocalizedExternal      string `json:"LocalizedExternal,omitempty"`
		ChannelLayout          string `json:"ChannelLayout,omitempty"`
		Channels               int    `json:"Channels,omitempty"`
		SampleRate             int    `json:"SampleRate,omitempty"`
	} `json:"MediaStreams"`
	VideoType string `json:"VideoType"`
	ImageTags struct {
		Primary string `json:"Primary"`
		Logo    string `json:"Logo"`
		Thumb   string `json:"Thumb"`
	} `json:"ImageTags"`
	BackdropImageTags []string `json:"BackdropImageTags"`
	ImageBlurHashes   struct {
		Backdrop struct {
			ThreeA7Fde3D3509249054Fd591B77468A69 string `json:"3a7fde3d3509249054fd591b77468a69"`
		} `json:"Backdrop"`
		Primary struct {
			Cf44375D10118801C4540811B46502B3 string `json:"cf44375d10118801c4540811b46502b3"`
		} `json:"Primary"`
		Logo struct {
			Nine9C0B052454Be28857C73Bcf22D12065 string `json:"99c0b052454be28857c73bcf22d12065"`
		} `json:"Logo"`
		Thumb struct {
			D203D5B15Ec7A905C9Cc0E415E34Ca81 string `json:"d203d5b15ec7a905c9cc0e415e34ca81"`
		} `json:"Thumb"`
	} `json:"ImageBlurHashes"`
	Chapters  []any `json:"Chapters"`
	Trickplay struct {
	} `json:"Trickplay"`
	LocationType string `json:"LocationType"`
	MediaType    string `json:"MediaType"`
	LockedFields []any  `json:"LockedFields"`
	LockData     bool   `json:"LockData"`
	Width        int    `json:"Width"`
	Height       int    `json:"Height"`
}
*/
