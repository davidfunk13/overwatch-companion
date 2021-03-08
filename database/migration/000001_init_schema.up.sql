CREATE TABLE `battletag` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` int NOT NULL,
  `name` varchar(100) NOT NULL,
  `urlName` varchar(100) NOT NULL,
  `blizzId` int NOT NULL,
  `level` int NOT NULL,
  `playerLevel` int NOT NULL,
  `platform` varchar(100) NOT NULL,
  `isPublic` bool NOT NULL,
  `portrait` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `session` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` int NOT NULL,
  `battletagId` int NOT NULL,
  `starting_sr_tank` int DEFAULT 0,
  `sr_tank` int DEFAULT 0,
  `starting_sr_damage` int DEFAULT 0,
  `sr_damage` int DEFAULT 0,
  `starting_sr_support` int DEFAULT 0,
  `sr_support` int DEFAULT 0,
  PRIMARY KEY (`id`)
);

CREATE TABLE `game` (
  `id` int NOT NULL AUTO_INCREMENT, 
  `userId` int NOT NULL,
  `sessionId` int NOT NULL,
  PRIMARY KEY (`id`)
);
