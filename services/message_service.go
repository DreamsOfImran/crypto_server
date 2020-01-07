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
	Message map[string]*models.Message
}

// NewMessageService method declaration
func NewMessageService() (MessageService, error) {
	return &messageService{
		Message: make(map[string]*models.Message),
	}, nil
}

func (ms messageService) ReceiveMessage(key string, msg string) (*models.Message, error) {
	result, err := models.Decrypt(key, msg)
	if err != nil {
		return nil, fmt.Errorf("Error in decrypting the message")
	}
	return result, nil
}

func (ms messageService) EncryptMessage(key string, msg string) (*models.Message, error) {
	result, _ := models.Encrypt(key, msg)

	return result, nil
}
