// 253. Meeting Rooms II
// https://leetcode.com/problems/meeting-rooms-ii/
// Sort by starting time. Maintain BST with ending time being the priority.
// First ending being removed and the new one comes.
// Maximum size of the priority queue is the required number of meeting rooms.
package main

import (
	"container/heap"
	"log"
	"sort"
)

func main() {
	// log.Println(minMeetingRooms([][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}))
	// log.Println(minMeetingRooms([][]int{[]int{26, 29}, []int{19, 26}, []int{19, 28}, []int{4, 19}, []int{4, 25}})) //3
	// log.Println(minMeetingRooms([][]int{[]int{2, 15}, []int{36, 45}, []int{9, 29}, []int{16, 23}, []int{4, 9}})) //2
	// log.Println(minMeetingRooms([][]int{[]int{6, 17}, []int{8, 9}, []int{11, 12}, []int{6, 9}})) //3
}

type Meetings [][]int

func (m Meetings) Len() int {
	return len(m)
}

func (m Meetings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Meetings) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

type Item struct {
	start    int
	priority int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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

func minMeetingRooms(intervals [][]int) int {
	mm := Meetings(intervals)
	sort.Sort(mm)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	log.Println(mm)
	max := len(pq)
	for i := 0; i < len(intervals); i++ {
		for len(pq) > 0 && pq[0].priority <= intervals[i][0] {
			_ = heap.Pop(&pq)
		}
		item := Item{intervals[i][0], intervals[i][1]}
		heap.Push(&pq, &item)
		if len(pq) > max {
			max = len(pq)
		}
	}
	return max
}
