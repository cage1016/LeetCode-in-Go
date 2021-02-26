/*
* @lc app=leetcode id=%!s(int=112) lang=golang
 */

package problem0112

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
func hasPathSum(root *TreeNode, targetSum int) bool {
	// method #1: recursive
	// if root == nil {
	// 	return false
	// }

	// if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 { // center
	// 	return true
	// }

	// return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) // left, right

	// method #2: iteration, preorder traversal
	if root == nil {
		return false
	}
	paths := []int{targetSum}
	stack := []*TreeNode{root}
	var top *TreeNode
	var ts int
	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		ts, paths = paths[len(paths)-1], paths[:len(paths)-1]
		if top.Left == nil && top.Right == nil && ts-top.Val == 0 {
			return true
		}
		// center

		if top.Right != nil { // right
			stack = append(stack, top.Right)
			paths = append(paths, ts-top.Val)
		}

		if top.Left != nil { // left
			stack = append(stack, top.Left)
			paths = append(paths, ts-top.Val)
		}
	}
	return false
}

// @lc code=end
