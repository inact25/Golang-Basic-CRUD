package usecases

import (
	"github.com/inact25/userbe/masters/api/models"
	"github.com/inact25/userbe/masters/api/repositories"
	"github.com/inact25/userbe/utils/validation"
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
	userData, err := u.userRepo.GetSpecificUser(user)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (u UserUseCaseImpl) AddNewUser(user *models.User) (string, error) {
	err := validation.CheckEmpty(user.IdentityID, user.UserName, user.UserBirth, user.UserJob, user.UserEducation)
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

func InitUserUseCase(userRepo repositories.UserRepositories) UserUsecases {
	return &UserUseCaseImpl{userRepo}
}
