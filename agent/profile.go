// package agent

// import (
// 	"strings"
// 	"wallet-soul-agent/utils"
// )

// // карта токенов → архетипов
// var tokenArchetypes = map[string]string{
// 	"SOL":  "Monk",
// 	"JUP":  "Strategist",
// 	"BONK": "Goblin",
// 	"BOME": "Meta Joker",
// 	"PEPE": "Memelord",
// 	"USDC": "Safeholder",
// 	"USDT": "Corporate Sleeper",
// 	"SHDW": "AI Priest",
// 	"PRCL": "Cultist",
// 	"WEN":  "Dreamer",
// }

// // запасной список архетипов
// var fallbackArchetypes = []string{
// 	"Nomad", "Shadow", "Ghost", "Invoker", "Wanderer", "Echo",
// }

// // генератор архетипа V2
// func DetectProfileV2(tokens []utils.Token) string {
// 	archetypeCount := make(map[string]int)

// 	for _, t := range tokens {
// 		if t.UiAmount <= 0 {
// 			continue
// 		}

// 		symbol := strings.ToUpper(t.Symbol)
// 		if arch, ok := tokenArchetypes[symbol]; ok {
// 			archetypeCount[arch]++
// 		}
// 	}

// 	// special case — если есть только SOL
// 	if len(archetypeCount) == 1 && archetypeCount["Monk"] > 0 {
// 		return "Monk"
// 	}

// 	// выбери наиболее "редкий" или сильный архетип
// 	var selected string
// 	var maxCount int
// 	for arch, count := range archetypeCount {
// 		if count >= maxCount && arch != "Monk" {
// 			maxCount = count
// 			selected = arch
// 		}
// 	}

// 	if selected != "" {
// 		return selected
// 	}

// 	return "Monk"
// }

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
