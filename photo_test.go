package pexels

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetPhoto(t *testing.T) {
	t.Parallel()
	for _, tc := range getPhotoTT {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetPhoto(tc.ID)
		if err != nil {
			t.Error(err)
		}
		if aResp := aResp.Common.StatusCode; aResp != tc.statusCode {
			t.Errorf("expected status code %d, got %d", tc.statusCode, aResp)
		}
		ePhoto := Photo{}
		json.Unmarshal([]byte(tc.respBody), &ePhoto)
		if aResp.Photo != ePhoto {
			t.Errorf("expected photo %v, got %v", ePhoto, aResp.Photo)
		}
	}
}

func TestGetCuratedPhotos(t *testing.T) {
	t.Parallel()
	for _, tc := range getCuratedPhotosTT {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetCuratedPhotos(tc.params)
		if err != nil {
			t.Errorf("Did not expect an error, got \"%s\"", err)
		}

		if aResp.Payload.Page != tc.params.Page {
			t.Errorf("expected page to be %d, got %d",
				aResp.Payload.Page, tc.params.Page)
		} else if aResp.Payload.PerPage != tc.params.PerPage {
			t.Errorf("expected per_page to be %d, got %d",
				aResp.Payload.PerPage, tc.params.PerPage)
		}

		ePayload := PhotoPayload{}
		json.Unmarshal([]byte(tc.respBody), &ePayload)
		if !cmp.Equal(aResp.Payload, ePayload) {
			t.Errorf("expected payload %v, got %v", ePayload, aResp.Payload)
		}
	}
}

func TestSearchPhotos(t *testing.T) {
	t.Parallel()
	for _, tc := range searchPhotosTT {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.SearchPhotos(tc.params)
		if err != nil && err.Error() == "Query is required" {
			continue
		} else if err != nil {
			t.Errorf("Did not expect an error, got \"%s\"", err)
		}
		if aResp.Common.StatusCode != tc.statusCode {
			t.Errorf("expected status code to be %d, got %d", tc.statusCode,
				aResp.Common.StatusCode)
		}

		ePage := tc.params.Page
		ePerPage := tc.params.PerPage
		// We need this because the API defaults to page 1 and per_page 1
		if len(aResp.Payload.Photos) != 0 {
			if aResp.Payload.Page != ePage {
				t.Errorf("expected page to be %d, got %d", aResp.Payload.Page, ePage)
			} else if aResp.Payload.PerPage != ePerPage {
				t.Errorf("expected per_page to be %d, got %d",
					aResp.Payload.PerPage, ePerPage)
			}
		}

		ePayload := PhotoPayload{}
		json.Unmarshal([]byte(tc.respBody), &ePayload)
		if !cmp.Equal(aResp.Payload, ePayload) {
			t.Errorf("expected payload %v, got %v", ePayload, aResp.Payload)
		}
	}
}

func TestFailedCuratedPhotos(t *testing.T) {
	t.Parallel()
	options := Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := Client{opts: options}
	params := CuratedPhotosParams{
		Page:    2,
		PerPage: 80,
	}

	zero, err := c.GetCuratedPhotos(&params)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, PhotosResponse{}) {
		t.Error("expected empty PhotosResponse struct")
	}
}

func TestFailedSearchPhotos(t *testing.T) {
	t.Parallel()
	options := Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := Client{opts: options}
	params := PhotoSearchParams{
		Query:       "Failure",
		Locale:      UK_UA,
		Orientation: Landscape,
		Size:        Medium,
		Color:       Red,
		Page:        1,
		PerPage:     1,
	}
	zero, err := c.SearchPhotos(&params)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, PhotosResponse{}) {
		t.Error("expected empty PhotosResponse struct")
	}
}

func TestPhotoIsMedia(t *testing.T) {
	t.Parallel()

	p := Photo{}
	p.isMedia()
	if p.MediaType() != photoType {
		t.Errorf("Expected photo media type to be %s, got %s", photoType,
			p.MediaType())
	}
}
