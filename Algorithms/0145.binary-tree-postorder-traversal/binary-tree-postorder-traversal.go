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
	if root == nil {
		return []int{}
	}

	res := postorderTraversal(root.Left)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, []int{root.Val}...)
	return res
}

// @lc code=end
