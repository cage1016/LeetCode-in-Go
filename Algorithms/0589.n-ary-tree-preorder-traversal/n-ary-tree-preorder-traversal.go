/*
* @lc app=leetcode id=589 lang=golang
 */

package problem0589

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

func preorder(root *Node) []int {
	// // recursive
	// if root == nil {
	// 	return []int{}
	// }

	// res := []int{root.Val}
	// for i := 0; i < len(root.Children); i++ {
	// 	res = append(res, preorder(root.Children[i])...)
	// }
	// return res

	// iteration, stack
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*Node{root}
	var top *Node
	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, top.Val)
		for i := len(top.Children) - 1; i >= 0; i-- {
			stack = append(stack, top.Children[i])
		}
	}
	return res
}

// @lc code=end
