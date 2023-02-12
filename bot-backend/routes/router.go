package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/middlewares"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func NewRouter() *gin.Engine {
	// create router
	router := gin.Default()
	router.Static("public", "./static")
	router.LoadHTMLGlob("templates/*.html")

	// add routes
	router.GET("/menu", middlewares.StartSession, Menu)
	router.POST("/init-menu-session", InitMenuSession)

	// diagnostic and setup endpoints
	router.GET("/env", Env)
	router.GET("/about-bot", AboutBot)
	router.POST("/init-bot", InitBot)

	// webhook
	router.POST("/bot"+utils.GetTelegramBotToken(), WebHook)

	return router
}
