package agent

import "wallet-soul-agent/utils"

func DetectProfile(tokens []utils.Token) string {
	hasMemes := false
	hasStable := false

	for _, t := range tokens {
		switch t.Symbol {
		case "BONK", "WIF", "DOGE":
			hasMemes = true
		case "USDC", "USDT", "DAI":
			hasStable = true
		}
	}

	switch {
	case hasMemes && hasStable:
		return "Balanced Degen"
	case hasMemes:
		return "Degen Explorer"
	case hasStable:
		return "Cautious Strategist"
	case len(tokens) == 0 || (len(tokens) == 1 && tokens[0].Symbol == "SOL" && tokens[0].UiAmount < 0.01):
		return "The Void"
	default:
		return "Unknown Archetype"
	}
}
