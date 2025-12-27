package usertenants

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserTenantsController struct {
	Service *UserTenantsService
}

func NewUserTenantsController(service *UserTenantsService) *UserTenantsController {
	return &UserTenantsController{Service: service}
}

func (c *UserTenantsController) GetUsers(ctx *gin.Context) {
	userTenants, err := c.Service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, userTenants)
}

func (c *UserTenantsController) CreateUser(ctx *gin.Context) {
	var userTenants UserTenants
	if err := ctx.ShouldBindJSON(&userTenants); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.Service.CreateUserTenants(&userTenants); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
