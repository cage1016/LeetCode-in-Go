/*
* @lc app=leetcode id=700 lang=golang
 */

package problem0700

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
func searchBST(root *TreeNode, val int) *TreeNode {
	// method #1: recursive
	// if root == nil {
	// 	return nil
	// }

	// if val == root.Val {
	// 	return root
	// }

	// if val > root.Val {
	// 	// right subtree
	// 	return searchBST(root.Right, val)
	// } else {
	// 	// left subtree
	// 	return searchBST(root.Left, val)
	// }

	// method #2: iteration
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var first *TreeNode
	for len(queue) > 0 {
		first, queue = queue[0], queue[1:]
		if first.Val == val {
			return first
		}

		if val > first.Val {
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		} else if val < first.Val {
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
		}
	}
	return nil
}

// @lc code=end
