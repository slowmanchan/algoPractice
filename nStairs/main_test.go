package main

import "testing"

func TestClimbStairs(t *testing.T) {
	type test struct {
		name string
		n    int
		want int
	}

	tests := []test{
		{
			"test 1", 2, 2,
		},
		{
			"test 2", 3, 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
