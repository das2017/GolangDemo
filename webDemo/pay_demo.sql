CREATE TABLE `pay_demo` (
  `id` int(64) NOT NULL AUTO_INCREMENT,
  `account_id` int(64) NOT NULL,
  `partner_id` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `user_id` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `amount` double DEFAULT '0',
  `outer_tradeno` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5024 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

