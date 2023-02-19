package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/aws"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/bot"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func InitBot(c *gin.Context) {
	// check headers to ensure it is legit request
	token := c.GetHeader("X-PSGNaviBot-Init-Token")
	tokenSecret := aws.GetStringParameter(
		utils.GetAWSParamStoreKeyName("init_token_secret"),
		"invalid_init_token_secret",
	)

	if hashed, err := utils.CreateHmacHexString(utils.GetAppEnv(), []byte(tokenSecret)); err != nil {
		c.Abort()
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	} else if hashed != token {
		c.Abort()
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	if _, err := bot.InitTelegramBot(); err != nil {
		c.Abort()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to setup Telegram bot API"})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  "ok",
				"version": utils.GetAppVersion(),
			},
		)
	}
}
