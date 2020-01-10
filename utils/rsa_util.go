package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// GenerateKey method declaration
func GenerateKey() (string, string) {

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	privateKeyStr, _ := changePrivateKeyToString(privateKey)
	publicKeyStr, _ := changePublicKeyToString(publicKey)

	return privateKeyStr, publicKeyStr
}

func changePrivateKeyToString(privateKey *rsa.PrivateKey) (string, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyString := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA Private Key",
			Bytes: privateKeyBytes,
		},
	)
	return string(privateKeyString), nil
}

func changePublicKeyToString(publicKey *rsa.PublicKey) (string, error) {
	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyString := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA public Key",
			Bytes: publicKeyBytes,
		},
	)
	return string(publicKeyString), nil
}

// StringToPublicKey method declaration
func StringToPublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}
	publicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	switch publicKey := publicKey.(type) {
	case *rsa.PublicKey:
		return publicKey, nil
	default:
		break
	}
	return nil, fmt.Errorf("Key type is not RSA")
}

// StringToPrivateKey declaration
func StringToPrivateKey(privateKeyStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return privateKey, nil
}
