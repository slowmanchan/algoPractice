package main

func removeElement(nums []int, val int) int {
	result := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[result] = nums[j]
			result++
		}
	}
	return result
}

// remove val from nums
// modify same array
// return the length of remaining vals

// EG. array = [3, 2, 2, 3] val = 3
// run through array
// 3 == 3 then do nothing
// 2 != 3 then set nums[0] to 2 why?
// we shifting the array to the left. keeping the
// deleted items in the front of the array
// 2 != 3 then set nums[1] to 2 why?
// we use result as the moving window because it keeps track
// of how many we've deleted which is also the right position the shift
// 3 == 3 then do nothing
// we end up with 2 deleted items
