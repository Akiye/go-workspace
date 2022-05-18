CREATE TABLE `users` (
  `id` int,
  `first_name` varchar (50) NOT NULL,
   `last_name` varchar (50) NOT NULL,
  `middle_name` varchar (50) NOT NULL,
  `email` varchar (50) NOT NULL,
  `title` varchar (50) NOT NULL,
  `display_name` varchar (50) NOT NULL,
  `height` varchar (50) NOT NULL,
  `weight` int,
  `body_fat` int,
  `target_weight` int,
  `test_date` date,
  `next_test_date` date,
  `physical_profile` varchar(12) NOT NULL,
`promotion_points` int,
`total_score` int,
`mdl` int,
`sdc` int,
`plk` int,
`mr2` int,
`hrp` int,
`spt` int
);
INSERT INTO `users` VALUES (4,'Jane','Smith','Teresa','rudi.shafakian@lmi.org','Soldier','Specialist Smith','60','117','20', '127','2022-03-31','2022-08-01','Permanent','60','343','55','58','61','0','51','56');
INSERT INTO `users` VALUES (5,'John','Williams','Thomas','rudi.shafakian@lmi.org','Soldier','Specialist Williams','75','223','14', '196','2022-03-31','2022-08-01','Permanent','60','343','55','58','61','0','51','56');
SELECT * FROM users;