package main

// O(n)
// we are creating a window of values and
// keeping track of the start and end of this window
// we also keep track of each letter and where the end bound is
// func lengthOfLongestSubstring(s string) int {
// 	// init a map to store the letters and index
// 	memo := map[string]int{}

// 	ans := 0

// 	// create a sliding window (i, j)
// 	start := 0
// 	for end := 0; end < len(s); end++ {
// 		letter := string(s[end])

// 		// if the letter is in our map
// 		// set i to the larger of the letters index or i
// 		// the window is the size of the largest substring
// 		if _, ok := memo[letter]; ok {
// 			start = max(memo[letter], start)
// 		}

// 		// the answer will be the larger of the window or
// 		// our stored answer
// 		ans = max(ans, end-start+1)

// 		// lets store the index of our high index and add 1
// 		// to "slide it"
// 		memo[letter] = end + 1
// 	}

// 	return ans
// }

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func lengthOfLongestSubstring(s string) int {
	memo := map[string]int{}
	ans := 0

	start := 0
	for end := 0; end < len(s); end++ {
		letter := string(s[end])

		if _, ok := memo[letter]; ok {
			start = max(memo[letter], start)
		}

		ans = max(ans, end-start+1)

		memo[letter] = end + 1
	}

	return ans
}
