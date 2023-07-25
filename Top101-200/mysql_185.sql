-- 表: Employee
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | id           | int     |
-- | name         | varchar |
-- | salary       | int     |
-- | departmentId | int     |
-- +--------------+---------+
-- Id是该表的主键列。
-- departmentId是Department表中ID的外键。
-- 该表的每一行都表示员工的ID、姓名和工资。它还包含了他们部门的ID。
 
-- 表: Department
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | name        | varchar |
-- +-------------+---------+
-- Id是该表的主键列。
-- 该表的每一行表示部门ID和部门名。
 

-- 公司的主管们感兴趣的是公司每个部门中谁赚的钱最多。一个部门的 高收入者 是指一个员工的工资在该部门的 不同 工资中 排名前三 。
-- 编写一个SQL查询，找出每个部门中 收入高的员工 。

--解析：利用滑动窗口根据所属部门进行分组，按照工资进行降序排列，得到一个表，然后对新表进行选择性查询即可得到工资前三高的员工。  

select Department,Employee,Salary
from 
	(select D.name Department,E.name Employee,E.salary Salary,dense_rank() over(partition by E.departmentId order by E.salary desc) as sm
		from Employee E,Department D
		where E.departmentId = D.id) as newtables
where sm<=3;