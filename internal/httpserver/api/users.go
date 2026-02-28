package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"opal/internal/jfstructs"
	"opal/internal/librarymgmt"
	"opal/internal/mediamgmt"
	"opal/internal/usermgmt"
	"path"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

func EndpointUsersPublic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	res := jfstructs.ResponseUsersPublic{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func EndpointUsersById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	aTokenString := getHeaderAuthToken(r)
	pathUserId := r.PathValue("id")

	isReqValid := false
	var aT usermgmt.AccessToken

	//TODO: token lock halts any other goroutines from reading currentTokens, should change in future
	usermgmt.CurrentTokensMutex.Lock()

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == aTokenString && pathUserId == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
			aT = t
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
	}

	res := jfstructs.CommonUser{
		Name:        aT.UserName,
		ServerID:    "STUBBED",
		ID:          aT.UserId,
		HasPassword: true,
		Configuration: jfstructs.CommonUsersConfiguration{
			SubtitleMode:               "Default",
			HidePlayedInLatest:         true,
			RememberAudioSelections:    true,
			RememberSubtitleSelections: true,
			EnableNextEpisodeAutoPlay:  true,
		},
		Policy: jfstructs.CommonUsersPolicy{
			IsAdministrator:                false,
			IsHidden:                       true,
			EnableLiveTvManagement:         false, //NOTE: live tv will probably never be supported, seems like alot of work for nothing
			EnableLiveTvAccess:             false,
			EnableMediaPlayback:            true,
			EnableAudioPlaybackTranscoding: true,
			EnableVideoPlaybackTranscoding: true,
			EnablePlaybackRemuxing:         true,
			EnableContentDownloading:       true,
			EnableMediaConversion:          true,
			EnableAllDevices:               true,
			EnableAllChannels:              true,
			EnableAllFolders:               true,
			LoginAttemptsBeforeLockout:     -1,
			MaxActiveSessions:              0,
			AuthenticationProviderID:       "Jellyfin.Server.Implementations.Users.DefaultAuthenticationProvider",
			PasswordResetProviderID:        "Jellyfin.Server.Implementations.Users.DefaultPasswordResetProvider",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func EndpointUsersItemsResume(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	aTokenString := getHeaderAuthToken(r)
	pathUserId := r.PathValue("id")

	isReqValid := false

	//TODO: token lock halts any other goroutines from reading currentTokens, should change in future
	usermgmt.CurrentTokensMutex.Lock()

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == aTokenString && pathUserId == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
	}

	//TODO: STUBBED
	res := []jfstructs.CommonItem{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func EndpointUsersItemsLatest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	aTokenString := getHeaderAuthToken(r)
	pathUserId := r.PathValue("id")

	isReqValid := false

	//TODO: token lock halts any other goroutines from reading currentTokens, should change in future
	usermgmt.CurrentTokensMutex.Lock()

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == aTokenString && pathUserId == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
	}

	query := r.URL.Query()

	rootUuid := query.Get("ParentId")
	nM := librarymgmt.NewestMedia[rootUuid]

	if nM == nil {
		http.Error(w, "Bad UUID", http.StatusBadRequest)
	}

	//http://localhost:7096/Users/ce1f6aa6-4e44-5414-8698-edaa9af48fb1/Items/Latest
	//Limit=16
	//Fields=PrimaryImageAspectRatio%2CPath
	//ImageTypeLimit=1
	//EnableImageTypes=Primary%2CBackdrop%2CThumb
	//ParentId=749718d8-2abc-512d-9ec7-c2e469a7dc11

	limit := 0
	limitStr := query.Get("Limit")
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Failed to parse Limit", http.StatusInternalServerError)
			return
		}
	}

	if limit > len(nM) || limit == 0 || limit < 0 {
		limit = len(nM)
	}

	limitedNewestMedia := nM[:limit]

	res := []jfstructs.CommonItem{}

	for _, media := range limitedNewestMedia {
		mediaItem := jfstructs.CommonItem{}

		switch media.Node.Type {
		case "movie":
			mediaItem = jfstructs.CommonItem{
				Name:     media.Node.MovieMetadata.Title,
				ServerID: "STUBBED",
				ID:       media.Node.Uuid,
				Path:     fmt.Sprintf("/Libraries/%s%s", media.Node.RootUuid, media.Node.Path),
				//PlayAccess:   "Full",
				MediaType:    "Video",
				LocationType: "FileSystem",
				VideoType:    "VideoFile",
			}
			mediaItem.ImageTags.Primary = "backdrop.png"

		case "tvshow":
			mediaItem = jfstructs.CommonItem{
				Name:         media.Node.TvshowMetadata.Name,
				ServerID:     "STUBBED",
				ID:           media.Node.Uuid,
				Path:         fmt.Sprintf("/Libraries/%s%s", media.Node.RootUuid, media.Node.Path),
				MediaType:    "Video",
				LocationType: "FileSystem",
				VideoType:    "VideoFile",
			}
			mediaItem.ImageTags.Primary = "backdrop.png"
		}

		res = append(res, mediaItem)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func EndpointUsersItemsByUuid(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	pathUserId := r.PathValue("id")
	aTokenString := getHeaderAuthToken(r)

	isReqValid := false

	//TODO: token lock halts any other goroutines from reading currentTokens, should change in future
	usermgmt.CurrentTokensMutex.Lock()

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == aTokenString && pathUserId == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
		return
	}

	itemUuid := r.PathValue("itemUuid")

	if librarymgmt.LibTreeMap[itemUuid] == nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	switch librarymgmt.LibTreeMap[itemUuid].Type {
	case "folder":
		EndpointUsersItemsByUuidFolder(w, r, itemUuid)
	case "tvshow":
		EndpointUsersItemsByUuidTvshow(w, r, itemUuid)
	case "movie":
		EndpointUsersItemsByUuidMovie(w, r, itemUuid)
	}

}

func EndpointUsersItemsByUuidFolder(w http.ResponseWriter, r *http.Request, itemUuid string) {
	itemNode := librarymgmt.LibTreeMap[itemUuid]

	res := jfstructs.ResponseUsersItemsByUuidFolder{
		Name:        itemNode.Name,
		ServerID:    "STUBBED",
		Id:          itemNode.Uuid,
		SortName:    strings.ToLower(itemNode.Name),
		Path:        strings.TrimSuffix(fmt.Sprintf("/Libraries/%s%s", itemNode.RootUuid, itemNode.Path), "/"),
		PlayAccess:  "Full",
		IsFolder:    true,
		CanDownload: true,
		Type:        "CollectionFolder",
		UserData: jfstructs.CommonItemUserData{
			//Note: key is supposed to be unstripped and itemid is supposed to be stripped in real responses
			Key:    itemNode.Uuid,
			ItemID: itemNode.Uuid,
		},
		ChildCount:              len(itemNode.Children),
		DisplayPreferencesID:    itemNode.Uuid, //???
		PrimaryImageAspectRatio: 1.7777777777777777,
		LocationType:            "FileSystem",
		MediaType:               "Unknown",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func EndpointUsersItemsByUuidTvshow(w http.ResponseWriter, r *http.Request, itemUuid string) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

/*
All:
"Codec": "hevc", //Codec identifier (e.g., "hevc", "aac", "PGSSUB", "ass")
"TimeBase": "1/1000",
"VideoRange": "SDR", //Set to "Unknown" on non-video/non-image streams
"VideoRangeType": "SDR", //Set to "Unknown" on non-video/non-image streams
"AudioSpatialFormat": "None", //Always set to none, TODO: check if surround sound streams set this
"DisplayTitle": "1080p HEVC SDR", //Friendly name for UI
"IsInterlaced": false,
"IsAVC": false,
"IsDefault": true, //Set if it's the primary track for that type
"IsForced": false,

//TODO: this generic field list is contaminated with some video only tags, filter out
Generic and Video:
"BitRate": 2480084, //Total bitrate (0 or null for most subtitles)
"BitDepth": 10, //Color depth (e.g., 8, 10)
"RefFrames": 1, //Number of reference frames for video/images
"IsHearingImpaired": false,
"Height": 1040, //Resolution height (0 for text-based subtitles)
"Width": 1920, //Resolution width (0 for text-based subtitles)
"AverageFrameRate": 23.976025,
"RealFrameRate": 23.976025,
"ReferenceFrameRate": 23.976025,
"Profile": "Main 10", //Codec profile (e.g., "Main 10", "LC", "Baseline")
"Type": "Video", //"Video", "Audio", "Subtitle", or "EmbeddedImage"
"AspectRatio": "1.85:1", //Width/Height ratio string
"Index": 0, //Stream index in file
"IsExternal": false, //True if file is external (like an .srt)
"IsTextSubtitleStream": false, //True for "ass", "subrip"; False for "PGSSUB" or video
"SupportsExternalStream": false,
"PixelFormat": "yuv420p10le", //Chroma format (e.g., yuv420p10le, yuvj420p)
"Level": 120, //Encoder level (e.g., 120 = Level 4.0)
"IsAnamorphic": false,

Audio and Subtitles:
"Language": "eng", //ISO 639-2 3-letter code
"LocalizedDefault": "Default",
"LocalizedExternal": "External",

Audio only:
"ChannelLayout": "stereo", //Speaker config (e.g., "stereo", "5.1")
"Channels": 2,
"SampleRate": 48000,

Subtitles only:
"LocalizedUndefined": "Undefined",
"LocalizedForced": "Forced",
"LocalizedHearingImpaired": "Hearing Impaired", //Note: Can be set even if "IsHearingImpaired" is false

EmbeddedImage only:
"ColorSpace": "bt470bg" //Color standard for thumbnails/covers
*/

func EndpointUsersItemsByUuidMovie(w http.ResponseWriter, r *http.Request, itemUuid string) {
	item := librarymgmt.LibTreeMap[itemUuid]
	libraryRecord := librarymgmt.AllLibrariesMap[item.RootUuid]

	if libraryRecord == nil {
		http.Error(w, "Library not found", http.StatusNotFound)
		return
	}

	movieFilePath := path.Join(libraryRecord.Path, item.Path)
	movieProbe, err := mediamgmt.Probe(movieFilePath)
	if err != nil {
		http.Error(w, "Media probe failed", http.StatusInternalServerError)
	}

	durationF64, err := strconv.ParseFloat(movieProbe.Format.Duration, 64)
	if err != nil {
		http.Error(w, "Failed to probe media duration", http.StatusInternalServerError)
	}
	durationTicks := durationF64 * 10_000_000 //The official jellyfin server is written in C#/.Net framework, a tick in .Net is a ten millionth of a second

	res := jfstructs.ResponseUsersItemsByUuidMovie{
		Name:            item.MovieMetadata.Title,
		OriginalTitle:   item.MovieMetadata.OriginalTitle,
		ServerID:        "STUBBED",
		ID:              item.Uuid,
		Overview:        item.MovieMetadata.Overview,
		Path:            fmt.Sprintf("/Libraries/%s%s", item.RootUuid, item.Path),
		Taglines:        []string{item.MovieMetadata.Tagline},
		RunTimeTicks:    int64(durationTicks),
		PlayAccess:      "Full",
		MediaType:       "Video",
		LocationType:    "FileSystem",
		VideoType:       "VideoFile",
		CanDownload:     true,
		CanDelete:       false,
		PremiereDate:    item.ReleasedTime,
		OfficialRating:  item.Rating,
		CommunityRating: item.MovieMetadata.VoteAverage,
	}
	res.ImageTags.Primary = "poster.png"
	res.ImageTags.Logo = "logo.png"
	res.ImageTags.Thumb = "backdrop.png"

	res.BackdropImageTags = []string{"backdrop.png"}

	res.PrimaryImageAspectRatio = item.MovieMetadata.PosterPrimaryAspectRatio

	imdbUrl := jfstructs.ResponseUsersItemsByUuidMovieExternalUrls{
		Name: "IMDB",
		URL:  fmt.Sprintf("https://www.imdb.com/title/%s", item.ImdbId),
	}
	res.ExternalUrls = append(res.ExternalUrls, imdbUrl)

	tmdbUrl := jfstructs.ResponseUsersItemsByUuidMovieExternalUrls{
		Name: "TMDB",
		URL:  fmt.Sprintf("https://www.themoviedb.org/movie/%d", item.MovieMetadata.ID),
	}
	res.ExternalUrls = append(res.ExternalUrls, tmdbUrl)

	//TODO: this maybe incorrect behaviour, jellyfin potentially just returns the file extension
	//NOTE: jellyfin's official demo server often returns "mov,mp4,m4a,3gp,3g2,mj2" but most selfhosted servers return "mp4"
	switch movieProbe.Format.FormatName {
	case "matroska,webm":
		res.Container = "mkv"
	case "mov,mp4,m4a,3gp,3g2,mj2":
		res.Container = "mp4"
	default:
		slog.Warn("unrecognised container format", "offendingMedia", item.MovieMetadata.Title, "containerFormat", movieProbe.Format.FormatName)
		res.Container = "Unknown" //Might be incorrect behaviour
	}

	for _, stream := range movieProbe.Streams {
		switch stream.CodecType {
		case "video":
			res.MediaStreams = append(res.MediaStreams, buildMediaStreamVideo(stream))
		case "audio":
			res.MediaStreams = append(res.MediaStreams, buildMediaStreamAudio(stream))
		case "subtitle":
			res.MediaStreams = append(res.MediaStreams, buildMediaStreamSubtitle(stream))
		}
	}

	resMediaSources := jfstructs.ResponseUsersItemsByUuidMovieMediaSources{
		Protocol:             "File",
		ID:                   item.Uuid,
		Path:                 res.Path,
		SupportsTranscoding:  true,
		SupportsDirectStream: false, //I dont think any modern jellyfin clients support/use direct stream, they either direct play or just use hls
		SupportsDirectPlay:   true,
		Type:                 "Default",
		Container:            res.Container,
	}

	//TODO: ???????
	resMediaSources.MediaStreams = res.MediaStreams
	res.MediaSources = append(res.MediaSources, resMediaSources)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func buildMediaStreamVideo(ffprobeStream mediamgmt.Stream) jfstructs.CommonStream {
	var jfStream jfstructs.CommonStream

	//TODO: make CodecName to human-readable string map

	jfStream.Index = ffprobeStream.Index
	jfStream.Codec = ffprobeStream.CodecName

	//TODO: default to sdr because i have no HDR media and therefore cannot implement HDR detection
	jfStream.VideoRange = "SDR"
	jfStream.VideoRangeType = "SDR"
	jfStream.AudioSpatialFormat = "None"
	jfStream.Language = "und"
	jfStream.Type = "Video"

	jfStream.TimeBase = ffprobeStream.TimeBase
	//TODO: Fix media resolution detection, simply printing the video height + "p" is incorrect
	jfStream.DisplayTitle = fmt.Sprintf("%dp %s %s", ffprobeStream.Height, ffprobeStream.CodecName, jfStream.VideoRangeType)

	if ffprobeStream.CodecName == "h264" {
		jfStream.IsAVC = true
		jfStream.CodecTag = "avc1"
	}

	return jfStream
}

func buildMediaStreamAudio(ffprobeStream mediamgmt.Stream) jfstructs.CommonStream {
	var jfStream jfstructs.CommonStream

	jfStream.Index = ffprobeStream.Index
	jfStream.Codec = ffprobeStream.CodecName
	jfStream.VideoRange = "Unknown"
	jfStream.VideoRangeType = "Unknown"
	jfStream.Type = "Audio"

	//TODO: make ffprobeStream.ChannelLayout start with an uppercase
	//TODO: make CodecName to human-readable string map

	if ffprobeStream.Tags.Title != "" {
		jfStream.DisplayTitle = fmt.Sprintf("%s - %s - %s", ffprobeStream.Tags.Title, ffprobeStream.CodecName, ffprobeStream.ChannelLayout)
	} else if ffprobeStream.Tags.Language != "" {
		iso639, err := language.Parse(ffprobeStream.Tags.Language)
		if err != nil {
			jfStream.DisplayTitle = fmt.Sprintf("%s - %s - %s", ffprobeStream.Tags.Language, ffprobeStream.CodecName, ffprobeStream.ChannelLayout)
		} else {
			jfStream.DisplayTitle = fmt.Sprintf("%s - %s - %s", display.English.Languages().Name(iso639), ffprobeStream.CodecName, ffprobeStream.ChannelLayout)
		}
	}

	return jfStream
}

func buildMediaStreamSubtitle(ffprobeStream mediamgmt.Stream) jfstructs.CommonStream {
	var jfStream jfstructs.CommonStream

	jfStream.Index = ffprobeStream.Index
	jfStream.Codec = ffprobeStream.CodecName
	jfStream.VideoRange = "Unknown"
	jfStream.VideoRangeType = "Unknown"
	jfStream.Type = "Subtitle"

	//TODO: make CodecName to human-readable string map

	if ffprobeStream.Tags.Title != "" {
		jfStream.DisplayTitle = fmt.Sprintf("%s - %s", ffprobeStream.Tags.Title, ffprobeStream.CodecName)
	} else if ffprobeStream.Tags.Language != "" {
		iso639, err := language.Parse(ffprobeStream.Tags.Language)
		if err != nil {
			jfStream.DisplayTitle = fmt.Sprintf("%s - %s", ffprobeStream.Tags.Language, ffprobeStream.CodecName)
		} else {
			jfStream.DisplayTitle = fmt.Sprintf("%s - %s", display.English.Languages().Name(iso639), ffprobeStream.CodecName)
		}
	}

	return jfStream
}

func EndpointUsersItems(w http.ResponseWriter, r *http.Request) {
	//TODO: respect a requsted client's fields instead of just sending them all

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	pathUserId := r.PathValue("id")
	aTokenString := getHeaderAuthToken(r)

	isReqValid := false

	//TODO: token lock halts any other goroutines from reading currentTokens, should change in future
	usermgmt.CurrentTokensMutex.Lock()

	for _, t := range usermgmt.CurrentTokens {
		if t.TokenString == aTokenString && pathUserId == t.UserId { //TODO: clients can only query their own user, allow admins to bypass this check
			isReqValid = true
		}
	}

	usermgmt.CurrentTokensMutex.Unlock()

	if !isReqValid {
		//TODO: should set WWW-Authenticate in header
		http.Error(w, "Permission denied", http.StatusUnauthorized)
	}

	query := r.URL.Query()

	parentId := query.Get("ParentId")
	node := librarymgmt.LibTreeMap[parentId]

	if node == nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	if node.Type != "folder" {
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}

	children := make([]*librarymgmt.TreeNode, len(node.Children))
	copy(children, node.Children)

	sortBy := query.Get("SortBy")
	sortOrder := strings.ToLower(query.Get("SortOrder")) //"Ascending" or "Descending"
	sort.Slice(children, func(i, j int) bool {

		if strings.Contains(sortBy, "IsFolder") {
			isFolderI := children[i].Type == "folder"
			isFolderJ := children[j].Type == "folder"

			if isFolderI != isFolderJ {
				return isFolderI
			}
		}

		//TODO: SortName isnt actually checked, its just always on
		nameI := strings.ToLower(children[i].Name)
		nameJ := strings.ToLower(children[j].Name)

		if sortOrder == "descending" {
			return nameI > nameJ
		}

		return nameI < nameJ
	})

	totalCount := len(children)
	startIndex, _ := strconv.Atoi(query.Get("StartIndex"))
	limit, _ := strconv.Atoi(query.Get("Limit"))

	if startIndex > len(children) {
		http.Error(w, "Malformed request", http.StatusBadRequest)
		return
	}

	//TODO: not 100% sure that this is correct behaviour
	if limit <= 0 {
		limit = totalCount
	}

	if limit <= 0 || startIndex+limit > totalCount {
		limit = totalCount - startIndex
	}

	resChildren := children[startIndex : startIndex+limit]

	resItems := make([]jfstructs.ResponseUsersItemsItem, 0)
	for _, child := range resChildren {
		cItem := jfstructs.ResponseUsersItemsItem{
			Name:     child.Name,
			ServerID: "STUBBED",
			ID:       child.Uuid,
		}

		switch child.Type {
		case "folder":
			cItem.Type = "Folder"
			cItem.IsFolder = true
		case "tvshow":
			cItem.Type = "Series"
			cItem.IsFolder = true
			cItem.Name = child.TvshowMetadata.Name
			cItem.ImageTags.Primary = "poster.png"
			cItem.PrimaryImageAspectRatio = child.TvshowMetadata.PosterPrimaryAspectRatio
		case "movie":
			cItem.Type = "Movie"
			cItem.IsFolder = false
			cItem.Name = child.MovieMetadata.Title
			cItem.ImageTags.Primary = "poster.png"
			cItem.PrimaryImageAspectRatio = child.MovieMetadata.PosterPrimaryAspectRatio
		}

		resItems = append(resItems, cItem)
	}

	res := jfstructs.ResponseUsersItems{
		TotalRecordCount: totalCount,
		StartIndex:       startIndex,
		Items:            resItems,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
