# [700. Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/)
You are given the <code>root</code> of a binary search tree (BST) and an integer <code>val</code>.

Find the node in the BST that the node&#39;s value equals <code>val</code> and return the subtree rooted with that node. If such a node does not exist, return <code>null</code>.



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/01/12/tree1.jpg)

<pre><strong>Input:</strong> root = [4,2,7,1,3], val = 2
<strong>Output:</strong> [2,1,3]

**Example 2:**
![](https://assets.leetcode.com/uploads/2021/01/12/tree2.jpg)

<pre><strong>Input:</strong> root = [4,2,7,1,3], val = 5
<strong>Output:</strong> []
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[1, 5000]</code>.
- <code>1 &lt;= Node.val &lt;= 10<sup>7</sup></code>
- <code>root</code> is a binary search tree.
- <code>1 &lt;= val &lt;= 10<sup>7</sup></code>


##  解题思路

1. recursive

    - 沒有要遍歷整顆 BST Tree，所以 recursive 就會需要有回傳值 func searchBST(root *TreeNode, val int) *TreeNode
    - 終止條件, root == nil return nil, root.Val == val return root
    - recursive 條件, BST. left tree < root val < right tree, 所以可以依照這個條件來決定搜尋左子樹還是右子樹

1. iteration

    - 需要判斷 root.Val == val，所以可以使用 preorder + queue
    - BST. left tree < root val < right tree, 來決定走左子樹還是右子樹，左右子樹不為 nil 入 queue

##  可能的變化

