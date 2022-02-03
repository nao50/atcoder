package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    int
		want int
	}{
		{"case 1", 15000, 65},
		{"case 2", 75000, 89},
		{"case 3", 200, 02},
	}

	for _, tt := range tests {
		// shadowing
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.m)
			if tt.want != got {
				t.Errorf("want: %v got: %v", tt.want, got)
			}
		})
	}
}
