package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang-jwt/jwt/v4"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/auth"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/bot"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/commands"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/middlewares"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

type AddFileRequest struct {
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc,omitempty" binding:"max=200"`
}

func DropboxAddFileRequest(c *gin.Context) {
	tokenCookie, _ := c.Get(middlewares.PsgNaviBotJwtName)
	token, ok := tokenCookie.(*jwt.Token)
	if !ok {
		c.Abort()
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"error": "Invalid token in cookie",
			},
		)
		return
	}

	jwtUtil := auth.NewJwtUtil(token)
	userName := jwtUtil.GetUserName()
	if userName == "" {
		c.Abort()
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"error": "Cookie token claims invalid",
			},
		)
		return
	}

	// check if user has permission to perform task
	adminManager := auth.NewAdminManager()
	if !adminManager.CanPerformTask(userName, auth.DomainDropbox, auth.AddFileRequest) {
		if userName == "" {
			c.Abort()
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "Not allowed to perform task",
				},
			)
			return
		}
	}

	// get add dropbox file request payload
	var payload AddFileRequest
	bindErr := c.BindJSON(&payload)
	if bindErr != nil {
		c.Abort()
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("Cannot create file request: %s", bindErr.Error()),
			},
		)
		return
	}

	// trigger add dropbox file request payload
	// Get Dropbox client
	dbx := utils.NewDropboxClient(
		utils.GetDropboxAppKey(),
		utils.GetDropboxAppSecret(),
		utils.GetDropboxRefreshToken(),
	)
	createdFileRequest, fileReqErr := dbx.CreateFileRequest(payload.Title, payload.Desc)

	if fileReqErr != nil {
		c.Abort()
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("Cannot create file request: %s", fileReqErr.Error()),
			},
		)
		return
	}

	// attempt to reply in chat via Telegram bot
	// get bot
	bot, botErr := bot.NewTelegramBot()
	if botErr == nil {
		msg := tgbotapi.NewMessage(
			int64(jwtUtil.GetUserID()),
			fmt.Sprintf("File request created! Please use this link: %s", createdFileRequest.URL),
		)

		commands.SendMessage(msg, bot)
	} else {
		log.Println(botErr)
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "",
	})
}
