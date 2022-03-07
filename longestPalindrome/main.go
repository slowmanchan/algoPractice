package main

func longestPalindrome(str string) string {
	palindrome := ""
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		temp := getPalindrome(runes, i, i)
		if len(temp) > len(palindrome) {
			palindrome = temp
		}
		temp = getPalindrome(runes, i, i+1)
		if len(temp) > len(palindrome) {
			palindrome = temp
		}
	}
	return palindrome
}

func getPalindrome(runes []rune, i, j int) string {
	for i >= 0 && j < len(runes) && runes[i] == runes[j] {
		i--
		j++
	}
	return string(runes[i+1 : j])
}
