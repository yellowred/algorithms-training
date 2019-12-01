// find max profit in stock prices array using Divide and Conquer algorithm design pattern
package main

import (
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	num := int(1e+7)
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr = append(arr, 1000+rand.Intn(4000))
	}
	startTime := time.Now()
	log.Println(dc(arr))
	endTime := time.Now()
	log.Println("Time:   ", endTime.Sub(startTime))
}

func bf(arr []int) int {

	max := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[i] > max {
				max = arr[j] - arr[i]
			}
		}
	}
	return max
}

func dc(arr []int) int {
	if len(arr) > 2 {
		var firstHalf = make([]int, len(arr)-len(arr)/2)
		var secondHalf = make([]int, len(arr)/2)
		copy(firstHalf, arr)
		copy(secondHalf, arr[len(arr)-len(arr)/2:len(arr)])
		max := math.Max(float64(dc(firstHalf)), float64(dc(secondHalf)))
		minIdx := 0
		maxIdx := 0
		for i := range arr {
			if arr[i] < arr[minIdx] {
				minIdx = i
			}
			if arr[i] > arr[maxIdx] {
				maxIdx = i
			}
		}
		if minIdx < len(arr)-len(arr)/2 && maxIdx > len(arr)/2 { // buy and sell in different subarrays
			max = float64(arr[maxIdx] - arr[minIdx])
		}
		return int(max)
	} else if len(arr) == 2 {
		return arr[1] - arr[0]
	}
	return 0
}
