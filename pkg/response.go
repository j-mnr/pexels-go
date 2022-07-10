package pexels

import (
	"net/http"
	"strconv"
)

// ResponseCommon holds the common values found inside of a HTTP response.
type ResponseCommon struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Header     http.Header `json:"headers"`
}

func (rc *ResponseCommon) convertHeaderToInt(h string) int {
	i, _ := strconv.Atoi(h)
	return i
}

// GetRateLimit returns the total request limit for the monthly period.
func (rc *ResponseCommon) GetRateLimit() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Limit"))
}

// GetRateLimitRemaining returns how many requests you have left that you can
// make for the monthly period.
func (rc *ResponseCommon) GetRateLimitRemaining() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Remaining"))
}

// GetRateLimitReset returns a UNIX timestamp of when the current monthly
// period will roll over
func (rc *ResponseCommon) GetRateLimitReset() int {
	return rc.convertHeaderToInt(rc.Header.Get("X-Ratelimit-Reset"))
}

type response struct {
	Common ResponseCommon
	Data   interface{}
}

func (r *response) copyCommon(rc *ResponseCommon) {
	rc.StatusCode = r.Common.StatusCode
	rc.Header = r.Common.Header
	rc.Status = r.Common.Status
}
