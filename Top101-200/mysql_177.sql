-- 表: Employee

-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | salary      | int  |
-- +-------------+------+
-- 在 SQL 中，id 是该表的主键。
-- 该表的每一行都包含有关员工工资的信息。
--  
-- 查询 Employee 表中第 n 高的工资。如果没有第 n 个最高工资，查询结果应该为 null 。

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    set N := N-1;
  RETURN (
      # Write your MySQL query statement below.
    select
        salary
    from 
        Employee 
    # 或者直接在上面去重也可以 distinct ，就不用group by了 这是分组后找第二个也可以的
    group by
        salary
    order by 
        salary desc 
    limit N,1
  );
END
