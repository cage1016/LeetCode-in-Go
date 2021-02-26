/*
* @lc app=leetcode id=%!s(int=113) lang=golang
 */

package problem0113

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
func pathSum(root *TreeNode, targetSum int) [][]int {
	// // method #1 recursive
	// if root == nil {
	// 	return [][]int{}
	// }

	// res := [][]int{}
	// paths := []int{}
	// var dfs func(root *TreeNode, targetSum int, paths []int, res *[][]int)
	// dfs = func(root *TreeNode, targetSum int, paths []int, res *[][]int) {
	// 	if root.Left == nil && root.Right == nil && root.Val == targetSum { // center
	// 		paths = append(paths, root.Val)
	// 		tmp := make([]int, len(paths))
	// 		// copy value only
	// 		copy(tmp, paths)
	// 		*res = append(*res, tmp)
	// 		return
	// 	}

	// 	if root.Left == nil && root.Right == nil {
	// 		return
	// 	}

	// 	if root.Left != nil { // left
	// 		dfs(root.Left, targetSum-root.Val, append(paths, root.Val), res) // targetSum, paths 皆有隱含的 back tracking
	// 	}

	// 	if root.Right != nil { // right
	// 		dfs(root.Right, targetSum-root.Val, append(paths, root.Val), res) // targetSum, paths 皆有隱含的 back tracking
	// 	}
	// }

	// dfs(root, targetSum, paths, &res)
	// return res

	// method #2, iteration, preorder
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	sum := []int{targetSum}
	path := [][]int{[]int{root.Val}}
	var first *TreeNode
	var s int
	var p []int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			s, sum = sum[0], sum[1:]
			p, path = path[0], path[1:]
			if first.Left == nil && first.Right == nil && s == first.Val {
				// add path to res
				tmp := make([]int, len(p))
				copy(tmp, p)
				res = append(res, tmp)
			}

			if first.Left != nil {
				queue = append(queue, first.Left)
				sum = append(sum, s-first.Val)
				tmp := make([]int, len(p)+1)
				copy(tmp, append(p, first.Left.Val))
				path = append(path, tmp)
			}

			if first.Right != nil {
				queue = append(queue, first.Right)
				sum = append(sum, s-first.Val)
				tmp := make([]int, len(p)+1)
				copy(tmp, append(p, first.Right.Val))
				path = append(path, tmp)
			}
		}
	}
	return res
}

// @lc code=end
