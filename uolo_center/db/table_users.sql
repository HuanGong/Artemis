CREATE TABLE `uolo_users` (
  `id`       varchar(40)  NOT NULL DEFAULT '',
  `email`    varchar(32)  DEFAULT NULL,
  `username` varchar(32)  DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for user_info
-- ----------------------------
CREATE TABLE `user_profile` (
  `id` varchar(40) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '用户主键Id',
  `username` varchar(24) NOT NULL COMMENT '用户名',
  `password` varchar(64) NOT NULL COMMENT '密文密码',
  `nickname` varchar(24) NOT NULL COMMENT '昵称',
  `avast` varchar(50) NOT NULL DEFAULT '' COMMENT '头像Id',
  `phone` varchar(50) NOT NULL DEFAULT '' COMMENT '用户手机号',
  `email` varchar(50) NOT NULL COMMENT '用户邮箱',
  `sex` int(2) NOT NULL DEFAULT '0' COMMENT '性别：0未知 1男，2女',
  `birth` date DEFAULT NULL COMMENT '用户年龄',
  `tags` bigint(20) DEFAULT NULL COMMENT '用户标签',
  `status` int(2) NOT NULL DEFAULT '0' COMMENT '逻辑删除位',
  `createtime` datetime NOT NULL COMMENT '创建时间',
  `updatetime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_username_password` (`username`,`password`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `uolo_admin` (
  `user_id`   varchar(64) character set utf8 collate utf8_bin NOT NULL,
  `level`     tinyint(2)  DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
