package main

import "fmt"

func main() {
	fmt.Println(removeDupes([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDupes(nums []int) int {
	// if the length of nums is 0 then
	// we have no dupes
	if len(nums) == 0 {
		return 0
	}

	// two pointers i and j
	// start j ahead of i
	low := 0
	for high := 1; high < len(nums); high++ {
		// if the number and j and i are not equal
		// increment i (the low pointer)
		// set the value at i to the j value
		if nums[high] != nums[low] {
			low++
			nums[low] = nums[high]
		}
	}

	// we return how many times the i pointer moved
	// plus 1 because the first value was not counted
	return low + 1
}
