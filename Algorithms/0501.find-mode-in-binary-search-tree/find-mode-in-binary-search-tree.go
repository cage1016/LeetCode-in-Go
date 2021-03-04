/*
* @lc app=leetcode id=501 lang=golang
 */

package problem0501

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
func findMode(root *TreeNode) []int {
	// // method #1: recursive
	// res := []int{}
	// var dfs func(root *TreeNode)
	// pre := -1 << 63
	// var count, maxCount int
	// dfs = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}

	// 	dfs(root.Left) // left

	// 	// center
	// 	if pre == -1<<63 {
	// 		count = 1
	// 	} else if pre == root.Val {
	// 		count++
	// 	} else {
	// 		count = 1
	// 	}

	// 	pre = root.Val

	// 	if count == maxCount {
	// 		res = append(res, root.Val)
	// 	}

	// 	if count > maxCount {
	// 		maxCount = count
	// 		res = []int{root.Val}
	// 	}

	// 	// right
	// 	dfs(root.Right)
	// }
	// dfs(root)

	// return res

	// method #2: iteration
	res := []int{}
	var count, maxCount int
	var top *TreeNode
	var NULL = -1 << 63
	var pre = NULL
	stack := []*TreeNode{}
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]

			if pre == NULL {
				count = 1
			} else if pre == top.Val {
				count++
			} else {
				count = 1
			}

			cur = top.Right
			pre = top.Val

			if count == maxCount {
				res = append(res, top.Val)
			}

			if count > maxCount {
				maxCount = count
				res = []int{top.Val}
			}
		}
	}
	return res
}

// @lc code=end
