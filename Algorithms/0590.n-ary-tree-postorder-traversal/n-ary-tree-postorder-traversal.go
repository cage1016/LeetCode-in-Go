/*
* @lc app=leetcode id=590 lang=golang
 */

package problem0590

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

func postorder(root *Node) []int {
	// // recursive
	// if root == nil {
	// 	return []int{}
	// }

	// res := []int{}
	// for i := 0; i < len(root.Children); i++ {
	// 	res = append(res, postorder(root.Children[i])...)
	// }
	// res = append(res, []int{root.Val}...)
	// return res

	// iteration
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*Node{root}
	var top *Node
	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, top.Val)
		// for loop 由小到大，以符合出 stack 順序為 right left
		for i := 0; i < len(top.Children); i++ {
			stack = append(stack, top.Children[i])
		}
	}
	return reverse(res)
}

func reverse(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// @lc code=end
