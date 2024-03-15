package services

import (
	"fmt"
	"log"

	"github.com/iki-rumondor/go-speech/internal/domain/layers/interfaces"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/models"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/request"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/response"
)

type MasterService struct {
	Repo interfaces.MasterInterface
}

func NewMasterService(repo interfaces.MasterInterface) *MasterService {
	return &MasterService{
		Repo: repo,
	}
}

func (s *MasterService) CreateClass(userUuid string, req *request.Class) error {
	var user models.User
	condition := fmt.Sprintf("uuid = '%s'", userUuid)
	if err := s.Repo.First(&user, condition); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	model := models.Class{
		TeacherID: user.Teacher.ID,
		Name:      req.Name,
		Code:      req.Code,
	}

	if err := s.Repo.Create(&model); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *MasterService) CreateDepartment(req *request.Department) error {

	model := models.Department{
		Name: req.Name,
	}

	if err := s.Repo.Create(&model); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *MasterService) UpdateDepartment(uuid string, req *request.Department) error {

	model := models.Department{
		Name: req.Name,
	}

	condition := fmt.Sprintf("uuid = '%s'", uuid)

	if err := s.Repo.Update(&model, condition); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *MasterService) DeleteDepartment(uuid string) error {

	var model models.Department
	condition := fmt.Sprintf("uuid = '%s'", uuid)
	if err := s.Repo.First(&model, condition); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	if err := s.Repo.Delete(&model, nil); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}
	return nil
}

func (s *MasterService) GetAllDepartment() (*[]response.Department, error) {

	var model []models.Department

	if err := s.Repo.Find(&model, "", "id"); err != nil {
		log.Println(err)
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Department
	for _, item := range model {
		resp = append(resp, response.Department{
			Uuid: item.Uuid,
			Name: item.Name,
		})
	}

	return &resp, nil
}

func (s *MasterService) GetDepartment(uuid string) (*response.Department, error) {

	var model models.Department
	condition := fmt.Sprintf("uuid = '%s'", uuid)
	if err := s.Repo.First(&model, condition); err != nil {
		log.Println(err)
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Department{
		Uuid: model.Uuid,
		Name: model.Name,
	}

	return &resp, nil
}
