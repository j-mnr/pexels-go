package pexels_test

var testCases = []struct {
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
