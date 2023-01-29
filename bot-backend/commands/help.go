package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Help(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		reply := "Need help? Here's the following tasks I can help with now:\n"

		commands := GetCommands(CATEGORY_TOP)
		for command, info := range commands {
			reply += GetCommandOneLinerDesc(command, info, true)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ParseMode = "HTML"

		SendMessage(msg, bot)
	}
}
