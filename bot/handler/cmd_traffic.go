package handler

import (
	"fmt"
	"strings"

	"v2board-telegram-bot/configs"
	"v2board-telegram-bot/errors"
	"v2board-telegram-bot/repository/mysql"
	userModel "v2board-telegram-bot/repository/mysql/v2_user"
	utils "v2board-telegram-bot/utils"
	"v2board-telegram-bot/utils/stringutil"

	tele "gopkg.in/telebot.v3"
)

func (b *V2boardBot) CmdTraffic() (string, string, string, BotCommandHandler) {
	command := "/traffic"
	desc := "查询流量信息"
	commandScope := tele.CommandScopeDefault
	return command, desc, commandScope, func(b *V2boardBot) {
		b.Bot.Handle(command, func(c tele.Context) error {
			replyPrefix := stringutil.GetRandomString(configs.Get().CustomPrefixPrompts)

			dbClient := mysql.GetDbClient().GetDb()

			userQb := userModel.NewQueryBuilder().WhereTelegramId(mysql.EqualPredicate, c.Sender().ID)
			userObj, err := userQb.QueryOne(dbClient)
			if err != nil {
				return errors.NewWithErr(configs.ErrDbGet, err)
			}

			if userObj == nil {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "没有查询到您的用户信息, 请先绑定账号!"))
			}

			remain := userObj.TransferEnable - (userObj.U + userObj.D)

			replyRows := []string{
				"🚥流量查询",
				"———————————————",
				fmt.Sprintf("订阅流量: %s", utils.TrafficConvert(userObj.TransferEnable)),
				fmt.Sprintf("已用上行: %s", utils.TrafficConvert(userObj.U)),
				fmt.Sprintf("已用下行: %s", utils.TrafficConvert(userObj.D)),
				fmt.Sprintf("剩余流量: %s", utils.TrafficConvert(remain)),
			}
			replyText := strings.Join(replyRows, "\n")

			return c.Reply(replyText)
		})
	}
}
