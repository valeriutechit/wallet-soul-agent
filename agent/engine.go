package agent

import (
	"fmt"
	"wallet-soul-agent/utils"
	"wallet-soul-agent/db"
)

func GenerateSoulReport(address string) SoulReport {
	if cached, err := db.GetCachedReport(address); err == nil && cached != nil {
		tokens, _ := utils.FetchTokens(address)
		fmt.Println("⚡ Using cached report for", address)

		return SoulReport{
			Address:    cached.Address,
			Tokens:     tokens,
			Profile:    cached.Profile,
			Reflection: cached.Reflection,
		}
	}

	tokens, err := utils.FetchTokens(address)
	if err != nil {
		return SoulReport{
			Address:    address,
			Tokens:     []utils.Token{},
			Profile:    "Unknown",
			Reflection: "This soul is too obscure to be understood.",
		}
	}

	profile := DetectProfile(tokens)
	reflection := GenerateReflection(tokens, profile)
	db.SaveReport(address, profile, reflection)

	return SoulReport{
		Address:    address,
		Tokens:     tokens,
		Profile:    profile,
		Reflection: reflection,
	}
}
