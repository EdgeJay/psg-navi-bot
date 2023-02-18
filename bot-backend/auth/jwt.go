package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/aws"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const tokenIssuer = "psgnavibot.sg"

func ParseToken(str string) (*jwt.Token, error) {
	rsaPublicKeyName := fmt.Sprintf("/psg_navi_bot/%s/rsa_public", utils.GetAppEnv())
	svc := aws.GetSSMServiceClient()
	param, ssmErr := aws.GetParameter(svc, &rsaPublicKeyName, true)
	if ssmErr != nil {
		log.Println("unable to get parameter", ssmErr)
		return nil, ssmErr
	}

	pemKey := *param.Parameter.Value
	pubKey, pemErr := jwt.ParseRSAPublicKeyFromPEM([]byte(pemKey))
	if pemErr != nil {
		log.Println("unable to parse pem key", pemErr)
		return nil, pemErr
	}

	token, tokenErr := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})

	if tokenErr != nil {
		return nil, tokenErr
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	if !claims.VerifyIssuer(tokenIssuer, true) {
		return nil, errors.New("invalid issuer")
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errors.New("expired token")
	}

	return token, nil
}

func GenerateToken(sub string, duration int) (string, error) {
	rsaPrivateKeyName := fmt.Sprintf("/psg_navi_bot/%s/rsa_private", utils.GetAppEnv())
	svc := aws.GetSSMServiceClient()
	param, ssmErr := aws.GetParameter(svc, &rsaPrivateKeyName, true)
	if ssmErr != nil {
		log.Println("unable to get parameter", ssmErr)
		return "", ssmErr
	}

	pemKey := *param.Parameter.Value
	pvtKey, pemErr := jwt.ParseRSAPrivateKeyFromPEM([]byte(pemKey))
	if pemErr != nil {
		log.Println("unable to parse pem key", pemErr)
		return "", pemErr
	}

	now := time.Now()

	uuid, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		log.Println("unable to get uuid for token", uuidErr)
		return "", uuidErr
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		ID:        uuid.String(),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(duration) * time.Second)),
		NotBefore: jwt.NewNumericDate(now),
		Issuer:    tokenIssuer,
		Subject:   sub,
	})

	signed, signErr := token.SignedString(pvtKey)
	if signErr != nil {
		log.Println("unable to sign token", signErr)
		return "", signErr
	}

	return signed, nil
}
