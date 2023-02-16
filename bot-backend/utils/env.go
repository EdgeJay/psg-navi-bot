package utils

import (
	"os"
	"strconv"

	"github.com/spf13/viper"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/aws"
)

func GetAppEnv() string {
	return os.Getenv("app_env")
}

func GetAppVersion() string {
	return os.Getenv("app_version")
}

func IsProductionEnv() bool {
	return GetAppEnv() == "prod"
}

func GetInteractionMode() string {
	return os.Getenv("interaction_mode")
}

func IsWebAppMode() bool {
	return GetInteractionMode() == "webapp"
}

func IsCommandsMode() bool {
	return GetInteractionMode() == "commands"
}

func GetTelegramBotToken() string {
	token := os.Getenv("bot_token")
	if token != "" {
		return token
	}
	// read from .env file as fallback
	return viper.GetString("bot_token")
}

func GetLambdaInvokeUrl() string {
	return os.Getenv("lambda_invoke_url")
}

func GetDropboxAppKey() string {
	return os.Getenv("dropbox_app_key")
}

func GetDropboxAppSecret() string {
	return os.Getenv("dropbox_app_secret")
}

func GetDropboxRefreshToken() string {
	// get from AWS SSM
	keyName := "/psg_navi_bot/dev/dropbox_refresh_token"
	return aws.GetStringParameter(keyName, os.Getenv("dropbox_refresh_token"))
}

func GetCookieDuration() int {
	duration, err := strconv.Atoi(os.Getenv("cookie_duration"))
	if err != nil {
		return 1200
	}
	return duration
}
