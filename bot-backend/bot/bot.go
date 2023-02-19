package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/commands"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

var bot *tgbotapi.BotAPI

type BotInfoResult struct {
	ID       int64  `json:"id"`
	Name     string `json:"first_name"`
	UserName string `json:"username"`
}

type BotInfo struct {
	Result BotInfoResult `json:"result"`
}

func setupWebHook(bot *tgbotapi.BotAPI) {
	// re-setup webhook
	newWebHook, err := tgbotapi.NewWebhookWithCert(
		utils.GetLambdaInvokeUrl()+"/bot"+utils.GetTelegramBotToken(),
		nil,
	)

	if err != nil {
		log.Println("unable to create new webhook", err)
	}

	_, err2 := bot.Request(newWebHook)
	if err != nil {
		log.Println("webhook request via bot failed", err2)
	}

	if !utils.IsProductionEnv() {
		existingWebHook, err3 := bot.GetWebhookInfo()
		log.Println(existingWebHook.URL, err3)
	}
}

func deleteCommands(bot *tgbotapi.BotAPI) error {
	// delete previously set commands
	delCmdCfg := tgbotapi.NewDeleteMyCommands()
	_, err := bot.Request(delCmdCfg)
	if err != nil {
		log.Println("Delete bot commands failed", err)
		return err
	}
	log.Println("Bot commands deleted")

	return nil
}

func setupCommands(bot *tgbotapi.BotAPI) error {
	delErr := deleteCommands(bot)
	if delErr != nil {
		return delErr
	}

	// get list of available commands
	commands := commands.GetCommands(commands.CATEGORY_ALL)
	botCommands := make([]tgbotapi.BotCommand, 0)

	for cmd, info := range commands {
		botCommands = append(botCommands, tgbotapi.BotCommand{
			Command:     "/" + cmd,
			Description: info.Description,
		})
	}

	cfg := tgbotapi.NewSetMyCommands(botCommands...)
	if _, err := bot.Request(cfg); err != nil {
		log.Println("Set bot commands failed", err)
		return err
	}
	log.Println("Bot commands registered")

	return nil
}

func NewTelegramBot() (*tgbotapi.BotAPI, error) {
	if bot != nil {
		return bot, nil
	}

	newBot, err := tgbotapi.NewBotAPI(utils.GetTelegramBotToken())
	if err != nil {
		return nil, err
	}

	bot = newBot
	newBot.Debug = !utils.IsProductionEnv()

	return bot, nil
}

func InitTelegramBot() (*tgbotapi.BotAPI, error) {
	if bot != nil {
		return bot, nil
	}

	newBot, err := tgbotapi.NewBotAPI(utils.GetTelegramBotToken())
	if err == nil {
		bot = newBot
		newBot.Debug = !utils.IsProductionEnv()
	}

	setupWebHook(newBot)
	setupCommands(newBot)

	/*
		if utils.IsCommandsMode() {
			setupCommands(newBot)
		} else if utils.IsWebAppMode() {
			setupWebApp(newBot)
		}
	*/

	return bot, err
}
