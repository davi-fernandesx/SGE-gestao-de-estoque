package dtos

import "time"

type CategoriasDTO struct {
	Nome      string `json:"nome"`
	Criacao   time.Time `json:"criacao"`
	Atualizao time.Time `json:"atualizao"`
}

func NewCategoriaDto(nome string) *CategoriasDTO {

	return &CategoriasDTO{
		Nome: nome,
	}
}
