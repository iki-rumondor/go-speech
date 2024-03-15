package handlers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-speech/internal/domain/layers/services"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/request"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/response"
	"github.com/iki-rumondor/go-speech/internal/utils"
)

type MasterHandler struct {
	Service *services.MasterService
}

func NewMasterHandler(service *services.MasterService) *MasterHandler {
	return &MasterHandler{
		Service: service,
	}
}

func (h *MasterHandler) CreateClass(c *gin.Context) {
	var body request.Class
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	if err := h.Service.CreateClass(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kelas Berhasil Ditambahkan"))
}

func (h *MasterHandler) CreateDepartment(c *gin.Context) {
	var body request.Department
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateDepartment(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Program Studi Berhasil Ditambahkan"))
}

func (h *MasterHandler) UpdateDepartment(c *gin.Context) {
	var body request.Department
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateDepartment(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Program Studi Berhasil Diperbarui"))
}

func (h *MasterHandler) DeleteDepartment(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteDepartment(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Program Studi Berhasil Dihapus"))
}

func (h *MasterHandler) GetAllDepartment(c *gin.Context) {

	resp, err := h.Service.GetAllDepartment()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MasterHandler) GetDepartment(c *gin.Context) {

	uuid := c.Param("uuid")
	resp, err := h.Service.GetDepartment(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}
