package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	t.Run("should return random crypto key", func(t *testing.T) {
		response := GenerateKey()

		assert.NotNil(t, response)
	})
}

func TestDecryptMessage(t *testing.T) {
	t.Run("should return error message for invalid key", func(t *testing.T) {
		key := []byte{1, 2, 3, 4, 5}
		encryptedMessage := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		response, err := DecryptMessage("2", key, encryptedMessage)

		assert.Nil(t, response)
		assert.Equal(t, err.Error(), "Invalid Key", "should have given invalid key")
	})

	t.Run("should return error message for authentication failure", func(t *testing.T) {
		key := []byte{60, 52, 55, 137,
			82, 65, 237, 150,
			110, 191, 78, 90,
			100, 238, 55, 230}
		encryptedMessage := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		response, err := DecryptMessage("2", key, encryptedMessage)

		assert.Nil(t, response)
		assert.Equal(t, err.Error(), "Message Authentication Failed", "should have given invalid message")
	})

	t.Run("should return decrypted message", func(t *testing.T) {
		key := []byte{60, 52, 55, 137,
			82, 65, 237, 150,
			110, 191, 78, 90,
			100, 238, 55, 230}
		encryptedMessage := []byte{98, 61, 113, 219,
			184, 114, 248, 167, 158,
			204, 67, 58, 44, 106,
			56, 118, 141, 61, 222,
			244, 138, 137, 6, 115,
			39, 178, 121, 56, 176,
			75, 130, 136, 228, 138,
			183, 136, 174, 84, 43}
		response, err := DecryptMessage("2", key, encryptedMessage)

		assert.Nil(t, err)
		assert.Equal(t, string(response), "Imran Basha", "should have passed with proper key")
	})
}
