package routes

import (
	"fmt"
	"net/http"

	"github.com/DreamsOfImran/crypto_server/models"
	"github.com/DreamsOfImran/crypto_server/services"
	"github.com/gin-gonic/gin"
)

// Handler method declaration
func Handler(engine *gin.Engine) {
	routes := engine.Group("/")
	agentService := services.NewAgentService()
	messageService := services.NewMessageService()

	routes.POST("/register/:id", func(ctx *gin.Context) {
		result, err := agentService.AddAgent(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": map[string]interface{}{"Error": fmt.Sprint(err)},
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	routes.POST("/send_message/:id", func(ctx *gin.Context) {
		message := &models.JSONMessage{}
		ctx.BindJSON(message)
		agent, err := agentService.FindByID(ctx.Param("id"))
		result, err := messageService.ReceiveMessage(agent.Key, message.EncryptedText)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": map[string]interface{}{"Error": fmt.Sprint(err)},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	routes.POST("/encrypt_message/:id", func(ctx *gin.Context) {
		message := &models.JSONMessage{}
		ctx.BindJSON(message)
		agent, err := agentService.FindByID(ctx.Param("id"))
		result, err := messageService.EncryptMessage(agent.Key, message.EncryptedText)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": map[string]interface{}{"Error": fmt.Sprint(err)},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})
}
