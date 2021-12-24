package pexels_test

import "github.com/JayMonari/go-pexels"

var clientNewTestCases = []struct {
	expectErr bool
	options   pexels.Options
}{
	{
		true,
		pexels.Options{}, // No API key
	},
	{
		false,
		pexels.Options{APIKey: "pexels-api-key"},
	},
}

var responseTestCases = []struct {
	description string
	value       string
	want        int
}{
	{
		description: "Works for negative numbers",
		value:       "-2147483647",
		want:        -2147483647,
	},
	{
		description: "Works for zero",
		value:       "0",
		want:        0,
	},
	{
		description: "Works for max int32",
		value:       "2147483647",
		want:        2147483647,
	},
}
