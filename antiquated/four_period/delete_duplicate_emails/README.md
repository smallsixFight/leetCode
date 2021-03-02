### 删除重复的电子邮箱（Delete Duplicate Emails）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/delete-duplicate-emails/)

#### 题目
编写一个 SQL 查询，来删除`Person`表中所有重复的电子邮箱，重复的邮箱里只保留 **Id** 最小的那个。
```
+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
```
例如，在运行你的查询语句之后，上面的 Person 表应返回以下几行:
```
+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+
```

提示：

- 执行 SQL 之后，输出是整个`Person`表。
- 使用`delete`语句。


#### 思路
第一个想法是，先查出重复且 id 不是最小的 id，然后进行删除。

然而可以转换一下，找出每个邮箱的最小 id，剩下都可以删除。

#### 题外话
这道题思路是简单的，但我还真不知道`Min`这些函数查出来的数据不能直接作用于 delete，还需要套多一层，学到了。

#### 答案
```sql
delete from Person where Id not in (
select Id from (
select MIN(Id) Id from Person group by Email) t
);
```