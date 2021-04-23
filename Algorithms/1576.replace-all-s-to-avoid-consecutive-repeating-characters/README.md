# [1576. Replace All ?&#39;s to Avoid Consecutive Repeating Characters](https://leetcode.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/)
Given a string <code>s</code>containing only lower case English letters and the &#39;?&#39; character, convert **all **the &#39;?&#39; characters into lower case letters such that the final string does not contain any **consecutive repeating **characters. You **cannot **modify the non &#39;?&#39; characters.

It is **guaranteed **that there are no consecutive repeating characters in the given string **except **for &#39;?&#39;.

Return the final string after all the conversions (possibly zero) have been made. If there is more than one solution, return any of them. It can be shown that an answer is always possible with the given constraints.



**Example 1:**


<pre><strong>Input:</strong> s = &#34;?zs&#34;
<strong>Output:</strong> &#34;azs&#34;
<strong>Explanation</strong>: There are 25 solutions for this problem. From &#34;azs&#34; to &#34;yzs&#34;, all are valid. Only &#34;z&#34; is an invalid modification as the string will consist of consecutive repeating characters in &#34;zzs&#34;.</pre>

**Example 2:**


<pre><strong>Input:</strong> s = &#34;ubv?w&#34;
<strong>Output:</strong> &#34;ubvaw&#34;
<strong>Explanation</strong>: There are 24 solutions for this problem. Only &#34;v&#34; and &#34;w&#34; are invalid modifications as the strings will consist of consecutive repeating characters in &#34;ubvvw&#34; and &#34;ubvww&#34;.
</pre>

**Example 3:**


<pre><strong>Input:</strong> s = &#34;j?qg??b&#34;
<strong>Output:</strong> &#34;jaqgacb&#34;
</pre>

**Example 4:**


<pre><strong>Input:</strong> s = &#34;??yw?ipkj?&#34;
<strong>Output:</strong> &#34;acywaipkja&#34;
</pre>



**Constraints:**


- <code>1 &lt;= s.length &lt;= 100</code>
- <code>s</code> contains only lower case English letters and <code>&#39;?&#39;</code>.


##  解题思路



##  可能的變化

