package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/inact25/Golang-Basic-CRUD/masters/api/models"
	"github.com/inact25/Golang-Basic-CRUD/masters/api/usecases"
	"github.com/inact25/Golang-Basic-CRUD/utils"
	"github.com/inact25/Golang-Basic-CRUD/utils/Rest"
	"net/http"
)

type UserHandler struct {
	UserUsecases usecases.UserUsecases
}

func (h UserHandler) GetAllUser(writer http.ResponseWriter, request *http.Request) {
	user, err := h.UserUsecases.GetAllUser()
	if err != nil {
		writer.Write([]byte("Data Not Found"))
	}
	var response = Rest.Res{Msg: "getAllUser", Data: user}
	byteOfCategory, err := json.Marshal(response)
	if err != nil {
		writer.Write([]byte("Something when Wrong"))
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(byteOfCategory)
}

func (h UserHandler) GetSpecificUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userFirstname := vars["firstname"]
	userModel := &models.User{}
	userModel.UserFirstName = userFirstname

	user, err := h.UserUsecases.GetSpecificUser(userModel)
	if err != nil {
		writer.Write([]byte("Data Not Found"))
	}
	var resp = Rest.Res{Msg: "getAllCategory", Data: user}
	byteOfCategory, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte("Something when Wrong"))
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(byteOfCategory)
}

func (h UserHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	user := &models.User{}
	getJsonDataCheck := json.NewDecoder(request.Body).Decode(&user)
	utils.ErrorCheck(getJsonDataCheck, "Print")
	_, err := h.UserUsecases.UpdateUser(user)
	utils.ErrorCheck(err, "Print")
	writer.Write([]byte("User Updated"))
}

func (h UserHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	user := &models.User{}
	user.UserID = vars["id"]
	_, err := h.UserUsecases.DeleteUser(user)
	if err != nil {
		writer.Write([]byte("Data Not Found"))
	}
	writer.Write([]byte("Data has been Deleted"))
}

func (h UserHandler) AddUser(writer http.ResponseWriter, request *http.Request) {
	user := &models.User{}
	getJsonDataCheck := json.NewDecoder(request.Body).Decode(&user)
	utils.ErrorCheck(getJsonDataCheck, "Print")
	_, err := h.UserUsecases.AddNewUser(user)
	utils.ErrorCheck(err, "Print")
	writer.Write([]byte("User Succesfully Added"))
}

func UserControl(r *mux.Router, service usecases.UserUsecases) {
	UserHandler := UserHandler{service}
	r.HandleFunc("/user", UserHandler.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{firstname}", UserHandler.GetSpecificUser).Methods(http.MethodGet)
	r.HandleFunc("/user", UserHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", UserHandler.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/user", UserHandler.AddUser).Methods(http.MethodPost)

}
