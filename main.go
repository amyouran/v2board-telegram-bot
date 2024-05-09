package main

import (
	"fmt"

	"v2board-telegram-bot/bot"
	"v2board-telegram-bot/logger"
	"v2board-telegram-bot/repository/mysql"
	"v2board-telegram-bot/repository/redis"
	"v2board-telegram-bot/utils/shutdown"
	"v2board-telegram-bot/utils/timeutil"

	"go.uber.org/zap"
)

func main() {

	// 初始化日志
	appLogger, err := logger.NewJSONLogger(
		logger.WithOutputInConsole(),
		logger.WithField("app", "v2board_bot"),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileRotationP("./logs/v2board_bot.log"),
	)
	if err != nil {
		panic(err)
	}

	// 替换全局日志对象
	zap.ReplaceGlobals(appLogger)

	// 初始化数据库连接（MySQL）
	dbRepo := mysql.GetDbClient()

	// 初始化Cache连接 (Redis)
	cacheRepo := redis.GetRedisClient()

	// 初始化并启动Bot服务
	v2boardBot := bot.GetBotClient()
	v2boardBot, err = bot.InitBot(v2boardBot)
	if err != nil {
		panic("Bot服务启动错误")
	}
	v2boardBot.Start()

	// 优雅关闭
	shutdown.Close(
		func() {
			// 关闭 bot poll server
			v2boardBot.Stop()

			// 关闭 db (仅支持读写)
			if err := dbRepo.DbClose(); err != nil {
				fmt.Printf("dbr close err: %s", err.Error())
			}

			// 关闭 cache 服务
			if err := cacheRepo.Close(); err != nil {
				fmt.Printf("cache close err: %s", err.Error())
			}
		},
	)
}
