package routes

import (
	"github.com/BerniceZTT/goadmin/controllers"
	"github.com/BerniceZTT/goadmin/repositories"
	"github.com/BerniceZTT/goadmin/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 初始化依赖
	userRepo := &repositories.UserRepository{}
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	api := r.Group("/api")
	{
		api.GET("/users/:id", userController.GetUser)
		api.POST("/users", userController.CreateUser)
	}

	return r
}
