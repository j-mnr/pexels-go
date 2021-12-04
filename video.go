package pexels

import (
	"errors"
	"fmt"
)

const videoPath = "/videos"

type Video struct {
	ID            uint64              `json:"id"`
	Width         uint16              `json:"width"`
	Height        uint16              `json:"height"`
	URL           string              `json:"url"`
	Image         string              `json:"image"`
	Duration      uint16              `json:"duration"` // In seconds
	User          PexelUser           `json:"user"`
	VideoFiles    []VideoFile         `json:"video_files"`
	VideoPictures []VideoPicture      `json:"video_pictures"`
	Type          CollectionMediaType `json:"type,omitempty"`
}

func (v *Video) isMedia()
func (v *Video) MediaType() string { return videoType }

type VideoResponse struct {
	Common ResponseCommon
	Video  Video
}

type VideoFile struct {
	ID       uint64  `json:"id"`
	Quality  Quality `json:"quality"` // SD or HD
	FileType string  `json:"file_type"`
	Width    uint16  `json:"width"`
	Height   uint16  `json:"height"`
	Link     string  `json:"link"`
}

type VideoPicture struct {
	ID      uint64 `json:"id"`
	Picture string `json:"picture"`
	NR      uint8  `json:"nr"` // Number of Records; Index of Object in Slice
}

type VideoPayload struct {
	Videos []Video `json:"videos"`
	Pagination
}

type VideosResponse struct {
	Common  ResponseCommon
	Payload VideoPayload
}

type Quality string

const (
	SD Quality = "sd"
	HD Quality = "hd"
)

type PopularVideoParams struct {
	MinWidth    uint16 `query:"min_width"`
	MinHeight   uint16 `query:"min_height"`
	MinDuration uint16 `query:"min_duration"` // In Seconds
	MaxDuration uint16 `query:"max_duration"` // In Seconds
	Page        uint16 `query:"page,1"`
	PerPage     uint8  `query:"per_page,15"` // Max: 80
}

type VideoSearchParams struct {
	Query string `query:"query"` // Required

	// Optional parameters
	Locale Locale `query:"locale"`
	// Landscape, Portrait, Square
	Orientation Orientation `query:"orientation"`
	// Large (24MP), Medium (12MP), Small (4MP)
	Size    Size   `query:"size"`
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

type PexelUser struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) GetVideo(videoID int) (VideoResponse, error) {
	resp, err := c.get(fmt.Sprintf("%s/videos/%d", videoPath, videoID), "", &Video{})
	if err != nil {
		return VideoResponse{}, err
	}

	vr := VideoResponse{Video: *resp.Data.(*Video)}
	resp.copyCommon(&vr.Common)
	return vr, nil
}

func (c *Client) GetPopularVideos(params *PopularVideoParams) (VideosResponse,
	error) {

	resp, err := c.get(fmt.Sprintf("%s/popular", videoPath), params, &VideoPayload{})
	if err != nil {
		return VideosResponse{}, err
	}

	popResp := VideosResponse{Payload: *resp.Data.(*VideoPayload)}
	resp.copyCommon(&popResp.Common)
	return popResp, nil
}

func (c *Client) SearchVideos(params *VideoSearchParams) (VideosResponse,
	error) {

	if params == nil || params.Query == "" {
		return VideosResponse{}, errors.New("Query is required")
	}
	resp, err := c.get(fmt.Sprintf("%s/search", videoPath), params, &VideoPayload{})
	if err != nil {
		return VideosResponse{}, err
	}

	vsr := VideosResponse{Payload: *resp.Data.(*VideoPayload)}
	resp.copyCommon(&vsr.Common)
	return vsr, nil
}
