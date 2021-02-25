/*
* @lc app=leetcode id=404 lang=golang
 */

package problem0404

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
func sumOfLeftLeaves(root *TreeNode) int {
	// method #1: recursive, postorder
	// if root == nil {
	// 	return 0
	// }

	// leftSum := sumOfLeftLeaves(root.Left)                                    // left
	// rightSum := sumOfLeftLeaves(root.Right)                                  // right
	// if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { // center
	// 	return leftSum + rightSum + root.Left.Val
	// }
	// return leftSum + rightSum

	// method #2: iteration
	res := 0
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	var first *TreeNode
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			if first.Left != nil && first.Left.Left == nil && first.Left.Right == nil {
				res += first.Left.Val
			}

			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}
	}
	return res
}

// @lc code=end
