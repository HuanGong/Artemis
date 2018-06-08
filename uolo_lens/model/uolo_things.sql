CREATE TABLE `things` (
  `uuid`        varchar(64)   NOT NULL,
  `owner`       varchar(64)   NOT NULL,
  `content`     varchar(256)  NOT NULL,
  `start_time`  datetime      NOT NULL,
  `archive_goal` bool         NOT NULL DEFAULT false,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `things_comment` (
  `thing_id`    VARCHAR(64)   NOT NULL,
  `comment`     varchar(256)  NOT NULL,
  PRIMARY KEY (`thing_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;