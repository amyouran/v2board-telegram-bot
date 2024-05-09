DROP 
  TABLE IF EXISTS `v2_telegram_checkin`;
CREATE TABLE `v2_telegram_checkin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键', 
  `user_tg_id` BIGINT(20) NOT NULL COMMENT '用户TGID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户id',  
  `award` 	BIGINT(20) NOT NULL COMMENT '奖励流量', 
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', 
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', 
  PRIMARY KEY (`id`), 
  INDEX `idx_user_tg_id` (`user_tg_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '签到记录表';