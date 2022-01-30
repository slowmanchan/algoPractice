package main

func main() {

}

func isValid(s string) bool {
	hashMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	reverseHash := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	if len(s)%2 != 0 {
		return false
	}
	if len(s) <= 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		strChar := string(s[i])
		if i == 0 && (strChar == ")" || strChar == "}" || strChar == "]") {
			return false
		}
		if val, ok := hashMap[string(s[i])]; ok {
			if i+1 > len(s)-1 {
				return false
			}
			if len(s)/2-1 < i {
				if string(s[i+1]) != val {
					return false
				}
			}
			if string(s[i+1]) != val && string(s[len(s)-i-1]) != val {
				return false
			}
		} else {
			if string(s[i-1]) != reverseHash[strChar] && string(s[len(s)-i-1]) != reverseHash[strChar] {
				return false
			}
		}

	}
	return true
}
