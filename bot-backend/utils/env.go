package utils

import (
	"os"

	"github.com/spf13/viper"
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
	return os.Getenv("dropbox_refresh_token")
}
