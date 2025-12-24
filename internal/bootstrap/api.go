package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/stolenleaf/device-metrics/internal/database"
	"github.com/stolenleaf/device-metrics/internal/routes/user"
	"log"
)

func SetupServer() *gin.Engine {
	database.NewConnection()
	database.DB.AutoMigrate(&user.User{})

	server := gin.Default()

	setupMiddleware(server)

	api := server.Group("/api")
	user.RegisterRoutes(api)

	return server
}

func setupMiddleware(r *gin.Engine) {
	log.Println("middleware started")
}
