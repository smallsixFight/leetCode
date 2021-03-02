### 体育馆的人流量（Human Traffic of Stadium）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/human-traffic-of-stadium/)

#### 题目
SQL 架构：
```sql
Create table If Not Exists stadium (id int, visit_date DATE NULL, people int)
Truncate table stadium
insert into stadium (id, visit_date, people) values ('1', '2017-01-01', '10')
insert into stadium (id, visit_date, people) values ('2', '2017-01-02', '109')
insert into stadium (id, visit_date, people) values ('3', '2017-01-03', '150')
insert into stadium (id, visit_date, people) values ('4', '2017-01-04', '99')
insert into stadium (id, visit_date, people) values ('5', '2017-01-05', '145')
insert into stadium (id, visit_date, people) values ('6', '2017-01-06', '1455')
insert into stadium (id, visit_date, people) values ('7', '2017-01-07', '199')
insert into stadium (id, visit_date, people) values ('8', '2017-01-08', '188')
```

 市建了一个新的体育馆，每日人流量信息被记录在这三列信息中：序号 (id)、日期 (visit_date)、 人流量 (people)。

请编写一个查询语句，找出人流量的高峰期。高峰期时，至少连续三行记录中的人流量不少于100。

例如，表`stadium`：
```
+------+------------+-----------+
| id   | visit_date | people    |
+------+------------+-----------+
| 1    | 2017-01-01 | 10        |
| 2    | 2017-01-02 | 109       |
| 3    | 2017-01-03 | 150       |
| 4    | 2017-01-04 | 99        |
| 5    | 2017-01-05 | 145       |
| 6    | 2017-01-06 | 1455      |
| 7    | 2017-01-07 | 199       |
| 8    | 2017-01-08 | 188       |
+------+------------+-----------+
```

对于上面的示例数据，输出为：
```
+------+------------+-----------+
| id   | visit_date | people    |
+------+------------+-----------+
| 5    | 2017-01-05 | 145       |
| 6    | 2017-01-06 | 1455      |
| 7    | 2017-01-07 | 199       |
| 8    | 2017-01-08 | 188       |
+------+------------+-----------+
```

提示：  
每天只有一行记录，日期随着 id 的增加而增加。

#### 思路
先查出符合条件的第一天的数据，然后将其后面三天的数据都查出来，然后在去重。

#### 答案
```sql
select distinct t2.* from (
select t1.id, t1.visit_date, IFNULL((case when t2.people >= 100 then 1 end) + (case when t3.people >= 100 then 1 end), 0) val
from stadium t1
inner join stadium t2
on t2.id = t1.id +1
inner join stadium t3
on t3.id = t1.id +2
where t1.people >= 100) t1, stadium t2
where t1.val > 0 and t2.id <= t1.id +2 and t2.id >= t1.id
order by id;
```