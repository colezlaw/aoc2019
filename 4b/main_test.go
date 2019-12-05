package main

import "testing"

func TestValid(t *testing.T) {
	tt := []struct {
		input    string
		expected bool
	}{
		{"11111", false},
		{"111111", false},
		{"223450", false},
		{"123789", false},
		{"112233", true},
		{"123444", false},
		{"111122", true},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			if isValidPassword(tc.input) != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, !tc.expected)
			}
		})
	}
}
