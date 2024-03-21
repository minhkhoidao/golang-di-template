package main

import (
	"golang-DI/containers"
	"golang-DI/routes"
)

func main() {
	// Create a new instance of our dependency container
	container := containers.NewContainer()

	// Retrieve the UserService from the container
	userService := container.UserService()
	userController := container.UserController()

	router := routes.UserRoute(userService, userController)

	router.Run(":8080")
}
