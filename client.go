package pexels

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

const (
	// Base URL for both Photos and Collections use VideoBaseURL if you need
	// videos
	PhotoBaseURL = "https://api.pexels.com/v1"
	VideoBaseURL = "https://api.pexels.com/videos"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	mu   sync.RWMutex
	opts Options
}

type Options struct {
	APIKey        string
	UserAgent     string
	HTTPClient    HTTPClient
	photosBaseURL string
	videoBaseURL  string
}

type ResponseCommon struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Header     http.Header `json:"headers"`
}

func (rc *ResponseCommon) convertHeaderToInt(h string) int {
	i, _ := strconv.Atoi(h)
	return i
}

func (rc *ResponseCommon) GetRateLimit() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Limit"))
}

func (rc *ResponseCommon) GetRateLimitRemaining() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Remaining"))
}

// Returns a UNIX timestamp of when the current monthly period will roll over
func (rc *ResponseCommon) GetRateLimitReset() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Reset"))
}

type Response struct {
	Common ResponseCommon
	Data   interface{}
}

func (r *Response) copyCommon(rc *ResponseCommon) {
	rc.StatusCode = r.Common.StatusCode
	rc.Header = r.Common.Header
	rc.Status = r.Common.Status
}

// Returns a new Pexels API client. If the Options provided does not contain an
// API Key an error will be returned
func NewClient(options Options) (*Client, error) {
	if options.APIKey == "" {
		return nil, errors.New("An API Key is required")
	}
	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}
	options.videoBaseURL = "https://api.pexels.com/"
	options.photosBaseURL = PhotoBaseURL
	client := &Client{
		opts: options,
	}
	return client, nil
}

func (c *Client) get(path string, reqData, respData interface{}) (Response,
	error) {

	response := Response{}
	if respData != nil {
		response.Data = respData
	}
	req, err := c.newRequest(path, reqData)
	if err != nil {
		return Response{}, err
	}
	err = c.doRequest(req, &response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func (c *Client) doRequest(req *http.Request, resp *Response) error {
	c.setRequestHeaders(req)
	httpResp, err := c.opts.HTTPClient.Do(req)
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
		url = c.opts.videoBaseURL + path
	} else {
		url = c.opts.photosBaseURL + path
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	opts := c.opts

	req.Header.Set("Authorization", opts.APIKey)
	req.Header.Set("Accept", "application/json")
	if opts.UserAgent != "" {
		req.Header.Set("User-Agent", opts.UserAgent)
	}
}

func (c *Client) SetUserAgent(userAgent string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.opts.UserAgent = userAgent
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
