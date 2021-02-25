/*
* @lc app=leetcode id=257 lang=golang
 */

package problem0257

import (
	"fmt"

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
func binaryTreePaths(root *TreeNode) []string {
	// method #1: recursive
	// if root == nil {
	// 	return []string{}
	// }

	// var res = []string{}
	// dfs(root, "", &res)
	// return res

	// method #2: iteration, preorder
	if root == nil {
		return []string{}
	}

	stack := []*TreeNode{root}
	paths := []string{fmt.Sprintf("%d", root.Val)}
	res := []string{}
	var first *TreeNode
	var path string
	for len(stack) > 0 {
		first, stack = stack[len(stack)-1], stack[:len(stack)-1] //center
		path, paths = paths[len(paths)-1], paths[:len(paths)-1]
		if first.Left == nil && first.Right == nil {
			// print path
			res = append(res, path)
		}

		if first.Right != nil { // right
			stack = append(stack, first.Right)
			paths = append(paths, path+"->"+fmt.Sprintf("%d", first.Right.Val))
		}
		if first.Left != nil { // left
			stack = append(stack, first.Left)
			paths = append(paths, path+"->"+fmt.Sprintf("%d", first.Left.Val))
		}
	}
	return res
}

func dfs(root *TreeNode, path string, res *[]string) {
	path += fmt.Sprintf("%d", root.Val) // center
	if root.Left == nil && root.Right == nil {
		// left node, add ans
		*res = append(*res, path)
		return
	}

	if root.Left != nil { // left
		dfs(root.Left, path+"->", res)
	}
	if root.Right != nil { // right
		dfs(root.Right, path+"->", res)
	}
}

// @lc code=end
