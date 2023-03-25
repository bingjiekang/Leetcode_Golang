-- 删除重复的电子邮箱
-- 表: Person
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | email       | varchar |
-- +-------------+---------+
-- id是该表的主键列。
-- 该表的每一行包含一封电子邮件。电子邮件将不包含大写字母。
-- 编写一个 SQL 删除语句来 删除 所有重复的电子邮件，只保留一个id最小的唯一电子邮件。

DELETE A from Person A,Person B where A.email = B.email and A.id > B.id
DELETE A from Person A left join  Person B on A.email = B.email where A.id > B.id
