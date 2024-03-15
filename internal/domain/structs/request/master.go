package request

type Class struct {
	Name string `json:"name" valid:"required~field nama tidak ditemukan"`
	Code string `json:"code" valid:"required~field kode tidak ditemukan, stringlength(8|8)~field kode harus 8 karakter"`
}

type Department struct {
	Name string `json:"name" valid:"required~field nama tidak ditemukan"`
}