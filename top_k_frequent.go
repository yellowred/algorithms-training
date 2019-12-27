// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/
// Priority Queue
package main

import (
	"container/heap"
	"log"
)

type Item struct {
	value    int // A number
	priority int // Count of occurencies of the number
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func main() {
	nums := []int{2, 3, 4, 1, 2, 3, 3, 4, 4, 4}
	log.Println(topKFrequent(nums, 3))
}

func topKFrequent(nums []int, k int) []int {

	// frequncy table: O(N)
	f := map[int]int{}
	i := 0
	for i < len(nums) {
		f[nums[i]]++
		i++
	}

	// construct heap: O(N*logN)
	pq := make(PriorityQueue, len(f))
	i = 0
	for el, freq := range f {
		item := Item{el, freq}
		pq[i] = &item
		i++
		log.Println(item)
	}
	heap.Init(&pq)

	// make answer: O(k*logN)
	ans := []int{}
	for i := 1; i <= k; i++ {
		item := heap.Pop(&pq).(*Item)
		ans = append(ans, item.value)
	}

	return ans
}
