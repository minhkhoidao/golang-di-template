package routes

import (
	"golang-DI/controllers"
	"golang-DI/services"

	"github.com/gin-gonic/gin"
)

func UserRoute(userService *services.UserService, userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("/users", userController.GetAllUsers)

	router.GET("/users/:id", userController.GetUserByID)

	return router
}
