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
			expected: []string{"", "", "", "leading", "and", "trailing", "spaces", "", "", ""},
		},
	}

	for _, c := range cases{
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual{
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("CleanInput(%v) == '%v' expected  %v", c.input, actual, c.expected)
			}
		}
	}
}