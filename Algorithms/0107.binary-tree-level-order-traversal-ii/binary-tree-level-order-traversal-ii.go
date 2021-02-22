/*
* @lc app=leetcode id=107 lang=golang
 */

package problem0107

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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	var first *TreeNode
	for len(queue) > 0 {
		size := len(queue)
		level := []int{}
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			level = append(level, first.Val)
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
		}
		res = append(res, level)
		level = []int{}
	}
	return reverse(res)
}

func reverse(a [][]int) [][]int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// @lc code=end
