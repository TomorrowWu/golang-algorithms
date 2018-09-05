# [解数独](https://leetcode-cn.com/problems/sudoku-solver/description/)

## 题目描述
编写一个程序，通过已填充的空格来解决数独问题

一个数独的解法需遵循如下规则：
1. 数字 1-9 在每一行只能出现一次。
2. 数字 1-9 在每一列只能出现一次。
3. 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


空白格用 '.' 表示。


Note:
- 给定的数独序列只包含数字 1-9 和字符 '.' 。
- 你可以假设给定的数独只有唯一解。
- 给定数独永远是 9x9 形式的。

示例：

![image](http://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

一个数独

![image](http://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png)

## 个人思路
我把数独框中，需要保证数字不重复的3×3小块，称为一个block。
由于数独需要保持每行，每列，每个block中的数字不重复。解题思路如下：
1. 依次往9个block中，分别填写1~9。
1. 如果block中已经存在n了，去填写下一个数。
1. 在可行的空位填好 n 后
    1. 如果后面的填写没有问题，返回true
    1. 如果后面的填写有问题，把n移入下一个可行的位置。
        1. n在这个block中，没有位置放了，返回false
1. 1～9都填写完了，就去填写下一个block分别填写1～9
1. 所有的block都填写完了，结束。

具体过程和细节，见程序及注释。

## 总结

