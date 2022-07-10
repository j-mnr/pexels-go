package pexels_test

import (
	"net/http"

	"github.com/JayMonari/pexels-go/pkg"
)

type testHTTPClient struct{}

func (c testHTTPClient) Do(*http.Request) (*http.Response, error) {
	return nil, nil
}

var newTT = map[string]struct {
	videoBaseURL string
	photoBaseURL string
	apiKey       string
	httpClient   pexels.HTTPClient
	err          bool
}{
	"error when no APIKey": {
		err: true,
	},
	"default values": {
		apiKey: "test-key",
	},
	"all values set correctly": {
		videoBaseURL: "https://videourl.com",
		photoBaseURL: "https://photourl.com",
		apiKey:       "test-key",
		httpClient:   testHTTPClient{},
	},
}

var rcTC = pexels.ResponseCommon{
	StatusCode: 200,
	Status:     "OK",
	Header: map[string][]string{
		"X-Ratelimit-Limit":     {"1000"},
		"X-Ratelimit-Remaining": {"1000"},
		"X-Ratelimit-Reset":     {"1000"},
	},
}
