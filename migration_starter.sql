create Database overwatch_companion;

use overwatch_companion;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` varchar(7),
  `name` varchar(30) NOT NULL,
  `email` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO user (id, name, email)
VALUES (2352, "IAmLorde", "poo@gmail.com");

INSERT INTO user (id, name, email)
VALUES (9148, "Fablorgian", "doctorflaborgian@gmail.com");

INSERT INTO user (id, name, email)
VALUES (6986, "Yourmom", "yourmom@gmail.com");

INSERT INTO user (id, name, email)
VALUES (5719, "TurdFergason", "thatsnotmyname@gmail.com");

select * from user;

select * from user where email = "doctorflaborgian@gmail.com";

DELETE FROM user WHERE id = 5719;
