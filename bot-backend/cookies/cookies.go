package cookies

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetLaxSameSiteCookie(c *gin.Context, params gin.H, cookieName, path, domain string, httpOnly bool) error {
	if cookieValue, err := json.Marshal(params); err != nil {
		return err
	} else {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie(cookieName, string(cookieValue), 1200, path, domain, true, httpOnly)
	}
	return nil
}

func SetStrictSameSiteCookie(c *gin.Context, params gin.H, cookieName, path, domain string, httpOnly bool) error {
	if cookieValue, err := json.Marshal(params); err != nil {
		return err
	} else {
		c.SetSameSite(http.SameSiteStrictMode)
		c.SetCookie(cookieName, string(cookieValue), 1200, path, domain, true, httpOnly)
	}
	return nil
}
