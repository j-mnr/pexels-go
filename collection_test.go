package pexels_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/JayMonari/go-pexels"
	"github.com/google/go-cmp/cmp"
)

func TestGetCollection(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		statusCode int
		opts       *pexels.Options
		params     *pexels.CollectionMediaParams
		respBody   string
	}{
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			&CollectionMediaParams{
				ID:      "g3aedup",
				Type:    VideoType,
				Page:    1,
				PerPage: 3,
			},
			`{"id":"g3aedup","Videos":[{"id":5409152,"width":1080,"height":1920,"url":"https://www.pexels.com/video/learning-spanish-simple-minimalism-5409152/","image":"https://images.pexels.com/videos/5409152/foreign-language-home-study-home-work-learning-5409152.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":15,"user":{"id":2946790,"name":"OlyaKobruseva","url":"https://www.pexels.com/@olyakobruseva"},"video_files":[{"id":1373504,"quality":"sd","file_type":"video/mp4","width":360,"height":640,"link":"https://player.vimeo.com/external/460194477.sd.mp4?s=96856e4d6f24eaa7b075ddf9d1bbbe21dc26b7da\u0026profile_id=164\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":2786191,"picture":"https://images.pexels.com/videos/5409152/pictures/preview-0.jpg","nr":0}],"type":"Video"},{"id":2759451,"width":1920,"height":1080,"url":"https://www.pexels.com/video/an-old-video-footage-of-people-enjoying-their-day-in-a-beach-2759451/","image":"https://images.pexels.com/videos/2759451/free-video-2759451.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":9,"user":{"id":1092752,"name":"LifeOnSuper8","url":"https://www.pexels.com/@life-on-super-8-1092752"},"video_files":[{"id":151386,"quality":"hd","file_type":"video/mp4","width":1280,"height":720,"link":"https://player.vimeo.com/external/352010608.hd.mp4?s=16e946fe241f74b03c68b7c73fc44d5e78ec0b10\u0026profile_id=174\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":370599,"picture":"https://images.pexels.com/videos/2759451/pictures/preview-0.jpg","nr":0}],"type":"Video"},{"id":2034302,"width":3840,"height":2160,"url":"https://www.pexels.com/video/an-old-phonograph-2034302/","image":"https://images.pexels.com/videos/2034302/free-video-2034302.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":19,"user":{"id":1054894,"name":"JakubGorajek","url":"https://www.pexels.com/@jakub-gorajek-1054894"},"video_files":[{"id":84782,"quality":"hd","file_type":"video/mp4","width":3840,"height":2160,"link":"https://player.vimeo.com/external/325039010.hd.mp4?s=0b149cb826b5811e593007d992715ebe0649290c\u0026profile_id=172\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":208651,"picture":"https://images.pexels.com/videos/2034302/pictures/preview-0.jpg","nr":0}],"type":"Video"}],"Photos":null,"total_results":3,"page":1,"per_page":15,"prev_page":"","next_page":""}`,
		},
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			&CollectionMediaParams{
				ID:      "c10vohh",
				Type:    PhotoType,
				Page:    1,
				PerPage: 3,
			},
			`{"page":1,"per_page":15,"media":[{"type":"Photo","id":5391353,"width":3456,"height":3456,"url":"https://www.pexels.com/photo/close-up-photography-of-green-plant-5391353/","photographer":"Antonio Prado","photographer_url":"https://www.pexels.com/@antonio-prado-1050855","photographer_id":1050855,"avg_color":"#373F38","src":{},"liked":false},{"type":"Photo","id":7120688,"width":3238,"height":4046,"url":"https://www.pexels.com/photo/woman-in-white-and-red-floral-shirt-standing-beside-white-flowers-7120688/","photographer":"TRAVELBLOG","photographer_url":"https://www.pexels.com/@travelblog-26954066","photographer_id":26954066,"avg_color":"#959486","src":{},"liked":false},{"type":"Photo","id":8043220,"width":2688,"height":4032,"url":"https://www.pexels.com/photo/woman-in-black-spaghetti-strap-top-8043220/","photographer":"Eman Genatilan","photographer_url":"https://www.pexels.com/@eman-genatilan-3459781","photographer_id":3459781,"avg_color":"#707360","src":{},"liked":false}],"total_results":3,"id":"c10vohh"}`,
		},
		{
			http.StatusNotFound,
			&pexels.Options{APIKey: "testAPIKey"},
			&pexels.CollectionMediaParams{
				ID: "",
			},
			`{"status":404,"error":"Not Found"}`,
		},
	}
	for _, tc := range tcs {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetCollection(tc.params)
		if err != nil && err.Error() == "Collection ID must be specified" {
			continue
		}

		if aResp := aResp.Common.StatusCode; aResp != tc.statusCode {
			t.Errorf("expected status code %d, got %d", tc.statusCode, aResp)
		}
		if cmp.Equal(aResp.Videos, []pexels.Video{}) {
			vids := []pexels.Video{}
			json.Unmarshal([]byte(tc.respBody), &vids)
			if !cmp.Equal(aResp.Videos, vids) {
				t.Errorf("expected videos: %v, got %v", vids, aResp.Videos)
			}
		} else if cmp.Equal(aResp.Photos, []pexels.Photo{}) {
			photos := []pexels.Photo{}
			json.Unmarshal([]byte(tc.respBody), &photos)
			if !cmp.Equal(aResp.Photos, photos) {
				t.Errorf("expected photos: %v, got %v", photos, aResp.Videos)
			}
		}
	}
}

func TestGetCollections(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		statusCode int
		opts       *pexels.Options
		respBody   string
	}{
		{
			http.StatusOK,
			&pexels.Options{APIKey: "testAPIKey"},
			`{"page":1,"per_page":1,"collections":[{"id":"c10vohh","title":"Test API Collection","description":null,"private":false,"media_count":3,"photos_count":3,"videos_count":0}],"total_results":1}`,
		},
		{
			http.StatusForbidden,
			&pexels.Options{APIKey: "invalid-APIKey"},
			`{"error": "Access to this API has been disallowed"}`,
		},
	}
	for _, tc := range tcs {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetCollections()
		if err != nil {
			t.Errorf("Did not expect an error, got \"%s\"", err)
		}
		eCollections := pexels.CollectionPayload{}
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

	options := pexels.Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := pexels.Client{
		opts: options,
	}
	params := pexels.CollectionMediaParams{
		ID:      "WONTWORK",
		Type:    pexels.VideoType,
		Page:    22,
		PerPage: 80,
	}
	zero, err := c.GetCollection(&params)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, pexels.MediaResponse{}) {
		t.Error("expected empty MediaResponse struct")
	}
}

func TestFailedGetCollections(t *testing.T) {
	t.Parallel()

	options := pexels.Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c, _ := pexels.New(options)

	zero, err := c.GetCollections()
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, pexels.CollectionsResponse{}) {
		t.Error("expected empty CollectionsResponse struct")
	}
}
