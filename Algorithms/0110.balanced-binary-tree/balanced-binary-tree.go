/*
* @lc app=leetcode id=%!s(int=110) lang=golang
 */

package problem0110

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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if GetDepth(root) == -1 {
		return false
	}
	return true
}

func GetDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := GetDepth(root.Left)   // left
	rightDepth := GetDepth(root.Right) // right

	if leftDepth == -1 {
		return -1
	}

	if rightDepth == -1 {
		return -1
	}

	if abs(leftDepth-rightDepth) > 1 { // center
		return -1
	}
	return max(leftDepth, rightDepth) + 1 // 以當前節點為根節點的最⼤⾼度
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
