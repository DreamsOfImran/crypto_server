package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAgentService(t *testing.T) {
	agentService := NewAgentService()

	t.Run("should add the new agent to the list", func(t *testing.T) {
		response, err := agentService.AddAgent("1")
		assert.Nil(t, err)
		assert.Equal(t, "1", response.AgentID)
	})

	t.Run("should return error for existing agent ID", func(t *testing.T) {
		response, err := agentService.AddAgent("1")
		assert.NotNil(t, response)
		assert.Equal(t, "1", response.AgentID)
		assert.Equal(t, "Agent ID already exists", err.Error())
	})

	t.Run("should find the agent from the list", func(t *testing.T) {
		response, err := agentService.FindByID("1")
		assert.Nil(t, err)
		assert.Equal(t, "1", response.AgentID)
		assert.NotEmpty(t, response.Key)
	})

	t.Run("should return error for invalid agent ID", func(t *testing.T) {
		response, err := agentService.FindByID("2")
		assert.Nil(t, response)
		assert.Equal(t, "Agent ID does not exists", err.Error())
	})
}
