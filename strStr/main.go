package main

// func strStr(haystack string, needle string) int {
// 	if len(needle) == 0 {
// 		return 0
// 	}

// 	for i := range haystack {
// 		if len(needle) > len(haystack[i:]) {
// 			return -1
// 		}
// 		if haystack[i:len(needle)+i] == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }

// find needle within haystack and return the index it starts at
// eg. needle = ll , haystack = hello
// go through each letter left to right
// check if the first letter matches the current letter
// if it does loop through the rest to check matches and
// keep track of the index in the result
// when there is no match reset the index to -1
// else bail out and check the next letter
// if the index is not -1 then return
// else keep going
func strStr(haystack string, needle string) int {
	// blank strings should return 0
	if needle == "" {
		return 0
	}

	// lets go through haystack
	for i := 0; i < len(haystack); i++ {
		// let check if the window size will exceed
		// the length of haystack ( oob )
		// our current i + the len of the needle will determine that
		if i+len(needle) > len(haystack) {
			return -1
		}

		// lets check the current window (current index to index + len of
		// needle)
		// if its needle , return the starting index where it
		// was found
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
