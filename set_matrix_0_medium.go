// Given a m x n matrix, if an element is 0, set its entire row and column to 0.
// Do it in-place.
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/777/
package main

import (
	"log"
)

func main() {
	matrix := [][]int{}
	setZeroes(matrix)
	log.Println(matrix)
}

func setZeroes(matrix [][]int) {
	col0 := false
	// remember matrix params to avoid recalc later
	cols := len(matrix[0])
	rows := len(matrix)

	// iterate over matrix and mark columns and rows to fill with zeroes using column 0 and row 0
	for i := 0; i < rows; i++ {
		// column 0 is for marks so we need save if we want to rewrite it later
		if matrix[i][0] == 0 {
			col0 = true
		}
		// iterate over all columns except column 0
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// start from opposite side of the column and row choosen for marks so we can overwrite
	// them in the process
	for i := rows - 1; i >= 0; i-- {
		// do not overwrite column 0
		for j := cols - 1; j >= 1; j-- {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		// set column 0 elemnt to 0 if we already processed the row
		// also we can't rely on matrix as 0,0 element is abigous
		if col0 {
			matrix[i][0] = 0
		}
	}
}
