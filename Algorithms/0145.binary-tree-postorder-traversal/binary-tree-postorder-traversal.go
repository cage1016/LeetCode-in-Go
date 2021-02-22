/*
* @lc app=leetcode id=%!s(int=145) lang=golang
 */

package problem0145

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
func postorderTraversal(root *TreeNode) []int {
	// // recursive
	// if root == nil {
	// 	return []int{}
	// }

	// res := postorderTraversal(root.Left)
	// res = append(res, postorderTraversal(root.Right)...)
	// res = append(res, []int{root.Val}...)
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
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	return reverse(res)
}

func reverse(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// @lc code=end
