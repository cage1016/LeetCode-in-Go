# [559. Maximum Depth of N-ary Tree](https://leetcode.com/problems/maximum-depth-of-n-ary-tree/)
Given a n-ary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

_Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples)._



**Example 1:**

![](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)


<pre><strong>Input:</strong> root = [1,null,3,2,4,null,5,6]
<strong>Output:</strong> 3
</pre>

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)


<pre><strong>Input:</strong> root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
<strong>Output:</strong> 5
</pre>



**Constraints:**


- The depth of the n-ary tree is less than or equal to <code>1000</code>.
- The total number of nodes is between <code>[0, 10<sup>4</sup>]</code>.


##  解题思路

- recursive: 基本上就是 104 的變型，二叉樹變縑 n 叉樹
- iteration: 基本上也就是 N 叉樹的層序遍歷，每多一個 level +1

##  可能的變化

