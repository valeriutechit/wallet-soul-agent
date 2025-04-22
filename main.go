package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"wallet-soul-agent/agent"
)

func main() {
	r := mux.NewRouter()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	if os.Getenv("TELEGRAM_BOT_TOKEN") == "" || os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatal("❌ Set TELEGRAM_BOT_TOKEN and OPENAI_API_KEY in .env")
	}

	// Запускаем Telegram бота
	agent.StartTelegramBot()

	r.HandleFunc("/wallet/{address}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]
		report := agent.GenerateSoulReport(address)

		// Текстовый вывод — красиво, поэтично
		fmt.Fprintf(w, "📍 Address: %s\n", report.Address)
		fmt.Fprintf(w, "🧠 Archetype: %s\n", report.Profile)
		fmt.Fprintf(w, "🪞 Reflection:\n%s\n", report.Reflection)
		fmt.Fprintf(w, "💎 Tokens:\n")
		for _, t := range report.Tokens {
			if t.UiAmount > 0 {
				fmt.Fprintf(w, " - %s: %.4f\n", t.Symbol, t.UiAmount)
			}
		}
	})

	r.HandleFunc("/api/wallet/{address}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]
		report := agent.GenerateSoulReport(address)

		// Установим тип ответа
		w.Header().Set("Content-Type", "application/json")

		// Сконвертируем структуру в JSON
		err := json.NewEncoder(w).Encode(report)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "👋 Welcome to Wallet Soul Agent. Try /wallet/{address}")
	})

	fmt.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

