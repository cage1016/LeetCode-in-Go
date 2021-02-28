/*
* @lc app=leetcode id=%!s(int=105) lang=golang
 */

package problem0105

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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	idx := Idx(inorder, root.Val)
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return root
}

func Idx(inorder []int, value int) int {
	for k, v := range inorder {
		if v == value {
			return k
		}
	}
	return 0
}

// @lc code=end
