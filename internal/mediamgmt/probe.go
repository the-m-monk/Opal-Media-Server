package mediamgmt

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Stream struct {
	//TODO: seperate out fields by type (video,audio,subtitles,etc)
	Index              int    `json:"index"`
	CodecName          string `json:"codec_name"`
	CodecLongName      string `json:"codec_long_name"`
	Profile            string `json:"profile"`
	CodecType          string `json:"codec_type"`
	CodecTagString     string `json:"codec_tag_string"`
	CodecTag           string `json:"codec_tag"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	CodedWidth         int    `json:"coded_width"`
	CodedHeight        int    `json:"coded_height"`
	ClosedCaptions     int    `json:"closed_captions"`
	FilmGrain          int    `json:"film_grain"`
	HasBFrames         int    `json:"has_b_frames"`
	SampleAspectRatio  string `json:"sample_aspect_ratio"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	PixFmt             string `json:"pix_fmt"`
	Level              int    `json:"level"`
	ColorRange         string `json:"color_range"`
	ChromaLocation     string `json:"chroma_location"`
	Refs               int    `json:"refs"`
	ViewIdsAvailable   string `json:"view_ids_available"`
	ViewPosAvailable   string `json:"view_pos_available"`
	RFrameRate         string `json:"r_frame_rate"`
	AvgFrameRate       string `json:"avg_frame_rate"`
	TimeBase           string `json:"time_base"`
	StartPts           int    `json:"start_pts"`
	StartTime          string `json:"start_time"`
	ExtradataSize      int    `json:"extradata_size"`
	Disposition        struct {
		Default         int `json:"default"`
		Dub             int `json:"dub"`
		Original        int `json:"original"`
		Comment         int `json:"comment"`
		Lyrics          int `json:"lyrics"`
		Karaoke         int `json:"karaoke"`
		Forced          int `json:"forced"`
		HearingImpaired int `json:"hearing_impaired"`
		VisualImpaired  int `json:"visual_impaired"`
		CleanEffects    int `json:"clean_effects"`
		AttachedPic     int `json:"attached_pic"`
		TimedThumbnails int `json:"timed_thumbnails"`
		NonDiegetic     int `json:"non_diegetic"`
		Captions        int `json:"captions"`
		Descriptions    int `json:"descriptions"`
		Metadata        int `json:"metadata"`
		Dependent       int `json:"dependent"`
		StillImage      int `json:"still_image"`
		Multilayer      int `json:"multilayer"`
	} `json:"disposition"`

	Tags struct {
		Title    string `json:"title"`
		Language string `json:"language"`
	} `json:"tags"`

	//Audio only fields
	SampleFmt      string `json:"sample_fmt"`
	SampleRate     string `json:"sample_rate"`
	Channels       int    `json:"channels"`
	ChannelLayout  string `json:"channel_layout"`
	BitsPerSample  int    `json:"bits_per_sample"`
	InitialPadding int    `json:"initial_padding"`
}

type FfprobeResult struct {
	Streams []Stream `json:"streams"`
	Format  struct {
		Filename       string `json:"filename"`
		NbStreams      int    `json:"nb_streams"`
		NbPrograms     int    `json:"nb_programs"`
		NbStreamGroups int    `json:"nb_stream_groups"`
		FormatName     string `json:"format_name"`
		FormatLongName string `json:"format_long_name"`
		StartTime      string `json:"start_time"`
		Duration       string `json:"duration"`
		Size           string `json:"size"`
		BitRate        string `json:"bit_rate"`
		ProbeScore     int    `json:"probe_score"`
		//Tags           struct {} `json:"tags"`
	} `json:"format"`
}

func Probe(path string) (*FfprobeResult, error) {
	ffprobeArgs := []string{
		"-v", "error",
		"-show_format",
		"-show_streams",
		"-print_format", "json",
		path,
	}

	ffprobeCmd := exec.Command("ffprobe", ffprobeArgs...)

	ffprobeOutput, err := ffprobeCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("ffprobe error: %s", err.Error())
	}

	var result FfprobeResult
	if err := json.Unmarshal(ffprobeOutput, &result); err != nil {
		return nil, fmt.Errorf("ffprobe error: %s", err.Error())
	}

	return &result, nil
}
