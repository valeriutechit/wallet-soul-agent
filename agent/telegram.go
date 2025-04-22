package agent

import (
	"fmt"
	"log"
	"os"

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

		address := update.Message.Text
		report := GenerateSoulReport(address)

		reply := fmt.Sprintf(`📍 Address: %s
🧠 Archetype: %s
🪞 Reflection: %s
💎 Tokens:
`, report.Address, report.Profile, report.Reflection)

		for _, t := range report.Tokens {
			reply += fmt.Sprintf("• %s: %.4f\n", t.Symbol, t.UiAmount)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}
}
