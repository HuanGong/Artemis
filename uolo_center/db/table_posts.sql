CREATE TABLE `posts` (
  `id`       INTEGER   AUTO_INCREMENT,
  `type`     varchar(100) NOT NULL,
  `title`    varchar(100) NOT NULL,
  `from`     varchar(128) NOT NULL DEFAULT '',
  `author`   varchar(100) NOT NULL,
  `summary`  varchar(256)  NOT NULL DEFAULT '',
  `status`   integer NOT NULL DEFAULT 0,
  `offsetpath` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;