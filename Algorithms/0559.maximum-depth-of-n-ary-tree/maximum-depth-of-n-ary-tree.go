/*
* @lc app=leetcode id=559 lang=golang
 */

package problem0559

type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	// // recursive
	// if root == nil {
	// 	return 0
	// }

	// max := 0
	// for i := 0; i < len(root.Children); i++ {
	// 	max = Max(max, maxDepth(root.Children[i]))
	// }
	// return 1 + max

	// iteration: queue
	if root == nil {
		return 0
	}

	res := 0
	queue := []*Node{root}
	var first *Node
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			first, queue = queue[0], queue[1:]
			for j := 0; j < len(first.Children); j++ {
				if first.Children[j] != nil {
					queue = append(queue, first.Children[j])
				}
			}
		}
		res += 1
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
