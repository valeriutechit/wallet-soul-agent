package agent

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

	solanaRegex := regexp.MustCompile(`^[1-9A-HJ-NP-Za-km-z]{32,44}$`)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := update.Message
		text := strings.TrimSpace(msg.Text)

		if text == "/start" {
			welcome := "👋 Welcome to Wallet Soul Agent!\n\nYou can:\n• Go to the website to analyze a wallet: https://wallet-soul-agent.vercel.app\n• Or just send a Solana address here to get its soul analyzed 🧠"
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, welcome))
			continue
		}

		if solanaRegex.MatchString(text) {
			report := GenerateSoulReport(text)

			reply := fmt.Sprintf("📍 Address: %s\n🧠 Archetype: %s\n🪞 Reflection:\n%s\n💎 Tokens:\n",
				report.Address, report.Profile, report.Reflection)

			for _, t := range report.Tokens {
				if t.UiAmount > 0 {
					reply += fmt.Sprintf("• %s: %.4f\n", t.Symbol, t.UiAmount)
				}
			}

			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, reply))
		} else {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "⚠️ Invalid address. Please send a valid Solana wallet address."))
		}
	}
}

