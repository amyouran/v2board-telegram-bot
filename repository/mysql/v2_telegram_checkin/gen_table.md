#### sql_dash_idcn_li.v2_telegram_checkin 
签到记录表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | user_tg_id | 用户TGID | bigint(20) | UNI | NO |  |  |
| 3 | user_id | 用户id | int(11) unsigned |  | NO |  |  |
| 4 | award | 奖励流量 | bigint(20) |  | NO |  |  |
| 5 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 6 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
