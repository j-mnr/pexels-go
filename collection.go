package pexels

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrMissingCollectionID = errors.New("a collection ID must be specified")

// Media is either Photo or Video.
type Media interface {
	MediaType() Type
	isMedia()
}

type Type interface{ enum() }

type typ string

func (typ) enum() {}

const (
	TypeVideo typ = "Video"
	TypePhoto typ = "Photo"
)

// Collection is the base data structure returned when consuming Pexels API
// Collection endpoints.
//
//nolint:tagliatelle
type Collection struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	MediaCount  uint16 `json:"media_count"`
	PhotosCount uint16 `json:"photos_count"`
	VideosCount uint16 `json:"videos_count"`
}

// MediaPayload is all of the Media (photos and videos) within a single
// collection.
type MediaPayload struct {
	ID    string  `json:"id"`
	Media []Media `json:"media"`
	Pagination
}

func (p *MediaPayload) UnmarshalJSON(raw []byte) error {
	var data struct {
		ID         string            `json:"id"`
		Media      []json.RawMessage `json:"media"`
		Pagination `json:"pagination"`
	}
	if err := json.Unmarshal(raw, &data); err != nil {
		return fmt.Errorf(wrapFmt, err)
	}
	ms := make([]Media, 0, len(data.Media))
	for _, rawMsg := range data.Media {
		m, err := decodeMediaFrom(rawMsg)
		if err != nil {
			return err
		}
		ms = append(ms, m)
	}
	p.ID = data.ID
	p.Media = ms
	p.Pagination = data.Pagination
	return nil
}

func decodeMediaFrom(data []byte) (Media, error) {
	var typeData struct {
		Type Type `json:"type"`
	}
	if err := json.Unmarshal(data, &typeData); err != nil {
		return nil, fmt.Errorf(wrapFmt, err)
	}
	var m Media
	switch typeData.Type {
	case TypeVideo:
		m = &Video{}
	case TypePhoto:
		m = &Photo{}
	default:
		return nil, ErrUnsupportedType
	}
	return m, nil
}

// CollectionPayload is all of the user's Collections.
type CollectionPayload struct {
	ID          string       `json:"id"`
	Collections []Collection `json:"collections"`
	Pagination
}

// MediaResponse is all media given back from a single collection, even though
// videos and photos are in the response, they may be empty slices if your
// collection doesn't have either.
type MediaResponse struct {
	Common ResponseCommon
	ID     string
	Videos []Video
	Photos []Photo
	Pagination
}

// CollectionResponse has a common attributes of a HTTP response and the
// received collection.
type CollectionResponse struct {
	Common     ResponseCommon
	Collection Collection
}

// CollectionsResponse has a common attributes of a HTTP response and the
// received collections.
type CollectionsResponse struct {
	Common  ResponseCommon
	Payload CollectionPayload
}

// CollectionParams allows you to pick which page to start at in your
// collections and how many per page you want.
type CollectionParams struct {
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

// CollectionMediaParams is the way to get back a single collection. If you're
// looking for a certain Media type (photos or videos) it can be specified
// here.
type CollectionMediaParams struct {
	ID string

	// Supported types are: videos, photos.
	Type    string `query:"type"`
	Page    uint16 `query:"page,1"`
	PerPage uint8  `query:"per_page,15"` // Max: 80
}

// GetCollection returns all the media based on parameters provided, within a
// single collection.
func (c *Client) GetCollection(
	params *CollectionMediaParams,
) (MediaResponse, error) {
	if params == nil || params.ID == "" {
		return MediaResponse{}, ErrMissingCollectionID
	}
	id := params.ID
	params.ID = ""
	resp, err := get(*c, fmt.Sprint("/collections/", id), params, &MediaPayload{})
	if err != nil {
		return MediaResponse{}, err
	}
	cr := MediaResponse{}
	for _, m := range resp.Data.Media {
		switch v := m.(type) {
		case *Photo:
			cr.Photos = append(cr.Photos, *v)
		case *Video:
			cr.Videos = append(cr.Videos, *v)
		}
	}
	cr.ID = resp.Data.ID
	cr.Pagination = resp.Data.Pagination
	resp.copyCommon(&cr.Common)
	return cr, nil
}

// GetCollections returns all of your collections.
func (c *Client) GetCollections() (CollectionsResponse, error) {
	resp, err := get(*c, "/collections", "", &CollectionPayload{})
	if err != nil {
		return CollectionsResponse{}, err
	}
	csr := CollectionsResponse{}
	csr.Payload = *resp.Data
	resp.copyCommon(&csr.Common)
	return csr, nil
}
