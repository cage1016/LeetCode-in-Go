/*
* @lc app=leetcode id=%!s(int=104) lang=golang
 */

package problem0104

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
func maxDepth(root *TreeNode) int {
	// method #1: recursive
	// if root == nil {
	// 	return 0
	// }

	// return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))

	// // method #2: iteration, 層序遍歷
	// if root == nil {
	// 	return 0
	// }

	// res := 0
	// queue := []*TreeNode{root}
	// var first *TreeNode
	// for len(queue) > 0 {
	// 	size := len(queue)
	// 	for i := 0; i < size; i++ {
	// 		first, queue = queue[0], queue[1:]
	// 		if first.Left != nil {
	// 			queue = append(queue, first.Left)
	// 		}
	// 		if first.Right != nil {
	// 			queue = append(queue, first.Right)
	// 		}
	// 	}
	// 	res += 1
	// }
	// return res

	// method #3

	res := 0
	if root == nil {
		return res
	}
	depth := 1
	GetDepth(root, depth, &res)
	return res
}

func GetDepth(root *TreeNode, depth int, res *int) {
	if depth > *res {
		*res = depth
	}

	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Left != nil {
		depth++
		GetDepth(root.Left, depth, res)
		depth--
	}

	if root.Right != nil {
		depth++
		GetDepth(root.Right, depth, res)
		depth++
	}
	return
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
