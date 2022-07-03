package main

// my solutions
func plusOne(digits []int) []int {
	// loop thru digits in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		// if theres a nine, set it to 0
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			// otherwise add one to the number
			// and were done
			digits[i]++
			return digits
		}
	}
	// if we have ALL zeroes, put one in the front
	return append([]int{1}, digits...)
}
