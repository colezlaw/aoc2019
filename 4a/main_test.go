package main

import "testing"

func TestValid(t *testing.T) {
	tt := []struct {
		input    string
		expected bool
	}{
		{"11111", false},
		{"111111", true},
		{"223450", false},
		{"123789", false},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			if isValidPassword(tc.input) != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, !tc.expected)
			}
		})
	}
}
