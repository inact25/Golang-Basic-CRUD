package models

type User struct {
	UserID        string `json:"userid"`
	UserFirstName string `json:"firstname"`
	UserLastName  string `json:"lastname"`
	UserAddress   string `json:"address"`
}
