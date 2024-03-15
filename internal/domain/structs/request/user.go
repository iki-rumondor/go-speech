package request

type User struct {
	Name           string `json:"name" valid:"required~field nama tidak ditemukan"`
	Email          string `json:"email" valid:"required~field email tidak ditemukan, email~Gunakan email yang valid"`
	Password       string `json:"password" valid:"required~field password tidak ditemukan"`
	RoleID         string `json:"role_id" valid:"required~field role tidak ditemukan"`
	Nip            string `json:"nip"`
	Nim            string `json:"nim"`
	DepartmentUuid string `json:"department_uuid"`
}

type SignIn struct {
	Email    string `json:"email" valid:"required~field email tidak ditemukan, email~Gunakan email yang valid"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}
