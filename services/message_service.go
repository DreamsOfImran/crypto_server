package services

import (
	"fmt"
	"github.com/DreamsOfImran/crypto_server/models"
)

// MessageService interface declaration
type MessageService interface {
	ReceiveMessage(key string, msg string) (*models.Message, error)
	EncryptMessage(key string, msg string) (*models.Message, error)
}

type messageService struct {
	Message      map[string]*models.Message
	AgentService AgentService
}

// NewMessageService method declaration
func NewMessageService(agentService AgentService) MessageService {
	return &messageService{
		Message:      make(map[string]*models.Message),
		AgentService: agentService,
	}
}

func (ms messageService) ReceiveMessage(agendID string, msg string) (*models.Message, error) {
	key, _ := ms.AgentService.FindByID(agendID)
	result, err := models.Decrypt(key.PrivateKey, msg)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return result, nil
}

func (ms messageService) EncryptMessage(agendID string, msg string) (*models.Message, error) {
	// agentService := NewAgentService()
	key, err := ms.AgentService.FindByID(agendID)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	result, err := models.Encrypt(key.PublicKey, msg)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return result, nil
}
