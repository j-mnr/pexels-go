package pexels_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JayMonari/go-pexels"
)

type mockHTTPClient struct {
	mockHandler http.HandlerFunc
}

func newMockHTTPClient(options *pexels.Options, mockHandler http.HandlerFunc) *pexels.Client {
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
