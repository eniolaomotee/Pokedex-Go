package internal

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{
			input : "Hello World",
			expected : []string{"hello", "world"},
		},
		{
			input : "Go is Awesome",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input : "tESTING is FuN",
			expected: []string{"testing", "is", "fun"},
		},
		{
			input : "   Leading and trailing spaces   ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
	}

	for _, c := range cases{
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Fatalf("len mismatch: got %v want %v (actual=%v expected=%v)", len(actual), len(c.expected), actual, c.expected)
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range c.expected{
			if actual[i] != c.expected[i]{
				t.Fatalf("CleanInput(%q) = %v, want %v", c.input, actual, c.expected)
			}
		}
	}
}