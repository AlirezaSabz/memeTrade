package telegram

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mod/internal/infrastructure/aggregator"
	"go.mod/internal/services"
	"gopkg.in/telebot.v4"
)

func Start() {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")

	bot, err := telebot.NewBot(telebot.Settings{
		Token:     telegramToken,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		ParseMode: telebot.ModeMarkdown,
	})
	if err != nil {
		log.Fatalf("failed to create telegram bot: %v", err)
	}

	bot.Handle("/turnon", turnOn)

	log.Println("Bot started...")
	bot.Start()
}

func turnOn(c telebot.Context) error {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	agg := aggregator.NewAggregator(httpClient)

	tokens, err := agg.Tokens()
	if err != nil {
		return c.Send(fmt.Sprintf("failed to get tokens: %v", err))
	}

	if len(tokens) == 0 || len(tokens[0].Pairs) == 0 {
		return c.Send("No token data available right now.")
	}

	triangleType, err := services.DetectTriangle(tokens[0].Pairs[0].Candles)
	if err != nil {
		return c.Send(fmt.Sprintf("error detecting triangle: %v", err))
	}

	msg := fmt.Sprintf("Triangle type for token `%s`\n Pair: `%s`\n Type: %s",
		tokens[0].Address,
		tokens[0].Pairs[0].Pair,
		triangleType,
	)
	return c.Send(msg)
}
