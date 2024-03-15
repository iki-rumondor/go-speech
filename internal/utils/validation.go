package utils

import (
	"mime/multipart"
	"strings"

	"github.com/iki-rumondor/go-speech/internal/domain/structs/response"
)

func IsExcelFile(file *multipart.FileHeader) error {

	if fileExt := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:]); fileExt != "xlsx" {
		return response.BADREQ_ERR("Ekstensi File Harus Berupa .xlsx")
	}

	return nil
}
