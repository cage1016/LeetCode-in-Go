# [501. Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)
Given a binary search tree (BST) with duplicates, find all the [mode(s)](https://en.wikipedia.org/wiki/Mode_(statistics)) (the most frequently occurred element) in the given BST.

Assume a BST is defined as follows:


- The left subtree of a node contains only nodes with keys **less than or equal to** the node&#39;s key.
- The right subtree of a node contains only nodes with keys **greater than or equal to** the node&#39;s key.
- Both the left and right subtrees must also be binary search trees.



For example:Given BST <code>[1,null,2,2]</code>,


<pre>   1
    \
     2
    /
   2
</pre>



return <code>[2]</code>.

**Note:** If a tree has more than one mode, you can return them in any order.

**Follow up:** Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).


##  解题思路

1. recursive

    BST inorder 為有序數列，所以可以以 inorder recursive 並使用 pre, cur 二個指針加上 count/maxCount 來判斷

1. iteration

    BST inorder iteration 遍歷，並使用 pre, cur 二個指針加上 count/maxCount 來判斷

##  可能的變化

