package auth

import "github.com/golang-jwt/jwt/v4"

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

func (u *JwtUtil) TokenIsValid() bool {
	return u.GetUserName() != ""
}
