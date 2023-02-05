package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/bot"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func AboutBot(c *gin.Context) {
	botToken := utils.GetTelegramBotToken()
	if res, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", botToken)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch bot info"})
	} else {
		defer res.Body.Close()
		var botInfo bot.BotInfo
		if err := json.NewDecoder(res.Body).Decode(&botInfo); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode bot info"})
		} else {
			c.JSON(http.StatusOK, botInfo)
		}
	}
}
