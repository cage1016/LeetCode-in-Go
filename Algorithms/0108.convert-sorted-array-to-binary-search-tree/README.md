# [108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)
Given an integer array <code>nums</code> where the elements are sorted in **ascending order**, convert _it to a **height-balanced** binary search tree_.

A **height-balanced** binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg)

<pre><strong>Input:</strong> nums = [-10,-3,0,5,9]
<strong>Output:</strong> [0,-3,9,-10,null,5]
<strong>Explanation:</strong> [0,-10,5,null,-3,null,9] is also accepted:
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg"/>
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2021/02/18/btree.jpg)

<pre><strong>Input:</strong> nums = [1,3]
<strong>Output:</strong> [3,1]
<strong>Explanation:</strong> [1,3] and [3,1] are both a height-balanced BSTs.
</pre>



**Constraints:**


- <code>1 &lt;= nums.length &lt;= 10<sup>4</sup></code>
- <code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code>
- <code>nums</code> is sorted in a **strictly increasing** order.


##  解题思路

1. recursive

    基本就是操作取中位數，切分左右，分別建立 Tree

##  可能的變化

