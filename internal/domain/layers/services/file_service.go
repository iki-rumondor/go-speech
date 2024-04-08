package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/iki-rumondor/go-speech/internal/domain/layers/interfaces"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/models"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/response"
	"github.com/iki-rumondor/go-speech/internal/utils"
)

type FileService struct {
	Repo interfaces.MasterInterface
}

func NewFileService(repo interfaces.MasterInterface) *FileService {
	return &FileService{
		Repo: repo,
	}
}

func (s *FileService) CreateVideo(pathFile, videoName, title, description, classUuid string) error {
	var class models.Class
	condition := fmt.Sprintf("uuid = '%s'", classUuid)
	if err := s.Repo.First(&class, condition); err != nil {
		return response.NOTFOUND_ERR("Kelas Tidak Ditemukan")
	}

	audioName := utils.GenerateRandomString(12)
	audioBasePath := "internal/files/audio"
	audioPath := filepath.Join(audioBasePath, fmt.Sprintf("%s.mp3", audioName))

	if err := os.MkdirAll(filepath.Dir(audioPath), 0750); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	cmd := exec.Command("ffmpeg", "-i", pathFile, "-vn", "-acodec", "libmp3lame", "-q:a", "4", audioPath)

	if err := cmd.Run(); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	defer func() {
		if err := os.Remove(audioPath); err != nil {
			log.Println(err.Error())
		}
	}()

	result, err := s.Repo.UploadAudio(audioPath)
	if err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	subtitle, err := s.Repo.AudioToSubtitleTranscript(result["upload_url"].(string))
	if err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	subtitleName := fmt.Sprintf("%s.vtt", utils.GenerateRandomString(12))
	subtitleBasePath := "internal/files/subtitle"
	subtitlePath := filepath.Join(subtitleBasePath, subtitleName)

	if err := os.MkdirAll(filepath.Dir(subtitlePath), 0750); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	if err := os.WriteFile(subtitlePath, subtitle, 0644); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	model := models.Video{
		ClassID:      class.ID,
		Title:        title,
		Description:  description,
		VideoName:    videoName,
		SubtitleName: subtitleName,
	}

	if err := s.Repo.Create(&model); err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *FileService) GetClassVideos(classUuid string) (*[]response.Video, error) {

	var class models.Class
	condition := fmt.Sprintf("uuid = '%s'", classUuid)
	if err := s.Repo.First(&class, condition); err != nil {
		log.Println(err)
		return nil, response.SERVICE_INTERR
	}

	var model []models.Video
	condition = fmt.Sprintf("class_id = '%d'", class.ID)
	if err := s.Repo.Find(&model, condition, "id"); err != nil {
		log.Println(err)
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Video
	for _, item := range model {
		resp = append(resp, response.Video{
			Uuid:         item.Uuid,
			Title:        item.Title,
			Description:  item.Description,
			VideoName:    item.VideoName,
			SubtitleName: item.SubtitleName,
		})
	}

	return &resp, nil
}
