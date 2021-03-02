# [98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)
Given the <code>root</code> of a binary tree, _determine if it is a valid binary search tree (BST)_.

A **valid BST** is defined as follows:


- The left subtree of a node contains only nodes with keys **less than** the node&#39;s key.
- The right subtree of a node contains only nodes with keys **greater than** the node&#39;s key.
- Both the left and right subtrees must also be binary search trees.



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg)

<pre><strong>Input:</strong> root = [2,1,3]
<strong>Output:</strong> true
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg)

<pre><strong>Input:</strong> root = [5,1,4,null,null,3,6]
<strong>Output:</strong> false
<strong>Explanation:</strong> The root node&#39;s value is 5 but its right child&#39;s value is 4.
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[1, 10<sup>4</sup>]</code>.
- <code>-2<sup>31</sup> &lt;= Node.val &lt;= 2<sup>31</sup> - 1</code>


##  解题思路

1. recursive

    沒有遍歷整個 Tree, 需是尋找樹中是否合法，所以 recursive 的 func 有回傳值
    BST 的基本定義是 left node < root < right node，所以在 recursive 中需要帶入 min, max 來檢查

1. iteration

    BST inorder 代表是有序數列，並不需要先轉出整個有序數列再判斷，可以在 iteration 中就檢查是否有為有序 (ps. <= bst 不能有重複的值)

##  可能的變化

