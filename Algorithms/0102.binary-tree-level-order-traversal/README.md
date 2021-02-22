# [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)
Given the <code>root</code> of a binary tree, return _the level order traversal of its nodes&#39; values_. (i.e., from left to right, level by level).



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

<pre><strong>Input:</strong> root = [3,9,20,null,null,15,7]
<strong>Output:</strong> [[3],[9,20],[15,7]]
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

size 記錄 current level 有幾個 child node
for loop 只需要跑幾次，並不需要使用其中的 index

##  可能的變化

