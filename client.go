package pexels

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	// BaseURL is the Pexels API starting point URL for all Media.
	BaseURL = "https://api.pexels.com/"
	// PhotoBaseURL is used to access both Photos and Collections.
	PhotoBaseURL = BaseURL + "v1"
	// VideoBaseURL is used to access Videos.
	VideoBaseURL = BaseURL + "videos"
)

// HTTPClient ...
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the Pexels API Client that allows you to interact with the Pexels
// endpoints for photos, videos, and collections.
type Client struct {
	APIKey string // required

	HTTPClient HTTPClient

	PhotoBaseURL string // Pre-set with pexels.New
	VideoBaseURL string // Pre-set with pexels.New
}

// New returns a Pexels API client. If WithAPIKey is not passed an error will
// be returned.
func New(opts ...Option) (*Client, error) {
	const baseURL = "https://api.pexels.com/"
	c := &Client{
		HTTPClient:   http.DefaultClient,
		PhotoBaseURL: PhotoBaseURL,
		VideoBaseURL: VideoBaseURL,
	}
	for _, o := range opts {
		o(c)
	}
	if c.APIKey == "" {
		return nil, errors.New("An API Key is required")
	}
	return c, nil
}

// Option are the options you can pass in when creating a new pexels Client.
// All Option function names start with `With`
type Option func(*Client)

// WithAPIKey ...
func WithAPIKey(key string) Option {
	return func(c *Client) { c.APIKey = key }
}

// WithHTTPClient ...
func WithHTTPClient(client HTTPClient) Option {
	return func(c *Client) { c.HTTPClient = client }
}

// WithPhotoBaseURL ...
func WithPhotoBaseURL(url string) Option {
	return func(c *Client) { c.PhotoBaseURL = url }
}

// WithVideoBaseURL ...
func WithVideoBaseURL(url string) Option {
	return func(c *Client) { c.VideoBaseURL = url }
}

func (c *Client) get(path string, reqData, respData interface{}) (response,
	error) {

	res := response{}
	if respData != nil {
		res.Data = respData
	}
	req, err := c.newRequest(path, reqData)
	if err != nil {
		return response{}, err
	}
	err = c.doRequest(req, &res)
	if err != nil {
		return response{}, err
	}
	return res, nil
}

func (c *Client) doRequest(req *http.Request, resp *response) error {
	c.setRequestHeaders(req)
	httpResp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	switch resp.Data.(type) {
	case *MediaPayload:
		var data struct {
			ID    string
			Media []json.RawMessage `json:"media"`
			Pagination
		}
		err = json.NewDecoder(httpResp.Body).Decode(&data)
		if err != nil {
			return err
		}
		ms := make([]Media, 0, len(data.Media))
		for _, rawMsg := range data.Media {
			m, err := decodeMediaFromRawMessage(rawMsg)
			if err != nil {
				return err
			}
			ms = append(ms, m)
		}
		resp.Data =
			MediaPayload{ID: data.ID, Media: ms, Pagination: data.Pagination}
	default:
		err = json.NewDecoder(httpResp.Body).Decode(resp.Data)
	}
	resp.Common.Header = httpResp.Header
	resp.Common.StatusCode = httpResp.StatusCode
	resp.Common.Status = httpResp.Status
	if err != nil {
		return err
	}
	return nil
}

func decodeMediaFromRawMessage(rawMsg []byte) (Media, error) {
	var typeData struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(rawMsg, &typeData); err != nil {
		return nil, err
	}
	var m Media
	switch typeData.Type {
	case videoType:
		m = &Video{}
	case photoType:
		m = &Photo{}
	default:
		return nil, errors.New("The type specified is cannot be Unmarshalled" +
			"into a Video or Photo struct")
	}
	return m, nil
}

func (c *Client) newRequest(path string, data interface{}) (*http.Request,
	error) {

	url := ""
	if strings.HasPrefix(path, "/videos") {
		url = c.VideoBaseURL + path
	} else {
		url = c.PhotoBaseURL + path
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return req, nil
	}
	query, err := buildQueryString(req, data)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query
	return req, nil
}

func (c *Client) setRequestHeaders(req *http.Request) {
	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("Accept", "application/json")
}

func buildQueryString(req *http.Request, v interface{}) (string, error) {
	isNil, err := isZero(v)
	if err != nil {
		return "", err
	} else if isNil {
		return "", nil
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
	return query.Encode(), nil
}

func isZero(v interface{}) (bool, error) {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable %v", t)
	}
	return v == reflect.Zero(t).Interface(), nil
}
