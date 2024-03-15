package config

import (
	"github.com/iki-rumondor/go-speech/internal/domain/layers/handlers"
	"github.com/iki-rumondor/go-speech/internal/domain/layers/repositories"
	"github.com/iki-rumondor/go-speech/internal/domain/layers/services"
	"gorm.io/gorm"
)

type Handlers struct {
	MasterHandler *handlers.MasterHandler
	UserHandler   *handlers.UserHandler
}

func GetAppHandlers(db *gorm.DB) *Handlers {

	master_repo := repositories.NewMasterInterface(db)
	master_service := services.NewMasterService(master_repo)
	master_handler := handlers.NewMasterHandler(master_service)

	user_repo := repositories.NewUserInterface(db)
	user_service := services.NewUserService(user_repo)
	user_handler := handlers.NewUserHandler(user_service)

	return &Handlers{
		MasterHandler: master_handler,
		UserHandler:   user_handler,
	}
}
