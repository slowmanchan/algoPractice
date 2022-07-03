package main

import "testing"

func TestIsPalindrom(t *testing.T) {
	type test struct {
		name string
		x    int
		want bool
	}

	tests := []test{
		{
			"test 1", -10, false,
		},
		{
			"test 2", 121, true,
		},
		{
			"test 3", 10, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestIsPalindromNumConv(t *testing.T) {
	type test struct {
		name string
		x    int
		want bool
	}

	tests := []test{
		{
			"test 1", -10, false,
		},
		{
			"test 2", 121, true,
		},
		{
			"test 3", 10, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindromeNoConv(tt.x)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}
