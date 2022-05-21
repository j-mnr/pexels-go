package pexels_test

import (
	"net/http"
	"testing"

	"github.com/JayMonari/pexels-go"
)

func TestNew(t *testing.T) {
	k := "REAL-KEY"
	c, err := pexels.New(pexels.WithAPIKey(k))
	switch {
	case err != nil:
		t.Errorf("Was not expecting error got %q", err)
	case c.APIKey != k:
		t.Errorf("%+v want %q got %q", c, c.APIKey, k)
	case c.PhotoBaseURL != pexels.PhotoBaseURL:
		t.Errorf("%+v want %q got %q", c, c.PhotoBaseURL, pexels.PhotoBaseURL)
	case c.VideoBaseURL != pexels.VideoBaseURL:
		t.Errorf("%+v want %q got %q", c, c.VideoBaseURL, pexels.VideoBaseURL)
	case c.HTTPClient != http.DefaultClient:
		t.Errorf("%+v want %+v got %+v", c, c.HTTPClient, http.DefaultClient)
	}
}
