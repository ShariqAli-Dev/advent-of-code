package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Run("base string", func(t *testing.T) {
		got := PartOne([]string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"})
		expected := 48
		if got != expected {
			t.Errorf("test %s failed: got %d, expected %d", "base string", got, expected)
		}
	})

}
