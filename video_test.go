package pexels_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/JayMonari/go-pexels"
	"github.com/google/go-cmp/cmp"
)

func TestGetVideo(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		statusCode int
		opts       pexels.Options
		ID         int
		respBody   string
	}{
		{
			http.StatusOK,
			pexels.Options{APIKey: "testAPIKey"},
			2499611,
			`{"id":2499611,"width":1080,"height":1920,"url":"https://www.pexels.com/video/2499611/","image":"https://images.pexels.com/videos/2499611/free-video-2499611.jpg?fit=crop&w=1200&h=630&auto=compress&cs=tinysrgb","duration":22,"user":{"id":680589,"name":"JoeyFarina","url":"https://www.pexels.com/@joey"},"video_files":[{"id":125004,"quality":"hd","file_type":"video/mp4","width":1080,"height":1920,"link":"https://player.vimeo.com/external/342571552.hd.mp4?s=6aa6f164de3812abadff3dde86d19f7a074a8a66&profile_id=175&oauth2_token_id=57447761"},{"id":125005,"quality":"sd","file_type":"video/mp4","width":540,"height":960,"link":"https://player.vimeo.com/external/342571552.sd.mp4?s=e0df43853c25598dfd0ec4d3f413bce1e002deef&profile_id=165&oauth2_token_id=57447761"},{"id":125006,"quality":"sd","file_type":"video/mp4","width":240,"height":426,"link":"https://player.vimeo.com/external/342571552.sd.mp4?s=e0df43853c25598dfd0ec4d3f413bce1e002deef&profile_id=139&oauth2_token_id=57447761"},{"id":125007,"quality":"hd","file_type":"video/mp4","width":720,"height":1280,"link":"https://player.vimeo.com/external/342571552.hd.mp4?s=6aa6f164de3812abadff3dde86d19f7a074a8a66&profile_id=174&oauth2_token_id=57447761"},{"id":125008,"quality":"sd","file_type":"video/mp4","width":360,"height":640,"link":"https://player.vimeo.com/external/342571552.sd.mp4?s=e0df43853c25598dfd0ec4d3f413bce1e002deef&profile_id=164&oauth2_token_id=57447761"},{"id":125009,"quality":"hls","file_type":"video/mp4","width":null,"height":null,"link":"https://player.vimeo.com/external/342571552.m3u8?s=53433233e4176eead03ddd6fea04d9fb2bce6637&oauth2_token_id=57447761"}],"video_pictures":[{"id":308178,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-0.jpg","nr":0},{"id":308179,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-1.jpg","nr":1},{"id":308180,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-2.jpg","nr":2},{"id":308181,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-3.jpg","nr":3},{"id":308182,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-4.jpg","nr":4},{"id":308183,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-5.jpg","nr":5},{"id":308184,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-6.jpg","nr":6},{"id":308185,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-7.jpg","nr":7},{"id":308186,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-8.jpg","nr":8},{"id":308187,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-9.jpg","nr":9},{"id":308188,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-10.jpg","nr":10},{"id":308189,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-11.jpg","nr":11},{"id":308190,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-12.jpg","nr":12},{"id":308191,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-13.jpg","nr":13},{"id":308192,"picture":"https://static-videos.pexels.com/videos/2499611/pictures/preview-14.jpg","nr":14}]}`,
		},
		{
			http.StatusNotFound,
			pexels.Options{APIKey: "testAPIKey"},
			0,
			`{"status":404,"error":"Not Found"}`,
		},
	}
	for _, tc := range tcs {
		c := newMockClient(tc.opts,
			newMockHandler(tc.statusCode, tc.respBody, nil))
		aResp, err := c.GetVideo(tc.ID)
		if err != nil {
			t.Error(err)
		}
		if aResp := aResp.Common.StatusCode; aResp != tc.statusCode {
			t.Errorf("expected status code %d, got %d", tc.statusCode, aResp)
		}
		eVideo := pexels.Video{}
		json.Unmarshal([]byte(tc.respBody), &eVideo)
		if !cmp.Equal(aResp.Video, eVideo) {
			t.Errorf("expected video %v, got %v", eVideo, aResp.Video)
		}
	}
}

func TestGetPopularVideos(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		statusCode int
		opts       pexels.Options
		params     *pexels.PopularVideoParams
		respBody   string
	}{
		{
			http.StatusOK,
			pexels.Options{APIKey: "testAPIKey"},
			&pexels.PopularVideoParams{
				MinWidth:    4096,
				MinHeight:   2160,
				MinDuration: 10,
				MaxDuration: 60,
				Page:        2,
				PerPage:     2,
			},
			`{"videos":[{"id":2953633,"width":4096,"height":2160,"url":"https://www.pexels.com/video/people-having-rest-in-a-ledge-2953633/","image":"https://images.pexels.com/videos/2953633/free-video-2953633.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":27,"user":{"id":1179532,"name":"KellyLacy","url":"https://www.pexels.com/@kelly-lacy-1179532"},"video_files":[{"id":187214,"quality":"sd","file_type":"video/mp4","width":640,"height":338,"link":"https://player.vimeo.com/external/360602969.sd.mp4?s=a77f0cd29f80c90a26ade4c5c8cdbf374287b28a\u0026profile_id=164\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":444963,"picture":"https://images.pexels.com/videos/2953633/pictures/preview-0.jpg","nr":0}]},{"id":1448735,"width":4096,"height":2160,"url":"https://www.pexels.com/video/video-of-forest-1448735/","image":"https://images.pexels.com/videos/1448735/free-video-1448735.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":32,"user":{"id":574687,"name":"RuvimMiksanskiy","url":"https://www.pexels.com/@digitech"},"video_files":[{"id":58649,"quality":"sd","file_type":"video/mp4","width":640,"height":338,"link":"https://player.vimeo.com/external/291648067.sd.mp4?s=7f9ee1f8ec1e5376027e4a6d1d05d5738b2fbb29\u0026profile_id=164\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":133236,"picture":"https://images.pexels.com/videos/1448735/pictures/preview-0.jpg","nr":0}]}],"total_results":33838,"page":2,"per_page":2,"prev_page":"https://api.pexels.com/v1/videos/popular/?max_duration=60\u0026min_duration=10\u0026min_height=2160\u0026min_width=4096\u0026page=1\u0026per_page=2","next_page":"https://api.pexels.com/v1/videos/popular/?max_duration=60\u0026min_duration=10\u0026min_height=2160\u0026min_width=4096\u0026page=3\u0026per_page=2"}`,
		},
	}
	for _, tc := range tcs {
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

		ePayload := pexels.VideoPayload{}
		json.Unmarshal([]byte(tc.respBody), &ePayload)
		if !cmp.Equal(aResp.Payload, ePayload) {
			t.Errorf("expected payload %v, got %v", ePayload, aResp.Payload)
		}
	}
}

func TestSearchVideos(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		statusCode int
		opts       pexels.Options
		params     *pexels.SearchVideoParams
		respBody   string
	}{
		{
			http.StatusOK,
			pexels.Options{APIKey: "testAPIKey"},
			&pexels.SearchVideoParams{Query: ""},
			``,
		},
		{
			http.StatusForbidden,
			pexels.Options{APIKey: "invalid-APIKey"},
			&pexels.SearchVideoParams{Query: "City Lights"},
			`{"error": "Access to this API has been disallowed"}`,
		},
		{
			http.StatusOK,
			pexels.Options{APIKey: "testAPIKey"},
			&pexels.SearchVideoParams{Query: "naasdfasdfasdfasdftasldkjfasdlkjfure"},
			`{"page":1,"per_page":1,"videos":[],"total_results":0,"url":"https://api-server.pexels.com/search/videos/naasdfasdfasdfasdftasldkjfasdlkjfure/"}18`,
		},
		{
			http.StatusOK,
			pexels.Options{APIKey: "testAPIKey"},
			&pexels.SearchVideoParams{
				Query:       "dogs",
				Locale:      pexels.UK_UA,
				Orientation: pexels.Landscape,
				Size:        pexels.Medium,
				Page:        4,
				PerPage:     3,
			},
			`{"videos":[{"id":3077158,"width":3840,"height":2160,"url":"https://www.pexels.com/uk-ua/video/3077158/","image":"https://images.pexels.com/videos/3077158/free-video-3077158.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":84,"user":{"id":102775,"name":"MagdaEhlers","url":"https://www.pexels.com/uk-ua/@magda-ehlers-pexels"},"video_files":[{"id":354864,"quality":"sd","file_type":"video/mp4","width":960,"height":540,"link":"https://player.vimeo.com/external/366179532.sd.mp4?s=fcfddee8ce5c72bcd9c6f3097ec48d530a0b8600\u0026profile_id=165\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":497782,"picture":"https://images.pexels.com/videos/3077158/pictures/preview-0.jpg","nr":0}]},{"id":3191251,"width":4096,"height":2160,"url":"https://www.pexels.com/uk-ua/video/3191251/","image":"https://images.pexels.com/videos/3191251/free-video-3191251.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":13,"user":{"id":1583460,"name":"Pressmaster","url":"https://www.pexels.com/uk-ua/@pressmaster"},"video_files":[{"id":253935,"quality":"sd","file_type":"video/mp4","width":960,"height":506,"link":"https://player.vimeo.com/external/371562597.sd.mp4?s=2d39d964e8af2dae0b0a0b303c04ea300c9be7fc\u0026profile_id=165\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":572772,"picture":"https://images.pexels.com/videos/3191251/pictures/preview-0.jpg","nr":0}]},{"id":2795691,"width":3840,"height":2160,"url":"https://www.pexels.com/uk-ua/video/2795691/","image":"https://images.pexels.com/videos/2795691/free-video-2795691.jpg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=630\u0026w=1200","duration":6,"user":{"id":1007758,"name":"EvgeniaKirpichnikova","url":"https://www.pexels.com/uk-ua/@evgenia-kirpichnikova-1007758"},"video_files":[{"id":157832,"quality":"hd","file_type":"video/mp4","width":1280,"height":720,"link":"https://player.vimeo.com/external/353554745.hd.mp4?s=e7aa9bc30aaa1618d0f1f076bec2ce45db61300d\u0026profile_id=174\u0026oauth2_token_id=57447761"}],"video_pictures":[{"id":385549,"picture":"https://images.pexels.com/videos/2795691/pictures/preview-0.jpg","nr":0}]}],"total_results":2879,"page":4,"per_page":3,"prev_page":"https://api.pexels.com/v1/videos/search/?locale=uk-UA\u0026orientation=landscape\u0026page=3\u0026per_page=3\u0026query=dogs\u0026size=medium","next_page":"https://api.pexels.com/v1/videos/search/?locale=uk-UA\u0026orientation=landscape\u0026page=5\u0026per_page=3\u0026query=dogs\u0026size=medium"}`,
		},
	}

	for _, tc := range tcs {
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
		ePayload := pexels.VideoPayload{}
		// We need tpexels.his because the API defaults to page 1 and per_page 1
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

	options := pexels.Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c, _ := pexels.New(options)
	zero, err := c.GetVideo(999999999999999)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, pexels.VideoResponse{}) {
		t.Error("expected empty pexels.VideoResponse struct")
	}
}

func TestFailedPopularVideos(t *testing.T) {
	t.Parallel()

	options := pexels.Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}

	c, _ := pexels.New(options)
	params := pexels.PopularVideoParams{
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
	} else if !cmp.Equal(zero, pexels.VideosResponse{}) {
		t.Error("expected empty pexels.VideosResponse struct")
	}
}

func TestFailedSearchVideos(t *testing.T) {
	t.Parallel()

	options := pexels.Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c, _ := pexels.New(options)
	params := pexels.SearchVideoParams{
		Query:       "Failure",
		Locale:      pexels.UK_UA,
		Orientation: pexels.Landscape,
		Size:        pexels.Medium,
		Page:        1,
		PerPage:     1,
	}

	zero, err := c.SearchVideos(&params)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	} else if !cmp.Equal(zero, pexels.VideosResponse{}) {
		t.Error("expected empty pexels.VideosResponse struct")
	}
}
