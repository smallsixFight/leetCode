### 查找重复的电子邮箱（Duplicate Emails）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/duplicate-emails/)

#### 题目
SQL 架构：
```sql
Create table If Not Exists Person (Id int, Email varchar(255))
Truncate table Person
insert into Person (Id, Email) values ('1', 'a@b.com')
insert into Person (Id, Email) values ('2', 'c@d.com')
insert into Person (Id, Email) values ('3', 'a@b.com')
```

编写一个 SQL 查询，查找`Person`表中所有重复的电子邮箱。

示例：
```
+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
```
根据以上输入，你的查询应返回以下结果：
```
+---------+
| Email   |
+---------+
| a@b.com |
+---------+
```

#### 思路
我想到两种写法，一种是用`having`，另外一种用子查询。

#### 题外话
子查询貌似快一些。

#### 答案
```sql
select 
    Email
from
    Person
group by
    Email
having 
    count(email) > 1;

-- two method
select Email from 
(select Email, count(Email) num from Person group by Email) t1 where t1.num > 1;
```
