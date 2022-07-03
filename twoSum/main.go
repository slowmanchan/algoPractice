package main

func main() {

}

// brute force solution
// loop - check each number (i)
// inner loop - check (i) against (j)
// O(n^2)

func bruteTwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); i++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 3, 2, 4 -- targ = 6
// memo
// find the diff between targ and number
// memoize the number we need plus the index
// if it appears then we got it
// O(n)
func twoSum(nums []int, target int) []int {
	memo := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if val, ok := memo[nums[i]]; ok {
			return []int{val, i}
		}
		memo[target-nums[i]] = i
	}

	return nil
}
