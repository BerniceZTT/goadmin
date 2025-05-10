package controllers

import (
	"fmt"
	"net/http"

	"github.com/BerniceZTT/goadmin/models"
	"github.com/BerniceZTT/goadmin/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserServiceInterface
}

func NewUserController(service services.UserServiceInterface) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var userID uint
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.service.GetUser(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
