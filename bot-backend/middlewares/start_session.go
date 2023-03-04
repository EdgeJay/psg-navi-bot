package middlewares

import (
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/auth"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/gin-gonic/gin"
)

func getSessionCookie(c *gin.Context) (*cookies.MenuSession, error) {
	str, err := c.Cookie(auth.PsgNaviBotSessionName)
	if err != nil {
		return nil, err
	}
	return auth.ParseCookieStringToMenuSession(str)
}

func StartSession(c *gin.Context) {
	// check if existing cookie existed
	// NOTE: err will be returned if cookie did not exist, already expired or checksum failed
	menuSession, err := getSessionCookie(c)

	if err != nil {
		// create and set new cookie for session
		sess, err := auth.StartMenuSession(c, auth.PsgNaviBotSessionName)

		if err != nil {
			c.Abort()
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error":   "Unable to start menu session",
					"details": err.Error(),
				},
			)
			return
		}

		// Save into context
		c.Set(auth.PsgNaviBotSessionName, sess)
	} else {
		c.Set(auth.PsgNaviBotSessionName, menuSession)
	}
}
