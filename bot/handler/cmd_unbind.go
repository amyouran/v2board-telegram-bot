package handler

import (
	"fmt"

	"v2board-telegram-bot/configs"
	"v2board-telegram-bot/errors"
	"v2board-telegram-bot/repository/mysql"
	userModel "v2board-telegram-bot/repository/mysql/v2_user"
	"v2board-telegram-bot/utils/stringutil"

	tele "gopkg.in/telebot.v3"
)

func (b *V2boardBot) CmdUnBlind() (string, string, string, BotCommandHandler) {
	command := "/unbind"
	desc := "解绑网站账户"
	commandScope := tele.CommandScopeAllPrivateChats
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

			updateData := map[string]interface{}{
				"TelegramId": nil,
			}

			err = userQb.Updates(dbClient, updateData)
			if err != nil {
				return errors.NewWithErr(configs.ErrDbSet, err)
			}

			return c.Reply("解绑成功了!")
		})
	}
}
