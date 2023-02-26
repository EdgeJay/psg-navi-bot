package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/auth"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/middlewares"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func NewRouter() *gin.Engine {
	// create router
	router := gin.Default()
	// router.Static("public", "./static")
	router.LoadHTMLGlob("templates/*.html")

	// menu
	router.GET("/menu-start", middlewares.StartSession, MenuStart)
	router.GET("/menu/:page", middlewares.CheckSession, MenuPage)
	router.POST("/init-menu-session", middlewares.CheckSession, middlewares.CheckCsrf, InitMenuSession)
	// dropbox
	router.POST(
		"/dbx-add-file-request",
		middlewares.CheckSession,
		middlewares.CheckCsrf,
		middlewares.CheckJwt,
		middlewares.GetCheckAdminPermissionFunc(auth.DomainDropbox, auth.AddFileRequest),
		DropboxAddFileRequest,
	)
	router.GET(
		"/dbx-list-file-requests",
		middlewares.CheckSession,
		middlewares.CheckCsrf,
		middlewares.CheckJwt,
		middlewares.GetCheckAdminPermissionFunc(auth.DomainDropbox, auth.ListFileRequests),
		DropboxListFileRequests,
	)

	// diagnostic and setup endpoints
	router.GET("/env", Env)
	router.GET("/about-bot", AboutBot)
	router.POST("/init-bot", InitBot)

	// webhook
	router.POST("/bot"+utils.GetTelegramBotToken(), WebHook)

	return router
}
