package pexels

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type mockHTTPClient struct {
	mockHandler http.HandlerFunc
}

func (mtc *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mtc.mockHandler)
	handler.ServeHTTP(rr, req)
	return rr.Result(), nil
}

func newMockClient(options *Options, mockHandler http.HandlerFunc) *Client {
	options.HTTPClient = &mockHTTPClient{mockHandler}
	return &Client{opts: *options}
}

func newMockHandler(statusCode int, json string, headers map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if headers != nil && len(headers) > 0 {
			for k, v := range headers {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(statusCode)
		w.Write([]byte(json))
	}
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		expectErr bool
		options   Options
	}{
		{true,
			Options{}, // No API key
		},
		{
			false,
			Options{
				APIKey:     "pexels-api-key",
				UserAgent:  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36",
				HTTPClient: &http.Client{},
			},
		},
		{
			false,
			Options{
				APIKey:    "pexels-api-key",
				UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36",
			},
		},
	}

	for _, tc := range testCases {
		client, err := NewClient(tc.options)
		if err != nil && !tc.expectErr {
			t.Errorf("Did not expect an error, got \"%s\"", err)
		} else if tc.expectErr {
			continue
		}
		opts := client.opts

		var actual, expected interface{}
		actual, expected = opts.APIKey, tc.options.APIKey
		if actual != expected {
			t.Errorf("expected %s to be \"%s\", got \"%s\"",
				"APIKey", actual, expected)
		}
		actual, expected = opts.UserAgent, tc.options.UserAgent
		if actual != expected {
			t.Errorf("expected %s to be \"%s\", got \"%s\"",
				"UserAgent", actual, expected)
		}
		actual, expected = opts.HTTPClient, tc.options.HTTPClient
		if actual != expected && (expected == &http.Client{} &&
			actual != http.DefaultClient) {
			t.Errorf("expected %s to be \"%s\", got \"%s\"",
				"HTTPClient", actual, expected)
		}
	}
}

func TestGetRateLimitHeaders(t *testing.T) {
	t.Parallel()

	tCases := []struct {
		statusCode      int
		options         *Options
		photoID         uint64
		respBody        string
		headerLimit     string
		headerRemaining string
		headerReset     string
	}{
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			2014422,
			`{"id":2014422,"width":3024,"height": 3024,"url":"https://www.pexels.com/photo/brown-rocks-during-golden-hour-2014422/"}`,
			"20000",
			"18000",
			"1625092515",
		},
		{
			http.StatusOK,
			&Options{APIKey: "testAPIKey"},
			2014422,
			`{"id":2014422,"width":3024,"height": 3024,"url":"https://www.pexels.com/photo/brown-rocks-during-golden-hour-2014422/"}`,
			"",
			"",
			"",
		},
	}
	for _, tc := range tCases {
		mockRespHeaders := make(map[string]string)

		if tc.headerLimit != "" {
			mockRespHeaders["X-Ratelimit-Limit"] = tc.headerLimit
			mockRespHeaders["X-Ratelimit-Remaining"] = tc.headerRemaining
			mockRespHeaders["X-Ratelimit-Reset"] = tc.headerReset
		}
		c := newMockClient(tc.options,
			newMockHandler(tc.statusCode, tc.respBody, mockRespHeaders))

		resp, err := c.GetPhoto(tc.photoID)
		if err != nil {
			t.Error(err)
		}

		expected, _ := strconv.Atoi(tc.headerLimit)
		if actual := resp.Common.GetRateLimit(); actual != expected {
			t.Errorf("expected \"X-Ratelimit-Limit\" to be \"%d\", got \"%d\"",
				expected, actual)
		}
		expected, _ = strconv.Atoi(tc.headerRemaining)
		if actual := resp.Common.GetRateLimitRemaining(); actual != expected {
			t.Errorf("expected \"X-Ratelimit-Remaining\" to be \"%d\", got \"%d\"",
				expected, actual)
		}
		expected, _ = strconv.Atoi(tc.headerReset)
		if actual := resp.Common.GetRateLimitReset(); actual != expected {
			t.Errorf("expected \"X-Ratelimit-Reset\" to be \"%d\", got \"%d\"",
				expected, actual)
		}
	}
}

func TestSetRequestHeaders(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		endpoint  string
		APIKey    string
		UserAgent string
	}{
		{"/photos/", "testAPIKey", "testUserAgent"},
		{"/videos/", "testAPIKey", ""},
		{"/collections/", "testAPIKey", "otherUserAgent"},
	}

	for _, tc := range tcs {
		client, err := NewClient(Options{
			APIKey:    tc.APIKey,
			UserAgent: tc.UserAgent,
		})
		if err != nil {
			t.Errorf("Did not expect an error, got \"%s\"", err.Error())
		}

		req, _ := http.NewRequest("GET", tc.endpoint, nil)
		client.setRequestHeaders(req)
		expected := client.opts.UserAgent
		if actual := req.Header.Get("User-Agent"); actual != expected {
			t.Errorf("expected \"User-Agent\" to be \"%s\", got \"%s\"",
				expected, actual)
		}
		expected = "application/json"
		if actual := req.Header.Get("Accept"); actual != expected {
			t.Errorf("expected \"Accept\" to be \"%s\", got \"%s\"",
				expected, actual)
		}
		expected = client.opts.APIKey
		if actual := req.Header.Get("Authorization"); actual != expected {
			t.Errorf("expected \"Authorization\" to be \"%s\", got \"%s\"",
				expected, actual)
		}
	}
}

func TestSetUserAgent(t *testing.T) {
	t.Parallel()

	c, err := NewClient(Options{APIKey: "testAPIKey"})
	if err != nil {
		t.Errorf("Did not expect an error, got \"%s\"", err.Error())
	}
	ua := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
	c.SetUserAgent(ua)
	if c.opts.UserAgent != ua {
		t.Errorf("expected accessToken to be \"%s\", got \"%s\"", ua,
			c.opts.UserAgent)
	}
}

func TestCopyCommon(t *testing.T) {
	t.Parallel()
	var sourceResp Response
	testStatusCode := 200
	testHeaders := http.Header{}
	testHeaders.Set("Content-Type", "application/json")
	testHeaders.Set("Accept", "application/json")
	sourceResp.Common.StatusCode = testStatusCode
	sourceResp.Common.Header = testHeaders

	var targetResp Response
	sourceResp.copyCommon(&targetResp.Common)
	if targetResp.Common.StatusCode != testStatusCode {
		t.Errorf("expected status code to be \"%d\", got \"%d\"", testStatusCode,
			targetResp.Common.StatusCode)
	}

	if !cmp.Equal(targetResp.Common.Header, sourceResp.Common.Header) {
		t.Errorf("expected headers to match")
	}
}

type badMockHTTPClient struct {
	mockHandler http.HandlerFunc
}

func (mtc *badMockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("BIG OOF dawg, looks like we found an error")
}

func TestFailedHTTPClientDoRequest(t *testing.T) {
	t.Parallel()
	options := Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c := Client{
		opts: options,
	}

	_, err := c.GetPhoto(12345)
	if err == nil {
		t.Error("expected error but got nil")
	}
	if err.Error() != "BIG OOF dawg, looks like we found an error" {
		t.Error("expected error does match return error")
	}
}

func TestDecodingBadJSON(t *testing.T) {
	t.Parallel()

	c := newMockClient(&Options{APIKey: "testAPIKey"},
		newMockHandler(http.StatusOK, `data":["key":"value"]}`, nil))
	_, err := c.GetPhoto(12345)
	if err == nil {
		t.Error("expected error but got nil")
	}
}

func TestBadNewRequest(t *testing.T) {
	c := newMockClient(&Options{}, newMockHandler(http.StatusTeapot, "", nil))
	data, err := c.newRequest("", nil)
	mockReq, _ := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Errorf("expected error but got nil and %v", data)
	}
	if !cmp.Equal(*data, *mockReq, cmpopts.IgnoreUnexported(http.Request{})) {
		t.Errorf("expected req to be %v, got %v", mockReq, data)
	}
	_, err = c.newRequest("/videos/videos/1234", PhotosResponse{})
	if err == nil {
		t.Error("expected error but got nil")
	}
}

func TestBadBuildQueryString(t *testing.T) {
	mockReq := &http.Request{}
	_, err := buildQueryString(mockReq, PhotosResponse{})
	if err == nil {
		t.Error("expected error but got nil")
	}
}

func TestBadIsZero(t *testing.T) {
	t.Parallel()

	eFalse, err := isZero(PhotosResponse{})
	if err == nil {
		t.Error("expected error but got nil")
	} else if eFalse != false {
		t.Error("expected false but got true?????????")
	}
}

func TestBadDecodeMediaFromRawMessage(t *testing.T) {
	t.Parallel()

	_, err := decodeMediaFromRawMessage([]byte("{\"data\": { \"notype\": []}}"))
	if err == nil {
		t.Error("expected error but got nil")
	}
	_, err = decodeMediaFromRawMessage([]byte("{\"type\": \"Collection\"}"))
	if err == nil {
		t.Error("expected error but got nil")
	}
}
