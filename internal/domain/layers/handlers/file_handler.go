package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-speech/internal/domain/layers/services"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/response"
	"github.com/iki-rumondor/go-speech/internal/utils"
)

type FileHandler struct {
	Service *services.FileService
}

func NewFileHandler(service *services.FileService) *FileHandler {
	return &FileHandler{
		Service: service,
	}
}

func (h *FileHandler) CreateVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	title := c.PostForm("title")
	if title == "" {
		utils.HandleError(c, response.BADREQ_ERR("Judul Tidak Ditemukan"))
		return
	}

	classUuid := c.PostForm("class_uuid")
	if classUuid == "" {
		utils.HandleError(c, response.BADREQ_ERR("Uuid Kelas Tidak Ditemukan"))
		return
	}

	description := c.PostForm("description")
	if description == "" {
		utils.HandleError(c, response.BADREQ_ERR("Deskripsi Tidak Ditemukan"))
		return
	}

	tempFolder := "internal/files/videos"
	videoName := utils.RandomFileName(file)
	pathFile := filepath.Join(tempFolder, videoName)

	if err := utils.SaveUploadedFile(file, pathFile); err != nil {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	if err := h.Service.CreateVideo(pathFile, videoName, title, description, classUuid); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Berhasil Menambahkan Video Pembelajaran"))
}

func (h *FileHandler) GetClassVideos(c *gin.Context) {
	classUuid := c.Param("uuid")
	resp, err := h.Service.GetClassVideos(classUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}
