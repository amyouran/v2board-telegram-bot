#### sql_dash_idcn_li.v2_user 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(11) | PRI | NO | auto_increment |  |
| 2 | invite_user_id |  | int(11) |  | YES |  |  |
| 3 | telegram_id |  | bigint(20) |  | YES |  |  |
| 4 | email |  | varchar(64) | UNI | NO |  |  |
| 5 | password |  | varchar(64) |  | NO |  |  |
| 6 | password_algo |  | char(10) |  | YES |  |  |
| 7 | password_salt |  | char(10) |  | YES |  |  |
| 8 | balance |  | int(11) |  | NO |  | 0 |
| 9 | discount |  | int(11) |  | YES |  |  |
| 10 | commission_type | 0: system 1: period 2: onetime | tinyint(4) |  | NO |  | 0 |
| 11 | commission_rate |  | int(11) |  | YES |  |  |
| 12 | commission_balance |  | int(11) |  | NO |  | 0 |
| 13 | t |  | int(11) |  | NO |  | 0 |
| 14 | u |  | bigint(20) |  | NO |  | 0 |
| 15 | d |  | bigint(20) |  | NO |  | 0 |
| 16 | transfer_enable |  | bigint(20) |  | NO |  | 0 |
| 17 | banned |  | tinyint(1) |  | NO |  | 0 |
| 18 | is_admin |  | tinyint(1) |  | NO |  | 0 |
| 19 | last_login_at |  | int(11) |  | YES |  |  |
| 20 | is_staff |  | tinyint(1) |  | NO |  | 0 |
| 21 | last_login_ip |  | int(11) |  | YES |  |  |
| 22 | uuid |  | varchar(36) |  | NO |  |  |
| 23 | group_id |  | int(11) |  | YES |  |  |
| 24 | plan_id |  | int(11) |  | YES |  |  |
| 25 | speed_limit |  | int(11) |  | YES |  |  |
| 26 | remind_expire |  | tinyint(4) |  | YES |  | 1 |
| 27 | remind_traffic |  | tinyint(4) |  | YES |  | 1 |
| 28 | token |  | char(32) |  | NO |  |  |
| 29 | expired_at |  | bigint(20) |  | YES |  | 0 |
| 30 | remarks |  | text |  | YES |  |  |
| 31 | created_at |  | int(11) |  | NO |  |  |
| 32 | updated_at |  | int(11) |  | NO |  |  |
