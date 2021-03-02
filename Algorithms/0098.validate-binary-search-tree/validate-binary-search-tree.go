/*
* @lc app=leetcode id=%!s(int=98) lang=golang
 */

package problem0098

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
func isValidBST(root *TreeNode) bool {
	// method #1: recursive
	// var dfs func(root, min, max *TreeNode) bool
	// dfs = func(root, min, max *TreeNode) bool {
	// 	if root == nil {
	// 		return true
	// 	}
	// 	if min != nil && root.Val <= min.Val {
	// 		return false
	// 	}
	// 	if max != nil && root.Val >= max.Val {
	// 		return false
	// 	}

	// 	return dfs(root.Left, min, root) || dfs(root.Right, root, max)
	// }

	// return dfs(root, nil, nil)

	// method #2: iteration, BST inorder 是有序 array
	if root == nil {
		return true
	}
	stack := []*TreeNode{}
	cur := root
	var top *TreeNode
	tmp := -1 << 63
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if top.Val <= tmp { // BST 不能重複
				return false
			} else {
				tmp = top.Val
			}
			cur = top.Right
		}
	}
	return true
}

// @lc code=end
