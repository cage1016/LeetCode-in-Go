/*
* @lc app=leetcode id=116 lang=golang
 */

package problem0116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	var first *Node
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			//
			if size == 1 || i == size-1 {
				first.Next = nil
			} else {
				first.Next = queue[0]
			}
		}
	}
	return root
}

// @lc code=end
