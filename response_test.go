package pexels_test

import (
	"net/http"
	"testing"

	"github.com/JayMonari/go-pexels"
)

const msg = `FAIL: %s
  Got %q
  Want %q
`

func newTestResponseCommon(n, v string) pexels.ResponseCommon {
	var rc = pexels.ResponseCommon{Header: http.Header{}}
	rc.Header.Add(n, v)
	return rc
}

func TestGetRateReset(t *testing.T) {
	for _, tt := range testCases {
		rc := newTestResponseCommon("X-Ratelimit-Reset", tt.value)
		if got := rc.GetRateLimitReset(); got != tt.want {
			t.Errorf(msg, tt.description, got, tt.want)
		}
		t.Logf("PASS: %s", tt.description)
	}
}

func TestGetRateRemaining(t *testing.T) {
	for _, tt := range testCases {
		rc := newTestResponseCommon("X-Ratelimit-Remaining", tt.value)
		if got := rc.GetRateLimitRemaining(); got != tt.want {
			t.Errorf(msg, tt.description, got, tt.want)
		}
		t.Logf("PASS: %s", tt.description)
	}
}

func TestGetRateLimit(t *testing.T) {
	for _, tt := range testCases {
		rc := newTestResponseCommon("X-Ratelimit-Limit", tt.value)
		if got := rc.GetRateLimit(); got != tt.want {
			t.Errorf(msg, tt.description, got, tt.want)
		}
		t.Logf("PASS: %s", tt.description)
	}
}
