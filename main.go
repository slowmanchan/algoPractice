package main

import "fmt"

func main() {
	letters := []string{"a", "b"}

	fmt.Println(pop(letters))
}

func pop(strs []string) ([]string, string) {
	if len(strs) > 0 {
		lastVal := strs[len(strs)-1]
		return strs[:len(strs)-1], lastVal
	} else {
		return strs, ""
	}
}
