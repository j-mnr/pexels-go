package pexels

import (
	"errors"
	"fmt"
)

const (
	videoEndpoint         = "/videos/videos/"
	popularVideosEndpoint = "/videos/popular"
	searchVideosEndpoint  = "/videos/search"
)

// Video is the base data structure returned when consuming Pexels API Video
// endpoints.
type Video struct {
	ID            uint64         `json:"id"`
	Width         uint16         `json:"width"`
	Height        uint16         `json:"height"`
	URL           string         `json:"url"`
	Image         string         `json:"image"`
	Duration      uint16         `json:"duration"` // In seconds
	User          PexelUser      `json:"user"`
	VideoFiles    []VideoFile    `json:"video_files"`
	VideoPictures []VideoPicture `json:"video_pictures"`
	Type          string         `json:"type,omitempty"`
}

func (v *Video) isMedia() {}

// MediaType is used to identify which type of Media a resource is.
// It always returns "Video"
func (v *Video) MediaType() string { return videoType }

// VideoResponse has a common attributes of an HTTP response and the
// received video.
type VideoResponse struct {
	Common ResponseCommon
	Video  Video
}

// PexelUser represents the videographer who shot the Video.
type PexelUser struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// VideoFile is a version of a Video.
type VideoFile struct {
	ID uint64 `json:"id"`
	// supported qualities are: sd, hd.
	Quality  string `json:"quality"`
	FileType string `json:"file_type"`
	Width    uint16 `json:"width"`
	Height   uint16 `json:"height"`
	Link     string `json:"link"`
}

// VideoPicture is a preview picture for a Video.
type VideoPicture struct {
	ID      uint64 `json:"id"`
	Picture string `json:"picture"`
	NR      uint8  `json:"nr"` // Index of Object in Slice
}

// VideoPayload is a slice of Video with Pagination.
type VideoPayload struct {
	Videos []Video `json:"videos"`
	Pagination
}

// VideosResponse has common values of an HTTP response and the received
// VideoPayload response.
type VideosResponse struct {
	Common  ResponseCommon
	Payload VideoPayload
}

// PopularVideoParams is paramaters that can be selected for when searching for
// specific popular videos on pexels.
type PopularVideoParams struct {
	MinWidth    uint16 `query:"min_width"`
	MinHeight   uint16 `query:"min_height"`
	MinDuration uint16 `query:"min_duration"` // In Seconds
	MaxDuration uint16 `query:"max_duration"` // In Seconds
	Page        uint16 `query:"page,1"`
	PerPage     uint8  `query:"per_page,15"` // Max: 80
}

// VideoSearchParams requires Query. A Query allows you to search for any topic
// that you would like to receive video information about.
type VideoSearchParams struct {
	Query string `query:"query"` // required

	General
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

// GetVideo returns a Video based on its ID. It does not return an error if the
// Video could not be found by its ID, only if something went wrong while
// getting the resource.
func (c *client) GetVideo(videoID uint64) (VideoResponse, error) {
	resp, err := c.get(fmt.Sprint(videoEndpoint, videoID), "", &Video{})
	if err != nil {
		return VideoResponse{}, err
	}
	vr := VideoResponse{Video: *resp.Data.(*Video)}
	resp.copyCommon(&vr.Common)
	return vr, nil
}

// GetPopularVideos returns the current popular pexels videos.
func (c *client) GetPopularVideos(params *PopularVideoParams) (VideosResponse,
	error) {

	resp, err := c.get(popularVideosEndpoint, params, &VideoPayload{})
	if err != nil {
		return VideosResponse{}, err
	}
	popResp := VideosResponse{Payload: *resp.Data.(*VideoPayload)}
	resp.copyCommon(&popResp.Common)
	return popResp, nil
}

// SearchVideos enables you to search the entire pexels database for any
// subject that you would like and receive videos on that subject.
func (c *client) SearchVideos(params *VideoSearchParams) (VideosResponse,
	error) {

	if params == nil || params.Query == "" {
		return VideosResponse{}, errors.New("Query is required")
	}
	resp, err := c.get(searchVideosEndpoint, params, &VideoPayload{})
	if err != nil {
		return VideosResponse{}, err
	}
	vsr := VideosResponse{Payload: *resp.Data.(*VideoPayload)}
	resp.copyCommon(&vsr.Common)
	return vsr, nil
}
