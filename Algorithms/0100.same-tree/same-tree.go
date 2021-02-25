/*
* @lc app=leetcode id=%!s(int=100) lang=golang
 */

package problem0100

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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// method #1: recursive
	// if p == nil && q == nil {
	// 	return true
	// } else if p == nil || q == nil {
	// 	return false
	// } else if p.Val != q.Val {
	// 	return false
	// }

	// return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

	// method #2: iteration
	queue := []*TreeNode{p, q}
	var leftNode, rightNode *TreeNode
	for len(queue) > 0 {
		leftNode, queue = queue[0], queue[1:]
		rightNode, queue = queue[0], queue[1:]

		if leftNode == nil && rightNode == nil {
			continue
		}

		if (leftNode == nil || rightNode == nil) || leftNode.Val != rightNode.Val {
			return false
		}

		queue = append(queue,
			leftNode.Left, rightNode.Left,
			leftNode.Right, rightNode.Right,
		)
	}

	return true
}

// @lc code=end
