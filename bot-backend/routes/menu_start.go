package routes

import (
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func MenuStart(c *gin.Context) {
	// find key in context to prove that session is properly set
	_, exists := c.Get(middlewares.PsgNaviBotSessionName)

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

	c.Redirect(http.StatusTemporaryRedirect, "/api/menu/home")
}
