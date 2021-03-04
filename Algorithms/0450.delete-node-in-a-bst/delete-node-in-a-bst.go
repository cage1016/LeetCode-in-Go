/*
* @lc app=leetcode id=%!s(int=450) lang=golang
 */

package problem0450

import (
	"github.com/cage1016/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	// method #1: recursive
	// if root == nil {
	// 	return root
	// }
	// if root.Val == key {
	// 	// one node
	// 	if root.Left == nil {
	// 		return root.Right
	// 	}
	// 	if root.Right == nil {
	// 		return root.Left
	// 	}

	// 	minNode := getMin(root.Right)
	// 	root.Val = minNode.Val
	// 	root.Right = deleteNode(root.Right, minNode.Val)
	// } else if root.Val > key {
	// 	root.Left = deleteNode(root.Left, key)
	// } else if root.Val < key {
	// 	root.Right = deleteNode(root.Right, key)
	// }
	// return root

	// method #2: iteration
	if root == nil {
		return root
	}

	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Val == key {
			break
		}
		pre = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if pre == nil {
		return fnDeleteNode(cur)
	}

	if pre.Left != nil && pre.Left.Val == key {
		pre.Left = fnDeleteNode(cur)
	}
	if pre.Right != nil && pre.Right.Val == key {
		pre.Right = fnDeleteNode(cur)
	}
	return root
}

func fnDeleteNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	cur := root.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	cur.Left = root.Left
	return root.Right
}

func getMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// @lc code=end
