package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
	"github.com/gin-gonic/gin"
)

const PsgNaviBotSessionName = "psg_navi_bot_session"

func getSessionCookie(c *gin.Context) (*cookies.MenuSession, error) {
	str, err := c.Cookie(PsgNaviBotSessionName)
	if err != nil {
		return nil, err
	}

	menuSession := cookies.MenuSession{}
	if err := menuSession.ParseJson(str); err != nil {
		return nil, err
	}

	// check cookie expiry
	if menuSession.IsExpired(utils.GetCookieDuration()) {
		log.Println("Session expired")
		return nil, errors.New("session expired")
	}

	// check cookie checksum validity
	if !menuSession.IsChecksumValid() {
		log.Println("Invalid session checksum")
		return nil, errors.New("session checksum invalid")
	}

	return &menuSession, nil
}

func StartSession(c *gin.Context) {
	// check if existing cookie existed
	// NOTE: err will be returned if cookie did not exist, already expired or checksum failed
	menuSession, err := getSessionCookie(c)

	if err != nil {
		// create new cookie for session
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

		// create cookie
		sess, sessErr := cookies.NewMenuSession()
		if sessErr != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error":   "Unable to start bot menu session",
					"details": sessErr.Error(),
				},
			)
		}

		// set cookie
		cookieErr := cookies.SetStrictSameSiteCookie(
			c,
			sess.Map(),
			PsgNaviBotSessionName,
			"/",
			domain,
			utils.GetCookieDuration(),
			true,
		)

		if cookieErr != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error":   "Unable to set menu session",
					"details": cookieErr.Error(),
				},
			)
		} else {
			// Save into context
			c.Set(PsgNaviBotSessionName, sess)
		}
	} else {
		c.Set(PsgNaviBotSessionName, menuSession)
	}
}

func CheckSession(c *gin.Context) {

}
