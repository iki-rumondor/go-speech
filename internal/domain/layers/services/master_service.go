package services

import "github.com/iki-rumondor/go-speech/internal/domain/layers/interfaces"

type MasterService struct {
	Repo interfaces.MasterInterface
}

func NewMasterService(repo interfaces.MasterInterface) *MasterService {
	return &MasterService{
		Repo: repo,
	}
}
