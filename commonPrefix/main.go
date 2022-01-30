package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	longestCommon := ""
	for i, s := range strs {
		if s == "" {
			return ""
		}
		if i == 0 {
			longestCommon = s
			continue
		}
		ii := 0
		if len(s) > len(longestCommon) {
			ii = len(longestCommon)
		} else {
			ii = len(s)
		}
		for i := ii - 1; i >= 0; i-- {
			nS := s[0 : i+1]
			if nS == longestCommon[0:i+1] {
				longestCommon = nS
				break
			}
			if i == 0 {
				return ""
			}
		}
	}
	return longestCommon
}
