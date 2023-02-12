package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/bot"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/middlewares"
)

type InitMenuSessionPayload struct {
	InitData string `json:"init_data"`
}

func InitMenuSession(c *gin.Context) {
	// find key in context to prove that session is properly set
	_, exists := c.Get(middlewares.PsgNaviBotSessionName)

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
	_, err := bot.UnMarshalWebAppInitData(payload.InitData)
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

	// create jwt token and save as cookie

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "ok",
		},
	)
}
