package usertenants

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {

	repo := &UserTenantsRepo{}
	service := NewUserTenantsService(repo)
	controller := NewUserTenantsController(service)
	r.GET("/usersTenants", controller.GetUsers)
	r.POST("/usersTenants", controller.CreateUser)
}
