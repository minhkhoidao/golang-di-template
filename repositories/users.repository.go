package repositories

// repository/user_repository.go

import "golang-DI/models"

type IUserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (models.User, error)
	// Add other methods as needed
}

// UserRepository is an implementation of the Repository interface for users
type UserRepository struct {
	// Add any dependencies needed for the UserRepository
	// For example, you might have a database connection or other services
}

// GetAll retrieves all users from the database
func (ur *UserRepository) GetAll() ([]models.User, error) {
	// Simulated database query, returns a list of users
	users := []models.User{
		{ID: "1", Name: "John"},
		{ID: "2", Name: "Alice"},
		// Add more users as needed
	}
	return users, nil
}

// GetByID retrieves a user by their ID from the database
func (ur *UserRepository) GetByID(id string) (models.User, error) {
	// Simulated database query, returns a single user by ID
	user := models.User{ID: id, Name: "Some User"}
	return user, nil
}
