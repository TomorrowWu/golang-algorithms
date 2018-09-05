# [无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/)

## 题目描述
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。

## [官方解答](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/)

## 个人思路
- 两个重复字符之间的子串,即形成的有效子串
- 使用map存已出现的字符以及该字符在数组中的索引

