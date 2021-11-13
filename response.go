package pexels

import (
	"net/http"
	"strconv"
)

// ResponseCommon has the status and status code along with all the headers of
// the response. These are common fields found in every HTTP response.
type ResponseCommon struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Header     http.Header `json:"header"`
}

// GetRateLimit returns the total request limit for the monthly period.
func (rc *ResponseCommon) GetRateLimit() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Limit"))
}

// GetRateLimitRemaining returns how many requests remain.
func (rc *ResponseCommon) GetRateLimitRemaining() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Remaining"))
}

// GetRateLimitReset returns a UNIX timestamp of when the currently monthly
// period will roll over.
func (rc *ResponseCommon) GetRateLimitReset() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Reset"))
}

func (rc *ResponseCommon) convertHeaderToInt(h string) int {
	i, _ := strconv.Atoi(h)
	return i
}

// Response holds common fields of an HTTP Response and the data requested.
type Response struct {
	Common ResponseCommon
	Data   interface{}
}

func (r *Response) copyCommon(rc *ResponseCommon) {
	rc.StatusCode = r.Common.StatusCode
	rc.Header = r.Common.Header
	rc.Status = r.Common.Status
}
