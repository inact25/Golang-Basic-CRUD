package repositories

import (
	"database/sql"
	"github.com/inact25/Golang-Basic-CRUD/masters/api/models"
	"github.com/inact25/Golang-Basic-CRUD/utils/queryDict"
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
		err := data.Scan(&users.UserID, &users.UserFirstName, &users.UserLastName, &users.UserAddress)
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
	data, err := u.db.Query(query, user.UserFirstName)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		users := models.User{}
		err := data.Scan(&users.UserID, &users.UserFirstName, &users.UserLastName, &users.UserAddress)
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
	if _, err := addUser.Exec(user.UserID, user.UserFirstName, user.UserLastName, user.UserAddress); err != nil {
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
	if _, err := putUser.Exec(user.UserFirstName, user.UserLastName, user.UserAddress, user.UserID); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (u UserRepoImpl) DeleteUser(user *models.User) (string, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return "", err
	}
	putCategories, err := u.db.Prepare(queryDict.DELETEUSER)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer putCategories.Close()
	if _, err := putCategories.Exec(user.UserID); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func InitUserRepoImpl(db *sql.DB) UserRepositories {
	return &UserRepoImpl{db: db}
}
