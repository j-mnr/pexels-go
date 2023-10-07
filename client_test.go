package pexels_test

import (
	"net/http"
	"testing"

	"github.com/j-mnr/pexels-go"
	"github.com/matryer/is"
)

const (
	testAPIKey = "XXXX"
)

func TestNew(t *testing.T) {
	t.Parallel()

	c, err := pexels.New(testAPIKey)
	is := is.New(t)
	is.NoErr(err)
	is.Equal(http.DefaultClient, c.HTTPClient)
	is.Equal(testAPIKey, c.APIKey)
	is.Equal(pexels.PhotoBaseURL, c.PhotoBaseURL)
	is.Equal(pexels.VideoBaseURL, c.VideoBaseURL)
}
