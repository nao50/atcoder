package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		h1   int
		h2   int
		want int
	}{
		{"case 1", 0, 0, 0},
		{"case 2", 1, 0, 1},
		{"case 3", 0, 1, -1},
	}

	for _, tt := range tests {
		// shadowing
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.h1, tt.h2)
			if tt.want != got {
				t.Errorf("want: %v got: %v", tt.want, got)
			}
		})
	}
}
