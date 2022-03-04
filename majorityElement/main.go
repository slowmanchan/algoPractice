package main

func majorityElement(nums []int) int {
	visited := map[int]int{}
	for _, n := range nums {
		visited[n] += 1
	}

	mostFrequent := 0
	result := 0
	for num, freq := range visited {
		if freq > mostFrequent {
			mostFrequent = freq
			result = num
		}
	}

	return result
}
