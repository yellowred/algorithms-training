package main

import (
	"log"
)

func main() {
	// [[1,1],[1,1],[2,3]]
	log.Println(maxPoints([][]int{[]int{5, 1}, []int{1, 1}, []int{1, 1}, []int{2, 3}, []int{2, 3}, []int{1, 1}, []int{3, 5}}))

	log.Println(maxPoints([][]int{[]int{0, 0}, []int{0, 0}}))

	//[[0,0],[94911151,94911150],[94911152,94911151]]
	log.Println(maxPoints([][]int{[]int{0, 0}, []int{94911151, 94911150}, []int{94911152, 94911151}}))
}

func isMatch(s string, p string) bool {
	si := 0
	pi := 0
	spec1 := "*"
	spec2 := "."

	log.Println(s, p)
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	// todo edge p=""

	if pi < len(p)-1 && p[pi+1] == spec1[0] {
		if isMatch(s[si:len(s)], p[pi+2:len(p)]) {
			return true
		}
		for si < len(s) && (s[si] == p[pi] || p[pi] == spec2[0]) {
			si++

			if isMatch(s[si:len(s)], p[pi+2:len(p)]) {
				return true
			}
		}
	} else if len(s) == 0 && len(p) > 0 || len(s) > 0 && len(p) == 0 {
		return false
	} else if s[si] == p[pi] || p[pi] == spec2[0] {
		si++
		pi++
		log.Println(s, p, si, pi)
		return isMatch(s[si:len(s)], p[pi:len(p)])
	} else {
		return false
	}

	if si == len(s) && pi == len(p) {
		return true
	}
	log.Println("finish", si, len(s), pi, len(p))
	return false
}

// 32. Longest Valid Parentheses
// Given a string containing just the characters '(' and ')', find the length
// of the longest valid (well-formed) parentheses substring.
func longestValidParentheses(s string) int {

	lp := "("
	rp := ")"

	i := 0

	max := 0
	stack := []int{}
	startIndex := -1
	for i < len(s) {
		if s[i] == lp[0] {
			stack = append(stack, i)
		} else if s[i] == rp[0] {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
				if len(stack) > 0 {
					startIndex = stack[len(stack)-1]
				}
				log.Println(i, startIndex)
				if max < i-startIndex {
					max = i - startIndex
				}
			}
		} else {
			panic("incorrect symbol")
		}
		i++
	}
	return max
}

// 149. Max Points on a Line
// Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.
func maxPoints(points [][]int) int {
	i := 0
	max := 0

	if len(points) == 1 {
		return 1
	}
	for i < len(points)-1 {
		duplicates := 0
		j := i + 1
		for j < len(points) {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				duplicates++
				if max < duplicates+1 {
					max = duplicates + 1
				}
			} else {
				s := slope(points[i], points[j])
				log.Println(i, j, s)
				l := j + 1
				counter := 2 + duplicates
				for l < len(points) {
					if slope(points[l], points[j]) == s || (points[l][0] == points[j][0] && points[l][1] == points[j][1]) {
						counter++
					}
					log.Println("l", slope(points[l], points[j]), s, l, counter)
					l++
				}
				if max < counter {
					max = counter
				}
			}
			j++
		}
		i++
	}
	return max
}

func slope(p1 []int, p2 []int) float64 {
	slope := float64(p1[1]-p2[1]) / float64(p1[0]-p2[0])
	if p1[0]-p2[0] == 0 {
		slope = 0
	}
	return slope
}
