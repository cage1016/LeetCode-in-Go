# [589. N-ary Tree Preorder Traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/)
Given an n-ary tree, return the _preorder_ traversal of its nodes&#39; values.

_Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples)._



**Follow up:**

Recursive solution is trivial, could you do it iteratively?



**Example 1:**

![](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)


<pre><strong>Input:</strong> root = [1,null,3,2,4,null,5,6]
<strong>Output:</strong> [1,3,5,6,2,4]
</pre>

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)


<pre><strong>Input:</strong> root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
<strong>Output:</strong> [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
</pre>



**Constraints:**


- The height of the n-ary tree is less than or equal to <code>1000</code>
- The total number of nodes is between <code>[0, 10^4]</code>


##  解题思路

- recursive 直接解
- iteration 方式，stack 方式為LIFO, 所以要注意入 stack 方式需由大到小，出 stack 方式才會符合 preorder

##  可能的變化

