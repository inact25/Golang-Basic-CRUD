package usecases

import (
	"errors"
	guuid "github.com/google/uuid"
	"github.com/inact25/Golang-Basic-CRUD/masters/api/models"
	"github.com/inact25/Golang-Basic-CRUD/masters/api/repositories"
	"github.com/inact25/Golang-Basic-CRUD/utils/validation"
)

type UserUseCaseImpl struct {
	userRepo repositories.UserRepositories
}

func (u UserUseCaseImpl) GetAllUser() ([]*models.User, error) {
	user, err := u.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserUseCaseImpl) GetSpecificUser(user *models.User) ([]*models.User, error) {
	if validation.IsStatusValid(user.UserFirstName) != true {
		return nil, errors.New("ERROR")
	}
	userData, err := u.userRepo.GetSpecificUser(user)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (u UserUseCaseImpl) AddNewUser(user *models.User) (string, error) {
	user.UserID = guuid.New().String()
	err := validation.CheckEmpty(user.UserID, user.UserFirstName, user.UserLastName, user.UserAddress)
	if err != nil {
		return "", err
	}
	userData, err := u.userRepo.AddNewUser(user)
	if err != nil {
		return "", err
	}
	return userData, nil
}

func (u UserUseCaseImpl) UpdateUser(user *models.User) (string, error) {
	err := validation.CheckEmpty(user)
	if err != nil {
		return "", err
	}
	userData, err := u.userRepo.UpdateUser(user)
	if err != nil {
		return "", err
	}
	return userData, nil
}

func (u UserUseCaseImpl) DeleteUser(user *models.User) (string, error) {
	err := validation.CheckEmpty(user.UserID)
	if err != nil {
		return "", err
	}
	userData, err := u.userRepo.DeleteUser(user)
	if err != nil {
		return "", err
	}
	return userData, nil
}

func InitUserUseCase(userRepo repositories.UserRepositories) UserUsecases {
	return &UserUseCaseImpl{userRepo}
}
