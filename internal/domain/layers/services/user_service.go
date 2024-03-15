package services

import (
	"fmt"
	"log"
	"strconv"

	"github.com/iki-rumondor/go-speech/internal/domain/layers/interfaces"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/models"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/request"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/response"
	"github.com/iki-rumondor/go-speech/internal/utils"
)

type UserService struct {
	Repo interfaces.UserInterface
}

func NewUserService(repo interfaces.UserInterface) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(req *request.User) error {

	var model models.User

	roleID, _ := strconv.Atoi(req.RoleID)

	if roleID == 2 {
		if req.DepartmentUuid == "" {
			return response.BADREQ_ERR("Field Department Tidak Ditemukan")
		}

		if req.Nip == "" {
			return response.BADREQ_ERR("Field Nip Tidak Ditemukan")
		}

		var department models.Department
		condition := fmt.Sprintf("uuid = '%s'", req.DepartmentUuid)
		if err := s.Repo.First(&department, condition); err != nil {
			log.Println(err)
			return response.SERVICE_INTERR
		}

		model = models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			RoleID:   uint(roleID),
			Teacher: &models.Teacher{
				Nip:          req.Nip,
				DepartmentID: department.ID,
			},
		}
	} else {
		if req.Nim == "" {
			return response.BADREQ_ERR("Field Nim Tidak Ditemukan")
		}

		model = models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			RoleID:   uint(roleID),
			Student: &models.Student{
				Nim: req.Nim,
			},
		}
	}

	if err := s.Repo.Create(&model); err != nil {
		log.Println(err)
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *UserService) GetUser(uuid string) (*response.User, error) {

	var model models.User
	condition := fmt.Sprintf("uuid = '%s'", uuid)
	if err := s.Repo.First(&model, condition); err != nil {
		log.Println(err)
		return nil, response.SERVICE_INTERR
	}

	resp := response.User{
		Uuid:  model.Uuid,
		Email: model.Email,
		Name:  model.Name,
		Role:  model.Role.Name,
	}

	return &resp, nil

}

func (s *UserService) GetRoles() (*[]response.Role, error) {

	var model []models.Role

	if err := s.Repo.Find(&model, "", "id"); err != nil {
		log.Println(err)
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Role
	for _, item := range model {
		resp = append(resp, response.Role{
			ID:   fmt.Sprintf("%d", item.ID),
			Name: item.Name,
		})
	}

	return &resp, nil

}

func (s *UserService) VerifyUser(req *request.SignIn) (map[string]string, error) {
	var user models.User
	condition := fmt.Sprintf("email = '%s'", req.Email)
	if err := s.Repo.First(&user, condition); err != nil {
		log.Println(err)
		return nil, response.NOTFOUND_ERR("Email atau Password Salah")
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return nil, response.NOTFOUND_ERR("Email atau Password Salah")
	}

	jwt, err := utils.GenerateToken(user.Uuid)
	if err != nil {
		return nil, err
	}

	resp := map[string]string{
		"token": jwt,
	}

	return resp, nil

}
