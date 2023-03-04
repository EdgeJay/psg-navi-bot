package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/aws"
)

func GetAWSParamStoreKeyName(name string) string {
	return fmt.Sprintf("/psg_navi_bot/%s/%s", GetAppEnv(), name)
}

func GetAppEnv() string {
	return os.Getenv("app_env")
}

func GetAppVersion() string {
	return os.Getenv("app_version")
}

func GetAppVersionSecret() string {
	return os.Getenv("app_version_secret")
}

func IsProductionEnv() bool {
	return GetAppEnv() == "prod"
}

func GetTelegramBotToken() string {
	token := aws.GetStringParameter(
		GetAWSParamStoreKeyName("telegram_api_token"),
		os.Getenv("bot_token"),
	)
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
	return aws.GetStringParameter(
		GetAWSParamStoreKeyName("dropbox_app_key"),
		os.Getenv("dropbox_app_key"),
	)
}

func GetDropboxAppSecret() string {
	return aws.GetStringParameter(
		GetAWSParamStoreKeyName("dropbox_app_secret"),
		os.Getenv("dropbox_app_secret"),
	)
}

func GetDropboxRefreshToken() string {
	return aws.GetStringParameter(
		GetAWSParamStoreKeyName("dropbox_refresh_token"),
		os.Getenv("dropbox_refresh_token"),
	)
}

func GetOpenAiApiKey() string {
	return aws.GetStringParameter(
		GetAWSParamStoreKeyName("openai_api_key"),
		os.Getenv("openai_api_key"),
	)
}

func GetCookieDuration() int {
	duration, err := strconv.Atoi(os.Getenv("cookie_duration"))
	if err != nil {
		return 1200
	}
	return duration
}
