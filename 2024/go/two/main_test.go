package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		reports  [][]int
		expected int
	}{
		{
			name: "all safe reports",
			reports: [][]int{
				{7, 6, 5, 4, 3, 2, 1},
				{1, 2, 3, 4, 5, 6, 8},
			},
			expected: 2,
		},
		{
			name: "all unsafe reports",
			reports: [][]int{
				{7, 10, 5, 4, 3, 2, 1},
				{100, 30, 10, 15},
			},
			expected: 0,
		},
		{
			name: "mixed safe & unsafe reports",
			reports: [][]int{
				{7, 6, 5, 4, 3, 2, 1},
				{100, 30, 10, 15},
			},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := PartOne(test.reports)
			if got != test.expected {
				t.Errorf("test %s failed: got %d, expected %d", test.name, got, test.expected)
			}
		})
	}
}
