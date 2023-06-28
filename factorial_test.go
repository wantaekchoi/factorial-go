package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		n        int64
		expected string
	}{
		{0, "1"},
		{1, "1"},
		{5, "120"},
		{10, "3628800"},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		result := factorial(tc.n).String()
		if result != tc.expected {
			t.Errorf("Factorial(%d) = %s, expected %s", tc.n, result, tc.expected)
		}
	}
}
