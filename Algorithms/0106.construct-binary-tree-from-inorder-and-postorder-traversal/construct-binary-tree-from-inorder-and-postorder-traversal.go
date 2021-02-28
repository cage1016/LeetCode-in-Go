/*
* @lc app=leetcode id=%!s(int=106) lang=golang
 */

package problem0106

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	idx := Idx(inorder, root.Val)
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])

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
