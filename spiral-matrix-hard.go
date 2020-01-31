// 54. Spiral Matrix
// https://leetcode.com/problems/spiral-matrix/
// hard
package main

import (
	"log"
	"math"
)

func main() {
	// log.Println(spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}))
	// log.Println(spiralOrder([][]int{[]int{1, 2}, []int{5, 6}, []int{9, 10}}))
	// log.Println(spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	// log.Println(spiralOrder([][]int{[]int{1}, []int{2}, []int{3}}))
	log.Println(spiralOrder([][]int{[]int{1, 2, 3}}))
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	ans := []int{}

	for dx := 0; dx < int(math.Min(float64((m+1)/2), float64((n+1)/2))); dx++ {
		ans1 := spiral(matrix, dx)
		log.Println(ans1)
		ans = append(ans, ans1...)
	}
	return ans
}

func spiral(matrix [][]int, dx int) []int {
	m := len(matrix)
	n := len(matrix[0])

	ans := []int{}

	i := 1
	dm := int(math.Min(float64(m/2), float64(dx)))
	dn := int(math.Min(float64(n/2), float64(dx)))
	im := 0
	in := 0
	spiralLength := n - dn*2 + m - dm*2 - 1
	if m-dm*2-1 > 0 {
		spiralLength += n - dn*2 - 1
	}
	if n-dn*2-1 > 0 {
		spiralLength += int(math.Max(float64(m-dm*2-2), 0))
	}
	log.Println("sum=", spiralLength, m, n, dm, dn)
	for i <= spiralLength {
		log.Println(dx, dm, dn, im, in, i)
		ans = append(ans, matrix[dm+im][dn+in])
		if i < n-dn*2 {
			in++
		} else if i < n-dn*2+m-dm*2-1 {
			im++
		} else if i < n-dn*2+m-dm*2-1+n-dn*2-1 {
			in--
		} else {
			im--
		}
		i++
	}
	return ans
}
