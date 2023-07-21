-- Employee 表：
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | salary      | int  |
-- +-------------+------+
-- 在 SQL 中，id 是这个表的主键。
-- 表的每一行包含员工的工资信息。
--  
-- 查询并返回 Employee 表中第二高的薪水 。如果不存在第二高的薪水，查询应该返回 null(Pandas 则返回 None) 。

-- distinct 去重函数，只能在Select语句中使用，要放在所有方法的最前面
-- order by 排序函数，asc 升序排列、desc降序排列
-- limit 获取指定行数据， limit x,y 跳过前x行（从第x行开始获取数据，不包括第x行），然后取y条数据
-- OFFSET子句来代替LIMIT子句，OFFSET子句指定从第几行开始获取记录。OFFSET x limit y 效果一样
-- 解决null问题：
-- 方法1：将表作为临时表select () as xxxx;
-- 方法2：“IFNULL” 函数，select IFNULL() as xxxx;

-- 解决方法:将原表排序后去重,然后选择第二条数据既可
Select (Select distinct salary from Employee order by salary desc limit 1,1) as SecondHighestSalary;

