# [226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)
Invert a binary tree.

**Example:**

Input:


<pre>     4
   /   \
  2     7
 / \   / \
1   3 6   9</pre>

Output:


<pre>     4
   /   \
  7     2
 / \   / \
9   6 3   1</pre>

**Trivia:**This problem was inspired by [this original tweet](https://twitter.com/mxcl/status/608682016205344768) by [Max Howell](https://twitter.com/mxcl):


<blockquote>Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.</blockquote>


##  解题思路

基本當前點翻轉 swap 子節點
可以前序遍歷/後序遍歷, recursive/iteration

##  可能的變化

