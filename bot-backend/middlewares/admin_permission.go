package middlewares

import (
	"net/http"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetCheckAdminPermissionFunc(domain, task string) func(c *gin.Context) {
	return func(c *gin.Context) {
		targetDomain := domain
		targetTask := task

		tokenCookie, _ := c.Get(PsgNaviBotJwtName)
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
		adminManager, err := auth.NewAdminManager()
		if err != nil {
			c.Abort()
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "Unable to fetch permissions",
				},
			)
			return
		}

		if !adminManager.CanPerformTask(userName, targetDomain, targetTask) {
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
}
