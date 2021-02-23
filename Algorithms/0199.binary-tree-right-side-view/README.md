# [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)
Given the <code>root</code> of a binary tree, imagine yourself standing on the **right side** of it, return _the values of the nodes you can see ordered from top to bottom_.



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/02/14/tree.jpg)

<pre><strong>Input:</strong> root = [1,2,3,null,5,null,4]
<strong>Output:</strong> [1,3,4]
</pre>

**Example 2:**


<pre><strong>Input:</strong> root = [1,null,3]
<strong>Output:</strong> [1,3]
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = []
<strong>Output:</strong> []
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 100]</code>.
- <code>-100 &lt;= Node.val &lt;= 100</code>


##  解题思路

進行層序遍歷，把最後一個值加到 res 中即可

##  可能的變化

