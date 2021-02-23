/*
* @lc app=leetcode id=429 lang=golang
 */

package problem0429

import "github.com/cage1016/LeetCode-in-Go/kit"

type Node kit.Node

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*Node{root}

	var first *Node
	for len(queue) > 0 {
		size := len(queue)
		level := []int{}
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			for j := 0; j < len(first.Children); j++ {
				if first.Children[j] != nil {
					queue = append(queue, first.Children[j])
				}
			}
			level = append(level, first.Val)
		}
		res = append(res, level)
		level = []int{}
	}
	return res
}

// @lc code=end
