package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	longestCommon := ""
	for i, s := range strs {
		if i == 0 {
			longestCommon = s
			continue
		}
		for i := len(longestCommon) - 1; i >= 0; i-- {
			if len(s) == 1 {
				longestCommon = s[0:1]
				break
			}
			if len(s) != len(longestCommon) {
				continue
			}
			if s[0:i] == longestCommon[0:i] {
				longestCommon = s[0:i]
				continue
			}
			if i == 0 {
				return ""
			}
		}
	}
	return longestCommon
}
