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
		TestData{A: []int{1}, B: []int{2, 3, 4, 5, 6, 7, 8}, Result: float64(4.5)},
		TestData{A: []int{2, 3, 4, 5, 6, 7, 8}, B: []int{1}, Result: float64(4.5)},
		TestData{A: []int{2, 3, 4}, B: []int{1}, Result: float64(2.5)},
		TestData{A: []int{1}, B: []int{-3, -2, -1}, Result: float64(-1.5)},
		TestData{A: []int{10}, B: []int{1, 2, 3, 4}, Result: float64(3)},
		TestData{A: []int{10}, B: []int{1, 2, 3}, Result: float64(2.5)},
		TestData{A: []int{1, 10, 20}, B: []int{2, 3, 12, 13, 14, 15}, Result: float64(12)},
		TestData{A: []int{20, 21, 22, 23, 24, 25}, B: []int{2, 12, 13, 14, 15}, Result: float64(20)},
		TestData{A: []int{5, 6}, B: []int{1, 2, 3, 4, 7, 8}, Result: float64(4.5)},
	}
	for _, testSet := range testData {
		if r := findMedianSortedArrays(testSet.A, testSet.B); r != testSet.Result {
			log.Fatalln("Incorrect result: ", r)
		} else {
			log.Println("Correct result: ", r)
		}
	}
	log.Println("Finish")
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

	steps := 0
	for {
		log.Println("i=", i, ", j=", j)
		if largestSide(nums1, nums2, i, j, f)<0 || (largestSide(nums1, nums2, i, j, f)==0 && (safeMax(nums1, nums2, i, j) > safeMin(nums1, nums2, i+1, j+1))) { // left side is bigger -- it donates
			if j<0 || i>=0 && nums1[i] > nums2[j] { // nums1 donates
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
			if j == len(nums2) -1 || i < len(nums1)-1 && nums1[i+1] < nums2[j+1] { // nums1 donates
				shift := int(float64(len(nums1)-i-1)/2)
				if shift == 0 {
					shift = 1
				}
				i = shift + i
			} else { // nums2 donates
				shift := int(float64(len(nums2)-j-1)/2)
				if shift == 0 {
					shift = 1
				}
				j = shift + j
			}
		} else { // balanced
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

func largestSide(nums1, nums2 []int, i, j, f int) int {
	if i+1+j+1+f > len(nums1)-i-1+len(nums2)-j-1 { // left
		return -1
	} else if i+1+j+1+f < len(nums1)-i-1+len(nums2)-j-1 {
		return 1
	}
	return 0
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


func centerPosition(nums []int) int {
	return int(math.Round(float64(len(nums))/2)) - 1
}
