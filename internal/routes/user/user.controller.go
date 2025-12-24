package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type USerController struct {
	Service *UserService
}

func NewUserController(service *UserService) *USerController {
	return &USerController{Service: service}
}

func (c *USerController) GetUsers(ctx *gin.Context) {
	users, err := c.Service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *USerController) CreateUser(ctx *gin.Context) {
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
