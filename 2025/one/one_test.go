package main

import (
	"fmt"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"base test", []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 3},
		{"empty input", []string{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("starting test for", tt.name)
			got := Answer(tt.input)
			if got != tt.want {
				t.Errorf("Answer() got = %d, want %d", got, tt.want)
			}
		})
	}
}
