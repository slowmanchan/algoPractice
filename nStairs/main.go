package main

// recursive
// brute force
// we are going to climb all
// 0(2^n)
// func climbStairs(n int) int {
// 	return climb(0, n)
// }

// func climb(i, n int) int {
// 	// the steps we take are greater then the num of steps
// 	// return 0 because theres no way to climb it
// 	if i > n {
// 		return 0
// 	}

// 	// if the number of steps we take is equal to the number of steps
// 	// that means theres one way to climb with that step combo
// 	if i == n {
// 		return 1
// 	}

// 	// remember, each call to climb branches out to TWO more
// 	// calls. this is how we check the different combos
// 	// of 1 step 2 step 1 step.
// 	return climb(i+1, n) + climb(i+2, n)
// }

// n = 3
//        (1, 3) 2           (2, 3)  1
//         / \                 / \
//      (2,3) (3,3) 1 + 1    (3,3) (4,3) 1 + 0
//       /   \
//     (3,3) (4,3) = 1

// if you think about it
// climb n stairs is just
// the fibonacci number
// f(i) = (i - 1) + (i - 2)
// to reach the ith step its a combo
// of each preceding i
// to get to the 3rd step we just add the ways
// to get to the first step AND the second
// to get 4th , we add the ways to get to the 3rd and 2nd
// instead of recounting like we did in the above brute force

func climbStairs(n int) int {
	// if we've got 1 step then only one way
	if n == 1 {
		return 1
	}

	// only 1 way for 1 step and 2 ways for 2 steps
	// we should init these values because the formula wont
	// work otherwise
	first := 1
	second := 2

	// start at 3 since we initialized the first 2 values already
	// if i steps exceeds the target then bailout
	// dont forget to swap first and second or we'd be doing
	// the same calculation over and over
	for i := 3; i <= n; i++ {
		third := first + second

		first = second
		second = third
	}

	return second
}
