package agent

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartTelegramBot() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("❌ TELEGRAM_BOT_TOKEN not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("🤖 Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := update.Message
		text := strings.TrimSpace(msg.Text)

		if strings.HasPrefix(text, "/start") {
			parts := strings.Split(text, " ")
			if len(parts) < 2 {
				bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "❗ Usage: /start <wallet_address>"))
				continue
			}

			address := parts[1]
			report := GenerateSoulReport(address)

			reply := fmt.Sprintf("📍 Address: %s\n🧠 Archetype: %s\n🪞 Reflection:\n%s\n💎 Tokens:\n",
				report.Address, report.Profile, report.Reflection)

			for _, t := range report.Tokens {
				if t.UiAmount > 0 {
					reply += fmt.Sprintf("• %s: %.4f\n", t.Symbol, t.UiAmount)
				}
			}

			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, reply))
		} else {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "👋 Send /start <wallet_address> to analyze a wallet"))
		}
	}
}
