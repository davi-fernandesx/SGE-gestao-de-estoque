package dtos


type CategoriasRequest struct {
	Nome      string `json:"nome"`

}

func NewCategoriaDto(nome string) *CategoriasRequest {

	return &CategoriasRequest{
		Nome: nome,
	}
}
