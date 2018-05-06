CREATE TABLE `articles` (
  `id`       INTEGER   AUTO_INCREMENT,
  `tag`      varchar(64)   NOT NULL DEFAULT '',
  `uuid`     varchar(128)  NOT NULL,
  `mime`     varchar(8)    NOT NULL,
  `title`    varchar(100)  NOT NULL,
  `origin`   varchar(1024) NOT NULL DEFAULT '',
  `author`   varchar(64)  NOT NULL,
  `summary`  varchar(256)  NOT NULL DEFAULT '',
  `rpath` 	 varchar(256)  NOT NULL,
  `count`    INTEGER NOT NULL DEFAULT 0,
  `status`   INTEGER NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;