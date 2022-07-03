package main

import (
	"strconv"
)

func main() {
	isPalindrome(12)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	intStr := strconv.Itoa(x)
	secondHalfMid := len(intStr) / 2

	if x%2 != 0 {
		secondHalfMid += 1
	}

	return intStr[0:len(intStr)/2] == intStr[secondHalfMid:]
}

// good answer
// no o(n) space
// no string conversion
func isPalindromeNoConv(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}
