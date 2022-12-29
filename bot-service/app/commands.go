package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func startCommand(a *An) {
	msg := tgbotapi.NewMessage(a.Cfg.ChatId, "Добрый день")
	msg.ReplyMarkup = keyboard

	if _, err := a.Bot.Send(msg); err != nil {
		log.Println("Failed to respond  %s", err)
	}
}

func pingAllCommand(a *An) {
	var body []string
	body = append(body, "📌 Проверка сервисов:")
	body = append(body, "🕠 "+ getCurrentTime())

	for title, url := range a.Cfg.Urls {
		statusCode := Ping(a.Client, url)
		body = append(body, FormatPing(title, statusCode))

		if statusCode == 200 {
			a.IgnoreUrl[url] = false
		}
	}

	respond(a, strings.Join(body, "\n"))
}

func respond(a *An, response string) {
	msg := tgbotapi.NewMessage(a.Cfg.ChatId, response)
	if _, err := a.Bot.Send(msg); err != nil {
		log.Println("Failed to respond  %s", err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// UI
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const BUTTON_PING_ALL = "🍔  Проверить сервисы"

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BUTTON_PING_ALL),
	))

func getCurrentTime() string {
	dt := time.Now()
	return fmt.Sprintf(dt.Format("01-02-2006 15:04:05"))
}
