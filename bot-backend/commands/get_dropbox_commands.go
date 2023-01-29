package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetDropboxCommands(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		reply := "Hi, the following Dropbox commands are available:\n"

		// Get all Dropbox related commands
		commands := GetCommands(CATEGORY_DROPBOX)

		// add inline keyboard
		kb := NewInlineKeyboard()
		buttonData := make([]KeyboardButtonData, len(commands))
		buttonIndex := 0

		for command, info := range commands {
			reply += GetCommandOneLinerDesc(command, info, true)

			if info.InlineShortcut != "" && info.InlineQueryData != "" {
				buttonData[buttonIndex] = KeyboardButtonData{
					Label: info.InlineShortcut,
					Data:  info.InlineQueryData,
				}
				buttonIndex++
			}
		}

		reply += "\nYou can also use the emoji buttons below to send commands."

		kb.AddRow(buttonData...)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ReplyMarkup = kb.Render()
		msg.ParseMode = "HTML"

		SendMessage(msg, bot)
	}
}
