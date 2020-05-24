### 第二高的薪水（Second Highest Salary）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/second-highest-salary/)

#### 题目
SQL 架构：
```sql
Create table If Not Exists Employee (Id int, Salary int)
Truncate table Employee
insert into Employee (Id, Salary) values ('1', '100')
insert into Employee (Id, Salary) values ('2', '200')
insert into Employee (Id, Salary) values ('3', '300')
```

编写一个 SQL 查询，获取`Employee`表中第二高的薪水（Salary）。

```
+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
```
例如上述`Employee`表，SQL 查询应该返回`200`作为第二高的薪水。如果不存在第二高的薪水，那么查询应返回`null`。
```
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
```


#### 思路
这道题需要考虑去重以及不存在的情况下要返回`null`。

#### 答案
```sql
select (
    select distinct Salary 
    from Employee 
    order by Salary desc 
    limit 1, 1 ) SecondHighestSalary 
from dual;
```

