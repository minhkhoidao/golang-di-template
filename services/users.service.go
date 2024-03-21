package services

import (
	"golang-DI/models"
	"golang-DI/repositories"
)

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{repository: repo}
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	return us.repository.GetAll()
}

// GetUserByID retrieves a user by their ID from the repository
func (us *UserService) GetUserByID(id string) (models.User, error) {
	return us.repository.GetByID(id)
}
