create database `mb_blockchain` ;

CREATE TABLE `mb_blockchain`.`mb_user_right` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户权益ID,自增主键',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `tg_id` varchar(128) NOT NULL DEFAULT '' COMMENT 'telegram id',
  `right_status` tinyint NOT NULL DEFAULT 0 COMMENT '权益状态(0不可用,1可用)',
  `ext_info` text COMMENT '扩展信息',
  `remark` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注信息',
  `del_flag` tinyint NOT NULL DEFAULT 0 COMMENT '删除标识(0未删除,1已删除)',
  `expire_time` datetime NOT NULL DEFAULT '1971-01-02 00:00:00' COMMENT '权益过期时间',
  `create_time` datetime NOT NULL DEFAULT '1971-01-02 00:00:00' COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT '1971-01-02 00:00:00' COMMENT '最近更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户权益表';


CREATE TABLE `mb_blockchain`.`mb_coin` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户权益ID,自增主键',
  `coin_name` varchar(32) NOT NULL DEFAULT '' COMMENT '币名称',
  `coin_price` varchar(32) NOT NULL DEFAULT '' COMMENT '最近价格',
  `market_tag` varchar(256) NOT NULL DEFAULT '' COMMENT '交易所标识',
  `ext_info` text COMMENT '扩展信息',
  `remark` varchar(1024) NOT NULL DEFAULT '' COMMENT '备注信息',
  `del_flag` tinyint NOT NULL DEFAULT 0 COMMENT '删除标识(0未删除,1已删除)',
  `list_time` datetime NOT NULL DEFAULT '1971-01-02 00:00:00' COMMENT '发币时间',
  `create_time` datetime NOT NULL DEFAULT '1971-01-02 00:00:00' COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT '1971-01-02 00:00:00' COMMENT '最近更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_coinname` (`coin_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='币种表';



