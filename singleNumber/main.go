package main

// hash table
func singleNumber(nums []int) int {
	visited := map[int]int{}

	for _, n := range nums {
		visited[n] += 1
	}

	for k, v := range visited {
		if v == 1 {
			return k
		}
	}

	return 0
}
