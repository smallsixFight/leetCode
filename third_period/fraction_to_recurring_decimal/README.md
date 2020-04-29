### 分数到小数（Fraction to Recurring Decimal）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/fraction-to-recurring-decimal/)

#### 题目
给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。

如果小数部分为循环小数，则将循环的部分括在括号内。

示例 1：
```
输入: numerator = 1, denominator = 2
输出: "0.5"
```
示例 2：
```
输入: numerator = 2, denominator = 1
输出: "2"
```
示例 3：
```
输入: numerator = 2, denominator = 3
输出: "0.(6)"
```

#### 思路
由于整数都是有理数，简单的四则运算不会得到无理数（查了资料），所以这道题的问题就在于当结果是无限循环小数时，找到循环的部分。

然后利用除法、求余数，使用 hashMap 分别对不同数字的索引进行记录以及对应的余数进行记录，只有数字相同和有相同的字符串并且余数相同，这才算找到循环的部分。具体看代码吧。

#### 伪代码
数学知识不够好，面向测试用例写代码，不写伪代码了。