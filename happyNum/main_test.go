package main

import "testing"

func TestHappy(t *testing.T) {
	type test struct {
		name string
		n    int
		want bool
	}

	tests := []test{
		{
			"test 1", 19, true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isHappy(tt.n)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}
