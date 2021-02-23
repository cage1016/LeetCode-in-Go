/*
* @lc app=leetcode id=226 lang=golang
 */

package problem0226

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
func invertTree(root *TreeNode) *TreeNode {
	// recursive
	// if root == nil {
	// 	return nil
	// }

	// root.Left, root.Right = root.Right, root.Left
	// invertTree(root.Left)
	// invertTree(root.Right)
	// return root

	// iteration, DFS, preorder
	// if root == nil {
	// 	return nil
	// }

	// stack := []*TreeNode{root}
	// var top *TreeNode
	// for len(stack) > 0 {
	// 	top, stack = stack[len(stack)-1], stack[:len(stack)-1]
	// 	top.Left, top.Right = top.Right, top.Left
	// 	if top.Left != nil {
	// 		stack = append(stack, top.Left)
	// 	}
	// 	if top.Right != nil {
	// 		stack = append(stack, top.Right)
	// 	}
	// }
	// return root

	// iteration, BFS
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var first *TreeNode
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			first.Left, first.Right = first.Right, first.Left
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}
	}
	return root
}

// @lc code=end
