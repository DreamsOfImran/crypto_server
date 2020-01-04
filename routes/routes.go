package routes

import (
	"fmt"
	"net/http"

	"github.com/DreamsOfImran/crypto_server/models"
	"github.com/DreamsOfImran/crypto_server/services"
	"github.com/gin-gonic/gin"
)

// Handler for gin routes
func Handler(engine *gin.Engine) {
	routes := engine.Group("/")

	routes.POST("/register/:id", func(ctx *gin.Context) {
		response := services.Register(ctx.Param("id"))
		ctx.JSON(http.StatusOK, gin.H{
			"ID":   fmt.Sprint(response.ID),
			"Key":  fmt.Sprintf(response.Key),
			"list": services.User,
		})
	})

	routes.POST("/send_message/:id", func(ctx *gin.Context) {
		message := &models.JSONMessage{}
		ctx.BindJSON(&message)
		response := services.SendMessage(ctx.Param("id"), message.EncryptedText)
		ctx.JSON(http.StatusOK, gin.H{
			"ID":                response.ID,
			"Decrypted Message": response.DecryptedText,
		})
	})

	routes.POST("/encrypt_message/:id", func(ctx *gin.Context) {
		message := &models.JSONMessage{}
		ctx.BindJSON(&message)
		response := services.EncryptMessage(ctx.Param("id"), message.EncryptedText)
		ctx.JSON(http.StatusOK, gin.H{
			"ID":                response.ID,
			"Encrypted Message": response.DecryptedText,
		})
	})
}
