DROP DATABASE IF EXISTS `tasks`;
CREATE DATABASE `tasks`;
USE `tasks`;

DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
    `id` INT AUTO_INCREMENT,
    `first_name` VARCHAR (128) NOT NULL,
    `last_name` VARCHAR (128) NOT NULL,
    `middle_name` VARCHAR (128) NOT NULL,
    `email` VARCHAR (128) NOT NULL,
    `title` VARCHAR (128) NOT NULL,
    `display_name` VARCHAR (128) NOT NULL,
    `height` VARCHAR (128) NOT NULL,
    `weight` INT,
    `body_fat` INT,
    `target_weight` INT,
    `test_date` DATE,
    `next_test_date` DATE,
    `physical_profile` VARCHAR(64) NOT NULL,
    `promotion_points` INT,
    `total_score` INT,
    `mdl` INT,
    `sdc` INT,
    `plk` INT,
    `mr2` INT,
    `hrp` INT,
    `spt` INT,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `users` (`first_name`, `last_name`, `middle_name`, `email`, `title`, `display_name`, `height`, `weight`, `body_fat`, `target_weight`, `test_date`, `next_test_date`, `physical_profile`, `promotion_points`, `total_score`, `mdl`, `sdc`, `plk`, `mr2`, `hrp`, `spt`)
VALUES ('Jane', 'Smith', 'Teresa', 'rudi.shafakian@lmi.org', 'Soldier', 'Specialist Smith', '60', '117', '20', '127', '2022-03-31', '2022-08-01', 'Permanent', '60', '343', '55', '58', '61', '0', '51', '56');
INSERT INTO `users` (`first_name`, `last_name`, `middle_name`, `email`, `title`, `display_name`, `height`, `weight`, `body_fat`, `target_weight`, `test_date`, `next_test_date`, `physical_profile`, `promotion_points`, `total_score`, `mdl`, `sdc`, `plk`, `mr2`, `hrp`, `spt`)
VALUES ('John', 'Williams', 'Thomas', 'rudi.shafakian@lmi.org', 'Soldier', 'Specialist Williams', '75', '223', '14', '196', '2022-03-31', '2022-08-01', 'Permanent', '60', '343', '55', '58', '61', '0', '51', '56');

CREATE TABLE IF NOT EXISTS `move` (
    `task_number` varchar (50) NOT NULL,
    `task_title` varchar(150) NOT NULL,
    `training_frequency` varchar(50) NOT NULL,
    PRIMARY KEY (`task_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `survive` (
    `task_number` varchar (50) NOT NULL,
    `task_title` varchar(150) NOT NULL,
    `training_frequency` varchar(50) NOT NULL,
    PRIMARY KEY (`task_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `shoot` (
    `task_number` varchar (50) NOT NULL,
    `task_title` varchar(150) NOT NULL,
    `training_frequency` varchar(50) NOT NULL,
    PRIMARY KEY (`task_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `communicate` (
    `task_number` varchar (50) NOT NULL,
    `task_title` varchar(150) NOT NULL,
    `training_frequency` varchar(50) NOT NULL,
    PRIMARY KEY (`task_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT IGNORE INTO `move` VALUES ('071-COM-0541', 'Perform Exterior Movement Techniques', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-0503', 'Move Over, Through, or Around Obstacles (Except Minefields)', 'QT');
INSERT IGNORE INTO `move` VALUES ('071-COM-1000', 'Identify Topographic Symbols on a Military Map', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-1001', 'Identify Terrain Features on a Map', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-1008', 'Measure distance on a Map', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-1002', 'Determine the Grid Coordinates of a Point on a Military Map', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-1005', 'Determine a Location on the Ground by Terrain Association', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-1012', 'Orient a Map to the Ground by Map-Terrain Association', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-1011', 'Orient a Map Using a Lensatic Compass', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-1003', 'Determine a Magnetic Azimuth Using a Lensatic Compass', 'AN');
INSERT IGNORE INTO `move` VALUES ('071-COM-1006', 'Navigate from One Point on the Ground to Another Point While Dismounted', 'SA');
INSERT IGNORE INTO `move` VALUES ('071-COM-0501', 'Move as a Member of a Team', 'SA');
INSERT IGNORE INTO `move` VALUES ('071-COM-0502', 'Move Under Direct Fire', 'SA');
INSERT IGNORE INTO `move` VALUES ('071-COM-0510', 'React to Indirect Fire While Dismounted', 'SA');
INSERT IGNORE INTO `move` VALUES ('071-COM-0513', 'Select Hasty Fighting Positions', 'SA');

INSERT IGNORE INTO `communicate` VALUES ('071-COM-2070', 'Operate SINCGARS Single-Channel (SC)', 'SA');
INSERT IGNORE INTO `communicate` VALUES ('071-COM-1022', 'Perform Voice Communications', 'AN');
INSERT IGNORE INTO `communicate` VALUES ('071-COM-4079', 'Send a Situation Report (SITREP)', 'AN');
INSERT IGNORE INTO `communicate` VALUES ('071-COM-4080', 'Send a Spot Report (SPOTREP)', 'AN');
INSERT IGNORE INTO `communicate` VALUES ('071-COM-0608', 'Use Visual Signaling Techniques', 'SA');

INSERT IGNORE INTO `survive` VALUES ('081-COM-1003', 'Perform First Aid to Clear an Object Stuck in theThroat of a ConsciousCasualty', 'AN');
INSERT IGNORE INTO `survive` VALUES ('081-COM-1005', 'Perform First Aid toPrevent or Control Shock', 'AN');
INSERT IGNORE INTO `survive` VALUES ('081-COM-1023', 'Perform First Aid toRestore Breathing and/or Pulse', 'AN');
INSERT IGNORE INTO `survive` VALUES ('081-COM-1032', 'Perform First Aid for aBleeding and/or Severed Extremity', 'AN');
INSERT IGNORE INTO `survive` VALUES ('081-COM-1046', 'Transport a Casualty', 'AN');
INSERT IGNORE INTO `survive` VALUES ('081-COM-1007', 'Perform First Aid for Burns', 'AN');
INSERT IGNORE INTO `survive` VALUES ('081-COM-1026', 'Perform First Aid for an Open Chest Wound', 'AN');
INSERT IGNORE INTO `survive` VALUES ('081-COM-0101', 'Request Medical Evacuation', 'AN');
INSERT IGNORE INTO `survive` VALUES ('052-COM-1270 ', 'React to Improvised Explosive Device (IED)', 'AN');
INSERT IGNORE INTO `survive` VALUES ('052-COM-1271', 'Identify Visual Indicators of an Improvised Device (IED)', 'AN');
INSERT IGNORE INTO `survive` VALUES ('071-COM-0804', 'Perform Surveillance without the Aid of Electronic Device', 'SA');
INSERT IGNORE INTO `survive` VALUES ('301-COM-1050', 'Report Information of Potential Intelligence Value', 'SASA');
INSERT IGNORE INTO `survive` VALUES ('071-COM-0815', 'Practice , Noise, Light, and Litter Discipline', 'SA');
INSERT IGNORE INTO `survive` VALUES ('071-COM-0801', 'Challenge Persons Entering your Area', 'AN');
INSERT IGNORE INTO `survive` VALUES ('071-COM-0512', 'React to Hand-to-Hand Combat', 'SA');
INSERT IGNORE INTO `survive` VALUES ('071-COM-4408', 'Construct an Individual Fighting Position', 'AN');
INSERT IGNORE INTO `survive` VALUES ('051-COM-1361', 'Camouflage Yourself and Individual Equipment', 'AN');
INSERT IGNORE INTO `survive` VALUES ('181-COM-1001', 'Conduct operations According to the Law of War', 'AN');
INSERT IGNORE INTO `survive` VALUES ('191-COM-0008', 'Search an Individual in a Tactical Environment', 'AN');
INSERT IGNORE INTO `survive` VALUES ('159-COM-2026', 'Identify Combatant and Non-Combatant Personnel & Hybrid Threats', 'AN');

INSERT IGNORE INTO `shoot` VALUES ('071-COM-0032', 'Maintain an M16-series Rifle/M4-series Rifle Carbine', 'AN');
INSERT IGNORE INTO `shoot` VALUES ('071-COM-0029', '071-COM-0029 Perform a Function Check on an M16-Series Rifle/M4 Series Carbine', 'QT');
INSERT IGNORE INTO `shoot` VALUES ('071-COM-0028', 'Load an M16- Series/M4 Series Carbine', 'SA');
INSERT IGNORE INTO `shoot` VALUES ('071-COM-0027', 'Unload an M16- Series/M4 Series Carbine', 'SA');
INSERT IGNORE INTO `shoot` VALUES ('071-COM-0030', 'Endage Targets with an M16- Series/M4 Series Carbine', 'SA');
INSERT IGNORE INTO `shoot` VALUES ('071-COM-0033', 'Correct Malfunctions of an M16- Series/M4 Series Carbine', 'QT');
INSERT IGNORE INTO `shoot` VALUES ('071-COM-0031', 'Zero an M16- Series/M4 Series Carbine', 'SA');
INSERT IGNORE INTO `shoot` VALUES ('071-COM-4401', 'Perform Safety Checks on Hand Grenades', 'AN');
INSERT IGNORE INTO `shoot` VALUES ('071-COM-4407', 'Employ Hand Grenades', 'AN');
