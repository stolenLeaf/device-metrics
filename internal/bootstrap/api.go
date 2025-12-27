package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stolenleaf/device-metrics/internal/database"
	"github.com/stolenleaf/device-metrics/internal/routes/user"
	"github.com/stolenleaf/device-metrics/internal/routes/usertenants"
)

func SetupServer() *gin.Engine {
	database.NewConnection()
	database.DB.AutoMigrate(&user.User{}, &usertenants.UserTenants{})

	server := gin.Default()

	setupMiddleware(server)

	api := server.Group("/api")
	user.RegisterRoutes(api)
	usertenants.RegisterRoutes(api)

	return server
}

func setupMiddleware(r *gin.Engine) {
	log.Println("middleware started")
}
