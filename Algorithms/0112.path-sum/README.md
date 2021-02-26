# [112. Path Sum](https://leetcode.com/problems/path-sum/)
Given the <code>root</code> of a binary tree and an integer <code>targetSum</code>, return <code>true</code> if the tree has a **root-to-leaf** path such that adding up all the values along the path equals <code>targetSum</code>.

A **leaf** is a node with no children.



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg)

<pre><strong>Input:</strong> root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
<strong>Output:</strong> true
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)

<pre><strong>Input:</strong> root = [1,2,3], targetSum = 5
<strong>Output:</strong> false
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [1,2], targetSum = 0
<strong>Output:</strong> false
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 5000]</code>.
- <code>-1000 &lt;= Node.val &lt;= 1000</code>
- <code>-1000 &lt;= targetSum &lt;= 1000</code>


##  解题思路

1. recursive

    結束條件為子葉節點(root.Left == nil && root.Right == nil) 再介斷 targetSum - root.Val 是否為0，recursive 中 targetSum - root.Val 則為隱性的 back tracking

1. iteration

    一樣使用前序遍歷，使用 stack, 另外需要額外記錄每一個節點的 TargetSum

##  可能的變化

