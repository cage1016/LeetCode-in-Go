# [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)
Given the <code>root</code> of a binary tree, _check whether it is a mirror of itself_ (i.e., symmetric around its center).



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg)

<pre><strong>Input:</strong> root = [1,2,2,3,4,4,3]
<strong>Output:</strong> true
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg)

<pre><strong>Input:</strong> root = [1,2,2,null,3,null,3]
<strong>Output:</strong> false
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[1, 1000]</code>.
- <code>-100 &lt;= Node.val &lt;= 100</code>


**Follow up:** Could you solve it both recursively and iteratively?

##  解题思路

- recursive: 確認單層的 logic, 就是比較二顆樹
  left.left 比 right.right, left.right 比 right.left
  單層節點中先排除空節點，一個為空節點就是 false，再來比較值

- iteration: queue / stack 都可以, 保持進出一致就可以了
  
##  可能的變化

