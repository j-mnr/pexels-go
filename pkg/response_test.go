package pexels_test

import (
	"testing"
)

func TestGetRateLimit(t *testing.T) {
	if v := rcTC.GetRateLimit(); v != 1000 {
		t.Errorf("expected: 1000; got: %d", v)
	}
}

func TestGetRateLimitRemaning(t *testing.T) {
	if v := rcTC.GetRateLimitRemaining(); v != 1000 {
		t.Errorf("expected: 1000; got: %d", v)
	}
}

func TestGetRateLimitReset(t *testing.T) {
	if v := rcTC.GetRateLimitReset(); v != 1000 {
		t.Errorf("expected: 1000; got: %d", v)
	}
}
