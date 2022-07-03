package main

import (
	"strings"
)

// My solution
// func lengthOfLastWord(s string) int {
// 	strParts := strings.Split(s, " ")
// 	for i := len(strParts) - 1; i >= 0; i-- {
// 		if strParts[i] != "" {
// 			return len(strParts[i])
// 		}
// 	}
// 	return 0
// }

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	return len(s) - strings.LastIndex(s, " ") - 1
}
