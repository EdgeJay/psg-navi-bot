package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/aws"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtTokenClaims struct {
	UserID int64 `json:"tgUserId"`
	jwt.RegisteredClaims
}

const tokenIssuer = "psgnavibot.sg"

func getRsaPublicKeyFromSSM() (*rsa.PublicKey, error) {
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
	return pubKey, nil
}

func validateToken(token *jwt.Token) error {
	if !token.Valid {
		return errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid claims")
	}

	if !claims.VerifyIssuer(tokenIssuer, true) {
		return errors.New("invalid issuer")
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return errors.New("expired token")
	}

	return nil
}

// Similar to ParseToken, but extracts RSA public key from file instead of AWS SSM
// This should only be used for unit testing purposes!
func ParseTokenFromFile(str string) (*jwt.Token, error) {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		return nil, errors.New("invalid working directory")
	}

	pemKey, err := os.ReadFile(wd + "/../certs/rsa_public.pem")
	if err != nil {
		return nil, errors.New("cannot read PEM file")
	}

	pubKey, pemErr := jwt.ParseRSAPublicKeyFromPEM(pemKey)
	if pemErr != nil {
		log.Println("unable to parse pem key", pemErr)
		return nil, pemErr
	}

	token, tokenErr := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})

	if tokenErr != nil {
		return nil, tokenErr
	}

	if err := validateToken(token); err != nil {
		return nil, err
	}

	return token, nil
}

func ParseToken(str string) (*jwt.Token, error) {
	pubKey, pemErr := getRsaPublicKeyFromSSM()
	if pemErr != nil {
		return nil, pemErr
	}

	token, tokenErr := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})

	if tokenErr != nil {
		return nil, tokenErr
	}

	if err := validateToken(token); err != nil {
		return nil, err
	}

	return token, nil
}

func GenerateToken(sub string, userId int64, duration int) (string, error) {
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
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, JwtTokenClaims{
		userId,
		jwt.RegisteredClaims{
			ID:        uuid.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(duration) * time.Second)),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    tokenIssuer,
			Subject:   sub,
		}})

	signed, signErr := token.SignedString(pvtKey)
	if signErr != nil {
		log.Println("unable to sign token", signErr)
		return "", signErr
	}

	return signed, nil
}
