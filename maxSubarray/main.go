package main

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for _, n := range nums[1:] {
		if currentSum+n > n {
			currentSum = currentSum + n
		} else {
			currentSum = n
		}

		if maxSum < currentSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// go through the array
// add up the next number, then do the same until the end
// do it for each number
// store results in a variable

// this is the way, start at the first

// [-2] current -2, max -2
// [-2, 1] current 1, max 1
// [-2, 1, -3] current -2, max 1
// [-2,1,-3,4] current 4, max 4
// [-2,1,-3,4,-1] current 3, max 4
// [-2,1,-3,4,-1,2] current 5, max 5
// [-2,1,-3,4,-1,2,1] current 6, max 6
// [-2,1,-3,4,-1,2,1,-5] current -1, max 6
// [-2,1,-3,4,-1,2,1,-5,4] current 4, max 6
