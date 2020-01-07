package services

import (
	"fmt"
	"github.com/DreamsOfImran/crypto_server/models"
)

// AgentService interface declaration
type AgentService interface {
	AddAgent(agentID string) (*models.Agent, error)
	FindByID(agentID string) (*models.Agent, error)
}

type agentService struct {
	Agents map[string]*models.Agent
}

// NewAgentService method declaration
func NewAgentService() AgentService {
	return &agentService{
		Agents: make(map[string]*models.Agent),
	}
}

func (c agentService) AddAgent(agentID string) (*models.Agent, error) {
	if c.Agents[agentID] != nil {
		return nil, fmt.Errorf("Agent ID already exists")
	}
	result := models.NewAgent(agentID)
	c.Agents[agentID] = result
	return result, nil
}

func (c agentService) FindByID(agentID string) (*models.Agent, error) {
	if c.Agents[agentID] == nil {
		return nil, fmt.Errorf("Agent ID does not exists")
	}
	return c.Agents[agentID], nil
}
