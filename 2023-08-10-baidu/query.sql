use test;

-- mysql> select * from students;
-- +----+------------+--------+-------------+-------+
-- | id | student_id | name   | course      | score |
-- +----+------------+--------+-------------+-------+
-- |  1 |  180512201 | John   | Chinese     |    90 |
-- |  2 |  180512201 | John   | Mathematics |    80 |
-- |  3 |  180512201 | John   | English     |    80 |
-- |  4 |  180512202 | Mary   | Chinese     |    60 |
-- |  5 |  180512202 | Mary   | Mathematics |   100 |
-- |  6 |  180512202 | Mary   | English     |    40 |
-- |  7 |  180512203 | Robert | Chinese     |    50 |
-- |  8 |  180512203 | Robert | Mathematics |    10 |
-- |  9 |  180512203 | Robert | English     |    20 |
-- +----+------------+--------+-------------+-------+

-- 用一条语句查出每个人的总分
SELECT student_id, name, SUM(score) as scores
FROM students
GROUP BY student_id, name
ORDER BY scores DESC;

-- mysql> select * from student_info;
-- +----+------------+--------------------+
-- | id | student_id | id_card            |
-- +----+------------+--------------------+
-- |  1 |  180512201 | 150430199901011699 |
-- |  2 |  180512202 | 150430200102021999 |
-- |  3 |  180512203 | 150430200011112222 |
-- +----+------------+--------------------+

-- 用一条语句查询每个人的姓名和身份证
SELECT DISTINCT name, id_card
FROM students AS a,
     student_info AS b
WHERE a.student_id = b.student_id;