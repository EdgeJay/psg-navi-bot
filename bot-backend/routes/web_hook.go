package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/bot"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/commands"
)

func WebHook(c *gin.Context) {
	// get bot
	if bot, err := bot.NewTelegramBot(); err != nil {
		log.Println("Webhook unable to init bot")
	} else {
		if update, err2 := bot.HandleUpdate(c.Request); err2 != nil {
			log.Println("Webhook unable to parse update")
		} else {
			if update.Message != nil {
				cmdStr := commands.ParseCommand(update)
				cmd := commands.GetCommandFunc(cmdStr)
				cmd(update, bot)
			} else if update.CallbackQuery != nil {
				/*
					// Flashes update.CallbackQuery.Data in Telegram window like a toast
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
					if _, err := bot.Request(callback); err != nil {
						log.Println(err)
					}
				*/

				queryInfo := commands.ParseCallbackQuery(update, bot)
				commands.HandleReplyToCommand(queryInfo, update, bot)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
