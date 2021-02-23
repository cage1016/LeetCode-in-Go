/*
* @lc app=leetcode id=515 lang=golang
 */

package problem0515

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
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	queue := []*TreeNode{root}

	var first *TreeNode
	for len(queue) > 0 {
		size := len(queue)
		max := -1 << 63
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			if first.Val > max {
				max = first.Val
			}
		}
		res = append(res, max)
	}
	return res
}

// @lc code=end
