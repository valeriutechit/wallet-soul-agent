package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"wallet-soul-agent/agent"
)

func main() {
	r := mux.NewRouter()

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

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "👋 Welcome to Wallet Soul Agent. Try /wallet/{address}")
	})

	fmt.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

