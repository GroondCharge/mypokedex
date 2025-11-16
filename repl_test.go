package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This Is a TeSt",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "TESTING ALL CAPS",
			expected: []string{"testing", "all", "caps"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("lengths not matching")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			fmt.Println(word, expectedWord)
			if word != expectedWord {
				t.Errorf("words not matching")
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
