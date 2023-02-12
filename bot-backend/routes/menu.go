package routes

import (
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/middlewares"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
	"github.com/gin-gonic/gin"
)

func Menu(c *gin.Context) {
	// find key in context to prove that session is properly set
	_, exists := c.Get(middlewares.PsgNaviBotSessionName)

	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Missing bot menu session",
			},
		)
	} else {
		// return HTML output
		c.HTML(http.StatusOK, "menu.html", gin.H{
			"version": utils.GetAppVersion(),
		})
	}
}
