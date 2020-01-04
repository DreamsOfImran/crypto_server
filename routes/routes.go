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
		response, err := services.Register(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": map[string]interface{}{"Error": fmt.Sprint(err)},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]interface{}{
				"ID":  response.ID,
				"Key": response.Key,
			},
		})
	})

	routes.POST("/send_message/:id", func(ctx *gin.Context) {
		message := &models.JSONMessage{}
		ctx.BindJSON(&message)
		response, err := services.SendMessage(ctx.Param("id"), message.EncryptedText)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data": map[string]interface{}{"Error": fmt.Sprint(err)},
			})
			return
		}

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
