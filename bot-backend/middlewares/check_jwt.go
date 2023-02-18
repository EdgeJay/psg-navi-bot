package middlewares

import (
	"fmt"
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/auth"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/gin-gonic/gin"
)

const PsgNaviBotJwtName = "psg_navi_bot_session"

func CheckJwt(c *gin.Context) {
	tokenStr, err := cookies.GetJwtCookie(c)
	if err != nil {
		c.Abort()
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"error": "Forbidden API access",
			},
		)
		return
	}

	token, tokenErr := auth.ParseToken(tokenStr)
	if tokenErr != nil {
		c.Abort()
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"error": fmt.Sprintf("Token validation failed: %s", tokenErr.Error()),
			},
		)
		return
	}

	c.Set(PsgNaviBotJwtName, token)
}
