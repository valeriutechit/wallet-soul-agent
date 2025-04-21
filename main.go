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

		soul := agent.AnalyzeWallet(address)
		fmt.Fprintln(w, soul)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "👋 Welcome to Wallet Soul Agent. Try /wallet/{address}")
	})

	fmt.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

