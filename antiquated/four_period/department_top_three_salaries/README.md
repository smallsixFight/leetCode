### 部门工资前三高的所有员工（Department Top Three Salaries）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/department-top-three-salaries/)

#### 题目
SQL 架构：
```sql
Create table If Not Exists Employee (Id int, Name varchar(255), Salary int, DepartmentId int)
Create table If Not Exists Department (Id int, Name varchar(255))
Truncate table Employee
insert into Employee (Id, Name, Salary, DepartmentId) values ('1', 'Joe', '85000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('2', 'Henry', '80000', '2')
insert into Employee (Id, Name, Salary, DepartmentId) values ('3', 'Sam', '60000', '2')
insert into Employee (Id, Name, Salary, DepartmentId) values ('4', 'Max', '90000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('5', 'Janet', '69000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('6', 'Randy', '85000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('7', 'Will', '70000', '1')
Truncate table Department
insert into Department (Id, Name) values ('1', 'IT')
insert into Department (Id, Name) values ('2', 'Sales')
```

`Employee`表包含所有员工信息，每个员工有其对应的工号`Id`，姓名`Name`，工资`Salary`和部门编号`DepartmentId`。

```
+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 85000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
| 5  | Janet | 69000  | 1            |
| 6  | Randy | 85000  | 1            |
| 7  | Will  | 70000  | 1            |
+----+-------+--------+--------------+
```

`Department`表包含公司所有部门的信息。

```
+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+
```

解释：

IT 部门中，Max 获得了最高的工资，Randy 和 Joe 都拿到了第二高的工资，Will 的工资排第三。销售部门（Sales）只有两名员工，Henry 的工资最高，Sam 的工资排第二。

#### 思路
用之前学过的用户变量，然后先将每个部门的所有员工按照工资从高到底排序，在加一个排名字段，超过 3 的直接用 0 表示，之后再套一层查询，将 0 的过滤掉就好。

#### 答案
```sql
select  d.Name Department, temp.Name Employee, temp.Salary Salary
from (
     select Name, Salary, DepartmentId,
            case when @preDept = e.DepartmentId and @count < 3 and @preSal != e.Salary and @preSal := e.Salary then @count := @count + 1
            when @preDept = e.DepartmentId and @count <= 3 and @preSal = e.Salary then @count
            when @preDept != e.DepartmentId and @preDept := null and @preSal := null then @count := 1
            when @preDept = e.DepartmentId and @count > 2 then 0
                when (@preDept := e.DepartmentId) is not null and (@preSal := e.Salary) is not null then @count := 1
            else @count := 0 end as total
     from Employee e,
          (select @preDept := null, @preSal := null, @count := 1) t order by e.DepartmentId, e.Salary desc
) temp
inner join Department d
on d.Id = temp.DepartmentId
where temp.total != 0;
```