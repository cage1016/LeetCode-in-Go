/*
* @lc app=leetcode id=101 lang=golang
 */

package problem0101

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// // recursive
	// return IsCompare(root.Left, root.Right)

	// // iteration: queue
	// queue := []*TreeNode{root.Left, root.Right}
	// var left, right *TreeNode
	// for len(queue) > 0 {
	// 	left, queue = queue[0], queue[1:]
	// 	right, queue = queue[0], queue[1:]

	// 	if left == nil && right == nil {
	// 		continue
	// 	}

	// 	if (left == nil || right == nil) || left.Val != right.Val {
	// 		return false
	// 	}

	// 	queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	// }
	// return true

	// iteration: stack
	stack := []*TreeNode{root.Left, root.Right}
	var left, right *TreeNode
	for len(stack) > 0 {
		left, stack = stack[len(stack)-1], stack[:len(stack)-1]
		right, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if left == nil && right == nil {
			continue
		}

		if (left == nil || right == nil) || left.Val != right.Val {
			return false
		}

		stack = append(stack, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}

func IsCompare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}
	return IsCompare(left.Left, right.Right) && IsCompare(left.Right, right.Left)
}

// @lc code=end
