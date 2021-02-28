package main

import (
	"github.com/inact25/userbe/configs"
	"github.com/inact25/userbe/masters/api"
	"github.com/inact25/userbe/utils"
)

func main() {
	conf := configs.NewAppConfig()
	db, err := configs.InitDB(conf)
	utils.ErrorCheck(err, "Print")
	myRoute := configs.CreateRouter()
	api.Init(myRoute, db)
	configs.RunServer(myRoute)
}
