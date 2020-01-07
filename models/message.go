package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Message struct declaration
type Message struct {
	EncryptedText string
	DecryptedText string
}

// Decrypt method declaration
func Decrypt(encodedKey string, encodedMsg string) (*Message, error) {
	key, err := base64.URLEncoding.DecodeString(encodedKey)
	if err != nil {
		return nil, fmt.Errorf("Invalid Key")
	}

	msg, err := base64.URLEncoding.DecodeString(encodedMsg)
	if err != nil {
		return nil, fmt.Errorf("Invalid Encoded Message")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := msg[:nonceSize], msg[nonceSize:]
	decryptedText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("Message Authentication Failed")
	}

	return &Message{
		DecryptedText: string(decryptedText),
		EncryptedText: encodedMsg,
	}, nil
}

// Encrypt method declaration
func Encrypt(encodedKey string, normalMsg string) (*Message, error) {
	key, err := base64.URLEncoding.DecodeString(encodedKey)
	if err != nil {
		return nil, fmt.Errorf("Invalid Key")
	}

	msg := []byte(normalMsg)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := gcm.Seal(nonce, nonce, msg, nil)
	encryptedMsg := base64.URLEncoding.EncodeToString(ciphertext)
	return &Message{
		EncryptedText: encryptedMsg,
		DecryptedText: normalMsg,
	}, nil
}
