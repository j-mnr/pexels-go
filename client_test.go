package pexels_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JayMonari/go-pexels"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type mockHTTPClient struct {
	mockHandler http.HandlerFunc
}

func newMockClient(options *pexels.Options, mockHandler http.HandlerFunc) *pexels.Client {
	c, _ := pexels.New(pexels.Options{HTTPClient: &mockHTTPClient{mockHandler}})
	return c
}

func (mhc *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	h := http.HandlerFunc(mhc.mockHandler)
	h.ServeHTTP(rr, req)
	return rr.Result(), nil
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

func TestNew(t *testing.T) {
	t.Parallel()
	for _, tc := range clientNewTestCases {
		_, err := pexels.New(tc.options)
		if err != nil && !tc.expectErr {
			t.Errorf("Did not expect an error, got \"%s\"", err)
		} else if tc.expectErr {
			continue
		}
	}
}

type badMockHTTPClient struct {
	mockHandler http.HandlerFunc
}

func (mtc *badMockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("ERROR: badMockHTTPClient")
}

func TestFailedHTTPClientDoRequest(t *testing.T) {
	t.Parallel()
	options := pexels.Options{
		APIKey:     "testAPIKey",
		HTTPClient: &badMockHTTPClient{newMockHandler(0, "", nil)},
	}
	c, _ := pexels.New(options)

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

	c := newMockClient(&pexels.Options{APIKey: "testAPIKey"},
		newMockHandler(http.StatusOK, `data":["key":"value"]}`, nil))
	_, err := c.GetPhoto(12345)
	if err == nil {
		t.Error("expected error but got nil")
	}
}
