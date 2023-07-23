-- 表：Logs

-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | num         | varchar |
-- +-------------+---------+
-- id 是这个表的主键。
--  

-- 编写一个 SQL 查询，查找所有至少连续出现三次的数字。

-- 返回的结果表中的数据可以按 任意顺序 排列。
-- 解析:将同一个增加到三份,对比相邻的三个id是否是递增,值是否想同,如果相同则列出来,为了防止4个相同的在一起导致的重复,所以需要去重
select distinct l1.num as ConsecutiveNums
from Logs l1,Logs l2,Logs l3
where l1.id+1 = l2.id and l2.id+1 = l3.id and l1.num = l2.num and l2.num = l3.num;
