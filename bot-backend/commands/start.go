package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func setupWebAppForUser(bot *tgbotapi.BotAPI, chatID int64) {
	domain, err := utils.GetLambdaInvokeUrlDomain()
	if err != nil {
		log.Fatal(err)
	}

	url := "https://" + domain + "/"
	cfg := NewSetChatMenuButtonConfig(url, chatID)
	if params, err := cfg.Params(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("set chat menu button params:", params)
		bot.MakeRequest(cfg.Method(), params)
	}
}

func Start(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Hi, nice to meet you! I am PSGNaviBot, and I am here to help with Dropbox requests and answer some NVPS PSG questions. To get started, try tapping on the \"Menu\" button or use /help command."
		SendMessage(msg, bot)

		// setup menu for user
		setupWebAppForUser(bot, update.Message.Chat.ID)
	}
}
