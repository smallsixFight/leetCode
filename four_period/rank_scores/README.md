### 分数排名（Rank Scores）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/rank-scores/)

#### 题目
SQL 架构：
```sql
Create table If Not Exists Scores (Id int, Score DECIMAL(3,2))
Truncate table Scores
insert into Scores (Id, Score) values ('1', '3.5')
insert into Scores (Id, Score) values ('2', '3.65')
insert into Scores (Id, Score) values ('3', '4.0')
insert into Scores (Id, Score) values ('4', '3.85')
insert into Scores (Id, Score) values ('5', '4.0')
insert into Scores (Id, Score) values ('6', '3.65')
```

编写一个 SQL 查询来实现分数排名。

如果两个分数相同，则两个分数排名（Rank）相同。请注意，平分后的下一个名次应该是下一个连续的整数值。换句话说，名次之间不应该有“间隔”。

```
+----+-------+
| Id | Score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+
```

例如，根据上述给定的`Scores`表，你的查询应该返回（按分数从高到低排列）：
```
+-------+------+
| Score | Rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+
```

#### 思路
首先，相同的分数排名是一样的，那么可以想到需要去重然后排序来获取该分数的排名。

#### 思路二
在看打看看到开窗函数的使用，速度真的快了很多，待学习。

#### 答案
```sql
select 
    Score, (select count(distinct score) from Scores where score >= t.score) 'Rank' 
from
    Scores t
order by
    score desc;

-- 使用开窗函数的答案
select score ,dense_rank() over(order by score desc) as 'rank' 
from scores;
```
