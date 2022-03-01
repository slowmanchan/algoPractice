package main

// my solutions
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		answer := digits[i] + 1
		if answer == 10 {
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}
		} else {
			digits[i] = answer
			return digits
		}
	}
	return digits
}
