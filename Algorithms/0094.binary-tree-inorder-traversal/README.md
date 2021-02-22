# [94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)
Given the <code>root</code> of a binary tree, return _the inorder traversal of its nodes&#39; values_.



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg)

<pre><strong>Input:</strong> root = [1,null,2,3]
<strong>Output:</strong> [1,3,2]
</pre>

**Example 2:**


<pre><strong>Input:</strong> root = []
<strong>Output:</strong> []
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [1]
<strong>Output:</strong> [1]
</pre>

**Example 4:**
![](https://assets.leetcode.com/uploads/2020/09/15/inorder_5.jpg)

<pre><strong>Input:</strong> root = [1,2]
<strong>Output:</strong> [2,1]
</pre>

**Example 5:**
![](https://assets.leetcode.com/uploads/2020/09/15/inorder_4.jpg)

<pre><strong>Input:</strong> root = [1,null,2]
<strong>Output:</strong> [1,2]
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 100]</code>.
- <code>-100 &lt;= Node.val &lt;= 100</code>



**Follow up:**

Recursive solution is trivial, could you do it iteratively?




##  解题思路

最艒單的就是直接使用 recursive 來解

##  可能的變化

- iteration
  - DFS