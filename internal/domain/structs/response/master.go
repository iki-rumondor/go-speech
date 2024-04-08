package response

type Class struct {
	Uuid              string `json:"uuid"`
	Name              string `json:"name"`
	Code              string `json:"code" `
	Teacher           string `json:"teacher" `
	TeacherDepartment string `json:"teacher_department" `
}

type Department struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Video struct {
	Uuid         string `json:"uuid"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	VideoName    string `json:"video_name"`
	SubtitleName string `json:"subtitle_name"`
}
