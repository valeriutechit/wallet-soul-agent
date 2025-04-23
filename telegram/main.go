package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"wallet-soul-agent/agent"
)

func loadEnv() {
	// Определим абсолютный путь до корня проекта
	rootDir, _ := os.Getwd()

	// Попробуем загрузить .env из корня
	err := godotenv.Load(filepath.Join(rootDir, ".env"))
	if err != nil {
		log.Println("⚠️ Could not load .env from root, using system env")
	}
}

func Start() {
	loadEnv()

	if os.Getenv("TELEGRAM_BOT_TOKEN") == "" || os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatal("❌ Missing TELEGRAM_BOT_TOKEN or OPENAI_API_KEY in env")
	}

	go agent.StartTelegramBot() // ✅ Telegram long polling в отдельной горутине

	// ✅ Railway ожидает, что сервер слушает PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Bot is alive")
	})
	log.Printf("🌐 Keep-alive server running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	Start()
}
