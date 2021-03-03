/*
* @lc app=leetcode id=530 lang=golang
 */

package problem0530

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
func getMinimumDifference(root *TreeNode) int {
	// // method #1: recursive
	// if root == nil {
	// 	return 0
	// }

	// var pre *TreeNode
	// var dfs func(root *TreeNode, array *[]int, res *int)
	// dfs = func(root *TreeNode, array *[]int, res *int) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	dfs(root.Left, array, res)        // left
	// 	*array = append(*array, root.Val) // center
	// 	if pre != nil {
	// 		*res = Min(*res, Abs((*array)[len((*array))-1]-(*array)[len(*array)-2]))
	// 	}
	// 	pre = root
	// 	dfs(root.Right, array, res) // right
	// }

	// var array []int
	// res := 1<<63 - 1 // MAX
	// dfs(root, &array, &res)
	// return res

	// method #2: iteration, inorder
	if root == nil {
		return 0
	}

	var pre *TreeNode
	stack := []*TreeNode{}
	cur := root
	res := 1<<63 - 1 // MAX
	array := []int{}
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
			array = append(array, cur.Val)
			if pre != nil {
				res = Min(res, Abs(array[len(array)-1]-array[len(array)-2]))
			}
			cur = cur.Right
			pre = cur
		}
	}
	return res
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
