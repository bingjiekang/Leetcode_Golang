-- 查找重复的电子邮箱
-- 表: Person
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | email       | varchar |
-- +-------------+---------+
-- Person 表:
-- +----+---------+
-- | id | email   |
-- +----+---------+
-- | 1  | a@b.com |
-- | 2  | c@d.com |
-- | 3  | a@b.com |
-- +----+---------+
-- id 是该表的主键列。
-- 此表的每一行都包含一封电子邮件。电子邮件不包含大写字母。
-- 编写一个 SQL 查询来报告所有重复的电子邮件。 请注意，可以保证电子邮件字段不为 NULL。

select Email from Person group by email having count(email)>1

