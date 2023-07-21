-- # Write your MySQL query statement below
-- 表： Weather

-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | recordDate    | date    |
-- | temperature   | int     |
-- +---------------+---------+
-- id 是这个表的主键
-- 该表包含特定日期的温度信息

-- # datediff(日期1, 日期2)：
-- # 得到的结果是日期1与日期2相差的天数。
-- # 如果日期1比日期2大，结果为正；如果日期1比日期2小，结果为负。

-- # timestampdiff(时间类型, 日期1, 日期2)
-- # 这个函数和上面diffdate的正、负号规则刚好相反。
-- # 日期1大于日期2，结果为负，日期1小于日期2，结果为正。
-- # cross join 是链表查询
select a.id from Weather as a cross join Weather as b on datediff(a.recordDate,b.recordDate) = 1 where a.temperature>b.temperature;