package repositories

import (
	"database/sql"
	"github.com/inact25/userbe/masters/api/models"
	"github.com/inact25/userbe/utils/queryDict"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (u UserRepoImpl) GetAllUser() ([]*models.User, error) {
	var dataUsers []*models.User
	query := queryDict.GETALLUSER
	data, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		users := models.User{}
		err := data.Scan(&users.IdentityID, &users.UserName, &users.UserBirth, &users.UserJob, &users.UserEducation)
		if err != nil {
			return nil, err
		}
		dataUsers = append(dataUsers, &users)
	}
	return dataUsers, nil
}

func (u UserRepoImpl) GetSpecificUser(user *models.User) (users []*models.User, err error) {
	var dataUsers []*models.User
	query := queryDict.GETSPECIFICUSER
	data, err := u.db.Query(query, user.IdentityID)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		users := models.User{}
		err := data.Scan(&users.IdentityID, &users.UserName, &users.UserBirth, &users.UserJob, &users.UserEducation)
		if err != nil {
			return nil, err
		}
		dataUsers = append(dataUsers, &users)
	}
	return dataUsers, nil
}

func (u UserRepoImpl) AddNewUser(user *models.User) (string, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return "", err
	}
	addUser, err := u.db.Prepare(queryDict.ADDNEWUSER)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addUser.Close()
	if _, err := addUser.Exec(user.IdentityID, user.UserName, user.UserBirth, user.UserJob, user.UserEducation); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (u UserRepoImpl) UpdateUser(user *models.User) (string, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return "", err
	}
	putUser, err := u.db.Prepare(queryDict.UPDATEUSER)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer putUser.Close()
	if _, err := putUser.Exec(user.UserName, user.UserBirth, user.UserJob, user.UserEducation); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func InitUserRepoImpl(db *sql.DB) UserRepositories {
	return &UserRepoImpl{db: db}
}
