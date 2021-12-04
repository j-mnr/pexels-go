package pexels

import (
	"errors"
	"fmt"
)

type Media interface {
	MediaType() string
	isMedia()
}

const (
	videoType = "Video"
	photoType = "Photo"
)

type Collection struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	MediaCount  uint16 `json:"media_count"`
	PhotosCount uint16 `json:"photos_count"`
	VideosCount uint16 `json:"videos_count"`
}

type MediaPayload struct {
	ID    string  `json:"id"`
	Media []Media `json:"media"`
	Pagination
}

type CollectionPayload struct {
	ID          string       `json:"id"`
	Collections []Collection `json:"collection"`
	Pagination
}

type MediaResponse struct {
	Common ResponseCommon
	ID     string
	Videos []Video
	Photos []Photo
	Pagination
}

type CollectionResponse struct {
	Common     ResponseCommon
	Collection Collection
}

type CollectionsResponse struct {
	Common  ResponseCommon
	Payload CollectionPayload
}

type CollectionParams struct {
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

type CollectionMediaType string

const (
	VideoType CollectionMediaType = "videos"
	PhotoType CollectionMediaType = "photos"
)

type CollectionMediaParams struct {
	ID   string
	Type CollectionMediaType `query:"type"`
	CollectionParams
}

func (c *Client) GetCollection(params *CollectionMediaParams) (MediaResponse, error) {
	if params == nil || params.ID == "" {
		return MediaResponse{}, errors.New("Collection ID must be specified")
	}

	ID := params.ID
	params.ID = ""
	resp, err := c.get(fmt.Sprint("/collections/", ID), params, &MediaPayload{})
	if err != nil {
		return MediaResponse{}, err
	}

	cr := MediaResponse{}

	for _, m := range resp.Data.(MediaPayload).Media {
		switch m.MediaType() {
		case photoType:
			cr.Photos = append(cr.Photos, *m.(*Photo))
		case videoType:
			cr.Videos = append(cr.Videos, *m.(*Video))
		}
	}
	cr.ID = resp.Data.(MediaPayload).ID
	cr.Pagination = resp.Data.(MediaPayload).Pagination
	resp.copyCommon(&cr.Common)
	return cr, nil
}

func (c *Client) GetCollections() (CollectionsResponse, error) {
	resp, err := c.get("/collections", "", &CollectionPayload{})
	if err != nil {
		return CollectionsResponse{}, err
	}

	csr := CollectionsResponse{}
	csr.Payload = *resp.Data.(*CollectionPayload)
	resp.copyCommon(&csr.Common)
	return csr, nil
}
