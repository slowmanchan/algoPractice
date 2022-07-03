package main

// two pointers again
// we want to shift the array down everytime we
// see and element not to be removed
func removeElement(nums []int, val int) int {
	// keep a low pointer
	low := 0

	// keep a high pointer
	for high := 0; high < len(nums); high++ {
		// if the high pointer doesnt equal our target
		// we set the low pointer value to the current value
		// and
		// we move the low pointer up one
		// we are basically shifting over the non target values to the
		// front and also keeping track of how many times we did this
		// ensures that if we get a run of the target, we can just keep moving
		// and if we hit a non target, we've kept the location (low pointer)
		// of the non target so we just replace the value there
		if nums[high] != val {
			nums[low] = nums[high]
			low++
		}
	}

	// we return the low pointer (which is also counting the number of)
	// non target values
	return low
}
