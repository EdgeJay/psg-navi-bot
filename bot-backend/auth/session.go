package auth

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

const PsgNaviBotSessionName = "psg_navi_bot_session"

func ParseCookieStringToMenuSession(str string) (*cookies.MenuSession, error) {
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

func StartMenuSession(c *gin.Context, cookieName string) (*cookies.MenuSession, error) {
	// create new cookie for session
	// get domain for cookie
	domain, err := utils.GetLambdaInvokeUrlDomain()
	if err != nil {
		return nil, err
	}

	// create cookie
	sess, sessErr := cookies.NewMenuSession()
	if sessErr != nil {
		return nil, sessErr
	}

	// set cookie
	cookieErr := cookies.SetStrictSameSiteCookie(
		c,
		sess.Map(),
		cookieName,
		"/",
		domain,
		utils.GetCookieDuration(),
		true,
	)

	if cookieErr != nil {
		return nil, cookieErr
	}

	return sess, nil
}
