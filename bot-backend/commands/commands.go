package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CATEGORY_ALL     = "all"
	CATEGORY_TOP     = "top"
	CATEGORY_DROPBOX = "dropbox"
)

type CommandFunc func(update *tgbotapi.Update, bot *tgbotapi.BotAPI)

type InlineQueryDataFunc func() string

type CommandInfo struct {
	Name            string
	Description     string
	Category        string
	InlineShortcut  string
	InlineQueryData string
	Func            CommandFunc
}

var mapping map[string]*CommandInfo

func SendMessage(msg tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Webhook unable to send message", err)
	}
}

func ParseCommand(update *tgbotapi.Update) string {
	if update.Message != nil && update.Message.IsCommand() {
		return update.Message.Command()
	}
	return ""
}

func GetCommandOneLinerDesc(command string, info *CommandInfo, addLineBreak bool) string {
	lineBreak := ""
	if addLineBreak {
		lineBreak = "\n"
	}

	desc := ""
	if info.InlineShortcut != "" {
		desc = fmt.Sprintf("%s | <b>Emoji button:</b> %s", info.Description, info.InlineShortcut)
	} else {
		desc = info.Description
	}

	return fmt.Sprintf("/%s (%s) : %s%s", command, info.Name, desc, lineBreak)
}

func GetCommands(category string) map[string]*CommandInfo {
	if mapping == nil {
		log.Println("Building command mapping...")
		mapping = make(map[string]*CommandInfo)
		mapping["help"] = &CommandInfo{
			Name:        "Help command",
			Description: "Get list of available commands",
			Category:    CATEGORY_TOP,
			Func:        Help,
		}
		mapping["dropbox"] = &CommandInfo{
			Name:        "Dropbox command",
			Description: "Get list of available Dropbox commands",
			Category:    CATEGORY_TOP,
			Func:        GetDropboxCommands,
		}
		mapping["makefilerequest"] = &CommandInfo{
			Name:            "Make file request command",
			Description:     "Make a new Dropbox file request",
			Category:        CATEGORY_DROPBOX,
			InlineShortcut:  "➕",
			InlineQueryData: GetMakeDropboxFileRequestInlineQueryData(),
			Func:            MakeDropboxFileRequest,
		}
		mapping["listfilerequests"] = &CommandInfo{
			Name:            "List all file requests command",
			Description:     "List all file requests",
			Category:        CATEGORY_DROPBOX,
			InlineShortcut:  "📃",
			InlineQueryData: GetListDropboxFileRequestsInlineQueryData(),
			Func:            GetDropboxFileRequests,
		}
		mapping["getfilerequest"] = &CommandInfo{
			Name:            "Get info on file request command",
			Description:     "Get info on a file request",
			Category:        CATEGORY_DROPBOX,
			InlineShortcut:  "🔍",
			InlineQueryData: GetDropboxFileRequestInfoInlineQueryData(),
			Func:            GetDropboxFileRequestInfo,
		}
	}

	if category == CATEGORY_ALL {
		return mapping
	}

	filteredMapping := make(map[string]*CommandInfo)
	for command, info := range mapping {
		if info.Category == category {
			filteredMapping[command] = info
		}
	}

	return filteredMapping
}

func GetCommandFunc(command string) CommandFunc {
	log.Println("Received command: ", command)

	commands := GetCommands(CATEGORY_ALL)

	if command == "start" {
		log.Println("Returning Start command...")
		return Start
	}

	if cmd := commands[command]; cmd != nil {
		return cmd.Func
	}

	log.Println("Returning default command...")
	return NotFoundCommand
}
