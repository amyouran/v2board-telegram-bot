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

A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.

在每个 Webhook 请求中，应该发送一个秘密令牌到头部“X-Telegram-Bot-Api-Secret-Token”，长度为 1-256 个字符。只允许使用字母 A-Z、a-z、数字 0-9、下划线（_）和连字符（-）。这个头部很有用，可以确保请求来自由你设置的 Webhook。

关于 CustomPrefixPrompts: 一个有趣的设置，Bot会随机挑选列表中的一个前缀来使用。

## 使用方法(只说明使用软件包，自行编译也差不多)

1. 下载软件包上传到服务器解压缩
2. 配置`configs.example.yaml`，完成后重命名为 `configs.yaml`
3. 安装Python依赖 `pip install -r requirements.txt`

4. 根据需求修改`config.yaml`

5. 配置并启动 `Linux crontab` 定时任务


## 推荐

- [简约、优雅的v2board主题](https://github.com/amyouran/V2b-Zero-Theme)
