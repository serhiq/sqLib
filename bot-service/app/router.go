package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func startsWith(prefix string, content string) bool {
	return (strings.Split(content, " ")[0] == prefix)
}

func commandRouter(m *tgbotapi.Message, a *An) {
	if startsWith("/start", m.Text) {
		go startCommand(a)
	} else if m.Text == BUTTON_PING_ALL {
		go pingAllCommand(a)
	}
}
