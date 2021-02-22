/*
* @lc app=leetcode id=%!s(int=144) lang=golang
 */

package problem0144

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
func preorderTraversal(root *TreeNode) []int {
	// recursive
	// if root == nil {
	// 	return []int{}
	// }

	// res := []int{root.Val}
	// res = append(res, preorderTraversal(root.Left)...)
	// res = append(res, preorderTraversal(root.Right)...)
	// return res

	// # iteration DFS, stack
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*TreeNode{root}

	var top *TreeNode
	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

// @lc code=end
