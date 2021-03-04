# [450. Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/)
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

- Search for a node to remove.
- If the node is found, delete the node.
**Follow up:** Can you solve it with time complexity <code>O(height of tree)</code>?



**Example 1:**
![](https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg)

<pre><strong>Input:</strong> root = [5,3,6,2,4,null,7], key = 3
<strong>Output:</strong> [5,4,6,2,null,null,7]
<strong>Explanation:</strong> Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it&#39;s also accepted.
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/04/del_node_supp.jpg"/>
</pre>

**Example 2:**


<pre><strong>Input:</strong> root = [5,3,6,2,4,null,7], key = 0
<strong>Output:</strong> [5,3,6,2,4,null,7]
<strong>Explanation:</strong> The tree does not contain a node with value = 0.
</pre>

**Example 3:**


<pre><strong>Input:</strong> root = [], key = 0
<strong>Output:</strong> []
</pre>



**Constraints:**


- The number of nodes in the tree is in the range <code>[0, 10<sup>4</sup>]</code>.
- <code>-10<sup>5</sup> &lt;= Node.val &lt;= 10<sup>5</sup></code>
- Each node has a **unique** value.
- <code>root</code> is a valid binary search tree.
- <code>-10<sup>5</sup> &lt;= key &lt;= 10<sup>5</sup></code>


##  解题思路



##  可能的變化

