/*
* @lc app=leetcode id=701 lang=golang
 */

package problem0701

import "github.com/cage1016/LeetCode-in-Go/kit"

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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// method #1: recursive
	// if root == nil {
	// 	return &TreeNode{Val: val}
	// }

	// if root.Val > val {
	// 	root.Left = insertIntoBST(root.Left, val)
	// } else {
	// 	root.Right = insertIntoBST(root.Right, val)
	// }
	// return root

	// method #2: iteration
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root
	for cur != nil {
		if cur.Val > val {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur.Left = &TreeNode{Val: val}
				break
			}
		} else if cur.Val < val {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur.Right = &TreeNode{Val: val}
				break
			}
		}
	}
	return root
}

// @lc code=end
