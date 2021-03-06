CREATE TABLE `battletag` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `userID` int NOT NULL,
  `battletag` varchar(100) NOT NULL,
  `platform` varchar(100) NOT NULL,
  `identifier` int 
  PRIMARY KEY (`id`)
);

CREATE TABLE `session` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `userID` int NOT NULL,
  `roleType` ENUM('DAMAGE', "SUPPORT", "TANK") NOT NULL,
  PRIMARY KEY (`id`)
);
