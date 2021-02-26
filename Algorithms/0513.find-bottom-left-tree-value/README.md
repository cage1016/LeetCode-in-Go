# [513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/)
Given the <code>root</code> of a binary tree, return the leftmost value in the last row of the tree.



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/12/14/tree1.jpg)

<pre><strong>Input:</strong> root = [2,1,3]
<strong>Output:</strong> 1
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2020/12/14/tree2.jpg)

<pre><strong>Input:</strong> root = [1,2,3,4,null,5,6,null,null,7]
<strong>Output:</strong> 7
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[1, 10<sup>4</sup>]</code>.
- <code>-2<sup>31</sup> &lt;= Node.val &lt;= 2<sup>31</sup> - 1</code>


##  解题思路

1. recursive
   
   前序遍歷, center, left, right
   最一層最左邊的節點就是最大深度的節點，每一次 depth > maxDepth 時順便更新一下最大節點的值

   depth++ recursive 的回溯可以使用明顯的方式 depth++/depth-- or 使用隱性的方式直接代入 dfs 方法中

1. iteration

    層序遍歷就很直覺，每層更新最左邊的值 i=0

##  可能的變化

