-- 表： Employee
-- +--------------+---------+
-- | 列名          | 类型    |
-- +--------------+---------+
-- | id           | int     |
-- | name         | varchar |
-- | salary       | int     |
-- | departmentId | int     |
-- +--------------+---------+
-- 在 SQL 中，id是此表的主键。
-- departmentId 是 Department 表中 id 的外键（在 Pandas 中称为 join key）。
-- 此表的每一行都表示员工的 id、姓名和工资。它还包含他们所在部门的 id。
-- 表： Department
-- +-------------+---------+
-- | 列名         | 类型    |
-- +-------------+---------+
-- | id          | int     |
-- | name        | varchar |
-- +-------------+---------+
-- 在 SQL 中，id 是此表的主键列。
-- 此表的每一行都表示一个部门的 id 及其名称。
-- 查找出每个部门中薪资最高的员工。
-- 按 任意顺序 返回结果表。
-- 解析:查找每个部门最高收入的人,先将每个部门最高的收入和对应的部门提取出来,然后在将这三个表连接,一起检索,判断工资和最高工资相同,且部门相同,输出对应部门名称,部门工资最高人员的名字,部门最高的工资
select B.name Department,A.name Employee,A.salary Salary
from Employee A,Department B,(select departmentId id,max(salary) salary
from Employee
group by departmentId) C
where C.id = A.departmentId and C.salary = A.salary and C.id = B.id;
