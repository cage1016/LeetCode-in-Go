# [590. N-ary Tree Postorder Traversal](https://leetcode.com/problems/n-ary-tree-postorder-traversal/)
Given an n-ary tree, return the _postorder_ traversal of its nodes&#39; values.

_Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples)._



**Follow up:**

Recursive solution is trivial, could you do it iteratively?



**Example 1:**

![](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)


<pre><strong>Input:</strong> root = [1,null,3,2,4,null,5,6]
<strong>Output:</strong> [5,6,3,2,4,1]
</pre>

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)


<pre><strong>Input:</strong> root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
<strong>Output:</strong> [2,6,14,11,7,3,12,8,4,13,9,10,5,1]
</pre>



**Constraints:**


- The height of the n-ary tree is less than or equal to <code>1000</code>
- The total number of nodes is between <code>[0, 10^4]</code>


##  解题思路

145 postorder 的 N 子樹變型
概念還是進行翻轉

prorder: center, left, right
postorder: left, right, center

所以使用 preorder 的方式，修改出 stack 的順序為 right, left
再進行整個 array reverse 得到答案

##  可能的變化

