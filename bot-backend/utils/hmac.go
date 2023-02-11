package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func CreateHmac(data string, secret []byte) ([]byte, error) {
	h := hmac.New(sha256.New, secret)
	if _, err := h.Write([]byte(data)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func CreateHmacHexString(data string, secret []byte) (string, error) {
	hmac, err := CreateHmac(data, secret)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hmac), nil
}
