package handler

import (
	"context"
	"fmt"
	"mysticboxbot-bot/apps/tgbot/sysconst"
	"time"

	"v2board-telegram-bot/configs"
	"v2board-telegram-bot/errors"
	"v2board-telegram-bot/repository/mysql"
	checkinModel "v2board-telegram-bot/repository/mysql/v2_telegram_checkin"
	userModel "v2board-telegram-bot/repository/mysql/v2_user"
	"v2board-telegram-bot/repository/redis"
	"v2board-telegram-bot/utils"
	"v2board-telegram-bot/utils/stringutil"

	tele "gopkg.in/telebot.v3"
)

func (b *V2boardBot) CmdCheckin() (string, string, string, BotCommandHandler) {
	command := "/checkin"
	desc := "签到领取流量"
	commandScope := tele.CommandScopeDefault
	return command, desc, commandScope, func(b *V2boardBot) {
		b.Bot.Handle(command, func(c tele.Context) error {
			replyPrefix := stringutil.GetRandomString(configs.Get().CustomPrefixPrompts)

			cacheClient := redis.GetRedisClient()
			chcheCheckinKey := fmt.Sprintf(configs.CacheTGCheckinKeyFormat, c.Sender().ID)

			redisQueryRes, err := cacheClient.Exists(context.Background(), chcheCheckinKey).Result()
			if err != nil {
				return errors.NewWithErr(sysconst.ErrCacheGet, err)
			}
			if redisQueryRes == 1 {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "已经签过到了!"))
			}

			dbClient := mysql.GetDbClient().GetDb()

			userQb := userModel.NewQueryBuilder().WhereTelegramId(mysql.EqualPredicate, c.Sender().ID)
			userObj, err := userQb.QueryOne(dbClient)
			if err != nil {
				return errors.NewWithErr(configs.ErrDbGet, err)
			}

			if userObj == nil {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "没有查询到您的用户信息, 请先绑定账号!"))
			}

			currentTime := time.Now()
			year, month, day := currentTime.Date()
			todayZero := time.Date(year, month, day, 0, 0, 0, 0, currentTime.Location())
			checkinQb := checkinModel.NewQueryBuilder().WhereUserTgId(mysql.EqualPredicate, c.Sender().ID).WhereCreatedAt(mysql.GreaterThanOrEqualPredicate, todayZero)
			checkinObj, err := checkinQb.QueryOne(dbClient)
			if err != nil {
				return errors.NewWithErr(configs.ErrDbGet, err)
			}
			if checkinObj != nil {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "已经签过到了!"))
			}

			// 开启事务
			dbClient.Begin()
			newCheckinObj := checkinModel.NewModel()
			newCheckinObj.UserId = userObj.Id
			newCheckinObj.UserTgId = userObj.TelegramId
			newCheckinObj.Award = utils.GenerateRandomNumber(configs.Get().CheckinAward.Min, configs.Get().CheckinAward.Max)

			_, err = newCheckinObj.Create(dbClient)
			if err != nil {
				dbClient.Rollback()
				return errors.NewWithErr(configs.ErrDbSet, err)
			}

			// 更新流量
			updateData := map[string]interface{}{
				"TransferEnable": userObj.TransferEnable + newCheckinObj.Award,
			}
			err = userQb.Updates(dbClient, updateData)
			if err != nil {
				dbClient.Rollback()
				return errors.NewWithErr(configs.ErrDbSet, err)
			}

			// 更新缓存
			// 计算明天 0 点的 Unix 时间戳
			tomorrow := time.Now().AddDate(0, 0, 1)
			tomorrowZeroTime := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())

			if err := cacheClient.SetEx(context.Background(), chcheCheckinKey, "-", time.Until(tomorrowZeroTime)).Err(); err != nil {
				dbClient.Rollback()
				return errors.NewWithErr(sysconst.ErrCacheSet, err)
			}

			// 提交事务
			dbClient.Commit()

			// 计算剩余流量
			remain := userObj.TransferEnable - (userObj.U + userObj.D) + newCheckinObj.Award

			replyText := fmt.Sprintf("签到成功, 获得%s流量, 剩余可用流量%s.", utils.TrafficConvert(newCheckinObj.Award), utils.TrafficConvert(remain))

			return c.Reply(replyText)
		})
	}
}
