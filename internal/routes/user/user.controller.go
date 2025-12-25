package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	Service *userService
}

func NewUserController(service *userService) *userController {
	return &userController{Service: service}
}

func (c *userController) GetUsers(ctx *gin.Context) {
	users, err := c.Service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.Service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
