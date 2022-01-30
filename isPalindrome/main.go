package main

import (
	"strconv"
)

func main() {

}

func isPalindrome(x int) bool {
	intStr := strconv.Itoa(x)
	halfWayLen := len(intStr) / 2

	secondHalfIndex := 0
	if len(intStr)%2 != 0 {
		secondHalfIndex = halfWayLen + 1
	} else {
		secondHalfIndex = halfWayLen
	}

	firstHalf := intStr[:halfWayLen]
	secondHalf := intStr[secondHalfIndex:]

	for i := 0; i < halfWayLen; i++ {
		if firstHalf[i] == secondHalf[halfWayLen-1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}
