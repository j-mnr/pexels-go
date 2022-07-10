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

// Photo is the base data structure returned when consuming Pexels API Photo
// endpoints.
type Photo struct {
	ID              uint64      `json:"id"`
	Width           uint16      `json:"width"`
	Height          uint16      `json:"height"`
	URL             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerURL string      `json:"photographer_url"`
	PhotographerID  uint64      `json:"photographer_id"`
	AvgColor        string      `json:"avg_color"`
	Type            string      `json:"type,omitempty"`
	Src             PhotoSource `json:"src"`

	Liked bool `json:"liked"`
}

func (p *Photo) isMedia() {}

// MediaType is used to identify which type of Media a resource is.
// It always returns "Photo"
func (p *Photo) MediaType() string { return photoType }

// PhotoResponse has a common attributes of an HTTP response and the
// received Photo.
type PhotoResponse struct {
	Common ResponseCommon
	Photo  Photo
}

// PhotoSource is an assortment of different image sizes that can be used to
// display a Photo.
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

// PhotoPayload is a slice of Photo with Pagination.
type PhotoPayload struct {
	Photos []Photo `json:"photos"`
	Pagination
}

// PhotosResponse has common values of an HTTP response and the received
// PhotoPayload response.
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

// PhotoSearchParams requires Query. It has all of the available parameters
// by which you can search for a photo.
type PhotoSearchParams struct {
	Query string `query:"query"` // required

	General
	// the supported pexels colors which you can search with are:
	// red, orange, yellow, green, turquoise, blue, violet, pink, brown, black,
	// gray, white.
	Color   string `query:"color"`
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

// GetPhoto retreives a photo by its ID found at the end of its URL
func (c *client) GetPhoto(photoID uint64) (PhotoResponse, error) {
	resp, err := c.get(fmt.Sprint(photoEndpoint, photoID), "", &Photo{})
	if err != nil {
		return PhotoResponse{}, err
	}
	pr := PhotoResponse{}
	pr.Photo = *resp.Data.(*Photo)
	resp.copyCommon(&pr.Common)
	return pr, nil
}

// GetCuratedPhotos retrieves the current Curated list, updated hourly by
// Pexels. If nil is passed it will default to the first page and return 15
// photos.
func (c *client) GetCuratedPhotos(params *CuratedPhotosParams) (PhotosResponse,
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

// SearchPhotos returns a slice of Photos 15 photos by defualt.
// The PhotoSearchParams.Query is required and SearchPhotos will return an
// error if it is nil.
func (c *client) SearchPhotos(params *PhotoSearchParams) (PhotosResponse,
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
