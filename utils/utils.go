package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// GenerateKey method declaration
func GenerateKey() string {
	key := make([]byte, 16)
	rand.Read(key)
	encodedKey := base64.URLEncoding.EncodeToString(key)
	return encodedKey
}

// DecryptMessage method declaration
func DecryptMessage(id string, key []byte, encryptedText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := encryptedText[:nonceSize], encryptedText[nonceSize:]
	decryptedText, err := gcm.Open(nil, nonce, ciphertext, nil)
	return decryptedText, err
}

// EncodeMessage method declaration
func EncodeMessage(key []byte, message []byte) string {
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := gcm.Seal(nonce, nonce, message, nil)
	enocdeText := base64.URLEncoding.EncodeToString(ciphertext)
	return enocdeText
}
