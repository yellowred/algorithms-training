// 285. Inorder Successor in BST
// medium
// Given a binary search tree and a node in it, find the in-order successor of that node in the BST.

package main

import (
	"fmt"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func main() {
	tree := buildTree([]int{2,1,3}, 0)
	print(tree)
	fmt.Println(inorderSuccessor(tree, tree.Left).Val)
}

var found bool
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    found = false
    return traverse(root, p)
}

func traverse(root, p *TreeNode) *TreeNode {
    var res *TreeNode

    if root == nil {
        return nil
    }

    //left
    res = traverse(root.Left, p)
    if res != nil {
        return res
    }

    // visit
	if root.Val == p.Val {
        found = true
    } else if found {
    	return root
    }

    //right
    return traverse(root.Right, p)
}

func buildTree(tree []int, idx int) *TreeNode {
	if idx>=len(tree) || tree[idx]<0 {
		return nil
	}
	node := TreeNode{
		Val: tree[idx],
		Left: buildTree(tree, (idx+1)*2-1),
		Right: buildTree(tree, (idx+1)*2),
	}
	return &node
}

func print(root *TreeNode) {
	if root != nil {
		log.Println("Node:", root.Val)
		print(root.Left)
		print(root.Right)
	} else {
		log.Println(root)
	}
}
