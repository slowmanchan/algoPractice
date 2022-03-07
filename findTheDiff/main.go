package main

import (
	"sort"
	"strings"
)

func findTheDifference(s string, t string) byte {
	strS := strings.Split(s, "")
	strB := strings.Split(t, "")

	sort.Strings(strS)
	sort.Strings(strB)

	for i := 0; i < len(strS); i++ {
		if strS[i] != strB[i] {
			return byte([]rune(strB[i])[0])
		}
	}

	return byte([]rune(strB[len(strB)-1])[0])
}

// simpler and more performant
func hashMapfindTheDifference(s string, t string) byte {
	hashMap := map[rune]int{}

	for _, c := range s {
		hashMap[c] += 1
	}

	for _, cc := range t {
		if hashMap[cc] == 0 {
			return byte(cc)
		}
		if _, ok := hashMap[cc]; ok {
			hashMap[cc] -= 1
		}
	}
	return 0
}
