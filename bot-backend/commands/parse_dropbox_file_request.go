package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func HandleDropboxFileRequest(update *tgbotapi.Update, bot *tgbotapi.BotAPI, params ...string) {
	if len(params) > 1 && params[0] == "yes" {
		fileRequestName := params[1]
		log.Printf("file request name: %s\n", fileRequestName)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Roger that! Creating Dropbox file request %s, give me a moment", fileRequestName),
		)
		SendMessage(msg, bot)

		// Get Dropbox client
		dbx := utils.NewDropboxClient(
			utils.GetDropboxAppKey(),
			utils.GetDropboxAppSecret(),
			utils.GetDropboxRefreshToken(),
		)

		if createdFileRequest, err := dbx.CreateFileRequest(fileRequestName, ""); err == nil {
			msg = tgbotapi.NewMessage(
				update.CallbackQuery.Message.Chat.ID,
				fmt.Sprintf("File request created! Please use this link: %s", createdFileRequest.URL),
			)
		} else {
			msg = tgbotapi.NewMessage(
				update.CallbackQuery.Message.Chat.ID,
				"Oops, something went wrong. I am unable to create file request now.",
			)
		}

		SendMessage(msg, bot)
	} else if len(params) > 1 && params[0] == "no" {
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Sorry, I don't understand the request.",
		)
		SendMessage(msg, bot)
	}
}
