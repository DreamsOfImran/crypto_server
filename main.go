package main

import (
	"github.com/DreamsOfImran/crypto_server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	routes.Handler(g)
	g.Run(":4040")
}
