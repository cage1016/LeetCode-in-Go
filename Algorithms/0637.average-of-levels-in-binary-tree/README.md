# [637. Average of Levels in Binary Tree](https://leetcode.com/problems/average-of-levels-in-binary-tree/)Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.
**Example 1:**


<pre><b>Input:</b>
    3
   / \
  9  20
    /  \
   15   7
<b>Output:</b> [3, 14.5, 11]
<b>Explanation:</b>
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
</pre>



**Note:**

- The range of node&#39;s value is in the range of 32-bit signed integer.



##  解题思路

層序遍歷，先加總每一層的值，再進行平均

##  可能的變化

