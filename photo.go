package pexels

import (
	"errors"
	"fmt"
)

const photoPath = "/v1"

type Photo struct {
	ID              uint64              `json:"id"`
	Width           uint16              `json:"width"`
	Height          uint16              `json:"height"`
	URL             string              `json:"url"`
	Photographer    string              `json:"photographer"`
	PhotographerURL string              `json:"photographer_url"`
	PhotographerID  uint64              `json:"photographer_id"`
	AvgColor        string              `json:"avg_color"`
	Type            CollectionMediaType `json:"type,omitempty"`
	Src             PhotoSource         `json:"src"`
	Liked           bool                `json:"liked"` // Optional
}

func (p *Photo) isMedia()
func (p *Photo) MediaType() string { return photoType }

type PhotoResponse struct {
	Common ResponseCommon
	Photo  Photo
}

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

type PhotoPayload struct {
	Photos []Photo `json:"photos"`
	Pagination
}

type PhotosResponse struct {
	Common  ResponseCommon
	Payload PhotoPayload
}

type CuratedPhotosParams struct {
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max 80
}

type SearchPhotosParams struct {
	Query string `query:"query"` // Required

	// Optional parameters
	Locale Locale `query:"locale"`
	// Landscape, Portrait, Square
	Orientation Orientation `query:"orientation"`
	// Large (24MP), Medium (12MP), Small (4MP)
	Size    Size   `query:"size"`
	Color   Color  `query:"color"`
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

type Color string

const (
	Red       Color = "red"
	Orange    Color = "orange"
	Yellow    Color = "yellow"
	Green     Color = "green"
	Turquoise Color = "turquoise"
	Blue      Color = "blue"
	Violet    Color = "violet"
	Pink      Color = "pink"
	Brown     Color = "brown"
	Black     Color = "black"
	Gray      Color = "gray"
	White     Color = "white"
)

// GetPhoto retrieves a photo by its ID
func (c *Client) GetPhoto(ID int) (PhotoResponse, error) {
	resp, err := c.get(fmt.Sprintf("%s/%s/%d", photoPath, "photos", ID), nil, &Photo{})
	if err != nil {
		return PhotoResponse{}, err
	}

	pr := PhotoResponse{}
	pr.Photo = *resp.Data.(*Photo)
	resp.copyCommon(&pr.Common)
	return pr, nil
}

func (c *Client) GetCuratedPhotos(params *CuratedPhotosParams) (PhotosResponse, error) {
	resp, err := c.get(fmt.Sprintf("%s/%s", photoPath, "curated"), params, &PhotoPayload{})
	if err != nil {
		return PhotosResponse{}, err
	}

	ppr := PhotosResponse{}
	ppr.Payload = *resp.Data.(*PhotoPayload)
	resp.copyCommon(&ppr.Common)
	return ppr, nil
}

func (c *Client) SearchPhotos(params *SearchPhotosParams) (PhotosResponse, error) {
	if params == nil || params.Query == "" {
		return PhotosResponse{}, errors.New("query is required")
	}

	resp, err := c.get(fmt.Sprintf("%s/%s", photoPath, "curated"), params, &PhotoPayload{})
	if err != nil {
		return PhotosResponse{}, err
	}

	ppr := PhotosResponse{}
	ppr.Payload = *resp.Data.(*PhotoPayload)
	resp.copyCommon(&ppr.Common)
	return ppr, nil
}
