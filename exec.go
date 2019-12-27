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
	nums := []int{4, 5, 3, 1, 2}
	queue := []int{}
	for k := range nums {
		PushToQueue(&queue, nums[k])
	}
	log.Println(queue)
	for i := 0; i < 5; i++ {
		log.Println(PopFromQueue(&queue))
	}
}

func PushToQueue(queue *[]int, x int) {
	*queue = append(*queue, x)
	j := len(*queue) - 1
	for {
		i := (j - 1) / 2 // parent
		if i == j || (*queue)[j] < (*queue)[i] {
			break
		}
		(*queue)[j], (*queue)[i] = (*queue)[i], (*queue)[j]
		j = i
	}
}

// 1. Replace the root node with the last element in the heap
// 2. Remove the last element
// 3. Swap (i.e. heapify) the new root with its child until the correct position has found (See MAX-HEAPIFY)
func PopFromQueue(queue *[]int) int {
	x := (*queue)[0]
	(*queue)[0] = (*queue)[len(*queue)-1]
	(*queue) = (*queue)[0 : len(*queue)-1]
	heapify(*queue)
	log.Println(*queue)
	return x
}

func heapify(queue []int) {
	i := 0
	n := len(queue)
	g := i
	for {
		l := i*2 + 1
		r := i*2 + 2
		if l < n && queue[i] < queue[l] {
			g = l
		} else {
			g = i
		}
		if r < n && queue[g] < queue[r] {
			g = r
		}
		if g != i {
			queue[i], queue[g] = queue[g], queue[i]
			i = g
		} else {
			break
		}
	}
}
