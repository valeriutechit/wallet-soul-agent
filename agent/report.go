package agent

import "wallet-soul-agent/utils"

type SoulReport struct {
	Address    string
	Tokens     []utils.Token
	Profile    string
	Reflection string
}
