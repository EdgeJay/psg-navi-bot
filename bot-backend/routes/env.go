package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func Env(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"task_root":         os.Getenv("LAMBDA_TASK_ROOT"),
		"app_env":           os.Getenv("app_env"),
		"lambda_invoke_url": utils.GetLambdaInvokeUrl(),
	})
}