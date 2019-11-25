package main

import (
	"container/heap"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var lists []*ListNode
	for _, list := range [][]int{[]int{1, 4, 5}, []int{1, 3, 4}, []int{2, 6}} {
		lists = append(lists, List(list))
	}
	printList(mergeKLists(lists))
}

func List(list []int) *ListNode {
	i := len(list) - 1
	var current *ListNode
	for i >= 0 {
		item := ListNode{Val: list[i], Next: current}
		log.Println(item)
		current = &item
		i--
	}
	return current
}

func printList(node *ListNode) {
	for node != nil {
		log.Println("NN", node, node.Val)
		node = node.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	pq := PriorityQueue{}
	for _, current := range lists {
		for current != nil {
			item := Item{value: current.Val, priority: current.Val}
			pq = append(pq, &item)
			current = current.Next
		}
	}
	heap.Init(&pq)
	var current *ListNode
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		ln := ListNode{Val: item.value, Next: current}
		current = &ln
	}
	return current
}

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}
