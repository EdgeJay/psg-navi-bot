package routes

import (
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
	"github.com/gin-gonic/gin"
)

func Menu(c *gin.Context) {
	// get domain for cookie
	domain, err := utils.GetLambdaInvokeUrlDomain()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Unable to fetch bot menu",
				"details": err.Error(),
			},
		)
		return
	}

	// set cookie
	sess, sessErr := cookies.NewMenuSession()
	if sessErr != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Unable to start bot menu session [1]",
				"details": sessErr.Error(),
			},
		)
	}

	cookieErr := cookies.SetStrictSameSiteCookie(c, sess.Map(), "psg_navi_bot_session", "/", domain, true)
	if cookieErr != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Unable to start bot menu session [2]",
				"details": cookieErr.Error(),
			},
		)
	} else {
		// return HTML output
		c.HTML(http.StatusOK, "menu.html", gin.H{
			"version": utils.GetAppVersion(),
		})
	}
}
