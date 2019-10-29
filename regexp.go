package main

import (
	"log"
	"regexp"
	"strings"
)

func main() {
	log.Println(isNumber("1e3"))
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
type Parenthesis struct {
    ch byte
    i int
}

func longestValidParentheses(s string) int {
	lp := "("
	rp := ")"

	i := 0

	max := 0
	stack := []Parenthesis{}
	startIndex := 0
	for i < len(s) {
		if s[i] == lp[0] {
	            stack = append(stack, Parenthesis{ch: lp[0], i: i})
	            log.Println("lp", stack)
		} else if s[i] == rp[0] {
			if len(stack) > 0 && stack[len(stack)-1].ch == lp[0] {
		                stack = stack[0 : len(stack)-1]
        	        	if len(stack) > 0 {
		                  startIndex = stack[len(stack)-1].i
                		} else {
				  startIndex = -1
				}
		                log.Println(i, startIndex, stack)
                		if max < i-startIndex {
		                    max = i - startIndex
                		}

		                log.Println(i, startIndex, stack, max)
			} else {
			  stack = append(stack, Parenthesis{ch: rp[0], i: i})
	                  log.Println("rp", stack)
			}
		} else {
			panic("incorrect symbol")
		}
		i++
	}
	return max
<<<<<<< HEAD
}
=======
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

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	r := regexp.MustCompile(`(^[^0-9]+$`)
	if !r.MatchString(s) {
		return false
	}
	vs := strings.Split(s, "e")
	if len(vs) > 2 {
		return false
	}
	return true
}
>>>>>>> 4747750d398f9329d4261bd32b59682072233186
