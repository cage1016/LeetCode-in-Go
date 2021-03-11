/*
* @lc app=leetcode id=669 lang=golang
 */

package problem0669

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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// // method #1: recursive
	// if root == nil {
	// 	return nil
	// }

	// if root.Val < low {
	// 	return trimBST(root.Right, low, high)
	// }

	// if root.Val > high {
	// 	return trimBST(root.Left, low, high)
	// }

	// root.Left = trimBST(root.Left, low, high)
	// root.Right = trimBST(root.Right, low, high)
	// return root

	// method #2: iteration
	if root == nil {
		return root
	}

	// 處理頭節點
	for root != nil && root.Val < low || root != nil && root.Val > high {
		if root.Val < low {
			root = root.Right // < low 往左走
		} else {
			root = root.Left // > high 往左走
		}
	}
	cur := root
	// root 已經在 low, high 區間，處左孩子 < low
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	cur = root

	// root 已經在 low, high 區間，處右孩子 > high
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}

// @lc code=end
