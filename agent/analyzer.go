package agent

import (
	"fmt"
	"strings"
	"wallet-soul-agent/utils"
)

func AnalyzeWallet(address string) string {
	tokens, err := utils.FetchTokens(address)
	if err != nil || len(tokens) == 0 {
		return "😶 This wallet is either empty or too mysterious to decode."
	}

	var symbols []string
	for _, t := range tokens {
		if t.UiAmount > 0 {
			symbols = append(symbols, t.Symbol)
		}
	}

	summary := strings.Join(symbols, ", ")
	return fmt.Sprintf("🧠 Address %s holds: %s\nThis soul leans toward chaos, memes, or maybe... enlightenment?", address, summary)
}
