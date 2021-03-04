# [235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow **a node to be a descendant of itself**).”



**Example 1:**
![](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

<pre><strong>Input:</strong> root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
<strong>Output:</strong> 6
<strong>Explanation:</strong> The LCA of nodes 2 and 8 is 6.
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

<pre><strong>Input:</strong> root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
<strong>Output:</strong> 2
<strong>Explanation:</strong> The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [2,1], p = 2, q = 1
<strong>Output:</strong> 2
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[2, 10<sup>5</sup>]</code>.
- <code>-10<sup>9</sup> &lt;= Node.val &lt;= 10<sup>9</sup></code>
- All <code>Node.val</code> are **unique**.
- <code>p != q</code>
- <code>p</code> and <code>q</code> will exist in the BST.


##  解题思路

不⽤使⽤回溯，⼆叉搜索樹⾃帶⽅向性，可以⽅便的從上向下查找⽬標區間，遇到⽬標區間內的節點，直接返回。 最後給出了對應的迭代法，⼆叉搜索樹的迭代法甚⾄⽐遞歸更容易理解，也是因為其有序性（⾃帶⽅向性），按照 ⽬標區間找就⾏了。

##  可能的變化

