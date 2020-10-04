package main

import (
	"github.com/inact25/Golang-Basic-CRUD/configs"
	"github.com/inact25/Golang-Basic-CRUD/masters/api"
	"github.com/inact25/Golang-Basic-CRUD/utils"
)

func main() {
	conf := configs.NewAppConfig()
	db, err := configs.InitDB(conf)
	utils.ErrorCheck(err, "Print")
	myRoute := configs.CreateRouter()
	api.Init(myRoute, db)
	configs.RunServer(myRoute)
}
