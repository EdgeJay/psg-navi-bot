package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Menu(c *gin.Context) {
	c.HTML(http.StatusOK, "menu.html", nil)
}
