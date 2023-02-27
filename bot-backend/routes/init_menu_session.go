package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/auth"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/bot"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/middlewares"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

const CsrfCookieName = "cs"

type InitMenuSessionPayload struct {
	InitData string `json:"init_data"`
}

func InitMenuSession(c *gin.Context) {
	sess, exists := c.Get(middlewares.PsgNaviBotSessionName)
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
	menuSession := (sess).(*cookies.MenuSession)

	// get payload
	var payload InitMenuSessionPayload
	if err := c.BindJSON(&payload); err != nil {
		c.Abort()
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid request payload",
			},
		)
		return
	}

	// Get WebAppInitData
	initData, err := bot.UnMarshalWebAppInitData(payload.InitData)
	if err != nil {
		log.Println("invalid init data", err)
		c.Abort()
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid init data",
			},
		)
		return
	}

	domain, domainErr := utils.GetLambdaInvokeUrlDomain()
	if domainErr != nil {
		c.Abort()
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Unable to setup jwt token",
				"details": domainErr.Error(),
			},
		)
		return
	}

	// create jwt token and save as cookie
	tokenDuration := utils.GetCookieDuration()
	token, tokenErr := auth.GenerateToken(initData.User.UserName, int64(initData.User.Id), tokenDuration)
	if tokenErr != nil {
		log.Println("unable to generate jwt token", tokenErr)
		c.Abort()
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Unable to generate jwt token",
			},
		)
		return
	}

	// save jwt token into cookie
	cookies.SetJwtCookie(c, token, domain, tokenDuration)

	// save csrf token for frontend app to pick up
	cookies.SetStrictSameSiteCookie(
		c,
		gin.H{
			"val": menuSession.Checksum,
			"ver": utils.GetAppVersion(),
		},
		CsrfCookieName,
		"/",
		domain,
		utils.GetCookieDuration(),
		false,
	)

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "ok",
		},
	)
}
