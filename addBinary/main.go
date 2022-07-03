package main

import (
	"strings"
)

//  0 + 0 = 0
// 	0 + 1 = 1
//  1 + 1 = 10.
// If the sum of 2 bits is greater than 1,
// we need to shift a column on the left. In decimal system, 1 + 1 = 2.
// Binary notation of 2 is 10 (1 * 2^1 + 0 * 2^0).
// So we keep 0 in the 1's column and shift (carry over) 1 to the 2's column.

// leet code
func addBinary(a string, b string) string {
	// padZeroes for the shorter string or
	// they wont line up for addition
	if len(a) > len(b) {
		b = padZeroes(b, len(a)-len(b))
	} else {
		a = padZeroes(a, len(b)-len(a))
	}

	// initialize a carry
	// and an answer slice
	carry := 0
	answer := []string{}

	// move through the first string in reverse
	// adding in math is always back to front
	for i := len(a) - 1; i >= 0; i-- {
		// if we have a 1, add it to the carry
		if a[i] == '1' {
			carry += 1
		}

		// same here
		if b[i] == '1' {
			carry += 1
		}

		// if theres a remainder here
		// that means we had 3 so we add
		// a 1. no remainder means we have a 0
		if carry%2 == 1 {
			answer = append(answer, "1")
		} else {
			answer = append(answer, "0")
		}

		// this will "consume" the carry
		// 1 gives us 0.5 which in int form is 0
		// 2 gives us 1
		// 3 also gives us 1
		// we want to carry over the 1 if we had a 3 or a 2
		// because 1 and 1 and 1 gives us 1 and a carry

		carry = carry / 2
	}

	// if we still have a carry, just add it on
	if carry == 1 {
		answer = append(answer, "1")
	}

	// reverse and join the array ( the answer is backwards )
	return strings.Join(reverse(answer), "")
}

func padZeroes(s string, n int) string {
	zeroes := ""
	for i := 0; i < n; i++ {
		zeroes += "0"
	}
	return zeroes + s
}

func reverse(strs []string) []string {
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
	return strs
}
