// 34. Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// medium
// Using binary searhc to find each edge.

package main

import (
	"log"
)

func main() {
	log.Println(searchRange([]int{1, 1, 1}, 1))
}

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	res[0] = fastSearch(nums, target, true)
	res[1] = fastSearch(nums, target, false) - 1

	if res[0] == len(nums) || nums[res[0]] != target {
		res[0] = -1
		res[1] = -1
	}
	return res
}

func fastSearch(nums []int, target int, left bool) int {
	l := 0
	// rightmost is plus one of the right index so that
	// leftmost can reach the right and can be returned,
	// otherwise l<r makes it impossible for l to be eq r
	r := len(nums)
	for l < r {
		mid := (r + l) / 2
		if nums[mid] > target || (left && nums[mid] == target) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	log.Println(l)
	return l
}
