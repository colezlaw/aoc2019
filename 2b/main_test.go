package main

import "testing"

func TestComputer(t *testing.T) {
	tt := []struct {
		title    string
		program  []int
		expected []int
	}{
		{"explained", []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{"example 1", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"example 2", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"example 3", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"example 4", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tc := range tt {
		t.Run(tc.title, func(t *testing.T) {
			actual := computer(tc.program)
			if len(actual) != len(tc.expected) {
				t.Fatalf("expected output len of %d, got %d", len(tc.expected), len(actual))
			}

			for i, x := range tc.expected {
				if actual[i] != x {
					t.Errorf("at index %d expected %d got %d", i, x, actual[i])
				}
			}
		})
	}
}

func testParseProgram(t *testing.T) {
	tt := []struct {
		title    string
		program  string
		expected []int
	}{
		{"explained", "1,9,10,3,2,3,11,0,99,30,40,50", []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}},
		{"example 1", "1,0,0,0,99", []int{1, 0, 0, 0, 99}},
		{"example 2", "2,3,0,3,99", []int{2, 3, 0, 3, 99}},
		{"example 3", "2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 0}},
		{"example 4", "1,1,1,4,99,5,6,0,99", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}},
	}

	for _, tc := range tt {
		t.Run(tc.title, func(t *testing.T) {
			actual := parse(tc.program)
			if len(actual) != len(tc.expected) {
				t.Fatalf("expected output len of %d, got %d", len(tc.expected), len(actual))
			}

			for i, x := range tc.expected {
				if actual[i] != x {
					t.Errorf("at index %d expected %d got %d", i, x, actual[i])
				}
			}
		})
	}
}
