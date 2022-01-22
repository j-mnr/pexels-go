package pexels

import (
	"errors"
	"fmt"
)

const (
	photoEndpoint         = "/photos/"
	searchPhotosEndpoint  = "/search"
	curatedPhotosEndpoint = "/curated"
)

// Photo is a JSON formatted version of a Pexels photo.
type Photo struct {
	ID              uint64              `json:"id"`
	Width           uint16              `json:"width"`
	Height          uint16              `json:"height"`
	URL             string              `json:"url"`
	Photographer    string              `json:"photographer"`
	PhotographerURL string              `json:"photographer_url"`
	PhotographerID  uint64              `json:"photographer_id"`
	AvgColor        string              `json:"avg_color"` // In hex e.g. #978E82
	Type            CollectionMediaType `json:"type,omitempty"`
	Src             PhotoSource         `json:"src"`   // URLs of images
	Liked           bool                `json:"liked"` // Optional
}

func (p *Photo) isMedia() {}

// MediaType always returns the string "Photo"
func (p *Photo) MediaType() string { return photoType }

// PhotoResponse has a common attributes of an HTTP response and the
// received photo.
type PhotoResponse struct {
	Common ResponseCommon
	Photo  Photo
}

// PhotoSource is an assortment of different image sizes that can be used to
// display a photo.
type PhotoSource struct {
	Original  string `json:"original"`
	Large2x   string `json:"large2x"`   // W 940px X H 650px DPR 1
	Large     string `json:"large"`     // W 940px X H 650px DPR 2
	Medium    string `json:"medium"`    // W scaled H 350px
	Small     string `json:"small"`     // W scaled H 130px
	Portrait  string `json:"portrait"`  // W 800px X H 1200px
	Landscape string `json:"landscape"` // W 1200px X H 627px
	Tiny      string `json:"tiny"`      // W 280px X H 200px
}

// PhotoPayload is a slice of photos with a pagination struct.
type PhotoPayload struct {
	Photos []Photo `json:"photos"`
	Pagination
}

// PhotosResponse has a common attributes of an HTTP response and the
// received photos.
type PhotosResponse struct {
	Common  ResponseCommon
	Payload PhotoPayload
}

// CuratedPhotosParams allows you to pick which page in your collections you
// start or how many per page you want.
type CuratedPhotosParams struct {
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max 80
}

// Color is the supported pexels' colors which you can search with.
type Color string

const (
	Red       = "red"
	Orange    = "orange"
	Yellow    = "yellow"
	Green     = "green"
	Turquoise = "turquoise"
	Blue      = "blue"
	Violet    = "violet"
	Pink      = "pink"
	Brown     = "brown"
	Black     = "black"
	Gray      = "gray"
	White     = "white"
)

// PhotoSearchParams requires a query. It has all of the available parameters
// by which you can search for a photo.
type PhotoSearchParams struct {
	Query string `query:"query"`

	// optional
	Locale Locale `query:"locale"`
	// Landscape, Portrait, Square
	Orientation Orientation `query:"orientation"`
	Size    Size   `query:"size"`
	Color   Color  `query:"color"`
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

// Retreives a photo by its ID found at the end of its URL
func (c *Client) GetPhoto(photoID uint64) (PhotoResponse, error) {
	resp, err := c.get(fmt.Sprint(photoEndpoint, photoID), "", &Photo{})
	if err != nil {
		return PhotoResponse{}, err
	}
	pr := PhotoResponse{}
	pr.Photo = *resp.Data.(*Photo)
	resp.copyCommon(&pr.Common)
	return pr, nil
}

// Retrieves the current Curated list, updated hourly by Pexels. If nil is
// passed it will default to the first page and return 15 photos.
func (c *Client) GetCuratedPhotos(params *CuratedPhotosParams) (PhotosResponse,
	error) {

	resp, err := c.get(curatedPhotosEndpoint, params, &PhotoPayload{})
	if err != nil {
		return PhotosResponse{}, err
	}
	cr := PhotosResponse{}
	cr.Payload = *resp.Data.(*PhotoPayload)
	resp.copyCommon(&cr.Common)
	return cr, nil
}

// Returns an error if the Query parameter has no value
func (c *Client) SearchPhotos(params *PhotoSearchParams) (PhotosResponse,
	error) {

	if params == nil || params.Query == "" {
		return PhotosResponse{}, errors.New("Query is required")
	}
	resp, err := c.get(searchPhotosEndpoint, params, &PhotoPayload{})
	if err != nil {
		return PhotosResponse{}, err
	}
	psr := PhotosResponse{}
	psr.Payload = *resp.Data.(*PhotoPayload)
	resp.copyCommon(&psr.Common)
	return psr, nil
}
