### 重新格式化部门表（Reformat Department Table）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/reformat-department-table/)

#### 题目
SQL 架构：
```sql
Create table If Not Exists Department (id int, revenue int, month varchar(5))
Truncate table Department
insert into Department (id, revenue, month) values ('1', '8000', 'Jan')
insert into Department (id, revenue, month) values ('2', '9000', 'Jan')
insert into Department (id, revenue, month) values ('3', '10000', 'Feb')
insert into Department (id, revenue, month) values ('1', '7000', 'Feb')
insert into Department (id, revenue, month) values ('1', '6000', 'Mar')
```

部门表`Department`：
```
+---------------+---------+
| Column Name   | Type    |
+---------------+---------+
| id            | int     |
| revenue       | int     |
| month         | varchar |
+---------------+---------+
(id, month) 是表的联合主键。
这个表格有关于每个部门每月收入的信息。
月份（month）可以取下列值 ["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"]。
```

编写一个 SQL 查询来重新格式化表，使得新的表中有一个部门 id 列和一些对应 每个月 的收入（revenue）列。

查询结果格式如下面的示例所示：

```
Department 表：
+------+---------+-------+
| id   | revenue | month |
+------+---------+-------+
| 1    | 8000    | Jan   |
| 2    | 9000    | Jan   |
| 3    | 10000   | Feb   |
| 1    | 7000    | Feb   |
| 1    | 6000    | Mar   |
+------+---------+-------+

查询得到的结果表：
+------+-------------+-------------+-------------+-----+-------------+
| id   | Jan_Revenue | Feb_Revenue | Mar_Revenue | ... | Dec_Revenue |
+------+-------------+-------------+-------------+-----+-------------+
| 1    | 8000        | 7000        | 6000        | ... | null        |
| 2    | 9000        | null        | null        | ... | null        |
| 3    | null        | 10000       | null        | ... | null        |
+------+-------------+-------------+-------------+-----+-------------+

注意，结果表有 13 列 (1个部门 id 列 + 12个月份的收入列)。
```

#### 思路
这道题的考点是行列转换。

#### 答案
```sql
select
    id, 
    sum(case month WHEN 'Jan' THEN revenue END) Jan_Revenue, 
    sum(case month WHEN 'Feb' THEN revenue END) Feb_Revenue, 
    sum(case month WHEN 'Mar' THEN revenue END) Mar_Revenue, 
    sum(case month WHEN 'Apr' THEN revenue END) Apr_Revenue, 
    sum(case month WHEN 'May' THEN revenue END) May_Revenue, 
    sum(case month WHEN 'Jun' THEN revenue END) Jun_Revenue, 
    sum(case month WHEN 'Jul' THEN revenue END) Jul_Revenue, 
    sum(case month WHEN 'Aug' THEN revenue END) Aug_Revenue, 
    sum(case month WHEN 'Sep' THEN revenue END) Sep_Revenue, 
    sum(case month WHEN 'Oct' THEN revenue END) Oct_Revenue, 
    sum(case month WHEN 'Nov' THEN revenue END) Nov_Revenue, 
    sum(case month WHEN 'Dec' THEN revenue END) Dec_Revenue 
from
    Department
group by id;

```
