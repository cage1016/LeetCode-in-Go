/*
* @lc app=leetcode id=%!s(int=235) lang=golang
 */

package problem0235

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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// method #1
	// if root.Val > p.Val && root.Val > q.Val { // left
	// 	return lowestCommonAncestor(root.Left, p, q)
	// } else if root.Val < p.Val && root.Val < q.Val { // right
	// 	return lowestCommonAncestor(root.Right, p, q)
	// } else {
	// 	return root
	// }

	// method #2
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

// @lc code=end
