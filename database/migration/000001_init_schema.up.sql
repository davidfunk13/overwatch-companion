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
  `roleType` ENUM('DAMAGE', "SUPPORT", "TANK") NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `game` (
  `id` int NOT NULL AUTO_INCREMENT, 
  `userId` int NOT NULL,
  `sessionId` int NOT NULL,
  PRIMARY KEY (`id`)
);
