CREATE TABLE `posts` (
  `id`       INTEGER      NOT NULL,
  `type`     varchar(100) NOT NULL,
  `title`    varchar(100) NOT NULL,
  `from`     varchar(128) NOT NULL DEFAULT '',
  `author`   varchar(100) NOT NULL,
  `summary`  varchar(256)  NOT NULL DEFAULT '',
  `offsetpath` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;