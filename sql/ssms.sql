-- phpMyAdmin SQL Dump
-- version 4.8.4
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2018-12-28 11:00:58
-- 服务器版本： 5.7.24
-- PHP 版本： 7.2.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `ssms`
--

-- --------------------------------------------------------

--
-- 表的结构 `clazz`
--

CREATE TABLE `clazz` (
  `id` int(11) NOT NULL,
  `clazz_name` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `gradeid` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `clazz`
--

INSERT INTO `clazz` (`id`, `clazz_name`, `gradeid`) VALUES
(1, '计科163', 3),
(2, '计科162', 3);

-- --------------------------------------------------------

--
-- 表的结构 `clazz_course_teacher`
--

CREATE TABLE `clazz_course_teacher` (
  `id` int(11) NOT NULL,
  `clazzid` int(11) DEFAULT NULL,
  `gradeid` int(11) DEFAULT NULL,
  `courseid` int(11) DEFAULT NULL,
  `teacherid` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `clazz_course_teacher`
--

INSERT INTO `clazz_course_teacher` (`id`, `clazzid`, `gradeid`, `courseid`, `teacherid`) VALUES
(1, 1, 3, 1, 1),
(2, 1, 3, 2, 2),
(3, 1, 3, 3, 3),
(4, 1, 3, 4, 4);

-- --------------------------------------------------------

--
-- 表的结构 `course`
--

CREATE TABLE `course` (
  `id` int(11) NOT NULL,
  `name` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `course`
--

INSERT INTO `course` (`id`, `name`) VALUES
(1, '算法'),
(2, '信息安全原理与技术'),
(3, 'java se'),
(4, 'java web'),
(5, '软件工程');

-- --------------------------------------------------------

--
-- 表的结构 `escore`
--

CREATE TABLE `escore` (
  `id` int(11) NOT NULL,
  `examid` int(11) DEFAULT NULL,
  `studentid` int(11) DEFAULT NULL,
  `courseid` int(11) DEFAULT NULL,
  `score` int(5) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `escore`
--

INSERT INTO `escore` (`id`, `examid`, `studentid`, `courseid`, `score`) VALUES
(1, 1, 1, 1, 100),
(2, 1, 1, 2, 100),
(3, 1, 1, 3, 80),
(4, 1, 1, 3, 70);

-- --------------------------------------------------------

--
-- 表的结构 `exam`
--

CREATE TABLE `exam` (
  `id` int(11) NOT NULL,
  `exam_name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `time` varchar(40) CHARACTER SET utf8 DEFAULT NULL,
  `gradeid` int(11) DEFAULT NULL,
  `classzid` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `exam`
--

INSERT INTO `exam` (`id`, `exam_name`, `time`, `gradeid`, `classzid`) VALUES
(1, '期中考试', '2018年-2019年第一学期', 3, 1),
(2, '期中考试', '2018年-2019年第一学期', 3, 2),
(3, '期末考试', '2016年-2017年第二学期', NULL, NULL);

-- --------------------------------------------------------

--
-- 表的结构 `grade`
--

CREATE TABLE `grade` (
  `id` int(11) NOT NULL,
  `grade_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `grade`
--

INSERT INTO `grade` (`id`, `grade_name`) VALUES
(1, '2014级'),
(2, '2015级'),
(3, '2016级'),
(4, '2017级');

-- --------------------------------------------------------

--
-- 表的结构 `student`
--

CREATE TABLE `student` (
  `id` int(11) NOT NULL,
  `number` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `sex` varchar(4) COLLATE utf8_unicode_ci DEFAULT NULL,
  `phone` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `qq` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `clazzid` int(11) DEFAULT NULL,
  `gradeid` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `student`
--

INSERT INTO `student` (`id`, `number`, `name`, `sex`, `phone`, `qq`, `clazzid`, `gradeid`) VALUES
(1, '20161514301', '张金高', '女', '13733823945', '2181550591@qq.com', 1, 3),
(2, '20161514302', '李鑫', '男', '13733823956', '2181550598@qq.com', 1, 3);

-- --------------------------------------------------------

--
-- 表的结构 `teacher`
--

CREATE TABLE `teacher` (
  `id` int(11) NOT NULL,
  `number` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `sex` varchar(4) COLLATE utf8_unicode_ci DEFAULT NULL,
  `phone` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `qq` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `teacher`
--

INSERT INTO `teacher` (`id`, `number`, `name`, `sex`, `phone`, `qq`) VALUES
(1, '12345678901', '朱家义', '男', '13733823947', '2181550599'),
(2, '12345678904', '金松林', '男', '13733823949', '2181550597');

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `account` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `password` varchar(90) COLLATE utf8_unicode_ci DEFAULT '111111',
  `name` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `type` tinyint(1) DEFAULT '2'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `account`, `password`, `name`, `type`) VALUES
(1, '20161514301', '46f94c8de14fb36680850768ff1b7f2a', '张金高', 2),
(2, '20161514302', '46f94c8de14fb36680850768ff1b7f2a', '李鑫', 2),
(3, '12345678901', '46f94c8de14fb36680850768ff1b7f2a', '朱家义', 1),
(4, '12345678904', '46f94c8de14fb36680850768ff1b7f2a', '金松林', 1),
(5, '55555555555', '46f94c8de14fb36680850768ff1b7f2a', '孙记风', 3),
(6, '88888888888', '46f94c8de14fb36680850768ff1b7f2a', '徐红', 3),
(7, '66666666666', '46f94c8de14fb36680850768ff1b7f2a', '周硕', 3);

--
-- 转储表的索引
--

--
-- 表的索引 `clazz`
--
ALTER TABLE `clazz`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `clazz_course_teacher`
--
ALTER TABLE `clazz_course_teacher`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `course`
--
ALTER TABLE `course`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `escore`
--
ALTER TABLE `escore`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `exam`
--
ALTER TABLE `exam`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `grade`
--
ALTER TABLE `grade`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `student`
--
ALTER TABLE `student`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `teacher`
--
ALTER TABLE `teacher`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `clazz`
--
ALTER TABLE `clazz`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `clazz_course_teacher`
--
ALTER TABLE `clazz_course_teacher`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `escore`
--
ALTER TABLE `escore`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `exam`
--
ALTER TABLE `exam`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `grade`
--
ALTER TABLE `grade`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `student`
--
ALTER TABLE `student`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `teacher`
--
ALTER TABLE `teacher`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
