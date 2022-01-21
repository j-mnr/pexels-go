package pexels

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetPhoto(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		statusCode int
		opts       *Options
		ID         uint64
		respBody   string
	}{
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			2014422,
			`{"id":2014422,"width":3024,"height":3024,"url":"https://www.pexels.com/photo/brown-rocks-during-golden-hour-2014422/","photographer":"Joey Farina","photographer_url":"https://www.pexels.com/@joey","photographer_id":680589,"avg_color":"#978E82","src":{"original":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg","large2x":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940","large":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940","medium":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350","small":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130","portrait":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800","landscape":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200","tiny":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"},"liked":false}`,
		},
		{
			http.StatusNotFound,
			&Options{APIKey: "testAPIKey"},
			0,
			`{"status":404,"error":"Not Found"}`,
		},
	}
	for _, tc := range tcs {
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

	tcs := []struct {
		statusCode int
		opts       *Options
		params     *CuratedPhotosParams
		respBody   string
	}{
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			&CuratedPhotosParams{Page: 2, PerPage: 5},
			`{"page":2,"per_page":5,"photos":[{"id":8536867,"width":4291,"height":5364,"url":"https://www.pexels.com/photo/silhouette-of-man-raising-his-hands-8536867/","photographer":"Trarete","photographer_url":"https://www.pexels.com/@trarete-73723862","photographer_id":73723862,"avg_color":"#405580","src":{"original":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg","large2x":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940","large":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940","medium":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350","small":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130","portrait":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800","landscape":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200","tiny":"https://images.pexels.com/photos/8536867/pexels-photo-8536867.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"},"liked":false}],"total_results":8000,"next_page":"https://api.pexels.com/v1/curated/?page=3\u0026per_page=5","prev_page":"https://api.pexels.com/v1/curated/?page=1\u0026per_page=5"}`,
		},
	}
	for _, tc := range tcs {
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

	tcs := []struct {
		statusCode int
		opts       *Options
		params     *PhotoSearchParams
		respBody   string
	}{
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			&PhotoSearchParams{Query: ""},
			``,
		},
		{
			http.StatusForbidden,
			&Options{APIKey: "invalid-APIKey"},
			&PhotoSearchParams{Query: "City Lights"},
			`{"error": "Access to this API has been disallowed"}`,
		},
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			&PhotoSearchParams{Query: "random-garbagepaosdjfpoijqwpoeijo!@#$%:;(*)"},
			`{"page":1,"per_page":1,"photos":[],"total_results":0,"url":"https://api-server.pexels.com/search/videos/random-garbagepaosdjfpoijqwpoeijo!@#$%:;(*)/"}`,
		},
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			&PhotoSearchParams{Query: "nature",
				Locale:      UK_UA,
				Orientation: Landscape,
				Size:        Medium,
				Color:       Red,
				Page:        1,
				PerPage:     1,
			},
			`{"page":1,"per_page":1,"photos":[{"id":130988,"width":6000,"height":3376,"url":"https://www.pexels.com/photo/red-flowers-130988/","photographer":"Mike","photographer_url":"https://www.pexels.com/@mikebirdy","photographer_id":20649,"avg_color":"#82251B","src":{"original":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg","large2x":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940","large":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940","medium":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350","small":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130","portrait":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800","landscape":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200","tiny":"https://images.pexels.com/photos/130988/pexels-photo-130988.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"},"liked":false}],"total_results":8000,"next_page":"https://api.pexels.com/v1/search/?color=red\u0026page=2\u0026per_page=1\u0026query=nature\u0026size=medium"}`,
		},
	}

	for _, tc := range tcs {
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
	c := Client{
		opts: options,
	}
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
	c := Client{
		opts: options,
	}
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
