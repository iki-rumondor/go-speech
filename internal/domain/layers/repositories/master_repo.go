package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	aai "github.com/AssemblyAI/assemblyai-go-sdk"
	"github.com/iki-rumondor/go-speech/internal/domain/layers/interfaces"
	"github.com/iki-rumondor/go-speech/internal/domain/structs/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MasterRepo struct {
	db *gorm.DB
}

func NewMasterInterface(db *gorm.DB) interfaces.MasterInterface {
	return &MasterRepo{
		db: db,
	}
}

func (r *MasterRepo) Create(model interface{}) error {
	return r.db.Create(model).Error
}

func (r *MasterRepo) Find(dest interface{}, condition, order string) error {
	return r.db.Preload(clause.Associations).Order(order).Find(dest, condition).Error
}

func (r *MasterRepo) First(dest interface{}, condition string) error {
	return r.db.Preload(clause.Associations).First(dest, condition).Error
}

func (r *MasterRepo) Update(model interface{}, condition string) error {
	return r.db.Where(condition).Updates(model).Error
}

func (r *MasterRepo) Delete(data interface{}, withAssociation []string) error {
	return r.db.Select(withAssociation).Delete(data).Error
}

func (r *MasterRepo) Distinct(model interface{}, column, condition string, dest *[]string) error {
	return r.db.Model(model).Distinct().Where(condition).Pluck(column, dest).Error
}

func (r *MasterRepo) Truncate(tableName string) error {
	return r.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
}

func (r *MasterRepo) FindStudentClasses(studentID uint, dest *[]models.Class) error {
	subQuery := r.db.Model(&models.ClassRequest{}).Where("student_id = ? AND status = ?", studentID, 2).Select("class_id")
	return r.db.Find(dest, "id IN (?)", subQuery).Error
}

func (r *MasterRepo) AudioToTextAPI(audioUrl string) error {
	assemblyKey := os.Getenv("ASSEMBLY_KEY")
	if assemblyKey == "" {
		return errors.New("assembly env not found")
	}

	apiUrl := "https://api.assemblyai.com/v2/transcript"

	values := map[string]string{"audio_url": audioUrl}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", assemblyKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	transcript, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Transcript:", string(transcript))

	return nil
}

func (r *MasterRepo) AudioToSubtitleTranscript(audioUrl string) ([]byte, error) {
	assemblyKey := os.Getenv("ASSEMBLY_KEY")
	if assemblyKey == "" {
		return nil, errors.New("assembly env not found")
	}

	ctx := context.Background()

	client := aai.NewClient(assemblyKey)

	transcript, err := client.Transcripts.TranscribeFromURL(ctx, audioUrl, &aai.TranscriptOptionalParams{
		// LanguageDetection: aai.Bool(true),
	})

	// log.Println(aai.ToString(transcript.Text))

	if err != nil {
		return nil, err
	}

	params := &aai.TranscriptGetSubtitlesOptions{
		CharsPerCaption: 32,
	}

	vtt, err := client.Transcripts.GetSubtitles(ctx, aai.ToString(transcript.ID), "vtt", params)
	if err != nil {
		return nil, err
	}

	return vtt, nil
}

func (r *MasterRepo) UploadAudio(audioPath string) (map[string]interface{}, error) {
	assemblyKey := os.Getenv("ASSEMBLY_KEY")
	if assemblyKey == "" {
		return nil, errors.New("assembly env not found")
	}

	apiUrl := "https://api.assemblyai.com/v2/upload"

	audioFile, err := os.ReadFile(audioPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(audioFile))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", assemblyKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
