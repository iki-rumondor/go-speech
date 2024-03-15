package handlers

import "github.com/iki-rumondor/go-speech/internal/domain/layers/services"

type MasterHandler struct {
	Service *services.MasterService
}

func NewMasterHandler(service *services.MasterService) *MasterHandler {
	return &MasterHandler{
		Service: service,
	}
}
