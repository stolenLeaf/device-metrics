package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {

	repo := &UserRepo{}
	service := NewUserService(repo)
	controller := NewUserController(service)
	r.GET("/users", controller.GetUsers)
	r.POST("/users", controller.CreateUser)
}
