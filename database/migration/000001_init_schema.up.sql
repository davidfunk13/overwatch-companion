CREATE TABLE `battletag` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `userID` int NOT NULL,
  `battletag` varchar(100) NOT NULL,
  `platform` varchar(100) NOT NULL,
  `identifier` int 
  PRIMARY KEY (`id`)
);
