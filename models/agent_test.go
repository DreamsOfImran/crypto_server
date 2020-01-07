package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAgent(t *testing.T) {
	t.Run("should add new agent with key", func(t *testing.T) {
		response := NewAgent("2")

		assert.Equal(t, "2", response.AgentID)
		assert.NotEmpty(t, response.Key)
	})
}
