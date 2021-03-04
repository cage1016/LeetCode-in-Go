/*
* @lc app=leetcode id=%!s(int=236) lang=golang
 */

package problem0236

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	leftNode := lowestCommonAncestor(root.Left, p, q)   // left
	rightNode := lowestCommonAncestor(root.Right, p, q) // right

	if leftNode != nil && rightNode != nil { // center
		return root
	}

	if leftNode == nil {
		return rightNode
	}
	return leftNode
}

// @lc code=end
