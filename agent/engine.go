package agent

import (
	"wallet-soul-agent/db"
	"wallet-soul-agent/utils"
)

func GenerateSoulReport(address string) SoulReport {
	// cached, _ := db.GetCachedReport(address)
	// if cached != nil {
	// 	fmt.Println("⚡ Using cached report for:", address)
	// 	return SoulReport{
	// 		Address:    cached.Address,
	// 		Profile:    cached.Profile,
	// 		Reflection: cached.Reflection,
	// 		Tokens:     []utils.Token{},
	// 	}
	// }

	tokens, _ := utils.FetchTokens(address)

	var balance float64
	for _, t := range tokens {
		if t.Symbol == "SOL" {
			balance = t.UiAmount
		}
	}

	profile, _ := ArchetypeFromBalance(balance)

	reflection := GenerateReflectionWithOpenAI(profile, balance)

	db.SaveReport(address, profile, reflection)

	return SoulReport{
		Address:    address,
		Profile:    profile,
		Reflection: reflection,
		Tokens:     tokens,
	}
}
