# [257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/)
Given a binary tree, return all root-to-leaf paths.

**Note:** A leaf is a node with no children.

**Example:**


<pre><strong>Input:</strong>

   1
 /   \
2     3
 \
  5

<strong>Output:</strong> [&#34;1-&gt;2-&gt;5&#34;, &#34;1-&gt;3&#34;]

<strong>Explanation:</strong> All root-to-leaf paths are: 1-&gt;2-&gt;5, 1-&gt;3
</pre>


##  解题思路

1. recursive: 前序遍歷

    貌似沒有看到回溯的邏輯，其實不然，回溯就隱藏在 traversal(cur->left, path + " ->" , result); 中的 path + " ->" 。 每次函數調⽤完，path依然是沒有加上"->" 的，這就是回溯了

    在 root.Left == nil && root.Right == nil 結束條件中處理答案

1. iteration: 前序遍歷

    除了前序遍歷使用的 stack 之外，還需要再加上 path 的 stack
    
    在 root.Left == nil && root.Right == nil 結束條件中處理答案

##  可能的變化

