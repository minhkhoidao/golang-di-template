// controller/user_controller.go
package controllers

import (
	"golang-DI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related HTTP requests
type UserController struct {
	userService *services.UserService
}

// NewUserController creates a new instance of UserController with the provided UserService
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetAllUsers handles GET request to retrieve all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID handles GET request to retrieve a user by ID
func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	// Convert id to integer, retrieve user by ID from userService, and return response
	c.JSON(http.StatusOK, gin.H{"user_id": id})
}
