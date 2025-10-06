package telegram

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mod/internal/domain"
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

	for _, token := range tokens {
		for _, pair := range token.Pairs {
			err := services.DetectTriangle(&pair)
			if err != nil {
				return c.Send(fmt.Sprintf("error detecting triangle: %v", err))
			}
		}
	}

	msg := fmt.Sprintf("Triangle type for token `%s`\n Pair: `%s`",
		tokens[0].Address,
		tokens[0].Pairs[0].Pair,
	)
	go monitorTrianglePatterns(&tokens, agg)

	return c.Send(msg)
}

func monitorTrianglePatterns(tokens *[]domain.Token, agg *aggregator.Aggregator) {
	for {
		time.Sleep(5 * time.Minute)
		if err := agg.FetchNewCandles(tokens); err != nil {
			log.Printf("failed to fetch new candles: %v", err)
			continue
		}
		for _, token := range *tokens {
			for _, pair := range token.Pairs {
				switch {
				case pair.TriangleType == domain.Ascending || pair.TriangleType == domain.Symmetrical:
					evaluateTriangleBreakout(&pair)
				default:
				}
			}
		}
	}
}

func evaluateTriangleBreakout(pair *domain.Pair) {
	lastCandle := pair.Candles[len(pair.Candles)-1]
	yUpper := pair.UpperTrendLine.LineEquation(len(pair.Candles) - 1)
	yLower := pair.LowerTrendLine.LineEquation(len(pair.Candles) - 1)

	switch {
	case lastCandle.Close > yUpper:
	case lastCandle.Close < yLower:
	default:
		// inside triangle — optional log or ignore	}
	}
}
