# [110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:


<blockquote>
<p>a binary tree in which the left and right subtrees of <em>every</em> node differ in height by no more than 1.</p>
</blockquote>



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

<pre><strong>Input:</strong> root = [3,9,20,null,null,15,7]
<strong>Output:</strong> true
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

<pre><strong>Input:</strong> root = [1,2,2,3,3,null,null,4,4]
<strong>Output:</strong> false
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = []
<strong>Output:</strong> true
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 5000]</code>.
- <code>-10<sup>4</sup> &lt;= Node.val &lt;= 10<sup>4</sup></code>


##  解题思路

因為要計算左右子樹，所以使用後序遍歷 left, right, center
當左右子樹的高度差絕對值>1，我們使用-1來表示

##  可能的變化

