package main

import (
	"log"
	"math"
)

type TestData struct {
	A      []int
	B      []int
	Result float64
}

func main() {
	testData := []TestData{
		TestData{A: []int{1, 2}, B: []int{}, Result: float64(1.5)},
		TestData{A: []int{1, 2, 3}, B: []int{}, Result: float64(2)},
		TestData{A: []int{1, 2}, B: []int{3, 4}, Result: float64(2.5)},
		TestData{A: []int{3, 4}, B: []int{1, 2}, Result: float64(2.5)},
		TestData{A: []int{1, 4, 6}, B: []int{3, 5}, Result: float64(4)},
		TestData{A: []int{1, 5, 9, 12}, B: []int{2, 4, 7, 8, 10, 15}, Result: float64(7.5)},
		TestData{A: []int{2, 4, 7, 8, 10, 15}, B: []int{1, 5, 9, 12}, Result: float64(7.5)},
		TestData{A: []int{1, 3}, B: []int{2}, Result: float64(2)},
		TestData{A: []int{3}, B: []int{-2, -1}, Result: float64(-1)},
		TestData{A: []int{2, 2, 2}, B: []int{2, 2, 2, 2}, Result: float64(2)},
		TestData{A: []int{1}, B: []int{2, 3}, Result: float64(2)},
		TestData{A: []int{1}, B: []int{2, 3, 4}, Result: float64(2.5)},
		TestData{A: []int{1}, B: []int{2, 3, 4, 5}, Result: float64(3)},
		TestData{A: []int{2, 3, 4}, B: []int{1}, Result: float64(2.5)},
		TestData{A: []int{1}, B: []int{-3, -2, -1}, Result: float64(-1.5)},
		TestData{A: []int{10}, B: []int{1, 2, 3, 4}, Result: float64(3)},
		TestData{A: []int{10}, B: []int{1, 2, 3}, Result: float64(2.5)},
		TestData{A: []int{1, 10, 20}, B: []int{2, 3, 12, 13, 14, 15}, Result: float64(12)},
		TestData{A: []int{20, 21, 22, 23, 24, 25}, B: []int{2, 12, 13, 14, 15}, Result: float64(20)},
	}
	for _, testSet := range testData {
		if r := findMedianSortedArrays2(testSet.A, testSet.B); r != testSet.Result {
			log.Fatalln("Incorrect result: ", r)
		} else {
			log.Println("Correct result: ", r)
		}
	}
	log.Println("Finish")
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	log.Println(nums1)
	log.Println(nums2)

	if len(nums1) == 0 {
		return arrayMedian(nums2)
	} else if len(nums2) == 0 {
		return arrayMedian(nums1)
	}

	// what diff in slices do we need before we are ready for median
	f := 1
	if math.Mod(float64(len(nums1)+len(nums2)), 2) == 0 {
		f = 0
	}
	log.Println(f)

	i := centerPosition(nums1)
	j := centerPosition(nums2)

	steps := 0
	for {
		log.Println("i=", i, ", j=", j)
		if i+1+j+1+f > len(nums1)-i-1+len(nums2)-j-1 { // left side is bigger -- it donates
			if nums1[i] > nums2[j] { // nums1 donates
				if i == 0 { // edge case -- indicate that nums1 has no elements for the left side
					i = -1
				} else {
					i = int(math.Round(float64(i+1)/2)) - 1
				}
			} else { // nums2 donates
				if j == 0 { // edge case -- indicate that nums2 has no elements for the left side
					j = -1
				} else {
					j = int(math.Round(float64(j+1)/2)) - 1
				}
			}
		} else if i+1+j+1+f < len(nums1)-i-1+len(nums2)-j-1 { // right side is bigger -- it donates
			if nums1[i+1] < nums2[j+1] { // nums1 donates
				i = int(math.Round(float64(len(nums1)-i-1)/2)) + i
			} else { // nums2 donates
				j = int(math.Round(float64(len(nums2)-j-1)/2)) + j
			}
		} else { // balanced
			// but is it balanced to median
			if safeMax(nums1, nums2, i, j) > safeMin(nums1, nums2, i+1, j+1) { // no
				panic("no balance!")
			}
			if f == 0 {
				return float64(safeMax(nums1, nums2, i, j)+safeMin(nums1, nums2, i+1, j+1)) / 2
			} else {
				// edge cases first -- nums1 or nums2 do not contribute to the right side
				return float64(safeMin(nums1, nums2, i+1, j+1))
			}
		}
		steps++
		if steps > 10 {
			panic("error")
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	log.Println(nums1)
	log.Println(nums2)

	if len(nums1) == 0 {
		return arrayMedian(nums2)
	} else if len(nums2) == 0 {
		return arrayMedian(nums1)
	}

	// what diff in slices do we need before we are ready for median
	f := 1
	if math.Mod(float64(len(nums1)+len(nums2)), 2) == 0 {
		f = 0
	}
	log.Println(f)

	i := centerPosition(nums1)
	j := centerPosition(nums2)

	// then we split to nums1[0:i-1] + nums2[0:j-1] as the left and rest as the right
	// and start to joggling elements left or right until we get parity
	for {
		log.Println("i=", i, "j=", j)
		if j > 10 || i > 10 {
			panic("error")
		}
		// median will be in the left half
		if i+j+f == len(nums1)-i+len(nums2)-j && (i == 0 || j == len(nums2) || j == 0 || i == len(nums1) || nums1[i-1] <= nums2[j]) {
			if f == 0 {
				var left int
				var right int
				// how nums1 and nums2 contribute to left side of the merged array
				if j > 0 && i > 0 {
					left = max(nums1[i-1], nums2[j-1])
					if i == len(nums1) {
						right = nums2[j]
					} else if j == len(nums2) {
						right = nums1[i]
					} else {
						right = min(nums1[i], nums2[j])
					}
				} else if j > 0 && i == 0 {
					left = nums2[j-1]
					if j == len(nums2) {
						right = nums1[0]
					} else {
						right = min(nums2[j], nums1[0])
					}
				} else if j == 0 && i > 0 {
					left = nums1[i-1]
					right = nums2[0]
				} else {
					panic("error!")
				}

				return float64(left+right) / 2
			}
			if i == len(nums1) && j < len(nums2) {
				return float64(nums2[j])
			} else if j == len(nums2) && i < len(nums1) {
				return float64(nums1[i])
			}
			return float64(min(nums1[i], nums2[j])) // because sorted in ascending order
		}
		if i == len(nums1) || nums2[j] < nums1[i] {
			j++
		} else {
			i++
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// some of the indexes must be valid
func safeMin(arr1, arr2 []int, i, j int) int {
	if i < 0 || i >= len(arr1) {
		return arr2[j]
	}
	if j < 0 || j >= len(arr2) {
		return arr1[i]
	}
	if arr1[i] < arr2[j] {
		return arr1[i]
	}
	return arr2[j]
}

func safeMax(arr1, arr2 []int, i, j int) int {
	if i < 0 || i >= len(arr1) {
		return arr2[j]
	}
	if j < 0 || j >= len(arr2) {
		return arr1[i]
	}
	if arr1[i] > arr2[j] {
		return arr1[i]
	}
	return arr2[j]
}

func arrayMedian(arr []int) float64 {
	if len(arr)%2 == 0 {
		return float64(arr[int(len(arr)/2)-1]+arr[int(len(arr)/2)]) / 2
	}
	return float64(arr[int(len(arr)/2)])
}

// After finding the center of nums1, we need to find where this central element could be placed in nums2
func findPositionOfElementInSortedArray(a int, nums []int) int {
	j := centerPosition(nums)
	if j == len(nums) {
		j--
	}
	for {
		if j == 0 || j == len(nums)-1 {
			return j
		}
		if a < nums[j] && a > nums[j-1] {
			return j
		}
		if a < nums[j] {
			j--
		} else if a > nums[j] {
			j++
		} else { // a == nums[j]
			return j
		}
	}
}

func centerPosition(nums []int) int {
	return int(math.Round(float64(len(nums))/2)) - 1
}
