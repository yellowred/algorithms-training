package main

import (
	"fmt"
	"log"
	"strings"
)

type TestData struct {
	A []int
	B []int
	K int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	testData := []TestData{
		TestData{A: []int{1, 2}, B: []int{}, K: 2},
		TestData{A: []int{1, 2, 3}, B: []int{}, K: 2},
	}
	for _, testSet := range testData {
		log.Println(testSet)
		var head *ListNode
		var prev *ListNode
		for _, elem := range testSet.A {
			newNode := ListNode{Val: elem}
			if head == nil {
				head = &newNode
			}
			if prev != nil {
				prev.Next = &newNode
				prev = prev.Next
			} else {
				prev = &newNode
			}
			// log.Println(newNode, head, prev)
		}
		printList(head)
		start := reverseKGroup(head, testSet.K)
		printList(start)
	}
	log.Println("Finish")
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var start *ListNode
	curHead := head
	curTail := head

	for {
		i := 1
		curTail = curHead
		log.Println("I", i, curTail)
		for i <= k && curTail.Next != nil {
			curTail = curTail.Next
			i++
		}

		log.Println("I", i, curTail)
		// not enough elements remain in the list
		if i < k {
			return start
		}

		nextHead := curTail.Next
		reverseK(curHead, k)
		// printList(curTail)
		curHead.Next = nextHead
		if start == nil {
			start = curTail
		}
	}
}

func reverseK(start *ListNode, k int) {
	log.Println("RevK", start, k)
	a := start
	var prev *ListNode
	for k > 0 {
		tmp := a.Next
		a.Next = prev
		prev = a
		a = tmp
		k--
	}
}

func printList(start *ListNode) {
	a := []string{}
	for start != nil {
		a = append(a, fmt.Sprintf("%d", start.Val))
		start = start.Next
	}
	log.Println("List:", strings.Join(a, "->"))
}
