package auth

import (
	"log"
	"time"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/aws"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const tokenIssuer = "psgnavibot.sg"

func GenerateToken(sub string, duration int) (string, error) {
	rsaPrivateKeyName := "/psg_navi_bot/dev/rsa_private"
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
