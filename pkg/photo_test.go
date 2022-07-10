package pexels_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	pexels "github.com/JayMonari/pexels-go/pkg"
)

var getPhotoTT = map[string]struct {
	statusCode int
	ID         uint64
	respBody   string
}{
	"Status OK": {
		http.StatusOK,
		2014422,
		`{"id":2014422,"width":3024,"height":3024,"url":"https://www.pexels.com/photo/brown-rocks-during-golden-hour-2014422/","photographer":"Joey Farina","photographer_url":"https://www.pexels.com/@joey","photographer_id":680589,"avg_color":"#978E82","src":{"original":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg","large2x":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=2\u0026h=650\u0026w=940","large":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=650\u0026w=940","medium":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=350","small":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026h=130","portrait":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=1200\u0026w=800","landscape":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026fit=crop\u0026h=627\u0026w=1200","tiny":"https://images.pexels.com/photos/2014422/pexels-photo-2014422.jpeg?auto=compress\u0026cs=tinysrgb\u0026dpr=1\u0026fit=crop\u0026h=200\u0026w=280"},"liked":false}`,
	},
	"Status Not Found": {
		http.StatusNotFound,
		0,
		`{"status":404,"error":"Not Found"}`,
	}}

func TestGetPhoto(t *testing.T) {
	t.Parallel()
	for name, tc := range getPhotoTT {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			s := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tc.respBody))
				w.WriteHeader(tc.statusCode)
			}))
			s.URL += "/video"
			s.Start()
			defer s.Close()
			c, err := pexels.New("TEST", pexels.WithPhotoBaseURL(s.URL))
			if err != nil {
				t.Fatal("expected err to be nil", err)
			}
			pr, err := c.GetPhoto(tc.ID)
			if err != nil {
				t.Fatal("expected err to be nil", err)
			}
			var want pexels.Photo
			json.Unmarshal([]byte(tc.respBody), &want)
			if got := pr.Photo; want != got {
				t.Fatalf("want: %+v\ngot: %+v", want, got)
			}
			t.Logf("%+v", pr.Photo)
			if pr.Common.StatusCode != tc.statusCode {
				t.Fatalf("want: %d\ngot: %d", tc.statusCode, pr.Common.StatusCode)
			}
		})
		// c := newMockClient(tc.opts,
		// 	newMockHandler(tc.statusCode, tc.respBody, nil))
		// aResp, err := c.GetPhoto(tc.ID)
		// if err != nil {
		// 	t.Error(err)
		// }
		// if aResp := aResp.Common.StatusCode; aResp != tc.statusCode {
		// 	t.Errorf("expected status code %d, got %d", tc.statusCode, aResp)
		// }
		// ePhoto := pexels.Photo{}
		// json.Unmarshal([]byte(tc.respBody), &ePhoto)
		// if aResp.Photo != ePhoto {
		// 	t.Errorf("expected photo %v, got %v", ePhoto, aResp.Photo)
		// }
	}
}
