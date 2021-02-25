# [404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/)
Find the sum of all left leaves in a given binary tree.

**Example:**


<pre>    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values <b>9</b> and <b>15</b> respectively. Return <b>24</b>.
</pre>




##  解题思路

先搞清楚定義，左子葉節點
> TreeNode.Left != nil && TreeNode.Left.Left == nil && TreeNode.Left.Right == nil

1. recursive

   因為要判斷子節點，所以需要使用 postorder, left, right, center
   
1. iteration

   前中後遍歷都可以，只要統計出子葉，再加總 TreeNode.Left != nil && TreeNode.Left.Left == nil && TreeNode.Left.Right == nil 的子葉值

##  可能的變化

