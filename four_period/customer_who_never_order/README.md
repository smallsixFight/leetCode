### 从不订购的客户（Customers Who Never Order）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/customers-who-never-order/)

#### 题目
SQL 架构：
```sql
Create table If Not Exists Customers (Id int, Name varchar(255))
Create table If Not Exists Orders (Id int, CustomerId int)
Truncate table Customers
insert into Customers (Id, Name) values ('1', 'Joe')
insert into Customers (Id, Name) values ('2', 'Henry')
insert into Customers (Id, Name) values ('3', 'Sam')
insert into Customers (Id, Name) values ('4', 'Max')
Truncate table Orders
insert into Orders (Id, CustomerId) values ('1', '3')
insert into Orders (Id, CustomerId) values ('2', '1')
```
某网站包含两个表，`Customers`表和`Orders`表。编写一个 SQL 查询，找出所有从不订购任何东西的客户。

`Customer`表：
```
+----+-------+
| Id | Name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+
```
`Order`表：
```
+----+------------+
| Id | CustomerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+
```
例如给定上述表格，你的查询应返回：
```
+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+
```

#### 思路
直接`not in`配上子查询。

#### 答案
```sql
select 
    Name Customers
from
    Customers 
where id not in (
    select CustomerId from Orders
);
```