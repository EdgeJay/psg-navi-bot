package env

import (
	"os"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/aws"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func GetOpenAiApiKey() string {
	return aws.GetStringParameter(
		utils.GetAWSParamStoreKeyName("openai_api_key"),
		os.Getenv("openai_api_key"),
	)
}
