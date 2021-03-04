# [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow **a node to be a descendant of itself**).”



**Example 1:**
![](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)

<pre><strong>Input:</strong> root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
<strong>Output:</strong> 3
<strong>Explanation:</strong> The LCA of nodes 5 and 1 is 3.
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)

<pre><strong>Input:</strong> root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
<strong>Output:</strong> 5
<strong>Explanation:</strong> The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [1,2], p = 1, q = 2
<strong>Output:</strong> 1
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[2, 10<sup>5</sup>]</code>.
- <code>-10<sup>9</sup> &lt;= Node.val &lt;= 10<sup>9</sup></code>
- All <code>Node.val</code> are **unique**.
- <code>p != q</code>
- <code>p</code> and <code>q</code> will exist in the tree.


##  解题思路

1. 求最⼩公共祖先，需要從底向上遍歷，那麼⼆叉樹，只能通過後序遍歷（即：回溯）實現從低向上的遍歷⽅ 式。

2. 在回溯的過程中，必然要遍歷整顆⼆叉樹，即使已經找到結果了，依然要把其他節點遍歷完，因為要使⽤遞歸 函數的返回值（也就是代碼中的left和right）做邏輯判斷。

3. 要理解如果返回值left為空，right不為空為什麼要返回right，為什麼可以⽤返回right傳給上⼀層結果。

本題沒有給出迭代法，因為迭代法不適合模擬回溯的過程。理解遞歸的解法就夠了。

##  可能的變化

