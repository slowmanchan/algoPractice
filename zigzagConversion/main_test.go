package main

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{"test 1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"test 2", "PAYPALISHIRING", 1, "PAYPALISHIRING"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convert(tt.s, tt.numRows)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
