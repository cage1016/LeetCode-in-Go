# [107. Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/)
Given the <code>root</code> of a binary tree, return _the bottom-up level order traversal of its nodes&#39; values_. (i.e., from left to right, level by level from leaf to root).



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

<pre><strong>Input:</strong> root = [3,9,20,null,null,15,7]
<strong>Output:</strong> [[15,7],[9,20],[3]]
</pre>

**Example 2:**


<pre><strong>Input:</strong> root = [1]
<strong>Output:</strong> [[1]]
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = []
<strong>Output:</strong> []
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 2000]</code>.
- <code>-1000 &lt;= Node.val &lt;= 1000</code>


##  解题思路

基本的思路就是層序遍歷的 reverse

要熟記簡單的 array reverse

##  可能的變化

