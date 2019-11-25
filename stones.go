package main

import (
	"log"
	"math"
)

func main() {
	// log.Println(mergeStones([]int{6, 4, 4, 6}, 2))
	log.Println(mergeStones([]int{69, 39, 79, 78, 16, 6, 36, 97, 79, 27, 14, 31, 4}, 2))
}

func mergeStones(stones []int, K int) int {
	var minCost int
	minCost = math.MaxInt32
	merge(stones, K, 0, &minCost)
	if minCost == math.MaxInt32 {
		return -1
	}
	return minCost
}

func merge(stones []int, k, cost int, minCost *int) {
	log.Println("S1", stones, cost, *minCost)
	if len(stones) == 1 && cost < *minCost {
		*minCost = cost
	}

	if len(stones) < k || math.Mod(float64(len(stones)-1), float64(k-1)) > 0 {
		return
	}

	var sum int
	for i := 0; i <= len(stones)-k; i++ {
		sum = 0
		for j := 0; j < k; j++ {
			sum += stones[i+j]
		}
		newStones := make([]int, len(stones)-k+1)
		copy(newStones, stones[:i])
		newStones[i] = sum
		if i+k < len(stones) {
			copy(newStones[i+1:], stones[i+k:len(stones)])
		}
		log.Println("S2", i, sum, newStones, stones)
		merge(newStones, k, cost+sum, minCost)
	}
}
