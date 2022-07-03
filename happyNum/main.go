package main

// func isHappy(n int) bool {
// 	// we have to use a cycle detecting algo
// 	// because some numbers just keep going

// 	// its kind of like a linked list
// 	// where the fast runner will eventually
// 	// catch up to slow. here we just get the next number
// 	// for fast so its ahead
// 	slow := n
// 	fast := getNext(n)

// 	// if fast === 1, its happy and we exit
// 	// if slow DOESNT equal fast we keep going
// 	// when it does we have a cycle and were done
// 	for fast != 1 && slow != fast {
// 		// advance each computation and store the result
// 		slow = getNext(slow)
// 		// fast will find the next 2 sums (to stay ahead)
// 		// a cycyle will be found if fast === slow
// 		fast = getNext(getNext(fast))
// 	}
// 	return fast == 1
// }

// func getNext(n int) int {
// 	sum := 0
// 	for n > 0 {
// 		// rip out the right most digit
// 		d := n % 10
// 		// reduce the rightmost digit
// 		// this is how we process a number from right to left
// 		// one by one
// 		n = n / 10
// 		sum += d * d
// 	}

// 	return sum
// }

func isHappy(n int) bool {
	slow := 0
	fast := getNext(n)

	for fast != 1 || slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return fast == 1
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		ones := n % 10
		n /= 10
		sum += ones * ones
	}
	return sum
}
