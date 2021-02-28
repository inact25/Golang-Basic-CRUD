package models

type User struct {
	IdentityID    string `json:"identity"`
	UserName      string `json:"name"`
	UserBirth     string `json:"birthdate"`
	UserJob       string `json:"job"`
	UserEducation string `json:"education"`
}
