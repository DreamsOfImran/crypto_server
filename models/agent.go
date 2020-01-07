package models

import (
	"crypto/rand"
	"encoding/base64"
)

// Agent struct declaration
type Agent struct {
	AgentID string
	Key     string
}

// NewAgent method declaration
func NewAgent(agentID string) *Agent {
	return &Agent{
		AgentID: agentID,
		Key:     generateKey(),
	}
}

func generateKey() string {
	key := make([]byte, 16)
	rand.Read(key)
	encodedKey := base64.URLEncoding.EncodeToString(key)
	return encodedKey
}
