package bot

import (
	"fmt"
	"reflect"
	"sync"

	"v2board-telegram-bot/bot/handler"
	"v2board-telegram-bot/bot/middleware"
	"v2board-telegram-bot/configs"
	"v2board-telegram-bot/errors"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

var (
	bot  *tele.Bot
	once sync.Once
)

func GetBotClient() *tele.Bot {
	once.Do(func() {
		botClient, err := NewBot()
		if err != nil {
			panic(err)
		}
		bot = botClient
	})

	return bot
}

func onBotError(err error, c tele.Context) {
	appLogger := zap.L()
	if c != nil {
		if _, ok := err.(errors.Error); ok {
			c.Send(err.Error())
		}
	} else {
		appLogger.Error("bot err", zap.Error(errors.WithStack(err)))
	}
}

func NewBot() (*tele.Bot, error) {
	bot, err := tele.NewBot(tele.Settings{
		Token: configs.Get().Bot.Token,
		Poller: &tele.Webhook{
			Endpoint:    &tele.WebhookEndpoint{PublicURL: configs.Get().Bot.PublicURL},
			DropUpdates: true,
			Listen:      fmt.Sprintf(":%d", configs.Get().Bot.ListenPort),
			SecretToken: configs.Get().Bot.Secret,
		},
		Verbose:     false,
		Synchronous: false,
		OnError:     onBotError,
		ParseMode:   "HTML",
	})
	if err != nil {
		panic(err)
	}

	return bot, err
}

func InitBot(Bot *tele.Bot) (*tele.Bot, error) {

	// 设置默认权限
	// tgBot.SetDefaultRights(tele.AdminRights(), false) // 群组默认权限
	// tgBot.SetDefaultRights(tele.AdminRights(), true)  // 频道默认权限

	// 清空命令
	err := Bot.DeleteCommands()
	if err != nil {
		panic("清空 bot commands 错误")
	}

	needSetCommandScopeDefault := make([]tele.Command, 0)
	needSetCommandScopeAllPrivateChats := make([]tele.Command, 0)
	needSetCommandScopeAllGroupChats := make([]tele.Command, 0)
	needSetCommandScopeAllChatAdmin := make([]tele.Command, 0)

	v2boardBot := handler.New(bot)

	// 开启全局中间件
	// v2boardBot.Bot.Use(middleware.Logger()) // 日志中间件
	v2boardBot.Bot.Use(middleware.AutoRespond())
	v2boardBot.Bot.Use(middleware.Recover())

	objType := reflect.TypeOf(v2boardBot)

	fmt.Println("Bot's NumMethod: ", objType.NumMethod())

	for i := 0; i < objType.NumMethod(); i++ {
		method := objType.Method(i)
		methodValue := reflect.ValueOf(v2boardBot).MethodByName(method.Name)
		callRes := methodValue.Call(nil)
		cmd, desc, scope, botHandler := callRes[0], callRes[1], callRes[2], callRes[3]
		// set endpoint
		if fn, ok := botHandler.Interface().(handler.BotCommandHandler); ok {
			fn(v2boardBot)
		} else {
			panic("bot endpoint set error")
		}
		// bot命令检查
		if cmd.String() != "" && desc.String() != "" && scope.String() != "" {
			tempCommand := tele.Command{Text: cmd.String()[1:], Description: desc.String()}
			switch scope.String() {
			case "all_private_chats":
				needSetCommandScopeAllPrivateChats = append(needSetCommandScopeAllPrivateChats, tempCommand)
			case "all_group_chats":
				needSetCommandScopeAllGroupChats = append(needSetCommandScopeAllGroupChats, tempCommand)
			case "all_chat_administrators":
				needSetCommandScopeAllChatAdmin = append(needSetCommandScopeAllChatAdmin, tempCommand)
			default:
				needSetCommandScopeDefault = append(needSetCommandScopeDefault, tempCommand)
			}
		}

	}

	// 添加默认范围命令
	fmt.Println("cmd default:", len(needSetCommandScopeDefault))
	if len(needSetCommandScopeDefault) > 0 {
		if err := v2boardBot.Bot.SetCommands(needSetCommandScopeDefault); err != nil {
			panic(err)
		}
	}

	// 添加所有私聊范围命令
	fmt.Println("cmd all_private_chats:", len(needSetCommandScopeAllPrivateChats))
	if len(needSetCommandScopeAllPrivateChats) > 0 {
		scope := tele.CommandScope{
			Type: tele.CommandScopeAllPrivateChats,
		}
		if err := v2boardBot.Bot.SetCommands(needSetCommandScopeAllPrivateChats, scope); err != nil {
			panic(err)
		}
	}

	// 添加所有群组范围命令
	fmt.Println("cmd all_group_chats:", len(needSetCommandScopeAllGroupChats))
	if len(needSetCommandScopeAllGroupChats) > 0 {
		scope := tele.CommandScope{
			Type: tele.CommandScopeAllGroupChats,
		}
		if err := v2boardBot.Bot.SetCommands(needSetCommandScopeAllGroupChats, scope); err != nil {
			panic(err)
		}
	}

	// 添加所有管理员范围命令
	fmt.Println("cmd all_chat_administrators:", len(needSetCommandScopeAllChatAdmin))
	if len(needSetCommandScopeAllChatAdmin) > 0 {
		scope := tele.CommandScope{
			Type: tele.CommandScopeAllChatAdmin,
		}
		if err := v2boardBot.Bot.SetCommands(needSetCommandScopeAllChatAdmin, scope); err != nil {
			panic(err)
		}
	}

	existComd, err := v2boardBot.Bot.Commands()
	if err != nil {
		return nil, errors.NewWithErr(configs.ErrBotApi, err)
	}
	for index := range existComd {
		fmt.Println(existComd[index].Text, ": ", existComd[index].Description)
	}

	return v2boardBot.Bot, nil
}
