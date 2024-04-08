package interfaces

type MasterInterface interface {
	Find(dest interface{}, condition, order string) error
	First(dest interface{}, condition string) error
	Truncate(tableName string) error
	Distinct(model interface{}, column, condition string, dest *[]string) error
	Create(data interface{}) error
	Update(data interface{}, condition string) error
	Delete(data interface{}, withAssociation []string) error
	UploadAudio(audioPath string) (map[string]interface{}, error)
	AudioToTextAPI(audioPath string) error
	AudioToSubtitleTranscript(audioUrl string) ([]byte, error)
}
