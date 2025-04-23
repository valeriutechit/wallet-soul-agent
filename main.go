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
	"wallet-soul-agent/db"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	_ = godotenv.Load()
	db.InitDB() // ← инициализация базы

	r := mux.NewRouter()

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file not found, using container env vars")
	}

	if os.Getenv("TELEGRAM_BOT_TOKEN") == "" || os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatal("❌ Missing TELEGRAM_BOT_TOKEN or OPENAI_API_KEY environment variables")
	}

	r.HandleFunc("/wallet/{address}", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w) // ← добавь сюда

		vars := mux.Vars(r)
		address := vars["address"]
		report := agent.GenerateSoulReport(address)

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
		enableCORS(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		vars := mux.Vars(r)
		address := vars["address"]
		report := agent.GenerateSoulReport(address)

		w.Header().Set("Content-Type", "application/json")
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

