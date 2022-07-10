package pexels_test

import (
	"net/http"
	"testing"

	"github.com/JayMonari/pexels-go/pkg"
)

func TestNew(t *testing.T) {
	t.Parallel()
	for name, tc := range newTT {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var ops []pexels.Option
			if tc.httpClient != nil {
				ops = opsAppend(ops, pexels.WithHTTPClient, tc.httpClient)
			}
			if tc.photoBaseURL != "" {
				ops = opsAppend(ops, pexels.WithPhotoBaseURL, tc.photoBaseURL)
			}
			if tc.videoBaseURL != "" {
				ops = opsAppend(ops, pexels.WithVideoBaseURL, tc.videoBaseURL)
			}

			c, err := pexels.New(tc.apiKey, ops...)
			switch {
			case tc.err:
				if err == nil {
					t.Error("expecting error got nil")
				}
			case tc.apiKey != c.APIKey:
				t.Errorf("%+v want %q got %q", c, c.APIKey, tc.apiKey)
			case tc.photoBaseURL != "":
				if c.PhotoBaseURL != tc.photoBaseURL {
					t.Errorf("%+v want %q got %q", c, c.PhotoBaseURL, tc.photoBaseURL)
				}
			case tc.videoBaseURL != "":
				if c.VideoBaseURL != tc.videoBaseURL {
					t.Errorf("%+v want %q got %q", c, c.VideoBaseURL, tc.videoBaseURL)
				}
			case tc.httpClient != nil:
				if c.HTTPClient != tc.httpClient {
					t.Errorf("%+v want %+v got %+v", c, c.HTTPClient, http.DefaultClient)
				}
			}
		})
	}
}

func opsAppend[T any](opts []pexels.Option, opt func(T) pexels.Option, v any) []pexels.Option {
	switch t := v.(type) {
	case string:
		if t != "" {
			if w, ok := v.(T); ok {
				return append(opts, opt(w))
			}
		}
	case pexels.HTTPClient:
		if t != nil {
			if w, ok := v.(T); ok {
				return append(opts, opt(w))
			}
		}
	}
	return opts
}
