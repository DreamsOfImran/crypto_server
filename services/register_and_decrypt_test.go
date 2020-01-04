package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	t.Run("should register user with new ID", func(t *testing.T) {
		response, err := Register("1")

		assert.Nil(t, err)
		assert.Equal(t, response.ID, "1", "Error in Registration method with new ID")
	})

	t.Run("should return error for existing ID", func(t *testing.T) {
		response, err := Register("1")

		assert.Nil(t, response)
		assert.Equal(t, err.Error(), "User ID already exists", "should have passed existing ID")
	})
}

func TestSendMessage(t *testing.T) {
	t.Run("should return error for un-registerd ID", func(t *testing.T) {
		response, err := SendMessage("2", "SomeEncodedMessage")

		assert.Nil(t, response)
		assert.Equal(t, err.Error(), "User ID does not exists", "should have entered existing ID")
	})

	t.Run("should return error in decoding message", func(t *testing.T) {
		response, err := SendMessage("1", "ErrorEncodedMessage")

		assert.Nil(t, response)
		assert.Equal(t, err.Error(), "Error in decoding the message", "should have entered proper encoded message")
	})

	t.Run("should return decrypted message", func(t *testing.T) {
		User["2"] = "PDQ3iVJB7ZZuv05aZO435g=="
		response, err := SendMessage("2", "Yj1x27hy-KeezEM6LGo4do093vSKiQZzJ7J5OLBLgojkireIrlQr")

		assert.Nil(t, err)
		assert.Equal(t, response.DecryptedText, "Imran Basha", "should have passed with proper data")
	})
}
