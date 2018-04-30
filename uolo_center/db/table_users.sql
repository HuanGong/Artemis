CREATE TABLE `uolo_users` (
  `id`       varchar(40)  NOT NULL DEFAULT '',
  `email`    varchar(32)  DEFAULT NULL,
  `username` varchar(32)  DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;