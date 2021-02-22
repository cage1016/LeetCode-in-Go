/*
* @lc app=leetcode id=%!s(int=94) lang=golang
 */

package problem0094

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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := inorderTraversal(root.Left)
	res = append(res, []int{root.Val}...)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// @lc code=end
