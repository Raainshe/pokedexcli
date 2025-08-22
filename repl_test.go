package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "blue greeN Yellow  ",
			expected: []string{"blue", "green", "yellow"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		fmt.Println("Length of actual: " + strconv.Itoa(len(actual)))
		fmt.Printf("%v\n", actual)
		fmt.Println("Length of expected: " + strconv.Itoa(len(c.expected)))
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match")
			t.FailNow()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Got '%s' but expected '%s'", word, expectedWord)
				t.Fail()
			}
		}
	}
}
