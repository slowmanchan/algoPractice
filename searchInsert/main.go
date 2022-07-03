package main

// bin search solution
func searchInsertBin(nums []int, target int) int {
	// low and high pointers
	// start and end
	low := 0
	high := len(nums) - 1

	for low <= high {
		// get the midpoint of the array (then subarrays)
		// start + (end-start)
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}

// non bin search solution
func searchInsert(nums []int, targ int) int {
	low := 0
	for high := 0; high < len(nums); high++ {
		if nums[high] == targ {
			return high
		}

		if nums[low] < targ {
			low++
		} else {
			return low
		}
	}

	return low
}
