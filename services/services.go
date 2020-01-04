package services

import (
	"encoding/base64"
	"fmt"

	"github.com/DreamsOfImran/crypto_server/models"
	"github.com/DreamsOfImran/crypto_server/utils"
)

// User map Declaration
var User = make(map[string]string)

// Register method declaration
func Register(id string) (*models.Registration, error) {
	if User[id] != "" {
		return nil, fmt.Errorf("User ID already exists")
	}

	User[id] = utils.GenerateKey()
	return &models.Registration{
		ID:  id,
		Key: User[id],
	}, nil
}

// SendMessage method declaration
func SendMessage(id string, encodedText string) (*models.Message, error) {
	if User[id] == "" {
		return nil, fmt.Errorf("User ID does not exists")
	}

	key, err := base64.URLEncoding.DecodeString(User[id])
	if err != nil {
		return nil, fmt.Errorf("Error in decoding the key")
	}

	decodedText, err := base64.URLEncoding.DecodeString(encodedText)
	if err != nil {
		return nil, fmt.Errorf("Error in decoding the message")
	}

	result, err := utils.DecryptMessage(id, key, decodedText)
	if err != nil {
		return nil, fmt.Errorf("Error in decrypting the message")
	}

	return &models.Message{
		ID:            id,
		DecryptedText: string(result),
	}, nil
}

// EncryptMessage method declaration
func EncryptMessage(id string, normalText string) *models.Message {
	decodedKey, _ := base64.URLEncoding.DecodeString(User[id])
	result := utils.EncodeMessage(decodedKey, []byte(normalText))

	return &models.Message{
		ID:            id,
		DecryptedText: result,
	}
}
