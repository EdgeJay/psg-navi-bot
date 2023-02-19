package commands

// SetChatMenuButtonConfig changes the bot's menu button in a private chat,
// or the default menu button.
//
// NOTE: The following code was cherry-picked from
// https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/configs.go#L2341
// as the code is not made available to TgBotApi V5 yet.

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebAppInfo struct {
	// URL is the HTTPS URL of a Web App to be opened with additional data as
	// specified in Initializing Web Apps.
	URL string `json:"url"`
}

// MenuButton describes the bot's menu button in a private chat.
type MenuButton struct {
	// Type is the type of menu button, must be one of:
	// - `commands`
	// - `web_app`
	// - `default`
	Type string `json:"type"`
	// Text is the text on the button, for `web_app` type.
	Text string `json:"text,omitempty"`
	// WebApp is the description of the Web App that will be launched when the
	// user presses the button for the `web_app` type.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

type SetChatMenuButtonConfig struct {
	ChatID          int64
	ChannelUsername string

	MenuButton *MenuButton
}

func NewSetChatMenuButtonConfig(url string, chatID int64) SetChatMenuButtonConfig {
	cfg := SetChatMenuButtonConfig{}
	cfg.ChatID = chatID
	cfg.MenuButton = &MenuButton{
		Type: "web_app",
		Text: "Start",
		WebApp: &WebAppInfo{
			URL: url,
		},
	}
	return cfg
}

func (config SetChatMenuButtonConfig) Method() string {
	return "setChatMenuButton"
}

func (config SetChatMenuButtonConfig) Params() (tgbotapi.Params, error) {
	params := make(tgbotapi.Params)

	if err := params.AddFirstValid("chat_id", config.ChatID, config.ChannelUsername); err != nil {
		return params, err
	}
	err := params.AddInterface("menu_button", config.MenuButton)

	return params, err
}
