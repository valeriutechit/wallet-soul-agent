package agent

import (
	"wallet-soul-agent/utils"
	"wallet-soul-agent/db"
)

func GenerateSoulReport(address string) SoulReport {
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
