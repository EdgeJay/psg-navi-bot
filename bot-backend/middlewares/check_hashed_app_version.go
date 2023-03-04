package middlewares

import (
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
	"github.com/gin-gonic/gin"
)

func CheckHashedAppVersion(c *gin.Context) {
	if c.GetHeader("X-PSGNaviBot-Version") != utils.GetAppVersion() {
		c.Abort()
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"error": "Invalid link",
			},
		)
		return
	}

	if !utils.VerifyAppVersionHmac(c.GetHeader("X-PSGNaviBot-Hash")) {
		c.Abort()
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"error": "Invalid hash",
			},
		)
		return
	}
}
