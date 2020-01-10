package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"

	"github.com/DreamsOfImran/crypto_server/utils"
)

// Message struct declaration
type Message struct {
	EncryptedText string
	DecryptedText string
}

// Decrypt method declaration
func Decrypt(privateKeyStr string, encodedMsg string) (*Message, error) {
	privateKey, _ := utils.StringToPrivateKey(privateKeyStr)
	encryptedMsg, _ := base64.URLEncoding.DecodeString(encodedMsg)
	label := []byte("")
	hash := sha256.New()

	plainText, _ := rsa.DecryptOAEP(
		hash,
		rand.Reader,
		privateKey,
		encryptedMsg,
		label,
	)

	return &Message{
		DecryptedText: string(plainText),
	}, nil
}

// Encrypt method declaration
func Encrypt(key string, normalMsg string) (*Message, error) {
	publicKey, _ := utils.StringToPublicKey(key)
	label := []byte("")
	hash := sha256.New()
	message := []byte(normalMsg)
	ciphertext, _ := rsa.EncryptOAEP(
		hash,
		rand.Reader,
		publicKey,
		message,
		label,
	)
	encodeCiphertext := base64.URLEncoding.EncodeToString(ciphertext)
	return &Message{
		EncryptedText: encodeCiphertext,
	}, nil
}
