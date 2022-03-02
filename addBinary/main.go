package main

// my solution
func addBinary(a string, b string) string {
	result := ""

	if len(a) > len(b) {
		for i := len(a) - len(b); i > 0; i-- {
			b = "0" + b
		}
	} else if len(a) < len(b) {
		for i := len(b) - len(a); i > 0; i-- {
			a = "0" + a
		}
	}
	carryOver := ""

	for i := len(a) - 1; i >= 0; i-- {
		letterA := string(a[i])
		letterB := string(b[i])

		if letterA == "0" && letterB == "0" {
			if carryOver == "1" {
				result = "1" + result
				carryOver = "0"
			} else {
				result = "0" + result
			}
		} else if letterA == "1" && letterB == "0" || letterB == "1" && letterA == "0" {
			if carryOver == "1" {
				result = "0" + result
			} else {
				result = "1" + result
			}
		} else {
			if carryOver == "1" {
				result = "1" + result
			} else {
				result = "0" + result
				carryOver = "1"
			}
		}

		if i == 0 && carryOver == "1" {
			result = "1" + result
		}
	}
	return result
}

//  0 + 0 = 0
// 	0 + 1 = 1
//  1 + 1 = 10.
// If the sum of 2 bits is greater than 1,
// we need to shift a column on the left. In decimal system, 1 + 1 = 2.
// Binary notation of 2 is 10 (1 * 2^1 + 0 * 2^0).
// So we keep 0 in the 1's column and shift (carry over) 1 to the 2's column.
