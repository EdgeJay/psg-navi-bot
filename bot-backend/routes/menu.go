package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func Menu(c *gin.Context) {
	if domain, err := utils.GetLambdaInvokeUrlDomain(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Unable to fetch bot menu",
				"details": err.Error(),
			},
		)
	} else {
		now := time.Now()

		params := gin.H{
			"invoke_time": now.UnixNano(),
		}

		if cookieValue, err := json.Marshal(params); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error":   "Unable to start bot menu session",
					"details": err.Error(),
				},
			)
		} else {
			// set cookie used to identify current session
			c.SetCookie("session_info", string(cookieValue), 1200, "/", domain, true, true)
			// generate JWT token that will be used to verify against other session info
			// return HTML content
			c.HTML(http.StatusOK, "menu.html", gin.H{
				"version": utils.GetAppVersion(),
			})
		}
	}
}
