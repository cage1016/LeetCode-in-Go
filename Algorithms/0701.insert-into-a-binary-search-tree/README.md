# [701. Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/)
You are given the <code>root</code> node of a binary search tree (BST) and a <code>value</code> to insert into the tree. Return _the root node of the BST after the insertion_. It is **guaranteed** that the new value does not exist in the original BST.

**Notice** that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return **any of them**.



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/10/05/insertbst.jpg)

<pre><strong>Input:</strong> root = [4,2,7,1,3], val = 5
<strong>Output:</strong> [4,2,7,1,3,5]
<strong>Explanation:</strong> Another accepted tree is:
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/05/bst.jpg"/>
</pre>

**Example 2:**


<pre><strong>Input:</strong> root = [40,20,60,10,30,50,70], val = 25
<strong>Output:</strong> [40,20,60,10,30,50,70,null,null,25]
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
<strong>Output:</strong> [4,2,7,1,3,5]
</pre>



**Constraints:**


- The number of nodes in the tree will be in the range <code>[0, 10<sup>4</sup>]</code>.
- <code>-10<sup>8</sup> &lt;= Node.val &lt;= 10<sup>8</sup></code>
- All the values <code>Node.val</code> are **unique**.
- <code>-10<sup>8</sup> &lt;= val &lt;= 10<sup>8</sup></code>
- It&#39;s **guaranteed** that <code>val</code> does not exist in the original BST.


##  解题思路

1. recursive

    recursive 到 root == nil，就回傳新值

1. iteration

    如同 235 一樣，利用 BST 的方向性由上往下，當左子樹還是右子樹為空時，建立新節點

##  可能的變化

