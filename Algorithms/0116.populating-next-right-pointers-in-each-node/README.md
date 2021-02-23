# [116. Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)
You are given a **perfect binary tree** where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:


<pre>struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
</pre>

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to <code>NULL</code>.

Initially, all next pointers are set to <code>NULL</code>.



**Follow up:**


- You may only use constant extra space.
- Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.



**Example 1:**

![](https://assets.leetcode.com/uploads/2019/02/14/116_sample.png)


<pre><strong>Input:</strong> root = [1,2,3,4,5,6,7]
<strong>Output:</strong> [1,#,2,3,#,4,5,6,7,#]
<strong>Explanation: </strong>Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with &#39;#&#39; signifying the end of each level.
</pre>



**Constraints:**


- The number of nodes in the given tree is less than <code>4096</code>.
- <code>-1000 &lt;= node.val &lt;= 1000</code>


##  解题思路

還是層序遍歷的概念

##  可能的變化

