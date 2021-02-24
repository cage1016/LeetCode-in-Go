# [222. Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/)
Given the <code>root</code> of a **complete** binary tree, return the number of the nodes in the tree.

According to **[Wikipedia](http://en.wikipedia.org/wiki/Binary_tree#Types_of_binary_trees)**, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between <code>1</code> and <code>2<sup>h</sup></code> nodes inclusive at the last level <code>h</code>.



**Example 1:**
![](https://assets.leetcode.com/uploads/2021/01/14/complete.jpg)

<pre><strong>Input:</strong> root = [1,2,3,4,5,6]
<strong>Output:</strong> 6
</pre>

**Example 2:**


<pre><strong>Input:</strong> root = []
<strong>Output:</strong> 0
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [1]
<strong>Output:</strong> 1
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 5 * 10<sup>4</sup>]</code>.
- <code>0 &lt;= Node.val &lt;= 5 * 10<sup>4</sup></code>
- The tree is guaranteed to be **complete**.


**Follow up:** Traversing the tree to count the number of nodes in the tree is an easy solution but with <code>O(n)</code> complexity. Could you find a faster algorithm?

##  解题思路

1. recursive

    就是簡單的 recursive binary tree traversal

1. iteration

    一樣是 binary tree traversal，只是換成 iteration 方式

2. recursive + iteration

    complete tree 一定有包含滿二叉樹 + 最一後層沒有滿 ![](/Algorithms/0222.count-complete-tree-nodes/pic.png), 可以分別找到左右子樹的高度，如果左右子樹的高度一樣，就可以直接計算出滿二叉樹的節點數，其他的就 recursive 計算

    ```
    注意(2<<1) 相當於2^2，所以 lh 初始為0
    ```

##  可能的變化

