# [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.



**Example 1:**


<pre><strong>Input:</strong> nums = [1,3,5,6], target = 5
<strong>Output:</strong> 2
</pre>

**Example 2:**


<pre><strong>Input:</strong> nums = [1,3,5,6], target = 2
<strong>Output:</strong> 1
</pre>

**Example 3:**


<pre><strong>Input:</strong> nums = [1,3,5,6], target = 7
<strong>Output:</strong> 4
</pre>

**Example 4:**


<pre><strong>Input:</strong> nums = [1,3,5,6], target = 0
<strong>Output:</strong> 0
</pre>

**Example 5:**


<pre><strong>Input:</strong> nums = [1], target = 0
<strong>Output:</strong> 0
</pre>



**Constraints:**


- <code>1 &lt;= nums.length &lt;= 10<sup>4</sup></code>
- <code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code>
- <code>nums</code> contains **distinct** values sorted in **ascending** order.
- <code>-10<sup>4</sup> &lt;= target &lt;= 10<sup>4</sup></code>


##  解题思路

先列出可能的目標值
- 目標值陣列之前
- 目標值等於陣列中某一個元素
- 目標值插入陣列中
- 目標值在陣列之後

1. 暴力法
2. 因為是有序的陣列，可以考慮使用 binary search

    使用 binary search 就必需注意區間的設置, 採左封右封/左封右開 即其對應的區間 left, right 的邊界

##  可能的變化

