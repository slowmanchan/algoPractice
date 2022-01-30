package main

import "testing"

func TestValidParens(t *testing.T) {
	var tests = []struct {
		name string
		x    string
		want bool
	}{
		{"test 9", "(}{)", false},
		{"test 10", "}{", false},
		{"test 10", "([}}])", false},
		{"test 9", "(])", false},
		{"test 9", "[[[]", false},
		{"test 4", "{[]}", true},
		{"test 1", "()", true},
		{"test 2", "()[]()", true},
		{"test 3", "(]", false},
		{"test 5", "{", false},
		{"test 6", "){", false},
		{"test 7", "(){}}{", false},
		{"test 8", ")(){}", false},
		{"test 8", "[(){}]", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lenOfLongword := isValid(tt.x)
			if lenOfLongword != tt.want {
				t.Errorf("got %v, want %v", lenOfLongword, tt.want)
			}
		})
	}
}
