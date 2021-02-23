# [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)
Given the <code>root</code> of a binary tree, return _its maximum depth_.

A binary tree&#39;s **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)

<pre><strong>Input:</strong> root = [3,9,20,null,null,15,7]
<strong>Output:</strong> 3
</pre>

**Example 2:**


<pre><strong>Input:</strong> root = [1,null,2]
<strong>Output:</strong> 2
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = []
<strong>Output:</strong> 0
</pre>

**Example 4:**


<pre><strong>Input:</strong> root = [0]
<strong>Output:</strong> 1
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 10<sup>4</sup>]</code>.
- <code>-100 &lt;= Node.val &lt;= 100</code>


##  解题思路

- recursive: 基本上就是計算左右子樹 count, 再取 Max
- iteration: 層序遍歷，每一個 level 加1即可

##  可能的變化

