/*
* @lc app=leetcode id=617 lang=golang
 */

package problem0617

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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// // method #1: recursive
	// if root1 == nil {
	// 	return root2
	// }

	// if root2 == nil {
	// 	return root1
	// }

	// root1.Val += root2.Val
	// if root1.Left != nil && root2.Left != nil {
	// 	root1.Left = mergeTrees(root1.Left, root2.Left)
	// } else if root1.Left == nil && root2.Left != nil {
	// 	root1.Left = root2.Left
	// }

	// if root1.Right != nil && root2.Right != nil {
	// 	root1.Right = mergeTrees(root1.Right, root2.Right)
	// } else if root1.Right == nil && root2.Right != nil {
	// 	root1.Right = root2.Right
	// }

	// return root1

	// method #2: iteration, preorder
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	queue := []*TreeNode{root1, root2}
	var top1, top2 *TreeNode
	for len(queue) > 0 {
		top1, queue = queue[0], queue[1:]
		top2, queue = queue[0], queue[1:]

		top1.Val += top2.Val
		if top1.Left != nil && top2.Left != nil {
			queue = append(queue, top1.Left)
			queue = append(queue, top2.Left)
		}

		if top1.Right != nil && top2.Right != nil {
			queue = append(queue, top1.Right)
			queue = append(queue, top2.Right)
		}

		if top1.Left == nil && top2.Left != nil {
			top1.Left = top2.Left
		}

		if top1.Right == nil && top2.Right != nil {
			top1.Right = top2.Right
		}
	}
	return root1
}

// @lc code=end
