CREATE TABLE `battletag` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` int NOT NULL,
  `battletag` varchar(100) NOT NULL,
  `platform` varchar(100) NOT NULL,
  `identifier` int,
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
