package pexels

import (
	"errors"
	"fmt"
)

// Media is either a Photo or a Video struct.
type Media interface {
	MediaType() string

	isMedia()
}

const (
	videoType = "Video"
	photoType = "Photo"
)

// Collection is a JSON formatted version of a Pexels collection.
type Collection struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	MediaCount  uint16 `json:"media_count"`
	PhotosCount uint16 `json:"photos_count"`
	VideosCount uint16 `json:"videos_count"`
}

// MediaPayload is all of the media (photos and videos) within a single
// collection.
type MediaPayload struct {
	ID    string  `json:"id"`
	Media []Media `json:"media"`
	Pagination
}

// CollectionPayload is all of the user's Collections
type CollectionPayload struct {
	ID          string       `json:"id"`
	Collections []Collection `json:"collections"`
	Pagination
}

// MediaResponse is all media given back from a single collection, even though
// videos and photos are in the response, it does not mean they will have any
// contents if your collection doesn't have either.
type MediaResponse struct {
	Common ResponseCommon
	ID     string
	Videos []Video
	Photos []Photo
	Pagination
}

// CollectionResponse has a common attributes of an HTTP response and the
// received collection.
type CollectionResponse struct {
	Common     ResponseCommon
	Collection Collection
}

// CollectionsResponse has a common attributes of an HTTP response and the
// received collections.
type CollectionsResponse struct {
	Common  ResponseCommon
	Payload CollectionPayload
}

// CollectionParams allows you to pick which page in your collections you start
// or how many per page you want.
type CollectionParams struct {
	// optional
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

// CollectionMediaType is a constant to which type of media you want.
type CollectionMediaType string

const (
  // VideoType ...
	VideoType CollectionMediaType = "videos"
  // PhotoType ...
	PhotoType CollectionMediaType = "photos"
)

// CollectionMediaParams is the way to get back a single collection. If you're
// looking for a certain Media type (photos or videos) it can be specified
// here.
type CollectionMediaParams struct {
	ID string

	// optional
	Type    CollectionMediaType `query:"type"`
	Page    uint16              `query:"page,1"`
	PerPage uint8               `query:"per_page,15"` // Max: 80
}

// GetCollection returns all the media based on parameters provided, within a
// single collection.
func (c *Client) GetCollection(params *CollectionMediaParams) (
	MediaResponse, error) {

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

// GetCollections returns all of your collections.
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
