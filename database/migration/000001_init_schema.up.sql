CREATE TABLE `battletag` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `urlName` varchar(100) NOT NULL,
  `blizzId` int NOT NULL,
  `level` int NOT NULL,
  `playerLevel` int NOT NULL,
  `platform` varchar(100) NOT NULL,
  `isPublic` bool NOT NULL,
  `portrait` varchar(100) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `session` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` varchar(100) NOT NULL,
  `battletagId` int NOT NULL,
  `starting_sr_tank` int NOT NULL,
  `sr_tank` int NOT NULL,
  `starting_sr_damage` int NOT NULL,
  `sr_damage` int NOT NULL,
  `starting_sr_support` int NOT NULL,
  `sr_support` int NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`battletagId`) REFERENCES `battletag` (`id`) ON DELETE CASCADE,
  PRIMARY KEY (`id`)
);

CREATE TABLE `game` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` varchar(100) NOT NULL,
  `battletagId` int NOT NULL,
  `sessionId` int NOT NULL,
  `location` ENUM(
    "BUSAN",
    "ILIOS",
    "LIJIANGTOWER",
    "NEPAL",
    "OASIS",
    "HANAMURA",
    "TEMPLEOFANUBIS",
    "VOLSKAYAINDUSTRIES",
    "DORADO",
    "HAVANA",
    "JUNKERTOWN",
    "RIALTO",
    "ROUTE66",
    "WATCHPOINTGIBRALTAR",
    "BLIZZARDWORLD",
    "EICHENWALDE",
    "HOLLYWOOD",
    "KINGSROW",
    "NUMBANI"
  ) NOT NULL,
  `role` ENUM ("TANK", "DAMAGE", "SUPPORT") NOT NULL,
  `sr_in` int NOT NULL,
  `sr_out` int NOT NULL,
  `match_outcome` ENUM("WIN", "LOSS", "DRAW") NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`battletagId`) REFERENCES `battletag` (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`sessionId`) REFERENCES `session` (`id`) ON DELETE CASCADE,
  PRIMARY KEY (`id`)
);