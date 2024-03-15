package config

import (
	"github.com/iki-rumondor/go-speech/internal/domain/layers/handlers"
	"github.com/iki-rumondor/go-speech/internal/domain/layers/repositories"
	"github.com/iki-rumondor/go-speech/internal/domain/layers/services"
	"gorm.io/gorm"
)

type Handlers struct {
	MasterHandler *handlers.MasterHandler
}

func GetAppHandlers(db *gorm.DB) *Handlers {

	master_repo := repositories.NewMasterInterface(db)
	master_service := services.NewMasterService(master_repo)
	master_handler := handlers.NewMasterHandler(master_service)

	return &Handlers{
		MasterHandler: master_handler,
	}
}
