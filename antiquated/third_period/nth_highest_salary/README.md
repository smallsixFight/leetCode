### 第 N 高的薪水（Nth Highest Salary）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/nth-highest-salary/)

#### 题目
编写一个 SQL 查询，获取`Employee`表中第 n 高的薪水（Salary）。
```
+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
```
例如上述`Employee`表，n = 2 时，应返回第二高的薪水`200`。如果不存在第 n 高的薪水，那么查询应返回`null`。
```
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+
```

#### 思路
突然冒出来一道数据库的题目，一个逗号写成中文，一直报错，真特么扎心。

#### 答案
```mysql
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET N = N -1;
    RETURN (
      # Write your MySQL query statement below.
      SELECT DISTINCT salary FROM employee ORDER BY salary DESC LIMIT N, 1
    );
END
```
