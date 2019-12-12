// Odd Even Linked List
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/784/
// medium
// Simple linked list management

package main

import (
	"log"
	"strconv"
)

func main() {
	head := list("12345")
	l3 := oddEvenList(head)
	printList(l3)
}

func list(str string) *ListNode {
	startNode := ListNode{}
	node := &startNode

	for i:=len(str)-1; i>=0; i-- {
		node.Val, _ = strconv.Atoi(string(str[i]))
		nodeNew := ListNode{}
		if i>0 {
			node.Next = &nodeNew
		}
		node = &nodeNew
	}
	return &startNode
}

func printList(node *ListNode) {
	log.Println("list=")
	log.Println(node)
	acc := ""
	for node != nil {
		str := strconv.Itoa(node.Val)
		acc = acc + str
		node = node.Next
	}
	log.Println(acc)
}

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var lastOdd, cursor *ListNode
    lastOdd = head
    currentEven := true
    cursor = head
    for cursor != nil && cursor.Next != nil {
        if !currentEven {
            // save link to the right element at insert
            firstEven := lastOdd.Next
            
            // save link to the right element at cursor
            next := cursor.Next.Next
            
            // insert
            lastOdd.Next = cursor.Next
            lastOdd.Next.Next = firstEven
            
            // manage cursorS
            lastOdd = lastOdd.Next
            cursor.Next = next
        } else {
            cursor = cursor.Next
        }
        currentEven = !currentEven
        
    }
    return head
}