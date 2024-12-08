package main

import (
	"testing"
)

// TestAddSum is a test function for AddSum.
func TestAddSum(t *testing.T) {
	testCases := []struct {
		num1     int
		num2     int
		expected int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{100, 200, 300},
		{-100, -200, -300},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			actual := AddSum(tc.num1, tc.num2)
			if actual != tc.expected {
				t.Errorf("AddSum(%d, %d) = %d; expected %d", tc.num1, tc.num2, actual, tc.expected)
			}
		})
	}
}
