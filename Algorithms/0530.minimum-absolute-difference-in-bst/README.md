# [530. Minimum Absolute Difference in BST](https://leetcode.com/problems/minimum-absolute-difference-in-bst/)
Given a binary search tree with non-negative values, find the minimum [absolute difference](https://en.wikipedia.org/wiki/Absolute_difference) between values of any two nodes.

**Example:**


<pre><b>Input:</b>

   1
    \
     3
    /
   2

<b>Output:</b>
1

<b>Explanation:</b>
The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
</pre>



**Note:**


- There are at least two nodes in this BST.
- This question is the same as 783: [https://leetcode.com/problems/minimum-distance-between-bst-nodes/](https://leetcode.com/problems/minimum-distance-between-bst-nodes/)


##  解题思路

1. recursive

   利用 BST inorder 是有序數列的特性，可以再 recursive 的過程中二二相減取最小

1. iteration

   也是利用 inorder 是有序數列的特性，可以再 iteration 的過程中二二相減取最小

##  可能的變化

