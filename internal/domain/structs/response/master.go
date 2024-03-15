package response

type Class struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code" `
}

type Department struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}
