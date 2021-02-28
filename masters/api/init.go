package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/inact25/userbe/masters/api/controllers"
	"github.com/inact25/userbe/masters/api/repositories"
	"github.com/inact25/userbe/masters/api/usecases"
)

func Init(r *mux.Router, db *sql.DB) {
	userRepo := repositories.InitUserRepoImpl(db)
	userUseCase := usecases.InitUserUseCase(userRepo)
	controllers.UserControl(r, userUseCase)
}
