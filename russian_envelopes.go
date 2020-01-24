// 354. Russian Doll Envelopes
// https://leetcode.com/problems/russian-doll-envelopes/
// You have a number of envelopes with widths and heights given as a pair of integers (w, h).
// One envelope can fit into another if and only if both the width and
// height of one envelope is greater than the width and height of the other envelope.
// What is the maximum number of envelopes can you Russian doll? (put one inside other)
// Note: Rotation is not allowed.

package main

import (
	"log"
	"sort"
)

type Envelopes [][]int

func (m Envelopes) Len() int {
	return len(m)
}

func (m Envelopes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Envelopes) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func main() {
	log.Println(maxEnvelopes([][]int{[]int{10, 8}, []int{1, 12}, []int{6, 15}, []int{2, 18}}))
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	var arr Envelopes
	arr = envelopes
	sort.Sort(arr)
	dp := make([]int, len(envelopes))
	dp[0] = 1
	var ans = 1
	for i := 1; i < len(arr); i++ {
		maxLis := 1
		for j := 0; j < i; j++ {
			if arr[j][0] < arr[i][0] && arr[j][1] < arr[i][1] {
				if maxLis < dp[j]+1 {
					maxLis = dp[j] + 1
				}
			}
		}
		dp[i] = maxLis
		if maxLis > ans {
			ans = maxLis
		}
	}
	return ans
}
