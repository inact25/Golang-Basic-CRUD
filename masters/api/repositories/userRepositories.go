package repositories

import "github.com/inact25/userbe/masters/api/models"

type UserRepositories interface {
	GetAllUser() ([]*models.User, error)
	GetSpecificUser(user *models.User) ([]*models.User, error)
	AddNewUser(user *models.User) (string, error)
	UpdateUser(user *models.User) (string, error)
}
