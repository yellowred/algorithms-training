// 94. Binary Tree Inorder Traversal
// https://leetcode.com/problems/binary-tree-inorder-traversal/
// medium
// Morris 1979
// http://codingrecipies.blogspot.com/2013/11/morris-inorder-traversal.html
// https://reader.elsevier.com/reader/sd/pii/0167642388900639?token=85D7F1F4F011AB443CEA96471C64F28C1654D56B9B8496DCD0A4311DCF16D7ECC71F86BAC6256505DF3977ADCEFA96F4

package main

import (
	"log"
)

func main() {
	log.Println("OK")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	curr := root
	var pre *TreeNode

	for curr != nil {
		if curr.Left == nil {
			res = append(res, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = curr
				curr = curr.Left
			} else {
				pre.Right = nil
				res = append(res, curr.Val)
				curr = curr.Right
			}
		}
	}

	return res
}
