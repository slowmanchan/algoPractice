package main

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
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
