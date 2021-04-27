# [669. Trim a Binary Search Tree](https://leetcode.com/problems/trim-a-binary-search-tree/)
Given the <code>root</code> of a binary search tree and the lowest and highest boundaries as <code>low</code> and <code>high</code>, trim the tree so that all its elements lies in <code>[low, high]</code>. Trimming the tree should **not** change the relative structure of the elements that will remain in the tree (i.e., any node&#39;s descendant should remain a descendant). It can be proven that there is a **unique answer**.

Return _the root of the trimmed binary search tree_. Note that the root may change depending on the given bounds.



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/09/09/trim1.jpg)

<pre><strong>Input:</strong> root = [1,0,2], low = 1, high = 2
<strong>Output:</strong> [1,null,2]
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2020/09/09/trim2.jpg)

<pre><strong>Input:</strong> root = [3,0,4,null,2,null,null,1], low = 1, high = 3
<strong>Output:</strong> [3,2,null,1]
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [1], low = 1, high = 2
<strong>Output:</strong> [1]
</pre>

**Example 4:**


<pre><strong>Input:</strong> root = [1,null,2], low = 1, high = 3
<strong>Output:</strong> [1,null,2]
</pre>

**Example 5:**


<pre><strong>Input:</strong> root = [1,null,2], low = 2, high = 4
<strong>Output:</strong> [2]
</pre>



**Constraints:**


- The number of nodes in the tree in the range <code>[1, 10<sup>4</sup>]</code>.
- <code>0 &lt;= Node.val &lt;= 10<sup>4</sup></code>
- The value of each node in the tree is **unique**.
- <code>root</code> is guaranteed to be a valid binary search tree.
- <code>0 &lt;= low &lt;= high &lt;= 10<sup>4</sup></code>


##  解题思路



##  可能的變化
