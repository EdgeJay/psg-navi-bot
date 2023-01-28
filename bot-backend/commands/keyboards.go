package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type KeyboardButtonData struct {
	Label string
	Data  string
}

type InlineKeyboard struct {
	rows [][]tgbotapi.InlineKeyboardButton
}

func NewYesNoKeyboard(yesData string, noData string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes", yesData),
			tgbotapi.NewInlineKeyboardButtonData("No", noData),
		),
	)
}

func NewInlineKeyboard() *InlineKeyboard {
	return &InlineKeyboard{
		rows: make([][]tgbotapi.InlineKeyboardButton, 0),
	}
}

func (kb *InlineKeyboard) AddRow(row ...KeyboardButtonData) *InlineKeyboard {
	buttons := make([]tgbotapi.InlineKeyboardButton, len(row))
	for buttonIndex, data := range row {
		buttons[buttonIndex] = tgbotapi.NewInlineKeyboardButtonData(data.Label, data.Data)
	}
	kb.rows = append(kb.rows, buttons)
	return kb
}

func (kb *InlineKeyboard) Render() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(kb.rows...)
}
