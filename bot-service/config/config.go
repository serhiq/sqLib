package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Urls     map[string]string
	Schedule float64
	Token    string
	ChatId   int64
}

func New() (*Config, error) {
	config := &Config{
		Urls:     make(map[string]string),
		Schedule: 1,
	}

	scheduleHoursEnv, _ := strconv.ParseFloat(os.Getenv("SCHEDULE_MIN"), 64)
	if scheduleHoursEnv != 0 {
		config.Schedule = scheduleHoursEnv
	}

	config.Urls = map[string]string{
		"service Broker": "http://localhost:7777",
	}

	token, tokenExists := os.LookupEnv("TELEGRAM_TOKEN")
	if !tokenExists {
		return nil, fmt.Errorf("%s is not set.  Exiting.", "TELEGRAM_TOKEN")
	}
	config.Token = token

	chatIdString, chatIdExists := os.LookupEnv("CHAT_ID")
	if !chatIdExists {
		return nil, fmt.Errorf("%s is not set.  Exiting.", "CHAT_ID")
	}

	chatId, err := strconv.ParseInt(chatIdString, 10, 64)
	if err != nil {
		panic(err)
	}
	config.ChatId = chatId

	return config, nil
}