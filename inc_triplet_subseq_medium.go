// Increasing Triplet Subsequence
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/781/
// medium
// Simple linked list management

package main

import (
	"log"
	"strconv"
)

func main() {
    log.Println(increasingTriplet[1,0,2,0,-1,0,1])
}


// Sol: find a number in an array that has two other numbers in the array that are smaller
// and the smaller of two comes first.
// The solution works because the task does not require to produce 3 numbers specifically,
// but only true/false that they exist.
func increasingTriplet(nums []int) bool {
    c1, c2 := math.MaxInt32, math.MaxInt32
    for i:=0; i<len(nums); i++ {
        if nums[i] <= c1 {
            c1 = nums[i]
        } else if nums[i] <= c2 {
            c2 = nums[i]
        } else {
            return true
        }
    }
    return false
}