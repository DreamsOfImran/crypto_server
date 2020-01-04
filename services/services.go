package services

import (
	"encoding/base64"

	"github.com/DreamsOfImran/crypto_server/models"
	"github.com/DreamsOfImran/crypto_server/utils"
)

// User map Declaration
var User = make(map[string]string)

// Register method declaration
func Register(id string) *models.Registration {
	User[id] = utils.GenerateKey()
	return &models.Registration{
		ID:  id,
		Key: User[id],
	}
}

// SendMessage method declaration
func SendMessage(id string, encodedText string) *models.Message {
	key, _ := base64.URLEncoding.DecodeString(User[id])
	decodedText, err := base64.URLEncoding.DecodeString(encodedText)
	if err != nil {
		panic(err.Error())
	}

	result := utils.DecryptMessage(id, key, decodedText)
	return &models.Message{
		ID:            id,
		DecryptedText: string(result),
	}
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
