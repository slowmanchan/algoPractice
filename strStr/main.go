package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := range haystack {
		if len(needle) > len(haystack[i:]) {
			return -1
		}
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

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
