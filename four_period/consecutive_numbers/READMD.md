### 连续出现的数字（Consecutive Numbers）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/consecutive-numbers/)

#### 题目
SQL 架构：
```sql
Create table If Not Exists Logs (Id int, Num int)
Truncate table Logs
insert into Logs (Id, Num) values ('1', '1')
insert into Logs (Id, Num) values ('2', '1')
insert into Logs (Id, Num) values ('3', '1')
insert into Logs (Id, Num) values ('4', '2')
insert into Logs (Id, Num) values ('5', '1')
insert into Logs (Id, Num) values ('6', '2')
insert into Logs (Id, Num) values ('7', '2')
```

编写一个 SQL 查询，查找所有至少连续出现三次的数字。
```
+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+
```

例如，给定上面的`Logs`表，`1`是唯一连续出现至少三次的数字。
```
+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
```

#### 思路
这道题考的连续出现次数的数字，我确实不知道该怎么做，虽然能够猜到怎么去做，但真的不会写，答案参考来自评论区的 [山水有相逢2333
的答案](https://leetcode-cn.com/problems/consecutive-numbers/comments/)。学习到了。

#### 答案
```sql
select distinct Num as ConsecutiveNums
from (
  select Num, 
    case 
      when @prev = Num then @count := @count + 1
      when (@prev := Num) is not null then @count := 1
    end as total
  from Logs, (select @prev := null, @count := null) as t
) as temp
where temp.total >= 3;
```
