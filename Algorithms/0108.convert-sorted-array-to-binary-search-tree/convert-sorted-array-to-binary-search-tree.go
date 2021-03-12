/*
* @lc app=leetcode id=108 lang=golang
 */

package problem0108

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
func sortedArrayToBST(nums []int) *TreeNode {
	// // method #1: recursive
	// if len(nums) == 0 {
	// 	return nil
	// }

	// mid := len(nums) / 2
	// root := &TreeNode{Val: nums[mid]}
	// root.Left = sortedArrayToBST(nums[:mid])
	// root.Right = sortedArrayToBST(nums[mid+1:])

	// return root

	// method #2: iteration
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{}
	queue := []*TreeNode{root}
	leftQueue := []int{0}
	rightQueue := []int{len(nums) - 1}
	var first *TreeNode
	var left, right int
	for len(queue) > 0 {
		first, queue = queue[0], queue[1:]
		left, leftQueue = leftQueue[0], leftQueue[1:]
		right, rightQueue = rightQueue[0], rightQueue[1:]
		mid := left + ((right - left) / 2)

		first.Val = nums[mid] // 中間節點

		if left <= mid-1 { // left
			first.Left = &TreeNode{}
			queue = append(queue, first.Left)
			leftQueue = append(leftQueue, left)
			rightQueue = append(rightQueue, mid-1)
		}

		if right >= mid+1 { // right
			first.Right = &TreeNode{}
			queue = append(queue, first.Right)
			leftQueue = append(leftQueue, mid+1)
			rightQueue = append(rightQueue, right)
		}

	}
	return root
}

// @lc code=end
