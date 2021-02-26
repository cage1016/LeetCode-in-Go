# [113. Path Sum II](https://leetcode.com/problems/path-sum-ii/)
Given the <code>root</code> of a binary tree and an integer <code>targetSum</code>, return all **root-to-leaf** paths where each path&#39;s sum equals <code>targetSum</code>.

A **leaf** is a node with no children.



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg)

<pre><strong>Input:</strong> root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
<strong>Output:</strong> [[5,4,11,2],[5,8,4,5]]
</pre>

**Example 2:**
![](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)

<pre><strong>Input:</strong> root = [1,2,3], targetSum = 5
<strong>Output:</strong> []
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [1,2], targetSum = 0
<strong>Output:</strong> []
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 5000]</code>.
- <code>-1000 &lt;= Node.val &lt;= 1000</code>
- <code>-1000 &lt;= targetSum &lt;= 1000</code>


##  解题思路

1. recursive

    需要遍歷 TreeNode, 所以 dfs func 不需要回傳值
    除了要計算 Path Sum 還要記錄有所過的 Path
    dfs(root.Left, targetSum-root.Val, append(paths, root.Val), res) // targetSum, paths 皆有隱含的 back tracking，再後在 root.Left == nil && root.Right == nil && root.Val == targetSum 中將 Path 存入 res

    > 需要注意 golang slice copy reference 問題

1. iteration: 層序遍歷

    使用 queue/sum/path 來記錄每一個節點所需的值
    > 需要注意 golang slice copy reference 問題

##  可能的變化

