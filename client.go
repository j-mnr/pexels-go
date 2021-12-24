package pexels

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

const APIBaseURL = "https://api.pexels.com"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Options struct {
	APIKey     string
	UserAgent  string
	HTTPClient HTTPClient
	baseURL    string
}

type Client struct {
	mu   sync.RWMutex
	opts Options
}

// New creates a Pexels API client. An API key must be provided or else an
// error will be returned and if new HTTPClient is provided the
// http.DefaultClient will be used.
func New(options Options) (*Client, error) {
	if options.APIKey == "" {
		return nil, errors.New("an API key is required to access the Pexels API")
	} else if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}
	options.baseURL = APIBaseURL
	return &Client{opts: options}, nil
}

func (c *Client) get(path string, reqData, respData interface{}) (Response, error) {
	resp := Response{}
	if respData != nil {
		resp.Data = respData
	}

	req, err := c.newRequest(path, reqData)
	if err != nil {
		return Response{}, err
	}

	err = c.doRequest(req, &resp)
	if err != nil {
		return Response{}, err
	}

	return resp, nil
}

func (c *Client) newRequest(path string, data interface{}) (*http.Request, error) {
	url := c.opts.baseURL + path

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	} else if data == nil {
		return req, nil
	}

	query, err := buildQueryString(req, data)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query
	return req, nil
}

func (c *Client) doRequest(req *http.Request, resp *Response) error {
	c.setRequestHeaders(req)
	httpResp, err := c.opts.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	err = json.NewDecoder(httpResp.Body).Decode(resp.Data)
	if err != nil {
		return err
	}

	resp.Common.Header = httpResp.Header
	resp.Common.StatusCode = httpResp.StatusCode
	resp.Common.Status = httpResp.Status
	return nil
}

func (c *Client) setRequestHeaders(req *http.Request) {
	opts := c.opts
	req.Header.Set("Authorization", opts.APIKey)
	req.Header.Set("Accept", "application/json")
	if opts.UserAgent != "" {
		req.Header.Set("User-Agent", opts.UserAgent)
	}
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
	vVal := reflect.ValueOf(v).Elem()
	for i := 0; i < vType.NumField(); i++ {
		defaultVal := ""

		tag := vType.Field(i).Tag.Get("query")
		if sl := strings.Split(tag, ","); len(sl) == 2 {
			tag, defaultVal = sl[0], sl[1]
		}

		fieldVal := fmt.Sprintf("%v", vVal.Field(i))
		if fieldVal == "" || fieldVal == "0" {
			if defaultVal == "" {
				continue
			}
			fieldVal = defaultVal
		}

		query.Add(tag, fieldVal)
	}

	return query.Encode(), nil
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
		return nil, errors.New("unknown type: " + typeData.Type)
	}
	return m, nil
}

func isZero(v interface{}) (bool, error) {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable %v", t)
	}
	return v == reflect.Zero(t).Interface(), nil
}
