package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/bot"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func InitBot(c *gin.Context) {
	if _, err := bot.InitTelegramBot(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to setup Telegram bot API"})
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":           "ok",
				"interaction_mode": utils.GetInteractionMode(),
			},
		)
	}
}
