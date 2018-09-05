# [Two Sum](https://leetcode-cn.com/problems/two-sum/description/)

##  两数相加
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

Example:
```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```
## [官方图解](https://leetcode-cn.com/problems/add-two-numbers/solution/)

## 个人思路
- 每一位上的两数x+y,临时变量carry记录对下一位的进位(比如满10进1)
- 当前位,sum=x+y+carry,sum % 10 即当前位的值,存到node节点中,carry = sum / 10即进下一位
- 注意:根据最后一位的carry,大于0时,比两个链表的长度多一位

