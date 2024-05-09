package handler

import (
	"fmt"
	"net/url"

	"v2board-telegram-bot/configs"
	"v2board-telegram-bot/errors"
	"v2board-telegram-bot/repository/mysql"
	userModel "v2board-telegram-bot/repository/mysql/v2_user"
	"v2board-telegram-bot/utils/stringutil"

	tele "gopkg.in/telebot.v3"
)

func (b *V2boardBot) CmdBlind() (string, string, string, BotCommandHandler) {
	command := "/bind"
	desc := "绑定网站账户"
	commandScope := tele.CommandScopeAllPrivateChats
	return command, desc, commandScope, func(b *V2boardBot) {
		b.Bot.Handle(command, func(c tele.Context) error {
			replyPrefix := stringutil.GetRandomString(configs.Get().CustomPrefixPrompts)
			if len(c.Args()) != 1 {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "参数有误, 请携带订阅地址发送."))
			}

			subscribeUrl := c.Args()[0]

			parsedURL, err := url.Parse(subscribeUrl)
			if err != nil {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "订阅地址无效."))
			}

			queryParams, err := url.ParseQuery(parsedURL.RawQuery)
			if err != nil {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "订阅地址无效."))
			}

			tokens, ok := queryParams["token"]

			if !ok {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "订阅地址无效."))
			}

			dbClient := mysql.GetDbClient().GetDb()

			userQb := userModel.NewQueryBuilder().WhereToken(mysql.EqualPredicate, tokens[0])
			userObj, err := userQb.QueryOne(dbClient)
			if err != nil {
				return errors.NewWithErr(configs.ErrDbGet, err)
			}

			if userObj == nil {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "用户不存在."))
			}

			if userObj.TelegramId != 0 {
				return c.Reply(fmt.Sprintf("%s, %s", replyPrefix, "该账号已经绑定了Telegram账号."))
			}

			updateData := map[string]interface{}{
				"TelegramId": c.Sender().ID,
			}

			err = userQb.Updates(dbClient, updateData)
			if err != nil {
				return errors.NewWithErr(configs.ErrDbSet, err)
			}

			return c.Send("绑定成功了!")
		})
	}
}
