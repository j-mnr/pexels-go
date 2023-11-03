package pexels

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var (
	ErrMissingAPIKey   = errors.New("an API Key is required")
	ErrUnsupportedType = errors.New("the type specified is cannot be unmarshalled into a Video or Photo")
)

const (
	// BaseURL is the Pexels API starting point URL for all Media.
	BaseURL = "https://api.pexels.com/"
	// RootPhotoURL is used to access both Photos and Collections.
	RootPhotoURL = BaseURL + "v1"
	// RootVideoURL is used to access Videos.
	RootVideoURL = BaseURL + "videos"

	wrapFmt = "pexels: %w"
)

// HTTPClient ...
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the Pexels API Client that allows you to interact with the Pexels
// endpoints for photos, videos, and collections.
type Client struct {
	apiKey string

	client HTTPClient

	RootPhotoURL string
	RootVideoURL string
}

// New returns a Pexels API client with the provided API key. If the API key is
// blank an error is returned.
func New(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, ErrMissingAPIKey
	}
	c := &Client{
		apiKey:    apiKey,
		client:    &http.Client{Timeout: time.Second},
		RootPhotoURL: RootPhotoURL,
		RootVideoURL: RootVideoURL,
	}
	for _, o := range opts {
		o(c)
	}
	return c, nil
}

// Option are the options you can pass in when creating a new pexels Client.
// All Option function names start with `With`.
type Option func(*Client)

func WithHTTPClient(c HTTPClient) Option {
	return func(cl *Client) { cl.client = c }
}

func get[T any](
	c Client, path string, reqData any, respData T,
) (response[T], error) {
	req, err := c.newRequest(path, reqData)
	if err != nil {
		return response[T]{}, err
	}

	res := response[T]{Data: respData}
	if err := doRequest(c, req, &res); err != nil {
		return response[T]{}, err
	}
	return res, nil
}

func doRequest[T any](c Client, req *http.Request, resp *response[T]) error {
	setRequestHeaders(req, c.apiKey)
	httpResp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf(wrapFmt, err)
	}
	defer httpResp.Body.Close()

	resp.Common.Header = httpResp.Header
	resp.Common.StatusCode = httpResp.StatusCode
	resp.Common.Status = httpResp.Status
	if err = json.NewDecoder(httpResp.Body).Decode(resp.Data); err != nil {
		return fmt.Errorf(wrapFmt, err)
	}
	return nil
}

func (c *Client) newRequest(path string, data any) (*http.Request, error) {
	url := c.RootPhotoURL + path
	if strings.HasPrefix(path, "/videos") {
		url = c.RootVideoURL + path
	}
	req, err := http.NewRequest(http.MethodGet, url, nil) //nolint:noctx
	if err != nil {
		return nil, fmt.Errorf(wrapFmt, err)
	}
	req.URL.RawQuery = buildQueryString(req, data)
	return req, nil
}

func setRequestHeaders(req *http.Request, apiKey string) {
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Accept", "application/json")
}

func buildQueryString(req *http.Request, v any) string {
	if isZero(v) {
		return ""
	}

	query := req.URL.Query()
	vType := reflect.TypeOf(v).Elem()
	vValue := reflect.ValueOf(v).Elem()
	for i := 0; i < vType.NumField(); i++ {
		var defaultValue string

		field := vType.Field(i)
		tag := field.Tag.Get("query")
		if strings.Contains(tag, ",") {
			tagSlice := strings.Split(tag, ",")
			tag = tagSlice[0]
			defaultValue = tagSlice[1]
		}

		// Add any scalar values as query params
		fieldValue := fmt.Sprintf("%v", vValue.Field(i))
		// If no value was set by the user, use the default value specified in
		// the struct tag
		if fieldValue == "" || fieldValue == "0" {
			if defaultValue == "" {
				continue
			}
			fieldValue = defaultValue
		}
		query.Add(tag, fieldValue)
	}
	return query.Encode()
}

func isZero[T comparable](v T) bool {
	var zero T
	return v == zero
}
