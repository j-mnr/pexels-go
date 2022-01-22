package pexels

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetCollection(t *testing.T) {
	t.Parallel()
	for _, tc := range getCollectionTT {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetCollection(tc.params)
		if err != nil && err.Error() == "Collection ID must be specified" {
			continue
		}
		if aResp := aResp.Common.StatusCode; aResp != tc.statusCode {
			t.Errorf("expected status code %d, got %d", tc.statusCode, aResp)
		}
		if cmp.Equal(aResp.Videos, []Video{}) {
			vids := []Video{}
			json.Unmarshal([]byte(tc.respBody), &vids)
			if !cmp.Equal(aResp.Videos, vids) {
				t.Errorf("expected videos: %v, got %v", vids, aResp.Videos)
			}
		} else if cmp.Equal(aResp.Photos, []Photo{}) {
			photos := []Photo{}
			json.Unmarshal([]byte(tc.respBody), &photos)
			if !cmp.Equal(aResp.Photos, photos) {
				t.Errorf("expected photos: %v, got %v", photos, aResp.Videos)
			}
		}
	}
}

func TestGetCollections(t *testing.T) {
	t.Parallel()
	for _, tc := range getCollectionsTT {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetCollections()
		if err != nil {
			t.Errorf("Did not expect an error, got \"%s\"", err)
		}
		eCollections := CollectionPayload{}
		json.Unmarshal([]byte(tc.respBody), &eCollections)
		if len(aResp.Payload.Collections) != 0 {
			if !cmp.Equal(eCollections, aResp.Payload) {
				t.Errorf("expected collections %v, got %v", eCollections,
					aResp.Payload)
			}
		}
	}
}

func TestFailedGetCollection(t *testing.T) {
	t.Parallel()
	options := Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := Client{opts: options}
	params := CollectionMediaParams{
		ID:      "WONTWORK",
		Type:    VideoType,
		Page:    22,
		PerPage: 80,
	}
	zero, err := c.GetCollection(&params)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, MediaResponse{}) {
		t.Error("expected empty MediaResponse struct")
	}
}

func TestFailedGetCollections(t *testing.T) {
	t.Parallel()
	options := Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := Client{opts: options}
	zero, err := c.GetCollections()
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, CollectionsResponse{}) {
		t.Error("expected empty CollectionsResponse struct")
	}
}
