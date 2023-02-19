package auth

import (
	"fmt"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

type JwtUtil struct {
	Token *jwt.Token
}

func NewJwtUtil(token *jwt.Token) *JwtUtil {
	return &JwtUtil{
		Token: token,
	}
}

func (u *JwtUtil) GetUserName() string {
	claims, ok := u.Token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return ""
	}

	return sub
}

func (u *JwtUtil) GetUserID() int64 {
	claims, ok := u.Token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("cannot cast to jwt.MapClaims")
		return 0
	}

	s := fmt.Sprintf("%.0f", claims["tgUserId"])
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Println(`cannot cast "tgUserId" to int64`)
		return 0
	}

	return int64(v)
}

func (u *JwtUtil) TokenIsValid() bool {
	return u.GetUserName() != ""
}
