## v2board-telegram-bot

**高性能、高负载、部署简单, 基于Golang实现了原版Bot的功能，并增加了群组签到随机赠送流量。**

<p>
  <a href="https://golang.org/doc/devel/release.html#go1.22"><img src="https://img.shields.io/badge/Go-1.22.0-blue.svg" alt="Go version 1.22.0"></a>
  <a href="https://github.com/v2board/v2board/tree/1.7.4"><img alt="Custom Badge" src="https://img.shields.io/badge/v2board-1.7.4-purple?style=flat-square""></a>
  <a href="https://t.me/zeroThemeGroup"><img alt="Telegram" src="https://img.shields.io/badge/交流群组-Telegram-blue?style=flat-square"></a>
</p>

## 介绍
根据互联网上的普遍论调，`golang`比`php`快5至25倍。由此推论，这个Bot比原版Bot快5至25倍。值得注意的是：在签到模块，本Bot使用了Redis缓存，这将进一步拉大两者之间的性能差距。
另一个值得称道的小细节是：本Bot向Telegram注册了自己的命令，在群组或私聊中输入`/`即可获得Bot可用命令的提示，这在原版Bot中似乎也是没有的。

![image](https://github.com/amyouran/v2board-telegram-bot/assets/150254537/5985726c-7ba8-4d61-9617-9ae22991c5db)

[测试Bot](https://t.me/zeroThemeGroup)

## 配置 Config.yaml
| Field                 | Desc                                                       | 
| --------------------- | ------------------------------------------------------------ | 
| Token          | bot的token                   |    
| PublicURL                | webhook的地址                               |      
| ListenPort             | Bot监听的端口                         |  
| Secret          | 验证请求是否来自Tg的密匙                   |
| CustomPrefixPrompts          | 自定义报错返回前缀                   |
| Max | 签到最高奖励流量(Byte) | 
| Min | 签到最低奖励流量(Byte) | 

> A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.

> 在每个 Webhook 请求中，应该发送一个秘密令牌到头部“X-Telegram-Bot-Api-Secret-Token”，长度为 1-256 个字符。只允许使用字母 A-Z、a-z、数字 0-9、下划线（_）和连字符（-）。这个头部很有用，可以确保请求来自由你设置的 Webhook。

关于 CustomPrefixPrompts: 一个有趣的设置，Bot会随机挑选列表中的一个前缀来使用。

## 使用方法(aapanel部署)
1. 创建数据库表, 进入v2board面板数据库, 执行以下命令:
````sql
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
````
1. 下载[最新软件包](https://github.com/amyouran/v2board-telegram-bot/releases)上传到服务器解压缩
2. 配置`configs`目录下的`configs.example.yaml`，完成后重命名为 `configs.yaml`
   
> PublicURL 配置为v2board域名/v2boardbot(例如: https://bot.xxxxx.top/v2boardbot), 保证bot.xxxxx.top可以访问v2board面板。
> ListenPort 选择一个空闲端口(例如：9996)

3. 打开v2board面板网站的配置增加反向代理，参考`PublicURL`和`ListenPort`修改成如下所示。
![image](https://github.com/amyouran/v2board-telegram-bot/assets/150254537/3551d8f9-2ff8-424f-8d0b-cac69a2619dc)

4. 配置后台运行, 按下方说明填写, 填写后点击Confirm添加即可运行。

> aaPanel 面板 > App Store > Tools
> 找到Supervisor进行安装，安装完成后点击设置 > Add Daemon按照如下填写, 
> 在 Name 填写 V2boardbot, 
> 在 Run User 选择 root (有权限的用户即可), 
> 在 Run Dir 选择 软件包上传的目录 , 
> 在 Start Command 填写 软件包上传的目录/main(例如上传到root目录下则是: /root/v2board-telegram-bot/main), 
> 在 Processes 填写 1 。

![image](https://github.com/amyouran/v2board-telegram-bot/assets/150254537/3abd3737-ba3a-4af9-838e-9f1a5ec226cf)

查看Log显示如上则为成功。

## 推荐

- [简约、优雅的v2board主题](https://github.com/amyouran/V2b-Zero-Theme)
- [v2board 动态倍率脚本](https://github.com/amyouran/v2board-dynamic-rate)

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2014-present, [Linki](https://t.me/is_linki)
