# [617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/)
You are given two binary trees <code>root1</code> and <code>root2</code>.

Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

Return _the merged tree_.

**Note:** The merging process must start from the root nodes of both trees.



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/02/05/merge.jpg)

<pre><strong>Input:</strong> root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
<strong>Output:</strong> [3,4,5,5,4,null,7]
</pre>

**Example 2:**


<pre><strong>Input:</strong> root1 = [1], root2 = [1,2]
<strong>Output:</strong> [2,2]
</pre>



**Constraints:**


- The number of nodes in both trees is in the range <code>[0, 2000]</code>.
- <code>-10<sup>4</sup> &lt;= Node.val &lt;= 10<sup>4</sup></code>


##  解题思路

1. recursive

    - recursive 參數, fn(root1 *TreeNode, root2 *TreeNode) *TreeNode
    - 終止條件
        t1 == nil, return t2
        t2 == nil, return t1
    - 值相加
    - recursive 構建左右子樹

1. iteration

    在求二叉樹對稱時就需要同時比較二顆樹。同時入 queue, dequeue 來比較

##  可能的變化

