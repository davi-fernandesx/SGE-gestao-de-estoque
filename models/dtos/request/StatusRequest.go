package dtos



type StatusRequest struct {
	Nome    string `json:"nome"`
	
}

func NewStatusDto(nome string) *StatusRequest {

	return &StatusRequest{
		Nome: nome,
	}
} 