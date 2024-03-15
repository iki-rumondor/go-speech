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
