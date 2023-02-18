package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/auth"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/bot"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/cookies"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

type InitMenuSessionPayload struct {
	InitData string `json:"init_data"`
}

func InitMenuSession(c *gin.Context) {
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
	token, tokenErr := auth.GenerateToken(initData.User.UserName, tokenDuration)
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

	cookies.SetJwtCookie(c, token, domain, tokenDuration)

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "ok",
		},
	)
}
