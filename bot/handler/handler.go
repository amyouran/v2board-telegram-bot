package handler

import (
	tele "gopkg.in/telebot.v3"
)

type V2boardBot struct {
	Bot *tele.Bot
}

func New(bot *tele.Bot) *V2boardBot {
	return &V2boardBot{
		Bot: bot,
	}
}

type BotCommandHandler func(b *V2boardBot)
