package main

import (
	"log"
)

func main() {
	log.Println(longestValidParentheses(")))()())"))
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
