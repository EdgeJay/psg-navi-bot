package routes

import (
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/middlewares"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
	"github.com/gin-gonic/gin"
)

func MenuStart(c *gin.Context) {
	// find key in context to prove that session is properly set
	sess, exists := c.Get(middlewares.PsgNaviBotSessionName)

	if !exists {
		c.Abort()
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "Missing bot menu session",
			},
		)
		return
	}

	menuSession := (sess).(*cookies.MenuSession)

	// return HTML output
	c.HTML(http.StatusOK, "start.html", gin.H{
		"token":   menuSession.Checksum,
		"version": utils.GetAppVersion(),
	})
}
