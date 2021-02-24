/*
* @lc app=leetcode id=%!s(int=111) lang=golang
 */

package problem0111

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
func minDepth(root *TreeNode) int {
	// // recursive
	// if root == nil {
	// 	return 0
	// }

	// ld := minDepth(root.Left)
	// rd := minDepth(root.Right)

	// if root.Left == nil && root.Right != nil {
	// 	return 1 + rd
	// }

	// if root.Left != nil && root.Right == nil {
	// 	return 1 + ld
	// }

	// return 1 + Min(ld, rd)

	// iteration, level
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}
	var first *TreeNode
	for len(queue) > 0 {
		size := len(queue)
		depth++
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			if first.Left == nil && first.Right == nil {
				return depth
			}
		}
	}
	return depth
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
