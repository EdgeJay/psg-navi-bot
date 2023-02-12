package middlewares

import (
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/gin-gonic/gin"
)

func CheckCsrf(c *gin.Context) {
	// find key in context to prove that session is properly set
	sess, exists := c.Get(PsgNaviBotSessionName)

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

	if menuSession.Checksum != c.GetHeader("X-PSGNaviBot-Csrf-Token") {
		c.Abort()
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "Invalid bot menu session",
			},
		)
		return
	}
}
