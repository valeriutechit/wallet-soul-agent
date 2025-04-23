package agent

import "wallet-soul-agent/utils"

type SoulReport struct {
	Address    string `json:"address"`
	Tokens     []utils.Token `json:"tokens"`
	Profile    string `json:"profile"`
	Reflection string `json:"reflection"`
}
