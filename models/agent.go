package models

import "github.com/DreamsOfImran/crypto_server/utils"

// Agent struct declaration
type Agent struct {
	AgentID    string
	PublicKey  string
	PrivateKey string
}

// NewAgent method declaration
func NewAgent(agentID string) *Agent {
	privKey, pubKey := utils.GenerateKey()
	return &Agent{
		AgentID:    agentID,
		PublicKey:  pubKey,
		PrivateKey: privKey,
	}
}
