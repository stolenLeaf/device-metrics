package main

import (
	"github.com/gin-gonic/gin"
	database "github.com/stolenleaf/device-metrics/internal/dataBase"
	"github.com/stolenleaf/device-metrics/internal/routes"
)

func main() {
	database.NewConnection()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":3001")
}
