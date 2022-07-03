package main

import "fmt"

func main() {
	// n := romanToInt("IVIX")
	// fmt.Println(n)
	n := romanToInt("MCMXCIV")
	fmt.Println(n)
}

func romanToInt(s string) int {
	numerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	prev := ""
	sum := 0

	for i := 0; i < len(s); i++ {
		letter := string(s[i])

		if prev == "I" && (letter == "V" || letter == "X") {
			sum += numerals[letter] - 2
		} else if prev == "X" && (letter == "L" || letter == "C") {
			sum += numerals[letter] - 20
		} else if prev == "C" && (letter == "D" || letter == "M") {
			sum += numerals[letter] - 200
		} else {
			sum += numerals[letter]
		}

		if letter == "I" || letter == "X" || letter == "C" {
			prev = letter
		}
	}

	return sum
}
