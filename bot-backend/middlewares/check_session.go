package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSession(c *gin.Context) {
	// check if existing cookie existed
	// NOTE: err will be returned if cookie did not exist, already expired or checksum failed
	menuSession, err := getSessionCookie(c)

	if err != nil {
		c.Abort()
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Unable to find session",
				"details": err.Error(),
			},
		)
		return
	}

	c.Set(PsgNaviBotSessionName, menuSession)
}
