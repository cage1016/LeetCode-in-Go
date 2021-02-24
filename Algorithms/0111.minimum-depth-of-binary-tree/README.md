# [111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/)
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

**Note:** A leaf is a node with no children.



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/10/12/ex_depth.jpg)

<pre><strong>Input:</strong> root = [3,9,20,null,null,15,7]
<strong>Output:</strong> 2
</pre>

**Example 2:**


<pre><strong>Input:</strong> root = [2,null,3,null,4,null,5,null,6]
<strong>Output:</strong> 5
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 10<sup>5</sup>]</code>.
- <code>-1000 &lt;= Node.val &lt;= 1000</code>


##  解题思路

1. recursive

基本概念跟求最大深度有點類似，不過不太一樣，如果直接使用 min 就會有誤區，最小深度都會是1，因為求解是 root 到最低的子節點

左子樹為空，右子樹不為空，最小深度 = 1 + 右子樹深度
左子樹為不空，右子樹為空，最小深度 = 1 + 左子樹深度

如果左右子樹都不為空，= 1 + 左右子樹的最小深度

2. iteration (使用層序遍歷)

每迭代一層深度+1，當當左右子樹都為 nil 時，就是最小深度

##  可能的變化

