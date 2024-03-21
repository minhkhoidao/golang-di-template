package containers

import (
	"golang-DI/controllers"
	"golang-DI/repositories"
	"golang-DI/services"
)

type Container struct{}

func NewContainer() *Container {
	return &Container{}
}

// UserRepository retrieves the UserRepository with its dependencies
func (c *Container) UserRepository() repositories.IUserRepository {
	// You could initialize dependencies here if needed
	return &repositories.UserRepository{}
}

// UserService retrieves the UserService with its dependencies
func (c *Container) UserService() *services.UserService {
	return services.NewUserService(c.UserRepository())
}

func (c *Container) UserController() *controllers.UserController {
	return controllers.NewUserController(c.UserService())
}
