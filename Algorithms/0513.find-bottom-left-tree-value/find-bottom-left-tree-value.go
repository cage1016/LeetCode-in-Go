/*
* @lc app=leetcode id=513 lang=golang
 */

package problem0513

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
func findBottomLeftValue(root *TreeNode) int {
	// method #1 recuresive
	// if root == nil {
	// 	return 0
	// }

	// maxDepthVal := root.Val
	// res := 0
	// dfs(root, &maxDepthVal, 0, &res)
	// return maxDepthVal

	// method #2 iteration
	res := 0
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	var first *TreeNode
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			if i == 0 {
				res = first.Val
			}
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}
	}
	return res
}

func dfs(root *TreeNode, maxDepthVal *int, depth int, res *int) {
	if root.Left == nil && root.Right == nil { // center
		if depth > *res {
			*res = depth
			*maxDepthVal = root.Val
		}
		return
	}

	if root.Left != nil { // left
		// depth++
		dfs(root.Left, maxDepthVal, depth+1, res) // 隱含 back tracking
		// depth--
	}

	if root.Right != nil { // right
		// depth++
		dfs(root.Right, maxDepthVal, depth+1, res)
		// depth--
	}
}

// @lc code=end
