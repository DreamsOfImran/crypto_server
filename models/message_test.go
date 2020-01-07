package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecryptMessage(t *testing.T) {
	t.Run("should decrypt the message", func(t *testing.T) {
		key, msg := "PDQ3iVJB7ZZuv05aZO435g==", "Yj1x27hy-KeezEM6LGo4do093vSKiQZzJ7J5OLBLgojkireIrlQr"
		response, err := Decrypt(key, msg)

		assert.Nil(t, err)
		assert.Equal(t, "Imran Basha", response.DecryptedText)
	})

	t.Run("should return error for invalid key", func(t *testing.T) {
		key, msg := "SomeRandomKey", "SomeRandomEncryptedMessage"
		response, err := Decrypt(key, msg)

		assert.Nil(t, response)
		assert.Equal(t, "Invalid Key", err.Error())
	})

	t.Run("should return error for invalid encrypted message", func(t *testing.T) {
		key, msg := "PDQ3iVJB7ZZuv05aZO435g==", "SomeRandomEncryptedMessage"
		response, err := Decrypt(key, msg)

		assert.Nil(t, response)
		assert.Equal(t, "Invalid Encoded Message", err.Error())
	})

	t.Run("should return error for wrong key", func(t *testing.T) {
		key, msg := "CyBVSYu9wlHQPoPyyJZ7pw==", "Yj1x27hy-KeezEM6LGo4do093vSKiQZzJ7J5OLBLgojkireIrlQr"
		response, err := Decrypt(key, msg)

		assert.Nil(t, response)
		assert.Equal(t, "Message Authentication Failed", err.Error())
	})
}

func TestEncryptMessage(t *testing.T) {
	t.Run("should encrypt the message", func(t *testing.T) {
		key, msg := "PDQ3iVJB7ZZuv05aZO435g==", "Imran Basha"
		response, err := Encrypt(key, msg)

		assert.Nil(t, err)
		assert.NotEmpty(t, response.EncryptedText)
	})

	t.Run("should return error for invalid key", func(t *testing.T) {
		key, msg := "SomeRandomKey", "Imran Basha"
		response, err := Encrypt(key, msg)

		assert.Nil(t, response)
		assert.Equal(t, "Invalid Key", err.Error())
	})
}
