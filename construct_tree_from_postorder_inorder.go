// 105. Construct Binary Tree from Preorder and Inorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// medium

package main

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	print(tree)
}

var idx int

func buildTree(preorder []int, inorder []int) *TreeNode {
	idx = -1
	return helper(0, len(inorder), preorder, inorder)
}

func helper(start, end int, preorder, inorder []int) *TreeNode {
	log.Println(idx, start, end)
	if start >= end {
		return nil
	}
	idx++
	root := TreeNode{Val: preorder[idx]}
	root.Left = helper(start, find(root.Val, inorder), preorder, inorder)
	root.Right = helper(find(root.Val, inorder)+1, end, preorder, inorder)
	return &root
}

func find(el int, arr []int) int {
	for k := range arr {
		if arr[k] == el {
			return k
		}
	}
	return -1
}

func print(root *TreeNode) {
	if root != nil {
		log.Println(root.Val)
		print(root.Left)
		print(root.Right)
	} else {
		log.Println(root)
	}
}
