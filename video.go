package pexels

import (
	"errors"
	"fmt"
)

var ErrMissingQuery = errors.New("query is required")

const (
	videoEndpoint         = "/videos/videos/"
	popularVideosEndpoint = "/videos/popular"
	searchVideosEndpoint  = "/videos/search"
)

// Video is the base data structure returned when consuming Pexels API Video
// endpoints.
//
//nolint:tagliatelle
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

func (Video) isMedia() {}

// MediaType is used to identify which type of Media a resource is.
// It always returns "Video".
func (Video) MediaType() Type { return TypeVideo }

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
	ID       uint64 `json:"id"`
	Quality  string `json:"quality"`   // supported qualities are: sd, hd.
	FileType string `json:"file_type"` //nolint:tagliatelle
	Width    uint16 `json:"width"`
	Height   uint16 `json:"height"`
	Link     string `json:"link"`
}

// VideoPicture is a preview picture for a Video.
type VideoPicture struct {
	ID      uint64 `json:"id"`
	Picture string `json:"picture"`
	NR      uint8  `json:"nr"` // Index of Picture in Slice
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

// PopularVideoParams is parameters that can be selected for when searching for
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
	// Query is required
	Query string `query:"query"`

	General
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

// GetVideo returns a Video based on its ID. It does not return an error if the
// Video could not be found by its ID, only if something went wrong while
// getting the resource.
func (c *Client) GetVideo(videoID uint64) (VideoResponse, error) {
	resp, err := get(*c, fmt.Sprint(videoEndpoint, videoID), "", &Video{})
	if err != nil {
		return VideoResponse{}, err
	}
	vr := VideoResponse{Video: *resp.Data}
	resp.copyCommon(&vr.Common)
	return vr, nil
}

// GetPopularVideos returns the current popular pexels videos.
func (c *Client) GetPopularVideos(
	pvp *PopularVideoParams,
) (VideosResponse, error) {
	resp, err := get(*c, popularVideosEndpoint, pvp, &VideoPayload{})
	if err != nil {
		return VideosResponse{}, err
	}
	popResp := VideosResponse{Payload: *resp.Data}
	resp.copyCommon(&popResp.Common)
	return popResp, nil
}

// SearchVideos enables you to search the entire pexels database for any
// subject that you would like and receive videos on that subject.
func (c *Client) SearchVideos(
	vsp *VideoSearchParams,
) (VideosResponse, error) {
	if vsp == nil || vsp.Query == "" {
		return VideosResponse{}, ErrMissingQuery
	}
	resp, err := get(*c, searchVideosEndpoint, vsp, &VideoPayload{})
	if err != nil {
		return VideosResponse{}, err
	}
	vsr := VideosResponse{Payload: *resp.Data}
	resp.copyCommon(&vsr.Common)
	return vsr, nil
}
