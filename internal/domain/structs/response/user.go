package response

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type Teacher struct {
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Nip        string `json:"nip"`
	Department string `json:"department"`
	Active     bool   `json:"status"`
}

type RequestClass struct {
	Uuid      string `json:"uuid"`
	ClassName string `json:"class_name"`
	ClassCode string `json:"class_code"`
	Teacher   string `json:"teacher"`
	Status    uint   `json:"status"`
	CreatedAt int64  `json:"created_at"`
}
