package utils

import "os"

func GetAppEnv() string {
	return os.Getenv("app_env")
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
	return os.Getenv("bot_token")
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
