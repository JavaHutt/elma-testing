package main

import "testing"

func TestIsValidURL(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{
			"numbers",
			"12345",
			false,
		},
		{
			"string",
			"abcde",
			false,
		},
		{
			"random",
			"asds23fdd",
			false,
		},
		{
			"nohttp",
			"google.com",
			false,
		},
		{
			"valid",
			"http://www.codewars.com",
			true,
		},
	}

	for _, test := range tests {
		if result := isValidURL(test.input); result != test.expected {
			t.Errorf("Test failed: %s inputted, %t expected, recieved: %t", test.input, test.expected, result)
		}
	}
}
