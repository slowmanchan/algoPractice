package main

import "math"

// CONSTRAINTS
// only uppers
// min len 1, max len 7

// store the result
// run through the columnTitle
// find the value in map (key = "A", v = "1")
// add result
// Next letter , previous letter**len(nums) + nextletter**len(nums)-i
// return the result

// THIS IS THE WAY
// multiply the results by 26 (A = 1 so 1 set of col, B = 2 so 2 sets of columns)
// first pass will always give 0 which is good, 1 leter only means incomplete column
// we then get the ascii value of the letter and just minus 65 (ascii value of A)
// then we know the position of each letter. We have to add one because it the letter position
// will start at zero. ( A = 0 but we want 1 , B = 1 but its 2 in actual letter positing)
// then we just add that to our result from left to right, much simpler.
// LEFT TO RIGHT NO MAPPING
func titleToNumber(columnTitle string) int {
	result := 0
	for _, c := range columnTitle {
		result = result * 26
		result += (int(c) - 65 + 1)
	}
	return result
}

// the excel sheet works like ABC >> A*, B
// A* , 1 col maxed so 26 ** 1,
// B, no col maxed so 26 ** 0, so its just 2 more col
// the math here is letter * 26**len-1-index because as we go right accros the title,
// the blocks are decreasing per letter. The left is the highest multiplier because letters are
// added for the most recent cols.
// store the result
// continue

// RIGHT TO LEFT WITH MAPPING
func RtoLtitleToNumber(columnTitle string) int {
	result := 0
	for i, c := range columnTitle {
		letter := string(c)
		result += alphabet[letter] * int(math.Pow(float64(26), float64(len(columnTitle)-1-i)))
	}
	return result
}

var alphabet = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}
