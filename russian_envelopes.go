// Top K Frequent Elements
// Priority Queue
//  - push
//  - pop
//  - len
package main

import (
	"log"
)

func main() {
	// log.Println(minMeetingRooms([][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}))
	// log.Println(maxEnvelopes([][]int{
	// 	[]int{5, 4},
	// 	[]int{6, 4},
	// 	[]int{3, 10},
	// 	[]int{6, 11},
	// 	[]int{2, 8},
	// 	[]int{1, 7}}))
	// log.Println(searchMatrix([][]int{[]int{1, 3, 5}}, 1))
	// log.Println(searchMatrix([][]int{
	// 	[]int{1, 2, 3, 4, 5},
	// 	[]int{6, 7, 8, 9, 10},
	// 	[]int{11, 12, 13, 14, 15},
	// 	[]int{16, 17, 18, 19, 20},
	// 	[]int{21, 22, 23, 24, 25}}, 5))
	// log.Println(searchMatrix([][]int{
	// 	[]int{3, 3, 8, 13, 13, 18},
	// 	[]int{4, 5, 11, 13, 18, 20},
	// 	[]int{9, 9, 14, 15, 23, 23},
	// 	[]int{13, 18, 22, 22, 25, 27},
	// 	[]int{18, 22, 23, 28, 30, 33},
	// 	[]int{21, 25, 28, 30, 35, 35},
	// 	[]int{24, 25, 33, 36, 37, 40}}, 21))
	log.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func maxEnvelopes(envelopes [][]int) int {
	return envelope(envelopes, -1)
}

func envelope(envelopes [][]int, currentEnvelope int) int {
	max := 0
	for i := 0; i < len(envelopes); i++ {
		if currentEnvelope < 0 || fitInside(envelopes[currentEnvelope], envelopes[i]) {
			lis := 1 + envelope(envelopes, i)
			if lis > max {
				max = lis
			}
		}
	}
	return max
}

func fitInside(currentEnvelope []int, newEnvelope []int) bool {
	if currentEnvelope[0] > newEnvelope[0] && currentEnvelope[1] > newEnvelope[1] {
		return true
	}
	return false
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	maxans := 1
	for i := 1; i < len(dp); i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && maxval < dp[j] {
				maxval = dp[j]
			}
		}
		dp[i] = maxval + 1
		if maxval > maxans {
			maxans = maxval
		}
	}
	log.Println(dp)
	return maxans
}
