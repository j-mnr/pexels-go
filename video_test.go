package pexels

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetVideo(t *testing.T) {
	t.Parallel()
	for _, tc := range getVideoTT {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetVideo(tc.ID)
		if err != nil {
			t.Error(err)
		}
		if aResp := aResp.Common.StatusCode; aResp != tc.statusCode {
			t.Errorf("expected status code %d, got %d", tc.statusCode, aResp)
		}
		eVideo := Video{}
		json.Unmarshal([]byte(tc.respBody), &eVideo)
		if !cmp.Equal(aResp.Video, eVideo) {
			t.Errorf("expected video %v, got %v", eVideo, aResp.Video)
		}
	}
}

func TestGetPopularVideos(t *testing.T) {
	t.Parallel()
	for _, tc := range getPopularVideosTT {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetPopularVideos(tc.params)
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
		for _, v := range aResp.Payload.Videos {
			if v.Duration < tc.params.MinDuration ||
				v.Duration > tc.params.MaxDuration {
				t.Errorf("expected duration in range %d - %d, got %d",
					tc.params.MinDuration, tc.params.MaxDuration, v.Duration)
			} else if v.Height < tc.params.MinHeight {
				t.Errorf("expected height to be greater than %d, got %d",
					tc.params.MinHeight, v.Height)
			} else if v.Width < tc.params.MinWidth {
				t.Errorf("expected width to be greater than %d, got %d",
					tc.params.MinWidth, v.Width)
			}
		}
		ePayload := VideoPayload{}
		json.Unmarshal([]byte(tc.respBody), &ePayload)
		if !cmp.Equal(aResp.Payload, ePayload) {
			t.Errorf("expected payload %v, got %v", ePayload, aResp.Payload)
		}
	}
}

func TestSearchVideos(t *testing.T) {
	t.Parallel()
	for _, tc := range searchVideosTT {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.SearchVideos(tc.params)
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
		ePayload := VideoPayload{}
		// We need this because the API defaults to page 1 and per_page 1
		if len(aResp.Payload.Videos) != 0 {
			if aResp.Payload.Page != ePage {
				t.Errorf("expected page to be %d, got %d", aResp.Payload.Page, ePage)
			} else if aResp.Payload.PerPage != ePerPage {
				t.Errorf("expected per_page to be %d, got %d",
					aResp.Payload.PerPage, ePerPage)
			}
			json.Unmarshal([]byte(tc.respBody), &ePayload)
			if !cmp.Equal(aResp.Payload, ePayload) {
				t.Errorf("expected payload %v, got %v", ePayload, aResp.Payload)
			}
			if len(aResp.Payload.Videos) != len(ePayload.Videos) {
				t.Errorf("expected %d videos, got %d", len(ePayload.Videos),
					len(aResp.Payload.Videos))
			}
		}
	}
}

func TestFailedGetVideo(t *testing.T) {
	t.Parallel()
	options := Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := Client{opts: options}
	zero, err := c.GetVideo(999999999999999)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, VideoResponse{}) {
		t.Error("expected empty VideoResponse struct")
	}
}

func TestFailedPopularVideos(t *testing.T) {
	t.Parallel()
	options := Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := Client{opts: options}
	params := PopularVideoParams{
		MinWidth:    4096,
		MinHeight:   2160,
		MinDuration: 10,
		MaxDuration: 60,
		Page:        2,
		PerPage:     2,
	}
	zero, err := c.GetPopularVideos(&params)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, VideosResponse{}) {
		t.Error("expected empty VideosResponse struct")
	}
}

func TestFailedSearchVideos(t *testing.T) {
	t.Parallel()
	options := Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := Client{opts: options}
	params := VideoSearchParams{
		Query:       "Failure",
		Locale:      UK_UA,
		Orientation: Landscape,
		Size:        Medium,
		Page:        1,
		PerPage:     1,
	}
	zero, err := c.SearchVideos(&params)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, VideosResponse{}) {
		t.Error("expected empty VideosResponse struct")
	}
}

func TestVideoIsMedia(t *testing.T) {
	t.Parallel()

	v := Video{}
	v.isMedia()
	if v.MediaType() != videoType {
		t.Errorf("Expected video media type to be %s, got %s", videoType,
			v.MediaType())
	}
}
