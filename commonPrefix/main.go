package main

func main() {

}

// vertical scanning
// flower
// fl
// flight

// move through and compare each column (same index) of each word
// and keep going until you encouter a letter thats not the same
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		// always use the first word to check against
		// check each character of the word
		c := strs[0][i]
		// check each word at the same index (vertical scan)
		for j := 1; j < len(strs); j++ {
			// if the outer length of the word is the same
			// as the word we are checking in the inner
			// OR the inner char doesnt match the outer one
			// return that previous prefix
			// if its blank then the longest common is blank
			// this covers all cases
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
