package commands

import (
	"encoding/json"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackQueryInfo struct {
	QueryType string
	Params    []string
}

func ParseCallbackQuery(update *tgbotapi.Update, bot *tgbotapi.BotAPI) *CallbackQueryInfo {
	if update.CallbackQuery != nil {
		// check if data is json string
		if strings.Index(update.CallbackQuery.Data, "{") == 0 {
			decoded := MakeDropboxFileRequestInlineQueryData{}
			log.Println(update.CallbackQuery.Data)
			err := json.Unmarshal([]byte(update.CallbackQuery.Data), &decoded)
			if err == nil {
				info := &CallbackQueryInfo{}
				info.QueryType = decoded.QueryType
				info.Params = []string{decoded.Data.RequestName}
			} else {
				log.Println(err)
			}
		} else {
			parts := strings.Split(update.CallbackQuery.Data, ",")
			if len(parts) > 0 {
				info := &CallbackQueryInfo{}
				info.QueryType = parts[0]

				if len(parts) > 1 {
					info.Params = parts[1:]
				} else {
					info.Params = []string{}
				}

				return info
			}
		}
	}

	return nil
}

func HandleReplyToCommand(queryInfo *CallbackQueryInfo, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if queryInfo == nil {
		log.Println("Invalid query info")
		return
	}

	log.Println("Received query: ", queryInfo.QueryType)

	switch queryInfo.QueryType {
	case "replytocommand":
		if len(queryInfo.Params) < 2 {
			log.Println("Insufficient params for replytocommand")
			return
		}
		switch queryInfo.Params[0] {
		case "makefilerequest":
			HandleDropboxFileRequest(update, bot, queryInfo.Params[1:]...)
		}
	case "make_file_request":
		if len(queryInfo.Params) < 1 {
			log.Println("Insufficient params for make_file_request")
			return
		}
		if queryInfo.Params[0] == "" {
			log.Println("Ask for file request name")

			msg := tgbotapi.NewMessage(
				update.CallbackQuery.Message.Chat.ID,
				"Please provide a file request name.",
			)

			kb := tgbotapi.NewOneTimeReplyKeyboard()
			kb.InputFieldPlaceholder = "/makefilerequest "
			msg.ReplyMarkup = kb

			SendMessage(msg, bot)
		}
	}
}
