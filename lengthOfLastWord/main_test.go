package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	strTests := []struct {
		name string
		s    string
		want int
	}{
		{
			"two word sentence",
			"hello world",
			5,
		},
		{
			"spaces front and back",
			"   fly me   to   the moon  ",
			4,
		},
		{
			"4 word sentence",
			"luffy is still joyboy",
			6,
		},
	}

	for _, tt := range strTests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLastWord(tt.s)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
