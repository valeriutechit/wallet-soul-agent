package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wallet-soul-agent/agent"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if os.Getenv("TELEGRAM_BOT_TOKEN") == "" || os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatal("❌ Missing TELEGRAM_BOT_TOKEN or OPENAI_API_KEY in env")
	}

	go agent.StartTelegramBot() // оставляем long polling

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Bot is alive")
	})

	log.Println("🌐 Starting keep-alive server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
