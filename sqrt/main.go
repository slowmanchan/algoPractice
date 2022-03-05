package main

// newton method
func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

// very slow
// func mySqrt(x int) int {
// 	if x < 2 {
// 		return x
// 	}

// 	left := 2
// 	right := x

// 	for left <= right {
// 		pivot := left + (right - left)
// 		num := pivot * pivot
// 		if num > x {
// 			right = pivot - 1
// 		} else if num < x {
// 			left = pivot + 1
// 		} else {
// 			return pivot
// 		}
// 	}
// 	return right
// }

// def mySqrt(self, x):
// if x < 2:
// 	return x

// left, right = 2, x // 2

// while left <= right:
// 	pivot = left + (right - left) // 2
// 	num = pivot * pivot
// 	if num > x:
// 		right = pivot -1
// 	elif num < x:
// 		left = pivot + 1
// 	else:
// 		return pivot

// return right
