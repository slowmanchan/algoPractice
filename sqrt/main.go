package main

func mySqrt(x int) int {
	result := 1
	for i := x / 2; i > 0; i-- {
		if i*i == x {
			return i
		}
	}
	return result
}
