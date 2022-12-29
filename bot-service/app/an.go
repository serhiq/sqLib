package app

import (
	"bot-service/config"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/go-resty/resty/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"time"
)

type An struct {
	Cfg *config.Config

	Bot    *tgbotapi.BotAPI
	Client *resty.Client

	IgnoreUrl map[string]bool
}

func New(cfg *config.Config) (*An, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}
	bot.Debug = false

	client := resty.New()

	an := An{
		Cfg:       cfg,
		Bot:       bot,
		Client:    client,
		IgnoreUrl: make(map[string]bool),
	}

	return &an, nil
}

func (a *An) Start() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(int(a.Cfg.Schedule)).Minutes().Do(PingAll, a)
	s.StartAsync()

	log.Printf("Bot online %s", a.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.Bot.GetUpdatesChan(u)

	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {

		if update.Message == nil { // ignore non-Message updates
			continue
		}

		//  Ignore all messages, from blacklist
		if a.Cfg.ChatId != update.Message.Chat.ID {
			continue
		}

		commandRouter(update.Message, a)
	}
}

func (a *An) Send(text string) {
	msg := tgbotapi.NewMessage(a.Cfg.ChatId, text)

	if _, err := a.Bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

var PingAll = func(a *An) {
	for title, url := range a.Cfg.Urls {

		if a.IgnoreUrl[url] {
			continue
		}

		statusCode := Ping(a.Client, url)
		if statusCode != 200 {
			a.IgnoreUrl[url] = true
			a.Send(FormatPing(title, statusCode))
		}
	}
}

func Ping(client *resty.Client, endpoint string) int {
	resp, _ := client.R().
		Get(endpoint)
	return resp.StatusCode()
}

func FormatPing(serviceName string, statusCode int) string {
	if statusCode != 200 {
		return fmt.Sprintf("%s: status check ERROR: %d != 200\n", serviceName, statusCode)
	}
	return fmt.Sprintf("%s: status check: Ok", serviceName)
}
