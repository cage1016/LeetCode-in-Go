/*
* @lc app=leetcode id=%!s(int=222) lang=golang
 */

package problem0222

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
func countNodes(root *TreeNode) int {
	// method #1: recursive
	// if root == nil {
	// 	return 0
	// }

	// return 1 + countNodes(root.Left) + countNodes(root.Right)

	// // iteration:
	// if root == nil {
	// 	return 0
	// }

	// method #2: basic tree node traversal
	// res := 0
	// queue := []*TreeNode{root}
	// var first *TreeNode
	// for len(queue) > 0 {
	// 	size := len(queue)
	// 	for i := 0; i < size; i++ {
	// 		first, queue = queue[0], queue[1:]
	// 		res++
	// 		if first.Left != nil {
	// 			queue = append(queue, first.Left)
	// 		}
	// 		if first.Right != nil {
	// 			queue = append(queue, first.Right)
	// 		}
	// 	}
	// }
	// return res

	// method #3
	if root == nil {
		return 0
	}

	lt := root.Left
	rt := root.Right
	var lh, rh int
	for lt != nil {
		lh++
		lt = lt.Left
	}
	for rt != nil {
		rh++
		rt = rt.Right
	}

	if lh == rh {
		return (2 << lh) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// @lc code=end
