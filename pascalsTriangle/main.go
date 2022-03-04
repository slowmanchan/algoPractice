package main

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for row := 0; row < numRows; row++ {
		triangle[row] = make([]int, row+1)

		left := 1
		right := row - 1

		triangle[row][left-1] = 1
		triangle[row][right+1] = 1

		for c := left; c <= right; c++ {
			triangle[row][c] = triangle[row-1][c-1] + triangle[row-1][c]
		}
	}
	return triangle
}

// pascals triangle
// each row is the sum of the above row
// for example 5
// start with 1
// second row is 1 - 1
// third row is 1 - 2 - 1
// forth row is 1 -3- 3- 1
// fifth row is 1 - 4 - 6 - 4 -1

// starts at 1 always
// first and last numbers are always 1

// eg 5
// iterate at most 5 times
//  length of the array = iteration , lets start at 1
// if length is greater then 1
// set first and end of array to 1

//6
// 1 i = 1
// 1 - 1 i = 2
// 1 - 2 - 1 i = 3
// 1 -3- 3- 1 i = 4
// 1 - 4 - 6 - 4 -1 i = 5
// 1 - 5 - 10 - 10 - 5 - 1 i = 6
// 1 - 6 - 15 - 20 - 15 - 6 - 1 = 7
